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
func (r *DNSRecordService) New(ctx context.Context, zoneID string, body DNSRecordNewParams, opts ...option.RequestOption) (res *DNSRecordNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
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
func (r *DNSRecordService) Update(ctx context.Context, zoneID string, dnsRecordID string, body DNSRecordUpdateParams, opts ...option.RequestOption) (res *DNSRecordUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *DNSRecordService) List(ctx context.Context, zoneID string, query DNSRecordListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DNSRecordListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *DNSRecordService) ListAutoPaging(ctx context.Context, zoneID string, query DNSRecordListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DNSRecordListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneID, query, opts...))
}

// Delete DNS Record
func (r *DNSRecordService) Delete(ctx context.Context, zoneID string, dnsRecordID string, opts ...option.RequestOption) (res *DNSRecordDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
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
func (r *DNSRecordService) Export(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records/export", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// DNS Record Details
func (r *DNSRecordService) Get(ctx context.Context, zoneID string, dnsRecordID string, opts ...option.RequestOption) (res *DNSRecordGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
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
func (r *DNSRecordService) Import(ctx context.Context, zoneID string, body DNSRecordImportParams, opts ...option.RequestOption) (res *DNSRecordImportResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordImportResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/import", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Scan for common DNS records on your domain and automatically add them to your
// zone. Useful if you haven't updated your nameservers yet.
func (r *DNSRecordService) Scan(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *DNSRecordScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordScanResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/scan", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DNSRecordNewResponseDNSRecordsARecord],
// [DNSRecordNewResponseDNSRecordsAaaaRecord],
// [DNSRecordNewResponseDNSRecordsCaaRecord],
// [DNSRecordNewResponseDNSRecordsCertRecord],
// [DNSRecordNewResponseDNSRecordsCnameRecord],
// [DNSRecordNewResponseDNSRecordsDnskeyRecord],
// [DNSRecordNewResponseDNSRecordsDsRecord],
// [DNSRecordNewResponseDNSRecordsHTTPSRecord],
// [DNSRecordNewResponseDNSRecordsLocRecord],
// [DNSRecordNewResponseDNSRecordsMxRecord],
// [DNSRecordNewResponseDNSRecordsNaptrRecord],
// [DNSRecordNewResponseDNSRecordsNsRecord],
// [DNSRecordNewResponseDNSRecordsPtrRecord],
// [DNSRecordNewResponseDNSRecordsSmimeaRecord],
// [DNSRecordNewResponseDNSRecordsSrvRecord],
// [DNSRecordNewResponseDNSRecordsSshfpRecord],
// [DNSRecordNewResponseDNSRecordsSvcbRecord],
// [DNSRecordNewResponseDNSRecordsTlsaRecord],
// [DNSRecordNewResponseDNSRecordsTxtRecord] or
// [DNSRecordNewResponseDNSRecordsUriRecord].
type DNSRecordNewResponse interface {
	implementsDNSRecordNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordNewResponse)(nil)).Elem(), "")
}

type DNSRecordNewResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordNewResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsARecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsARecord]
type dnsRecordNewResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsARecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsARecordType string

const (
	DNSRecordNewResponseDNSRecordsARecordTypeA DNSRecordNewResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsARecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsARecordMeta]
type dnsRecordNewResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsARecordTTLNumber].
type DNSRecordNewResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsARecordTTLNumber1 DNSRecordNewResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordNewResponseDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsAaaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsAaaaRecord]
type dnsRecordNewResponseDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsAaaaRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordNewResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordNewResponseDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsAaaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsAaaaRecordMeta]
type dnsRecordNewResponseDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsAaaaRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsAaaaRecordTTLNumber1 DNSRecordNewResponseDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordNewResponseDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsCaaRecord]
type dnsRecordNewResponseDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsCaaRecord) implementsDNSRecordNewResponse() {}

// Components of a CAA record.
type DNSRecordNewResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                          `json:"value"`
	JSON  dnsRecordNewResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCaaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCaaRecordData]
type dnsRecordNewResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsCaaRecordType string

const (
	DNSRecordNewResponseDNSRecordsCaaRecordTypeCaa DNSRecordNewResponseDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCaaRecordMeta]
type dnsRecordNewResponseDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsCaaRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsCaaRecordTTLNumber1 DNSRecordNewResponseDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordNewResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsCertRecord]
type dnsRecordNewResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsCertRecord) implementsDNSRecordNewResponse() {}

// Components of a CERT record.
type DNSRecordNewResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                          `json:"type"`
	JSON dnsRecordNewResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCertRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCertRecordData]
type dnsRecordNewResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsCertRecordType string

const (
	DNSRecordNewResponseDNSRecordsCertRecordTypeCert DNSRecordNewResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCertRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCertRecordMeta]
type dnsRecordNewResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsCertRecordTTLNumber1 DNSRecordNewResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordNewResponseDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCnameRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsCnameRecord]
type dnsRecordNewResponseDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsCnameRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsCnameRecordType string

const (
	DNSRecordNewResponseDNSRecordsCnameRecordTypeCname DNSRecordNewResponseDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCnameRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCnameRecordMeta]
type dnsRecordNewResponseDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsCnameRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsCnameRecordTTLNumber1 DNSRecordNewResponseDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordNewResponseDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDnskeyRecordJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsDnskeyRecord]
type dnsRecordNewResponseDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsDnskeyRecord) implementsDNSRecordNewResponse() {}

// Components of a DNSKEY record.
type DNSRecordNewResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                             `json:"public_key"`
	JSON      dnsRecordNewResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDnskeyRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordNewResponseDNSRecordsDnskeyRecordData]
type dnsRecordNewResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordNewResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordNewResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDnskeyRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordNewResponseDNSRecordsDnskeyRecordMeta]
type dnsRecordNewResponseDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsDnskeyRecordTTLNumber1 DNSRecordNewResponseDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordNewResponseDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsDsRecord]
type dnsRecordNewResponseDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsDsRecord) implementsDNSRecordNewResponse() {}

// Components of a DS record.
type DNSRecordNewResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                        `json:"key_tag"`
	JSON   dnsRecordNewResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDsRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsDsRecordData]
type dnsRecordNewResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsDsRecordType string

const (
	DNSRecordNewResponseDNSRecordsDsRecordTypeDs DNSRecordNewResponseDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsDsRecordMeta]
type dnsRecordNewResponseDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsDsRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsDsRecordTTLNumber1 DNSRecordNewResponseDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordNewResponseDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsHTTPSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsHTTPSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsHTTPSRecord]
type dnsRecordNewResponseDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsHTTPSRecord) implementsDNSRecordNewResponse() {}

// Components of a HTTPS record.
type DNSRecordNewResponseDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                            `json:"value"`
	JSON  dnsRecordNewResponseDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsHTTPSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsHTTPSRecordData]
type dnsRecordNewResponseDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsHTTPSRecordType string

const (
	DNSRecordNewResponseDNSRecordsHTTPSRecordTypeHTTPS DNSRecordNewResponseDNSRecordsHTTPSRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsHTTPSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsHTTPSRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsHTTPSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsHTTPSRecordMeta]
type dnsRecordNewResponseDNSRecordsHTTPSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsHTTPSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsHTTPSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsHTTPSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsHTTPSRecordTTLNumber1 DNSRecordNewResponseDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordNewResponseDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsLocRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsLocRecord]
type dnsRecordNewResponseDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsLocRecord) implementsDNSRecordNewResponse() {}

// Components of a LOC record.
type DNSRecordNewResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordNewResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordNewResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                         `json:"size"`
	JSON dnsRecordNewResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsLocRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsLocRecordData]
type dnsRecordNewResponseDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordNewResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordNewResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordNewResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordNewResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordNewResponseDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordNewResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordNewResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordNewResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordNewResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordNewResponseDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordNewResponseDNSRecordsLocRecordType string

const (
	DNSRecordNewResponseDNSRecordsLocRecordTypeLoc DNSRecordNewResponseDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsLocRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsLocRecordMeta]
type dnsRecordNewResponseDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsLocRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsLocRecordTTLNumber1 DNSRecordNewResponseDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsMxRecord]
type dnsRecordNewResponseDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsMxRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsMxRecordType string

const (
	DNSRecordNewResponseDNSRecordsMxRecordTypeMx DNSRecordNewResponseDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsMxRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsMxRecordMeta]
type dnsRecordNewResponseDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsMxRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsMxRecordTTLNumber1 DNSRecordNewResponseDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordNewResponseDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNaptrRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsNaptrRecord]
type dnsRecordNewResponseDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsNaptrRecord) implementsDNSRecordNewResponse() {}

// Components of a NAPTR record.
type DNSRecordNewResponseDNSRecordsNaptrRecordData struct {
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
	Service string                                            `json:"service"`
	JSON    dnsRecordNewResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNaptrRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsNaptrRecordData]
type dnsRecordNewResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordNewResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordNewResponseDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNaptrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsNaptrRecordMeta]
type dnsRecordNewResponseDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsNaptrRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsNaptrRecordTTLNumber1 DNSRecordNewResponseDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsNsRecord]
type dnsRecordNewResponseDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsNsRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsNsRecordType string

const (
	DNSRecordNewResponseDNSRecordsNsRecordTypeNs DNSRecordNewResponseDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsNsRecordMeta]
type dnsRecordNewResponseDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsNsRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsNsRecordTTLNumber1 DNSRecordNewResponseDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsPtrRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsPtrRecord]
type dnsRecordNewResponseDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsPtrRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsPtrRecordType string

const (
	DNSRecordNewResponseDNSRecordsPtrRecordTypePtr DNSRecordNewResponseDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsPtrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsPtrRecordMeta]
type dnsRecordNewResponseDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsPtrRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsPtrRecordTTLNumber1 DNSRecordNewResponseDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordNewResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSmimeaRecord]
type dnsRecordNewResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSmimeaRecord) implementsDNSRecordNewResponse() {}

// Components of a SMIMEA record.
type DNSRecordNewResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                            `json:"usage"`
	JSON  dnsRecordNewResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSmimeaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordNewResponseDNSRecordsSmimeaRecordData]
type dnsRecordNewResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordNewResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordNewResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSmimeaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordNewResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordNewResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordNewResponseDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSrvRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsSrvRecord]
type dnsRecordNewResponseDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSrvRecord) implementsDNSRecordNewResponse() {}

// Components of a SRV record.
type DNSRecordNewResponseDNSRecordsSrvRecordData struct {
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
	Weight float64                                         `json:"weight"`
	JSON   dnsRecordNewResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSrvRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSrvRecordData]
type dnsRecordNewResponseDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSrvRecordType string

const (
	DNSRecordNewResponseDNSRecordsSrvRecordTypeSrv DNSRecordNewResponseDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSrvRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSrvRecordMeta]
type dnsRecordNewResponseDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSrvRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSrvRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordNewResponseDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSshfpRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsSshfpRecord]
type dnsRecordNewResponseDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSshfpRecord) implementsDNSRecordNewResponse() {}

// Components of a SSHFP record.
type DNSRecordNewResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                           `json:"type"`
	JSON dnsRecordNewResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSshfpRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSshfpRecordData]
type dnsRecordNewResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordNewResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordNewResponseDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSshfpRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSshfpRecordMeta]
type dnsRecordNewResponseDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSshfpRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSshfpRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordNewResponseDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSvcbRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsSvcbRecord]
type dnsRecordNewResponseDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSvcbRecord) implementsDNSRecordNewResponse() {}

// Components of a SVCB record.
type DNSRecordNewResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                           `json:"value"`
	JSON  dnsRecordNewResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSvcbRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSvcbRecordData]
type dnsRecordNewResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordNewResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordNewResponseDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSvcbRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSvcbRecordMeta]
type dnsRecordNewResponseDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSvcbRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSvcbRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordNewResponseDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTlsaRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsTlsaRecord]
type dnsRecordNewResponseDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsTlsaRecord) implementsDNSRecordNewResponse() {}

// Components of a TLSA record.
type DNSRecordNewResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                          `json:"usage"`
	JSON  dnsRecordNewResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTlsaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsTlsaRecordData]
type dnsRecordNewResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordNewResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordNewResponseDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTlsaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsTlsaRecordMeta]
type dnsRecordNewResponseDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsTlsaRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsTlsaRecordTTLNumber1 DNSRecordNewResponseDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTxtRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsTxtRecord]
type dnsRecordNewResponseDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsTxtRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsTxtRecordType string

const (
	DNSRecordNewResponseDNSRecordsTxtRecordTypeTxt DNSRecordNewResponseDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTxtRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsTxtRecordMeta]
type dnsRecordNewResponseDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsTxtRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsTxtRecordTTLNumber1 DNSRecordNewResponseDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordNewResponseDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsUriRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsUriRecord]
type dnsRecordNewResponseDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsUriRecord) implementsDNSRecordNewResponse() {}

// Components of a URI record.
type DNSRecordNewResponseDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                         `json:"weight"`
	JSON   dnsRecordNewResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsUriRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsUriRecordData]
type dnsRecordNewResponseDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsUriRecordType string

const (
	DNSRecordNewResponseDNSRecordsUriRecordTypeUri DNSRecordNewResponseDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsUriRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsUriRecordMeta]
type dnsRecordNewResponseDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsUriRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsUriRecordTTLNumber1 DNSRecordNewResponseDNSRecordsUriRecordTTLNumber = 1
)

// Union satisfied by [DNSRecordUpdateResponseDNSRecordsARecord],
// [DNSRecordUpdateResponseDNSRecordsAaaaRecord],
// [DNSRecordUpdateResponseDNSRecordsCaaRecord],
// [DNSRecordUpdateResponseDNSRecordsCertRecord],
// [DNSRecordUpdateResponseDNSRecordsCnameRecord],
// [DNSRecordUpdateResponseDNSRecordsDnskeyRecord],
// [DNSRecordUpdateResponseDNSRecordsDsRecord],
// [DNSRecordUpdateResponseDNSRecordsHTTPSRecord],
// [DNSRecordUpdateResponseDNSRecordsLocRecord],
// [DNSRecordUpdateResponseDNSRecordsMxRecord],
// [DNSRecordUpdateResponseDNSRecordsNaptrRecord],
// [DNSRecordUpdateResponseDNSRecordsNsRecord],
// [DNSRecordUpdateResponseDNSRecordsPtrRecord],
// [DNSRecordUpdateResponseDNSRecordsSmimeaRecord],
// [DNSRecordUpdateResponseDNSRecordsSrvRecord],
// [DNSRecordUpdateResponseDNSRecordsSshfpRecord],
// [DNSRecordUpdateResponseDNSRecordsSvcbRecord],
// [DNSRecordUpdateResponseDNSRecordsTlsaRecord],
// [DNSRecordUpdateResponseDNSRecordsTxtRecord] or
// [DNSRecordUpdateResponseDNSRecordsUriRecord].
type DNSRecordUpdateResponse interface {
	implementsDNSRecordUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordUpdateResponse)(nil)).Elem(), "")
}

type DNSRecordUpdateResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsARecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsARecord]
type dnsRecordUpdateResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsARecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsARecordType string

const (
	DNSRecordUpdateResponseDNSRecordsARecordTypeA DNSRecordUpdateResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsARecordMeta]
type dnsRecordUpdateResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsARecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsARecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsAaaaRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsAaaaRecord]
type dnsRecordUpdateResponseDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsAaaaRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordUpdateResponseDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsAaaaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsAaaaRecordMeta]
type dnsRecordUpdateResponseDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsAaaaRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsAaaaRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordUpdateResponseDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCaaRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsCaaRecord]
type dnsRecordUpdateResponseDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsCaaRecord) implementsDNSRecordUpdateResponse() {}

// Components of a CAA record.
type DNSRecordUpdateResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                             `json:"value"`
	JSON  dnsRecordUpdateResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCaaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCaaRecordData]
type dnsRecordUpdateResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsCaaRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsCaaRecordTypeCaa DNSRecordUpdateResponseDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCaaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCaaRecordMeta]
type dnsRecordUpdateResponseDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsCaaRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsCaaRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordUpdateResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCertRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsCertRecord]
type dnsRecordUpdateResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsCertRecord) implementsDNSRecordUpdateResponse() {}

// Components of a CERT record.
type DNSRecordUpdateResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                             `json:"type"`
	JSON dnsRecordUpdateResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCertRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCertRecordData]
type dnsRecordUpdateResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsCertRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsCertRecordTypeCert DNSRecordUpdateResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCertRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCertRecordMeta]
type dnsRecordUpdateResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsCertRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCnameRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsCnameRecord]
type dnsRecordUpdateResponseDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsCnameRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsCnameRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsCnameRecordTypeCname DNSRecordUpdateResponseDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCnameRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCnameRecordMeta]
type dnsRecordUpdateResponseDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsCnameRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsCnameRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordUpdateResponseDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDnskeyRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsDnskeyRecord]
type dnsRecordUpdateResponseDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsDnskeyRecord) implementsDNSRecordUpdateResponse() {}

// Components of a DNSKEY record.
type DNSRecordUpdateResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                `json:"public_key"`
	JSON      dnsRecordUpdateResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDnskeyRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsDnskeyRecordData]
type dnsRecordUpdateResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordUpdateResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDnskeyRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsDnskeyRecordMeta]
type dnsRecordUpdateResponseDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordUpdateResponseDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsDsRecord]
type dnsRecordUpdateResponseDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsDsRecord) implementsDNSRecordUpdateResponse() {}

// Components of a DS record.
type DNSRecordUpdateResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                           `json:"key_tag"`
	JSON   dnsRecordUpdateResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDsRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsDsRecordData]
type dnsRecordUpdateResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsDsRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsDsRecordTypeDs DNSRecordUpdateResponseDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsDsRecordMeta]
type dnsRecordUpdateResponseDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsDsRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsDsRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordUpdateResponseDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsHTTPSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsHTTPSRecord]
type dnsRecordUpdateResponseDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsHTTPSRecord) implementsDNSRecordUpdateResponse() {}

// Components of a HTTPS record.
type DNSRecordUpdateResponseDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                               `json:"value"`
	JSON  dnsRecordUpdateResponseDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsHTTPSRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsHTTPSRecordData]
type dnsRecordUpdateResponseDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsHTTPSRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsHTTPSRecordTypeHTTPS DNSRecordUpdateResponseDNSRecordsHTTPSRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsHTTPSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsHTTPSRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsHTTPSRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsHTTPSRecordMeta]
type dnsRecordUpdateResponseDNSRecordsHTTPSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsHTTPSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsHTTPSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordUpdateResponseDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsLocRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsLocRecord]
type dnsRecordUpdateResponseDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsLocRecord) implementsDNSRecordUpdateResponse() {}

// Components of a LOC record.
type DNSRecordUpdateResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordUpdateResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordUpdateResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                            `json:"size"`
	JSON dnsRecordUpdateResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsLocRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsLocRecordData]
type dnsRecordUpdateResponseDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordUpdateResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordUpdateResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordUpdateResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordUpdateResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordUpdateResponseDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordUpdateResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordUpdateResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordUpdateResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordUpdateResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordUpdateResponseDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordUpdateResponseDNSRecordsLocRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsLocRecordTypeLoc DNSRecordUpdateResponseDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsLocRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsLocRecordMeta]
type dnsRecordUpdateResponseDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsLocRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsLocRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsMxRecord]
type dnsRecordUpdateResponseDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsMxRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsMxRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsMxRecordTypeMx DNSRecordUpdateResponseDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsMxRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsMxRecordMeta]
type dnsRecordUpdateResponseDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsMxRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsMxRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordUpdateResponseDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNaptrRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsNaptrRecord]
type dnsRecordUpdateResponseDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsNaptrRecord) implementsDNSRecordUpdateResponse() {}

// Components of a NAPTR record.
type DNSRecordUpdateResponseDNSRecordsNaptrRecordData struct {
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
	Service string                                               `json:"service"`
	JSON    dnsRecordUpdateResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNaptrRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsNaptrRecordData]
type dnsRecordUpdateResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordUpdateResponseDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNaptrRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsNaptrRecordMeta]
type dnsRecordUpdateResponseDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsNaptrRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsNaptrRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsNsRecord]
type dnsRecordUpdateResponseDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsNsRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsNsRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsNsRecordTypeNs DNSRecordUpdateResponseDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsNsRecordMeta]
type dnsRecordUpdateResponseDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsNsRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsNsRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsPtrRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsPtrRecord]
type dnsRecordUpdateResponseDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsPtrRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsPtrRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsPtrRecordTypePtr DNSRecordUpdateResponseDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsPtrRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsPtrRecordMeta]
type dnsRecordUpdateResponseDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsPtrRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsPtrRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordUpdateResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSmimeaRecord]
type dnsRecordUpdateResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSmimeaRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SMIMEA record.
type DNSRecordUpdateResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                               `json:"usage"`
	JSON  dnsRecordUpdateResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSmimeaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSmimeaRecordData]
type dnsRecordUpdateResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordUpdateResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSmimeaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordUpdateResponseDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSrvRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSrvRecord]
type dnsRecordUpdateResponseDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSrvRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SRV record.
type DNSRecordUpdateResponseDNSRecordsSrvRecordData struct {
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
	Weight float64                                            `json:"weight"`
	JSON   dnsRecordUpdateResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSrvRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSrvRecordData]
type dnsRecordUpdateResponseDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSrvRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSrvRecordTypeSrv DNSRecordUpdateResponseDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSrvRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSrvRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSrvRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSrvRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordUpdateResponseDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSshfpRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSshfpRecord]
type dnsRecordUpdateResponseDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSshfpRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SSHFP record.
type DNSRecordUpdateResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                              `json:"type"`
	JSON dnsRecordUpdateResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSshfpRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSshfpRecordData]
type dnsRecordUpdateResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordUpdateResponseDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSshfpRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSshfpRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSshfpRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSshfpRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordUpdateResponseDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSvcbRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSvcbRecord]
type dnsRecordUpdateResponseDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSvcbRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SVCB record.
type DNSRecordUpdateResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                              `json:"value"`
	JSON  dnsRecordUpdateResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSvcbRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSvcbRecordData]
type dnsRecordUpdateResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordUpdateResponseDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSvcbRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSvcbRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSvcbRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSvcbRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordUpdateResponseDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTlsaRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsTlsaRecord]
type dnsRecordUpdateResponseDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsTlsaRecord) implementsDNSRecordUpdateResponse() {}

// Components of a TLSA record.
type DNSRecordUpdateResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                             `json:"usage"`
	JSON  dnsRecordUpdateResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTlsaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsTlsaRecordData]
type dnsRecordUpdateResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordUpdateResponseDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTlsaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsTlsaRecordMeta]
type dnsRecordUpdateResponseDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsTlsaRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsTlsaRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTxtRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsTxtRecord]
type dnsRecordUpdateResponseDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsTxtRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsTxtRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsTxtRecordTypeTxt DNSRecordUpdateResponseDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTxtRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsTxtRecordMeta]
type dnsRecordUpdateResponseDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsTxtRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsTxtRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordUpdateResponseDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsUriRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsUriRecord]
type dnsRecordUpdateResponseDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsUriRecord) implementsDNSRecordUpdateResponse() {}

// Components of a URI record.
type DNSRecordUpdateResponseDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                            `json:"weight"`
	JSON   dnsRecordUpdateResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsUriRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsUriRecordData]
type dnsRecordUpdateResponseDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsUriRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsUriRecordTypeUri DNSRecordUpdateResponseDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsUriRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsUriRecordMeta]
type dnsRecordUpdateResponseDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsUriRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsUriRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsUriRecordTTLNumber = 1
)

// Union satisfied by [DNSRecordListResponseDNSRecordsARecord],
// [DNSRecordListResponseDNSRecordsAaaaRecord],
// [DNSRecordListResponseDNSRecordsCaaRecord],
// [DNSRecordListResponseDNSRecordsCertRecord],
// [DNSRecordListResponseDNSRecordsCnameRecord],
// [DNSRecordListResponseDNSRecordsDnskeyRecord],
// [DNSRecordListResponseDNSRecordsDsRecord],
// [DNSRecordListResponseDNSRecordsHTTPSRecord],
// [DNSRecordListResponseDNSRecordsLocRecord],
// [DNSRecordListResponseDNSRecordsMxRecord],
// [DNSRecordListResponseDNSRecordsNaptrRecord],
// [DNSRecordListResponseDNSRecordsNsRecord],
// [DNSRecordListResponseDNSRecordsPtrRecord],
// [DNSRecordListResponseDNSRecordsSmimeaRecord],
// [DNSRecordListResponseDNSRecordsSrvRecord],
// [DNSRecordListResponseDNSRecordsSshfpRecord],
// [DNSRecordListResponseDNSRecordsSvcbRecord],
// [DNSRecordListResponseDNSRecordsTlsaRecord],
// [DNSRecordListResponseDNSRecordsTxtRecord] or
// [DNSRecordListResponseDNSRecordsUriRecord].
type DNSRecordListResponse interface {
	implementsDNSRecordListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordListResponse)(nil)).Elem(), "")
}

type DNSRecordListResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordListResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsARecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsARecord]
type dnsRecordListResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsARecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsARecordType string

const (
	DNSRecordListResponseDNSRecordsARecordTypeA DNSRecordListResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsARecordMeta]
type dnsRecordListResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsARecordTTLNumber].
type DNSRecordListResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsARecordTTLNumber1 DNSRecordListResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordListResponseDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsAaaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsAaaaRecord]
type dnsRecordListResponseDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsAaaaRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordListResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordListResponseDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsAaaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsAaaaRecordMeta]
type dnsRecordListResponseDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsAaaaRecordTTLNumber].
type DNSRecordListResponseDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsAaaaRecordTTLNumber1 DNSRecordListResponseDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordListResponseDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsCaaRecord]
type dnsRecordListResponseDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsCaaRecord) implementsDNSRecordListResponse() {}

// Components of a CAA record.
type DNSRecordListResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                           `json:"value"`
	JSON  dnsRecordListResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCaaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCaaRecordData]
type dnsRecordListResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsCaaRecordType string

const (
	DNSRecordListResponseDNSRecordsCaaRecordTypeCaa DNSRecordListResponseDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCaaRecordMeta]
type dnsRecordListResponseDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsCaaRecordTTLNumber].
type DNSRecordListResponseDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsCaaRecordTTLNumber1 DNSRecordListResponseDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordListResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsCertRecord]
type dnsRecordListResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsCertRecord) implementsDNSRecordListResponse() {}

// Components of a CERT record.
type DNSRecordListResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                           `json:"type"`
	JSON dnsRecordListResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCertRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCertRecordData]
type dnsRecordListResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsCertRecordType string

const (
	DNSRecordListResponseDNSRecordsCertRecordTypeCert DNSRecordListResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCertRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCertRecordMeta]
type dnsRecordListResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordListResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsCertRecordTTLNumber1 DNSRecordListResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordListResponseDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCnameRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCnameRecord]
type dnsRecordListResponseDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsCnameRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsCnameRecordType string

const (
	DNSRecordListResponseDNSRecordsCnameRecordTypeCname DNSRecordListResponseDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCnameRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsCnameRecordMeta]
type dnsRecordListResponseDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsCnameRecordTTLNumber].
type DNSRecordListResponseDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsCnameRecordTTLNumber1 DNSRecordListResponseDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordListResponseDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDnskeyRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsDnskeyRecord]
type dnsRecordListResponseDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsDnskeyRecord) implementsDNSRecordListResponse() {}

// Components of a DNSKEY record.
type DNSRecordListResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                              `json:"public_key"`
	JSON      dnsRecordListResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDnskeyRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsDnskeyRecordData]
type dnsRecordListResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordListResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordListResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDnskeyRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsDnskeyRecordMeta]
type dnsRecordListResponseDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordListResponseDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsDnskeyRecordTTLNumber1 DNSRecordListResponseDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordListResponseDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsDsRecord]
type dnsRecordListResponseDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsDsRecord) implementsDNSRecordListResponse() {}

// Components of a DS record.
type DNSRecordListResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                         `json:"key_tag"`
	JSON   dnsRecordListResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDsRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsDsRecordData]
type dnsRecordListResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsDsRecordType string

const (
	DNSRecordListResponseDNSRecordsDsRecordTypeDs DNSRecordListResponseDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsDsRecordMeta]
type dnsRecordListResponseDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsDsRecordTTLNumber].
type DNSRecordListResponseDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsDsRecordTTLNumber1 DNSRecordListResponseDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordListResponseDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsHTTPSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsHTTPSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsHTTPSRecord]
type dnsRecordListResponseDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsHTTPSRecord) implementsDNSRecordListResponse() {}

// Components of a HTTPS record.
type DNSRecordListResponseDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                             `json:"value"`
	JSON  dnsRecordListResponseDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsHTTPSRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsHTTPSRecordData]
type dnsRecordListResponseDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsHTTPSRecordType string

const (
	DNSRecordListResponseDNSRecordsHTTPSRecordTypeHTTPS DNSRecordListResponseDNSRecordsHTTPSRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsHTTPSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsHTTPSRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsHTTPSRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsHTTPSRecordMeta]
type dnsRecordListResponseDNSRecordsHTTPSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsHTTPSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordListResponseDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsHTTPSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsHTTPSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsHTTPSRecordTTLNumber1 DNSRecordListResponseDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordListResponseDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsLocRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsLocRecord]
type dnsRecordListResponseDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsLocRecord) implementsDNSRecordListResponse() {}

// Components of a LOC record.
type DNSRecordListResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordListResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordListResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                          `json:"size"`
	JSON dnsRecordListResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsLocRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsLocRecordData]
type dnsRecordListResponseDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordListResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordListResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordListResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordListResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordListResponseDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordListResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordListResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordListResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordListResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordListResponseDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordListResponseDNSRecordsLocRecordType string

const (
	DNSRecordListResponseDNSRecordsLocRecordTypeLoc DNSRecordListResponseDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsLocRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsLocRecordMeta]
type dnsRecordListResponseDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsLocRecordTTLNumber].
type DNSRecordListResponseDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsLocRecordTTLNumber1 DNSRecordListResponseDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsMxRecord]
type dnsRecordListResponseDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsMxRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsMxRecordType string

const (
	DNSRecordListResponseDNSRecordsMxRecordTypeMx DNSRecordListResponseDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsMxRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsMxRecordMeta]
type dnsRecordListResponseDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsMxRecordTTLNumber].
type DNSRecordListResponseDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsMxRecordTTLNumber1 DNSRecordListResponseDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordListResponseDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNaptrRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsNaptrRecord]
type dnsRecordListResponseDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsNaptrRecord) implementsDNSRecordListResponse() {}

// Components of a NAPTR record.
type DNSRecordListResponseDNSRecordsNaptrRecordData struct {
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
	Service string                                             `json:"service"`
	JSON    dnsRecordListResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNaptrRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsNaptrRecordData]
type dnsRecordListResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordListResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordListResponseDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNaptrRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsNaptrRecordMeta]
type dnsRecordListResponseDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsNaptrRecordTTLNumber].
type DNSRecordListResponseDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsNaptrRecordTTLNumber1 DNSRecordListResponseDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsNsRecord]
type dnsRecordListResponseDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsNsRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsNsRecordType string

const (
	DNSRecordListResponseDNSRecordsNsRecordTypeNs DNSRecordListResponseDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsNsRecordMeta]
type dnsRecordListResponseDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsNsRecordTTLNumber].
type DNSRecordListResponseDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsNsRecordTTLNumber1 DNSRecordListResponseDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsPtrRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsPtrRecord]
type dnsRecordListResponseDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsPtrRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsPtrRecordType string

const (
	DNSRecordListResponseDNSRecordsPtrRecordTypePtr DNSRecordListResponseDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsPtrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsPtrRecordMeta]
type dnsRecordListResponseDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsPtrRecordTTLNumber].
type DNSRecordListResponseDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsPtrRecordTTLNumber1 DNSRecordListResponseDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordListResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSmimeaRecord]
type dnsRecordListResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSmimeaRecord) implementsDNSRecordListResponse() {}

// Components of a SMIMEA record.
type DNSRecordListResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                             `json:"usage"`
	JSON  dnsRecordListResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSmimeaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsSmimeaRecordData]
type dnsRecordListResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordListResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordListResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSmimeaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordListResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordListResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordListResponseDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSrvRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsSrvRecord]
type dnsRecordListResponseDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSrvRecord) implementsDNSRecordListResponse() {}

// Components of a SRV record.
type DNSRecordListResponseDNSRecordsSrvRecordData struct {
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
	Weight float64                                          `json:"weight"`
	JSON   dnsRecordListResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSrvRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSrvRecordData]
type dnsRecordListResponseDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSrvRecordType string

const (
	DNSRecordListResponseDNSRecordsSrvRecordTypeSrv DNSRecordListResponseDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSrvRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSrvRecordMeta]
type dnsRecordListResponseDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSrvRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSrvRecordTTLNumber1 DNSRecordListResponseDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordListResponseDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSshfpRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSshfpRecord]
type dnsRecordListResponseDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSshfpRecord) implementsDNSRecordListResponse() {}

// Components of a SSHFP record.
type DNSRecordListResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                            `json:"type"`
	JSON dnsRecordListResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSshfpRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsSshfpRecordData]
type dnsRecordListResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordListResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordListResponseDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSshfpRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsSshfpRecordMeta]
type dnsRecordListResponseDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSshfpRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSshfpRecordTTLNumber1 DNSRecordListResponseDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordListResponseDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSvcbRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsSvcbRecord]
type dnsRecordListResponseDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSvcbRecord) implementsDNSRecordListResponse() {}

// Components of a SVCB record.
type DNSRecordListResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                            `json:"value"`
	JSON  dnsRecordListResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSvcbRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSvcbRecordData]
type dnsRecordListResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordListResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordListResponseDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSvcbRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSvcbRecordMeta]
type dnsRecordListResponseDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSvcbRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSvcbRecordTTLNumber1 DNSRecordListResponseDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordListResponseDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTlsaRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsTlsaRecord]
type dnsRecordListResponseDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsTlsaRecord) implementsDNSRecordListResponse() {}

// Components of a TLSA record.
type DNSRecordListResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                           `json:"usage"`
	JSON  dnsRecordListResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTlsaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsTlsaRecordData]
type dnsRecordListResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordListResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordListResponseDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTlsaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsTlsaRecordMeta]
type dnsRecordListResponseDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsTlsaRecordTTLNumber].
type DNSRecordListResponseDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsTlsaRecordTTLNumber1 DNSRecordListResponseDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTxtRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsTxtRecord]
type dnsRecordListResponseDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsTxtRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsTxtRecordType string

const (
	DNSRecordListResponseDNSRecordsTxtRecordTypeTxt DNSRecordListResponseDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTxtRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsTxtRecordMeta]
type dnsRecordListResponseDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsTxtRecordTTLNumber].
type DNSRecordListResponseDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsTxtRecordTTLNumber1 DNSRecordListResponseDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordListResponseDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsUriRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsUriRecord]
type dnsRecordListResponseDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsUriRecord) implementsDNSRecordListResponse() {}

// Components of a URI record.
type DNSRecordListResponseDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                          `json:"weight"`
	JSON   dnsRecordListResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsUriRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsUriRecordData]
type dnsRecordListResponseDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsUriRecordType string

const (
	DNSRecordListResponseDNSRecordsUriRecordTypeUri DNSRecordListResponseDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsUriRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsUriRecordMeta]
type dnsRecordListResponseDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsUriRecordTTLNumber].
type DNSRecordListResponseDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsUriRecordTTLNumber1 DNSRecordListResponseDNSRecordsUriRecordTTLNumber = 1
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

// Union satisfied by [DNSRecordGetResponseDNSRecordsARecord],
// [DNSRecordGetResponseDNSRecordsAaaaRecord],
// [DNSRecordGetResponseDNSRecordsCaaRecord],
// [DNSRecordGetResponseDNSRecordsCertRecord],
// [DNSRecordGetResponseDNSRecordsCnameRecord],
// [DNSRecordGetResponseDNSRecordsDnskeyRecord],
// [DNSRecordGetResponseDNSRecordsDsRecord],
// [DNSRecordGetResponseDNSRecordsHTTPSRecord],
// [DNSRecordGetResponseDNSRecordsLocRecord],
// [DNSRecordGetResponseDNSRecordsMxRecord],
// [DNSRecordGetResponseDNSRecordsNaptrRecord],
// [DNSRecordGetResponseDNSRecordsNsRecord],
// [DNSRecordGetResponseDNSRecordsPtrRecord],
// [DNSRecordGetResponseDNSRecordsSmimeaRecord],
// [DNSRecordGetResponseDNSRecordsSrvRecord],
// [DNSRecordGetResponseDNSRecordsSshfpRecord],
// [DNSRecordGetResponseDNSRecordsSvcbRecord],
// [DNSRecordGetResponseDNSRecordsTlsaRecord],
// [DNSRecordGetResponseDNSRecordsTxtRecord] or
// [DNSRecordGetResponseDNSRecordsUriRecord].
type DNSRecordGetResponse interface {
	implementsDNSRecordGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordGetResponse)(nil)).Elem(), "")
}

type DNSRecordGetResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsARecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsARecord]
type dnsRecordGetResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsARecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsARecordType string

const (
	DNSRecordGetResponseDNSRecordsARecordTypeA DNSRecordGetResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsARecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsARecordMeta]
type dnsRecordGetResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsARecordTTLNumber].
type DNSRecordGetResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsARecordTTLNumber1 DNSRecordGetResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsAaaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsAaaaRecord]
type dnsRecordGetResponseDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsAaaaRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordGetResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordGetResponseDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsAaaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsAaaaRecordMeta]
type dnsRecordGetResponseDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsAaaaRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsAaaaRecordTTLNumber1 DNSRecordGetResponseDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordGetResponseDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCaaRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsCaaRecord]
type dnsRecordGetResponseDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsCaaRecord) implementsDNSRecordGetResponse() {}

// Components of a CAA record.
type DNSRecordGetResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                          `json:"value"`
	JSON  dnsRecordGetResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCaaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCaaRecordData]
type dnsRecordGetResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsCaaRecordType string

const (
	DNSRecordGetResponseDNSRecordsCaaRecordTypeCaa DNSRecordGetResponseDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCaaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCaaRecordMeta]
type dnsRecordGetResponseDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsCaaRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsCaaRecordTTLNumber1 DNSRecordGetResponseDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordGetResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsCertRecord]
type dnsRecordGetResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsCertRecord) implementsDNSRecordGetResponse() {}

// Components of a CERT record.
type DNSRecordGetResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                          `json:"type"`
	JSON dnsRecordGetResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCertRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCertRecordData]
type dnsRecordGetResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsCertRecordType string

const (
	DNSRecordGetResponseDNSRecordsCertRecordTypeCert DNSRecordGetResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCertRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCertRecordMeta]
type dnsRecordGetResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsCertRecordTTLNumber1 DNSRecordGetResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCnameRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsCnameRecord]
type dnsRecordGetResponseDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsCnameRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsCnameRecordType string

const (
	DNSRecordGetResponseDNSRecordsCnameRecordTypeCname DNSRecordGetResponseDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCnameRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCnameRecordMeta]
type dnsRecordGetResponseDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsCnameRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsCnameRecordTTLNumber1 DNSRecordGetResponseDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordGetResponseDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDnskeyRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsDnskeyRecord]
type dnsRecordGetResponseDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsDnskeyRecord) implementsDNSRecordGetResponse() {}

// Components of a DNSKEY record.
type DNSRecordGetResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                             `json:"public_key"`
	JSON      dnsRecordGetResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDnskeyRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseDNSRecordsDnskeyRecordData]
type dnsRecordGetResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordGetResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordGetResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDnskeyRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseDNSRecordsDnskeyRecordMeta]
type dnsRecordGetResponseDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsDnskeyRecordTTLNumber1 DNSRecordGetResponseDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordGetResponseDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDsRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsDsRecord]
type dnsRecordGetResponseDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsDsRecord) implementsDNSRecordGetResponse() {}

// Components of a DS record.
type DNSRecordGetResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                        `json:"key_tag"`
	JSON   dnsRecordGetResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDsRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsDsRecordData]
type dnsRecordGetResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsDsRecordType string

const (
	DNSRecordGetResponseDNSRecordsDsRecordTypeDs DNSRecordGetResponseDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsDsRecordMeta]
type dnsRecordGetResponseDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsDsRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsDsRecordTTLNumber1 DNSRecordGetResponseDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordGetResponseDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsHTTPSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsHTTPSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsHTTPSRecord]
type dnsRecordGetResponseDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsHTTPSRecord) implementsDNSRecordGetResponse() {}

// Components of a HTTPS record.
type DNSRecordGetResponseDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                            `json:"value"`
	JSON  dnsRecordGetResponseDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsHTTPSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsHTTPSRecordData]
type dnsRecordGetResponseDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsHTTPSRecordType string

const (
	DNSRecordGetResponseDNSRecordsHTTPSRecordTypeHTTPS DNSRecordGetResponseDNSRecordsHTTPSRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsHTTPSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsHTTPSRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsHTTPSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsHTTPSRecordMeta]
type dnsRecordGetResponseDNSRecordsHTTPSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsHTTPSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsHTTPSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsHTTPSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsHTTPSRecordTTLNumber1 DNSRecordGetResponseDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordGetResponseDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsLocRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsLocRecord]
type dnsRecordGetResponseDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsLocRecord) implementsDNSRecordGetResponse() {}

// Components of a LOC record.
type DNSRecordGetResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordGetResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordGetResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                         `json:"size"`
	JSON dnsRecordGetResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsLocRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsLocRecordData]
type dnsRecordGetResponseDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordGetResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordGetResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordGetResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordGetResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordGetResponseDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordGetResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordGetResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordGetResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordGetResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordGetResponseDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordGetResponseDNSRecordsLocRecordType string

const (
	DNSRecordGetResponseDNSRecordsLocRecordTypeLoc DNSRecordGetResponseDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsLocRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsLocRecordMeta]
type dnsRecordGetResponseDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsLocRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsLocRecordTTLNumber1 DNSRecordGetResponseDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsMxRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsMxRecord]
type dnsRecordGetResponseDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsMxRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsMxRecordType string

const (
	DNSRecordGetResponseDNSRecordsMxRecordTypeMx DNSRecordGetResponseDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsMxRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsMxRecordMeta]
type dnsRecordGetResponseDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsMxRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsMxRecordTTLNumber1 DNSRecordGetResponseDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordGetResponseDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNaptrRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsNaptrRecord]
type dnsRecordGetResponseDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsNaptrRecord) implementsDNSRecordGetResponse() {}

// Components of a NAPTR record.
type DNSRecordGetResponseDNSRecordsNaptrRecordData struct {
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
	Service string                                            `json:"service"`
	JSON    dnsRecordGetResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNaptrRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsNaptrRecordData]
type dnsRecordGetResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordGetResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordGetResponseDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNaptrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsNaptrRecordMeta]
type dnsRecordGetResponseDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsNaptrRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsNaptrRecordTTLNumber1 DNSRecordGetResponseDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNsRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsNsRecord]
type dnsRecordGetResponseDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsNsRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsNsRecordType string

const (
	DNSRecordGetResponseDNSRecordsNsRecordTypeNs DNSRecordGetResponseDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsNsRecordMeta]
type dnsRecordGetResponseDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsNsRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsNsRecordTTLNumber1 DNSRecordGetResponseDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsPtrRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsPtrRecord]
type dnsRecordGetResponseDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsPtrRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsPtrRecordType string

const (
	DNSRecordGetResponseDNSRecordsPtrRecordTypePtr DNSRecordGetResponseDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsPtrRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsPtrRecordMeta]
type dnsRecordGetResponseDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsPtrRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsPtrRecordTTLNumber1 DNSRecordGetResponseDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordGetResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSmimeaRecord]
type dnsRecordGetResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSmimeaRecord) implementsDNSRecordGetResponse() {}

// Components of a SMIMEA record.
type DNSRecordGetResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                            `json:"usage"`
	JSON  dnsRecordGetResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSmimeaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseDNSRecordsSmimeaRecordData]
type dnsRecordGetResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordGetResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordGetResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSmimeaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordGetResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordGetResponseDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSrvRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsSrvRecord]
type dnsRecordGetResponseDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSrvRecord) implementsDNSRecordGetResponse() {}

// Components of a SRV record.
type DNSRecordGetResponseDNSRecordsSrvRecordData struct {
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
	Weight float64                                         `json:"weight"`
	JSON   dnsRecordGetResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSrvRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSrvRecordData]
type dnsRecordGetResponseDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSrvRecordType string

const (
	DNSRecordGetResponseDNSRecordsSrvRecordTypeSrv DNSRecordGetResponseDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSrvRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSrvRecordMeta]
type dnsRecordGetResponseDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSrvRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSrvRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordGetResponseDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSshfpRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsSshfpRecord]
type dnsRecordGetResponseDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSshfpRecord) implementsDNSRecordGetResponse() {}

// Components of a SSHFP record.
type DNSRecordGetResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                           `json:"type"`
	JSON dnsRecordGetResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSshfpRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSshfpRecordData]
type dnsRecordGetResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordGetResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordGetResponseDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSshfpRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSshfpRecordMeta]
type dnsRecordGetResponseDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSshfpRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSshfpRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordGetResponseDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSvcbRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsSvcbRecord]
type dnsRecordGetResponseDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSvcbRecord) implementsDNSRecordGetResponse() {}

// Components of a SVCB record.
type DNSRecordGetResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                           `json:"value"`
	JSON  dnsRecordGetResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSvcbRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSvcbRecordData]
type dnsRecordGetResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordGetResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordGetResponseDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSvcbRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSvcbRecordMeta]
type dnsRecordGetResponseDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSvcbRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSvcbRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordGetResponseDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTlsaRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsTlsaRecord]
type dnsRecordGetResponseDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsTlsaRecord) implementsDNSRecordGetResponse() {}

// Components of a TLSA record.
type DNSRecordGetResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                          `json:"usage"`
	JSON  dnsRecordGetResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTlsaRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsTlsaRecordData]
type dnsRecordGetResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordGetResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordGetResponseDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTlsaRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsTlsaRecordMeta]
type dnsRecordGetResponseDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsTlsaRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsTlsaRecordTTLNumber1 DNSRecordGetResponseDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTxtRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsTxtRecord]
type dnsRecordGetResponseDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsTxtRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsTxtRecordType string

const (
	DNSRecordGetResponseDNSRecordsTxtRecordTypeTxt DNSRecordGetResponseDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTxtRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsTxtRecordMeta]
type dnsRecordGetResponseDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsTxtRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsTxtRecordTTLNumber1 DNSRecordGetResponseDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordGetResponseDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsUriRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsUriRecord]
type dnsRecordGetResponseDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsUriRecord) implementsDNSRecordGetResponse() {}

// Components of a URI record.
type DNSRecordGetResponseDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                         `json:"weight"`
	JSON   dnsRecordGetResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsUriRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsUriRecordData]
type dnsRecordGetResponseDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsUriRecordType string

const (
	DNSRecordGetResponseDNSRecordsUriRecordTypeUri DNSRecordGetResponseDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsUriRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsUriRecordMeta]
type dnsRecordGetResponseDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsUriRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsUriRecordTTLNumber1 DNSRecordGetResponseDNSRecordsUriRecordTTLNumber = 1
)

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

// This interface is a union satisfied by one of the following:
// [DNSRecordNewParamsDNSRecordsARecord], [DNSRecordNewParamsDNSRecordsAaaaRecord],
// [DNSRecordNewParamsDNSRecordsCaaRecord],
// [DNSRecordNewParamsDNSRecordsCertRecord],
// [DNSRecordNewParamsDNSRecordsCnameRecord],
// [DNSRecordNewParamsDNSRecordsDnskeyRecord],
// [DNSRecordNewParamsDNSRecordsDsRecord],
// [DNSRecordNewParamsDNSRecordsHTTPSRecord],
// [DNSRecordNewParamsDNSRecordsLocRecord], [DNSRecordNewParamsDNSRecordsMxRecord],
// [DNSRecordNewParamsDNSRecordsNaptrRecord],
// [DNSRecordNewParamsDNSRecordsNsRecord], [DNSRecordNewParamsDNSRecordsPtrRecord],
// [DNSRecordNewParamsDNSRecordsSmimeaRecord],
// [DNSRecordNewParamsDNSRecordsSrvRecord],
// [DNSRecordNewParamsDNSRecordsSshfpRecord],
// [DNSRecordNewParamsDNSRecordsSvcbRecord],
// [DNSRecordNewParamsDNSRecordsTlsaRecord],
// [DNSRecordNewParamsDNSRecordsTxtRecord],
// [DNSRecordNewParamsDNSRecordsUriRecord].
type DNSRecordNewParams interface {
	ImplementsDNSRecordNewParams()
}

type DNSRecordNewParamsDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsARecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordNewParamsDNSRecordsARecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsARecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsARecordType string

const (
	DNSRecordNewParamsDNSRecordsARecordTypeA DNSRecordNewParamsDNSRecordsARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsARecordTTLNumber].
type DNSRecordNewParamsDNSRecordsARecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsARecordTTL()
}

type DNSRecordNewParamsDNSRecordsARecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsARecordTTLNumber1 DNSRecordNewParamsDNSRecordsARecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsAaaaRecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordNewParamsDNSRecordsAaaaRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsAaaaRecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsAaaaRecordType string

const (
	DNSRecordNewParamsDNSRecordsAaaaRecordTypeAaaa DNSRecordNewParamsDNSRecordsAaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsAaaaRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsAaaaRecordTTL()
}

type DNSRecordNewParamsDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsAaaaRecordTTLNumber1 DNSRecordNewParamsDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data param.Field[DNSRecordNewParamsDNSRecordsCaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsCaaRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsCaaRecord) ImplementsDNSRecordNewParams() {

}

// Components of a CAA record.
type DNSRecordNewParamsDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordNewParamsDNSRecordsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsCaaRecordType string

const (
	DNSRecordNewParamsDNSRecordsCaaRecordTypeCaa DNSRecordNewParamsDNSRecordsCaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsCaaRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsCaaRecordTTL()
}

type DNSRecordNewParamsDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsCaaRecordTTLNumber1 DNSRecordNewParamsDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data param.Field[DNSRecordNewParamsDNSRecordsCertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsCertRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsCertRecord) ImplementsDNSRecordNewParams() {

}

// Components of a CERT record.
type DNSRecordNewParamsDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordNewParamsDNSRecordsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsCertRecordType string

const (
	DNSRecordNewParamsDNSRecordsCertRecordTypeCert DNSRecordNewParamsDNSRecordsCertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsCertRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsCertRecordTTL()
}

type DNSRecordNewParamsDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsCertRecordTTLNumber1 DNSRecordNewParamsDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsCnameRecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordNewParamsDNSRecordsCnameRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsCnameRecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsCnameRecordType string

const (
	DNSRecordNewParamsDNSRecordsCnameRecordTypeCname DNSRecordNewParamsDNSRecordsCnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsCnameRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsCnameRecordTTL()
}

type DNSRecordNewParamsDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsCnameRecordTTLNumber1 DNSRecordNewParamsDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordNewParamsDNSRecordsDnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsDnskeyRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsDnskeyRecord) ImplementsDNSRecordNewParams() {

}

// Components of a DNSKEY record.
type DNSRecordNewParamsDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DNSRecordNewParamsDNSRecordsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsDnskeyRecordType string

const (
	DNSRecordNewParamsDNSRecordsDnskeyRecordTypeDnskey DNSRecordNewParamsDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsDnskeyRecordTTL()
}

type DNSRecordNewParamsDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsDnskeyRecordTTLNumber1 DNSRecordNewParamsDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data param.Field[DNSRecordNewParamsDNSRecordsDsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsDsRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsDsRecord) ImplementsDNSRecordNewParams() {

}

// Components of a DS record.
type DNSRecordNewParamsDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r DNSRecordNewParamsDNSRecordsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsDsRecordType string

const (
	DNSRecordNewParamsDNSRecordsDsRecordTypeDs DNSRecordNewParamsDNSRecordsDsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsDsRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsDsRecordTTL()
}

type DNSRecordNewParamsDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsDsRecordTTLNumber1 DNSRecordNewParamsDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data param.Field[DNSRecordNewParamsDNSRecordsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsHTTPSRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsHTTPSRecord) ImplementsDNSRecordNewParams() {

}

// Components of a HTTPS record.
type DNSRecordNewParamsDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordNewParamsDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsHTTPSRecordType string

const (
	DNSRecordNewParamsDNSRecordsHTTPSRecordTypeHTTPS DNSRecordNewParamsDNSRecordsHTTPSRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsHTTPSRecordTTL()
}

type DNSRecordNewParamsDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsHTTPSRecordTTLNumber1 DNSRecordNewParamsDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data param.Field[DNSRecordNewParamsDNSRecordsLocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsLocRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsLocRecord) ImplementsDNSRecordNewParams() {

}

// Components of a LOC record.
type DNSRecordNewParamsDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordNewParamsDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordNewParamsDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
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

func (r DNSRecordNewParamsDNSRecordsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordNewParamsDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordNewParamsDNSRecordsLocRecordDataLatDirectionN DNSRecordNewParamsDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordNewParamsDNSRecordsLocRecordDataLatDirectionS DNSRecordNewParamsDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordNewParamsDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordNewParamsDNSRecordsLocRecordDataLongDirectionE DNSRecordNewParamsDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordNewParamsDNSRecordsLocRecordDataLongDirectionW DNSRecordNewParamsDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordNewParamsDNSRecordsLocRecordType string

const (
	DNSRecordNewParamsDNSRecordsLocRecordTypeLoc DNSRecordNewParamsDNSRecordsLocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsLocRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsLocRecordTTL()
}

type DNSRecordNewParamsDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsLocRecordTTLNumber1 DNSRecordNewParamsDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsMxRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsMxRecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsMxRecordType string

const (
	DNSRecordNewParamsDNSRecordsMxRecordTypeMx DNSRecordNewParamsDNSRecordsMxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsMxRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsMxRecordTTL()
}

type DNSRecordNewParamsDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsMxRecordTTLNumber1 DNSRecordNewParamsDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[DNSRecordNewParamsDNSRecordsNaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsNaptrRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsNaptrRecord) ImplementsDNSRecordNewParams() {

}

// Components of a NAPTR record.
type DNSRecordNewParamsDNSRecordsNaptrRecordData struct {
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

func (r DNSRecordNewParamsDNSRecordsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsNaptrRecordType string

const (
	DNSRecordNewParamsDNSRecordsNaptrRecordTypeNaptr DNSRecordNewParamsDNSRecordsNaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsNaptrRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsNaptrRecordTTL()
}

type DNSRecordNewParamsDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsNaptrRecordTTLNumber1 DNSRecordNewParamsDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsNsRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsNsRecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsNsRecordType string

const (
	DNSRecordNewParamsDNSRecordsNsRecordTypeNs DNSRecordNewParamsDNSRecordsNsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsNsRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsNsRecordTTL()
}

type DNSRecordNewParamsDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsNsRecordTTLNumber1 DNSRecordNewParamsDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsPtrRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsPtrRecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsPtrRecordType string

const (
	DNSRecordNewParamsDNSRecordsPtrRecordTypePtr DNSRecordNewParamsDNSRecordsPtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsPtrRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsPtrRecordTTL()
}

type DNSRecordNewParamsDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsPtrRecordTTLNumber1 DNSRecordNewParamsDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordNewParamsDNSRecordsSmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsSmimeaRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsSmimeaRecord) ImplementsDNSRecordNewParams() {

}

// Components of a SMIMEA record.
type DNSRecordNewParamsDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordNewParamsDNSRecordsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsSmimeaRecordType string

const (
	DNSRecordNewParamsDNSRecordsSmimeaRecordTypeSmimea DNSRecordNewParamsDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSmimeaRecordTTL()
}

type DNSRecordNewParamsDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsSmimeaRecordTTLNumber1 DNSRecordNewParamsDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data param.Field[DNSRecordNewParamsDNSRecordsSrvRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsSrvRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsSrvRecord) ImplementsDNSRecordNewParams() {

}

// Components of a SRV record.
type DNSRecordNewParamsDNSRecordsSrvRecordData struct {
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

func (r DNSRecordNewParamsDNSRecordsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsSrvRecordType string

const (
	DNSRecordNewParamsDNSRecordsSrvRecordTypeSrv DNSRecordNewParamsDNSRecordsSrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsSrvRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSrvRecordTTL()
}

type DNSRecordNewParamsDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsSrvRecordTTLNumber1 DNSRecordNewParamsDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[DNSRecordNewParamsDNSRecordsSshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsSshfpRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsSshfpRecord) ImplementsDNSRecordNewParams() {

}

// Components of a SSHFP record.
type DNSRecordNewParamsDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordNewParamsDNSRecordsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsSshfpRecordType string

const (
	DNSRecordNewParamsDNSRecordsSshfpRecordTypeSshfp DNSRecordNewParamsDNSRecordsSshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsSshfpRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSshfpRecordTTL()
}

type DNSRecordNewParamsDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsSshfpRecordTTLNumber1 DNSRecordNewParamsDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[DNSRecordNewParamsDNSRecordsSvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsSvcbRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsSvcbRecord) ImplementsDNSRecordNewParams() {

}

// Components of a SVCB record.
type DNSRecordNewParamsDNSRecordsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordNewParamsDNSRecordsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsSvcbRecordType string

const (
	DNSRecordNewParamsDNSRecordsSvcbRecordTypeSvcb DNSRecordNewParamsDNSRecordsSvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsSvcbRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsSvcbRecordTTL()
}

type DNSRecordNewParamsDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsSvcbRecordTTLNumber1 DNSRecordNewParamsDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[DNSRecordNewParamsDNSRecordsTlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsTlsaRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsTlsaRecord) ImplementsDNSRecordNewParams() {

}

// Components of a TLSA record.
type DNSRecordNewParamsDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordNewParamsDNSRecordsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsTlsaRecordType string

const (
	DNSRecordNewParamsDNSRecordsTlsaRecordTypeTlsa DNSRecordNewParamsDNSRecordsTlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsTlsaRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsTlsaRecordTTL()
}

type DNSRecordNewParamsDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsTlsaRecordTTLNumber1 DNSRecordNewParamsDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsTxtRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsTxtRecord) ImplementsDNSRecordNewParams() {

}

// Record type.
type DNSRecordNewParamsDNSRecordsTxtRecordType string

const (
	DNSRecordNewParamsDNSRecordsTxtRecordTypeTxt DNSRecordNewParamsDNSRecordsTxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsTxtRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsTxtRecordTTL()
}

type DNSRecordNewParamsDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsTxtRecordTTLNumber1 DNSRecordNewParamsDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordNewParamsDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data param.Field[DNSRecordNewParamsDNSRecordsUriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[DNSRecordNewParamsDNSRecordsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordNewParamsDNSRecordsUriRecordTTL] `json:"ttl"`
}

func (r DNSRecordNewParamsDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordNewParamsDNSRecordsUriRecord) ImplementsDNSRecordNewParams() {

}

// Components of a URI record.
type DNSRecordNewParamsDNSRecordsUriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordNewParamsDNSRecordsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordNewParamsDNSRecordsUriRecordType string

const (
	DNSRecordNewParamsDNSRecordsUriRecordTypeUri DNSRecordNewParamsDNSRecordsUriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordNewParamsDNSRecordsUriRecordTTLNumber].
type DNSRecordNewParamsDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordNewParamsDNSRecordsUriRecordTTL()
}

type DNSRecordNewParamsDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordNewParamsDNSRecordsUriRecordTTLNumber1 DNSRecordNewParamsDNSRecordsUriRecordTTLNumber = 1
)

type DNSRecordNewResponseEnvelope struct {
	Errors   []DNSRecordNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordNewResponse                   `json:"result,required"`
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

// This interface is a union satisfied by one of the following:
// [DNSRecordUpdateParamsDNSRecordsARecord],
// [DNSRecordUpdateParamsDNSRecordsAaaaRecord],
// [DNSRecordUpdateParamsDNSRecordsCaaRecord],
// [DNSRecordUpdateParamsDNSRecordsCertRecord],
// [DNSRecordUpdateParamsDNSRecordsCnameRecord],
// [DNSRecordUpdateParamsDNSRecordsDnskeyRecord],
// [DNSRecordUpdateParamsDNSRecordsDsRecord],
// [DNSRecordUpdateParamsDNSRecordsHTTPSRecord],
// [DNSRecordUpdateParamsDNSRecordsLocRecord],
// [DNSRecordUpdateParamsDNSRecordsMxRecord],
// [DNSRecordUpdateParamsDNSRecordsNaptrRecord],
// [DNSRecordUpdateParamsDNSRecordsNsRecord],
// [DNSRecordUpdateParamsDNSRecordsPtrRecord],
// [DNSRecordUpdateParamsDNSRecordsSmimeaRecord],
// [DNSRecordUpdateParamsDNSRecordsSrvRecord],
// [DNSRecordUpdateParamsDNSRecordsSshfpRecord],
// [DNSRecordUpdateParamsDNSRecordsSvcbRecord],
// [DNSRecordUpdateParamsDNSRecordsTlsaRecord],
// [DNSRecordUpdateParamsDNSRecordsTxtRecord],
// [DNSRecordUpdateParamsDNSRecordsUriRecord].
type DNSRecordUpdateParams interface {
	ImplementsDNSRecordUpdateParams()
}

type DNSRecordUpdateParamsDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsARecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsARecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsARecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsARecordType string

const (
	DNSRecordUpdateParamsDNSRecordsARecordTypeA DNSRecordUpdateParamsDNSRecordsARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsARecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsARecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsARecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsARecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsARecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsARecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsAaaaRecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsAaaaRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsAaaaRecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsAaaaRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsAaaaRecordTypeAaaa DNSRecordUpdateParamsDNSRecordsAaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsAaaaRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsAaaaRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsAaaaRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsCaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsCaaRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsCaaRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a CAA record.
type DNSRecordUpdateParamsDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordUpdateParamsDNSRecordsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsCaaRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsCaaRecordTypeCaa DNSRecordUpdateParamsDNSRecordsCaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsCaaRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCaaRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsCaaRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsCertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsCertRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsCertRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a CERT record.
type DNSRecordUpdateParamsDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordUpdateParamsDNSRecordsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsCertRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsCertRecordTypeCert DNSRecordUpdateParamsDNSRecordsCertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsCertRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCertRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsCertRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsCnameRecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsCnameRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsCnameRecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsCnameRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsCnameRecordTypeCname DNSRecordUpdateParamsDNSRecordsCnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsCnameRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsCnameRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsCnameRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsDnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsDnskeyRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsDnskeyRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a DNSKEY record.
type DNSRecordUpdateParamsDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DNSRecordUpdateParamsDNSRecordsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsDnskeyRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsDnskeyRecordTypeDnskey DNSRecordUpdateParamsDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsDnskeyRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsDnskeyRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsDsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsDsRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsDsRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a DS record.
type DNSRecordUpdateParamsDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r DNSRecordUpdateParamsDNSRecordsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsDsRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsDsRecordTypeDs DNSRecordUpdateParamsDNSRecordsDsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsDsRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsDsRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsDsRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsHTTPSRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsHTTPSRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsHTTPSRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsHTTPSRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a HTTPS record.
type DNSRecordUpdateParamsDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordUpdateParamsDNSRecordsHTTPSRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsHTTPSRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsHTTPSRecordTypeHTTPS DNSRecordUpdateParamsDNSRecordsHTTPSRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsHTTPSRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsLocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsLocRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsLocRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a LOC record.
type DNSRecordUpdateParamsDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordUpdateParamsDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordUpdateParamsDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
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

func (r DNSRecordUpdateParamsDNSRecordsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordUpdateParamsDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordUpdateParamsDNSRecordsLocRecordDataLatDirectionN DNSRecordUpdateParamsDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordUpdateParamsDNSRecordsLocRecordDataLatDirectionS DNSRecordUpdateParamsDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordUpdateParamsDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordUpdateParamsDNSRecordsLocRecordDataLongDirectionE DNSRecordUpdateParamsDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordUpdateParamsDNSRecordsLocRecordDataLongDirectionW DNSRecordUpdateParamsDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordUpdateParamsDNSRecordsLocRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsLocRecordTypeLoc DNSRecordUpdateParamsDNSRecordsLocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsLocRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsLocRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsLocRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsMxRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsMxRecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsMxRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsMxRecordTypeMx DNSRecordUpdateParamsDNSRecordsMxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsMxRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsMxRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsMxRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsNaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsNaptrRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsNaptrRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a NAPTR record.
type DNSRecordUpdateParamsDNSRecordsNaptrRecordData struct {
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

func (r DNSRecordUpdateParamsDNSRecordsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsNaptrRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsNaptrRecordTypeNaptr DNSRecordUpdateParamsDNSRecordsNaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsNaptrRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsNaptrRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsNaptrRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsNsRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsNsRecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsNsRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsNsRecordTypeNs DNSRecordUpdateParamsDNSRecordsNsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsNsRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsNsRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsNsRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsPtrRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsPtrRecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsPtrRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsPtrRecordTypePtr DNSRecordUpdateParamsDNSRecordsPtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsPtrRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsPtrRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsPtrRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsSmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsSmimeaRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsSmimeaRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a SMIMEA record.
type DNSRecordUpdateParamsDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordUpdateParamsDNSRecordsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsSmimeaRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsSmimeaRecordTypeSmimea DNSRecordUpdateParamsDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSmimeaRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsSmimeaRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsSrvRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsSrvRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsSrvRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a SRV record.
type DNSRecordUpdateParamsDNSRecordsSrvRecordData struct {
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

func (r DNSRecordUpdateParamsDNSRecordsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsSrvRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsSrvRecordTypeSrv DNSRecordUpdateParamsDNSRecordsSrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsSrvRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSrvRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsSrvRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsSshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsSshfpRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsSshfpRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a SSHFP record.
type DNSRecordUpdateParamsDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordUpdateParamsDNSRecordsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsSshfpRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsSshfpRecordTypeSshfp DNSRecordUpdateParamsDNSRecordsSshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsSshfpRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSshfpRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsSshfpRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsSvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsSvcbRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsSvcbRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a SVCB record.
type DNSRecordUpdateParamsDNSRecordsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordUpdateParamsDNSRecordsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsSvcbRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsSvcbRecordTypeSvcb DNSRecordUpdateParamsDNSRecordsSvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsSvcbRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsSvcbRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsSvcbRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsTlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsTlsaRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsTlsaRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a TLSA record.
type DNSRecordUpdateParamsDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordUpdateParamsDNSRecordsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsTlsaRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsTlsaRecordTypeTlsa DNSRecordUpdateParamsDNSRecordsTlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsTlsaRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsTlsaRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsTlsaRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsTxtRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsTxtRecord) ImplementsDNSRecordUpdateParams() {

}

// Record type.
type DNSRecordUpdateParamsDNSRecordsTxtRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsTxtRecordTypeTxt DNSRecordUpdateParamsDNSRecordsTxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsTxtRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsTxtRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsTxtRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordUpdateParamsDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsUriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsUriRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsUriRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a URI record.
type DNSRecordUpdateParamsDNSRecordsUriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordUpdateParamsDNSRecordsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsUriRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsUriRecordTypeUri DNSRecordUpdateParamsDNSRecordsUriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsUriRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsUriRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsUriRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsUriRecordTTLNumber = 1
)

type DNSRecordUpdateResponseEnvelope struct {
	Errors   []DNSRecordUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordUpdateResponse                   `json:"result,required"`
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

type DNSRecordGetResponseEnvelope struct {
	Errors   []DNSRecordGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordGetResponse                   `json:"result,required"`
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
