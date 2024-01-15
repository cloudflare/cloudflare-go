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

// ZoneDNSRecordService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneDNSRecordService] method
// instead.
type ZoneDNSRecordService struct {
	Options []option.RequestOption
	Exports *ZoneDNSRecordExportService
	Imports *ZoneDNSRecordImportService
	Scans   *ZoneDNSRecordScanService
}

// NewZoneDNSRecordService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDNSRecordService(opts ...option.RequestOption) (r *ZoneDNSRecordService) {
	r = &ZoneDNSRecordService{}
	r.Options = opts
	r.Exports = NewZoneDNSRecordExportService(opts...)
	r.Imports = NewZoneDNSRecordImportService(opts...)
	r.Scans = NewZoneDNSRecordScanService(opts...)
	return
}

// DNS Record Details
func (r *ZoneDNSRecordService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing DNS record.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *ZoneDNSRecordService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneDNSRecordUpdateParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete DNS Record
func (r *ZoneDNSRecordService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneDNSRecordDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
func (r *ZoneDNSRecordService) DNSRecordsForAZoneNewDNSRecord(ctx context.Context, zoneIdentifier string, body ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *ZoneDNSRecordService) DNSRecordsForAZoneListDNSRecords(ctx context.Context, zoneIdentifier string, query ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParams, opts ...option.RequestOption) (res *shared.Page[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records", zoneIdentifier)
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

// Patch DNS Record
func (r *ZoneDNSRecordService) Patch(ctx context.Context, zoneIdentifier string, identifier string, body ZoneDNSRecordPatchParams, opts ...option.RequestOption) (res *DNSRecord, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DNSRecord struct {
	Errors   []DNSRecordError   `json:"errors"`
	Messages []DNSRecordMessage `json:"messages"`
	Result   DNSRecordResult    `json:"result"`
	// Whether the API call was successful
	Success DNSRecordSuccess `json:"success"`
	JSON    dnsRecordJSON    `json:"-"`
}

// dnsRecordJSON contains the JSON metadata for the struct [DNSRecord]
type dnsRecordJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordError struct {
	Code    int64              `json:"code,required"`
	Message string             `json:"message,required"`
	JSON    dnsRecordErrorJSON `json:"-"`
}

// dnsRecordErrorJSON contains the JSON metadata for the struct [DNSRecordError]
type dnsRecordErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordMessage struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    dnsRecordMessageJSON `json:"-"`
}

// dnsRecordMessageJSON contains the JSON metadata for the struct
// [DNSRecordMessage]
type dnsRecordMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DNSRecordResultMpBuiT95ARecord],
// [DNSRecordResultMpBuiT95AaaaRecord], [DNSRecordResultMpBuiT95CaaRecord],
// [DNSRecordResultMpBuiT95CertRecord], [DNSRecordResultMpBuiT95CnameRecord],
// [DNSRecordResultMpBuiT95DnskeyRecord], [DNSRecordResultMpBuiT95DsRecord],
// [DNSRecordResultMpBuiT95HTTPsRecord], [DNSRecordResultMpBuiT95LocRecord],
// [DNSRecordResultMpBuiT95MxRecord], [DNSRecordResultMpBuiT95NaptrRecord],
// [DNSRecordResultMpBuiT95NsRecord], [DNSRecordResultMpBuiT95PtrRecord],
// [DNSRecordResultMpBuiT95SmimeaRecord], [DNSRecordResultMpBuiT95SrvRecord],
// [DNSRecordResultMpBuiT95SshfpRecord], [DNSRecordResultMpBuiT95SvcbRecord],
// [DNSRecordResultMpBuiT95TlsaRecord], [DNSRecordResultMpBuiT95TxtRecord] or
// [DNSRecordResultMpBuiT95UriRecord].
type DNSRecordResult interface {
	implementsDNSRecordResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordResult)(nil)).Elem(), "")
}

type DNSRecordResultMpBuiT95ARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95ARecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95ARecordMeta `json:"meta"`
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
	Ttl DNSRecordResultMpBuiT95ARecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95ARecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95ARecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95ARecord]
type dnsRecordResultMpBuiT95ARecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95ARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95ARecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95ARecordType string

const (
	DNSRecordResultMpBuiT95ARecordTypeA DNSRecordResultMpBuiT95ARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95ARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                 `json:"source"`
	JSON   dnsRecordResultMpBuiT95ARecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95ARecordMetaJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95ARecordMeta]
type dnsRecordResultMpBuiT95ARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95ARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95ARecordTtlNumber].
type DNSRecordResultMpBuiT95ARecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95ARecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95ARecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95ARecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95ARecordTtlNumber1 DNSRecordResultMpBuiT95ARecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95AaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95AaaaRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95AaaaRecordMeta `json:"meta"`
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
	Ttl DNSRecordResultMpBuiT95AaaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95AaaaRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95AaaaRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95AaaaRecord]
type dnsRecordResultMpBuiT95AaaaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95AaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95AaaaRecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95AaaaRecordType string

const (
	DNSRecordResultMpBuiT95AaaaRecordTypeAaaa DNSRecordResultMpBuiT95AaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95AaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsRecordResultMpBuiT95AaaaRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95AaaaRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95AaaaRecordMeta]
type dnsRecordResultMpBuiT95AaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95AaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95AaaaRecordTtlNumber].
type DNSRecordResultMpBuiT95AaaaRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95AaaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95AaaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95AaaaRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95AaaaRecordTtlNumber1 DNSRecordResultMpBuiT95AaaaRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95CaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordResultMpBuiT95CaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95CaaRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95CaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95CaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95CaaRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CaaRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95CaaRecord]
type dnsRecordResultMpBuiT95CaaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95CaaRecord) implementsDNSRecordResult() {}

// Components of a CAA record.
type DNSRecordResultMpBuiT95CaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                   `json:"value"`
	JSON  dnsRecordResultMpBuiT95CaaRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CaaRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95CaaRecordData]
type dnsRecordResultMpBuiT95CaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95CaaRecordType string

const (
	DNSRecordResultMpBuiT95CaaRecordTypeCaa DNSRecordResultMpBuiT95CaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95CaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsRecordResultMpBuiT95CaaRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CaaRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95CaaRecordMeta]
type dnsRecordResultMpBuiT95CaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95CaaRecordTtlNumber].
type DNSRecordResultMpBuiT95CaaRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95CaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95CaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95CaaRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95CaaRecordTtlNumber1 DNSRecordResultMpBuiT95CaaRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95CertRecord struct {
	// Components of a CERT record.
	Data DNSRecordResultMpBuiT95CertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95CertRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95CertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95CertRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95CertRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CertRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95CertRecord]
type dnsRecordResultMpBuiT95CertRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95CertRecord) implementsDNSRecordResult() {}

// Components of a CERT record.
type DNSRecordResultMpBuiT95CertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                   `json:"type"`
	JSON dnsRecordResultMpBuiT95CertRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CertRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95CertRecordData]
type dnsRecordResultMpBuiT95CertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95CertRecordType string

const (
	DNSRecordResultMpBuiT95CertRecordTypeCert DNSRecordResultMpBuiT95CertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95CertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsRecordResultMpBuiT95CertRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CertRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95CertRecordMeta]
type dnsRecordResultMpBuiT95CertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95CertRecordTtlNumber].
type DNSRecordResultMpBuiT95CertRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95CertRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95CertRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95CertRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95CertRecordTtlNumber1 DNSRecordResultMpBuiT95CertRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95CnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95CnameRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95CnameRecordMeta `json:"meta"`
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
	Ttl DNSRecordResultMpBuiT95CnameRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95CnameRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CnameRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95CnameRecord]
type dnsRecordResultMpBuiT95CnameRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95CnameRecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95CnameRecordType string

const (
	DNSRecordResultMpBuiT95CnameRecordTypeCname DNSRecordResultMpBuiT95CnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95CnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsRecordResultMpBuiT95CnameRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95CnameRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95CnameRecordMeta]
type dnsRecordResultMpBuiT95CnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95CnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95CnameRecordTtlNumber].
type DNSRecordResultMpBuiT95CnameRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95CnameRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95CnameRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95CnameRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95CnameRecordTtlNumber1 DNSRecordResultMpBuiT95CnameRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95DnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordResultMpBuiT95DnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95DnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95DnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95DnskeyRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95DnskeyRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95DnskeyRecordJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95DnskeyRecord]
type dnsRecordResultMpBuiT95DnskeyRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95DnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95DnskeyRecord) implementsDNSRecordResult() {}

// Components of a DNSKEY record.
type DNSRecordResultMpBuiT95DnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                      `json:"public_key"`
	JSON      dnsRecordResultMpBuiT95DnskeyRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95DnskeyRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95DnskeyRecordData]
type dnsRecordResultMpBuiT95DnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95DnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95DnskeyRecordType string

const (
	DNSRecordResultMpBuiT95DnskeyRecordTypeDnskey DNSRecordResultMpBuiT95DnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95DnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsRecordResultMpBuiT95DnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95DnskeyRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95DnskeyRecordMeta]
type dnsRecordResultMpBuiT95DnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95DnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95DnskeyRecordTtlNumber].
type DNSRecordResultMpBuiT95DnskeyRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95DnskeyRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95DnskeyRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95DnskeyRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95DnskeyRecordTtlNumber1 DNSRecordResultMpBuiT95DnskeyRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95DsRecord struct {
	// Components of a DS record.
	Data DNSRecordResultMpBuiT95DsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95DsRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95DsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95DsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95DsRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95DsRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95DsRecord]
type dnsRecordResultMpBuiT95DsRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95DsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95DsRecord) implementsDNSRecordResult() {}

// Components of a DS record.
type DNSRecordResultMpBuiT95DsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                 `json:"key_tag"`
	JSON   dnsRecordResultMpBuiT95DsRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95DsRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95DsRecordData]
type dnsRecordResultMpBuiT95DsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95DsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95DsRecordType string

const (
	DNSRecordResultMpBuiT95DsRecordTypeDs DNSRecordResultMpBuiT95DsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95DsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                  `json:"source"`
	JSON   dnsRecordResultMpBuiT95DsRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95DsRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95DsRecordMeta]
type dnsRecordResultMpBuiT95DsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95DsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95DsRecordTtlNumber].
type DNSRecordResultMpBuiT95DsRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95DsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95DsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95DsRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95DsRecordTtlNumber1 DNSRecordResultMpBuiT95DsRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95HTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordResultMpBuiT95HTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95HTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95HTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95HTTPsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95HTTPsRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95HTTPsRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95HTTPsRecord]
type dnsRecordResultMpBuiT95HTTPsRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95HTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95HTTPsRecord) implementsDNSRecordResult() {}

// Components of a HTTPS record.
type DNSRecordResultMpBuiT95HTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                     `json:"value"`
	JSON  dnsRecordResultMpBuiT95HTTPsRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95HTTPsRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95HTTPsRecordData]
type dnsRecordResultMpBuiT95HTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95HTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95HTTPsRecordType string

const (
	DNSRecordResultMpBuiT95HTTPsRecordTypeHTTPs DNSRecordResultMpBuiT95HTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95HTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsRecordResultMpBuiT95HTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95HTTPsRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95HTTPsRecordMeta]
type dnsRecordResultMpBuiT95HTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95HTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95HTTPsRecordTtlNumber].
type DNSRecordResultMpBuiT95HTTPsRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95HTTPsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95HTTPsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95HTTPsRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95HTTPsRecordTtlNumber1 DNSRecordResultMpBuiT95HTTPsRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95LocRecord struct {
	// Components of a LOC record.
	Data DNSRecordResultMpBuiT95LocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95LocRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95LocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95LocRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95LocRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95LocRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95LocRecord]
type dnsRecordResultMpBuiT95LocRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95LocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95LocRecord) implementsDNSRecordResult() {}

// Components of a LOC record.
type DNSRecordResultMpBuiT95LocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordResultMpBuiT95LocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordResultMpBuiT95LocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                  `json:"size"`
	JSON dnsRecordResultMpBuiT95LocRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95LocRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95LocRecordData]
type dnsRecordResultMpBuiT95LocRecordDataJSON struct {
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

func (r *DNSRecordResultMpBuiT95LocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordResultMpBuiT95LocRecordDataLatDirection string

const (
	DNSRecordResultMpBuiT95LocRecordDataLatDirectionN DNSRecordResultMpBuiT95LocRecordDataLatDirection = "N"
	DNSRecordResultMpBuiT95LocRecordDataLatDirectionS DNSRecordResultMpBuiT95LocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordResultMpBuiT95LocRecordDataLongDirection string

const (
	DNSRecordResultMpBuiT95LocRecordDataLongDirectionE DNSRecordResultMpBuiT95LocRecordDataLongDirection = "E"
	DNSRecordResultMpBuiT95LocRecordDataLongDirectionW DNSRecordResultMpBuiT95LocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordResultMpBuiT95LocRecordType string

const (
	DNSRecordResultMpBuiT95LocRecordTypeLoc DNSRecordResultMpBuiT95LocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95LocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsRecordResultMpBuiT95LocRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95LocRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95LocRecordMeta]
type dnsRecordResultMpBuiT95LocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95LocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95LocRecordTtlNumber].
type DNSRecordResultMpBuiT95LocRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95LocRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95LocRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95LocRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95LocRecordTtlNumber1 DNSRecordResultMpBuiT95LocRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95MxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95MxRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95MxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95MxRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95MxRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95MxRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95MxRecord]
type dnsRecordResultMpBuiT95MxRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95MxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95MxRecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95MxRecordType string

const (
	DNSRecordResultMpBuiT95MxRecordTypeMx DNSRecordResultMpBuiT95MxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95MxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                  `json:"source"`
	JSON   dnsRecordResultMpBuiT95MxRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95MxRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95MxRecordMeta]
type dnsRecordResultMpBuiT95MxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95MxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95MxRecordTtlNumber].
type DNSRecordResultMpBuiT95MxRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95MxRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95MxRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95MxRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95MxRecordTtlNumber1 DNSRecordResultMpBuiT95MxRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95NaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordResultMpBuiT95NaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95NaptrRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95NaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95NaptrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95NaptrRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95NaptrRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95NaptrRecord]
type dnsRecordResultMpBuiT95NaptrRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95NaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95NaptrRecord) implementsDNSRecordResult() {}

// Components of a NAPTR record.
type DNSRecordResultMpBuiT95NaptrRecordData struct {
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
	Service string                                     `json:"service"`
	JSON    dnsRecordResultMpBuiT95NaptrRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95NaptrRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95NaptrRecordData]
type dnsRecordResultMpBuiT95NaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95NaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95NaptrRecordType string

const (
	DNSRecordResultMpBuiT95NaptrRecordTypeNaptr DNSRecordResultMpBuiT95NaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95NaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsRecordResultMpBuiT95NaptrRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95NaptrRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95NaptrRecordMeta]
type dnsRecordResultMpBuiT95NaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95NaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95NaptrRecordTtlNumber].
type DNSRecordResultMpBuiT95NaptrRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95NaptrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95NaptrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95NaptrRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95NaptrRecordTtlNumber1 DNSRecordResultMpBuiT95NaptrRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95NsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95NsRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95NsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95NsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95NsRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95NsRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95NsRecord]
type dnsRecordResultMpBuiT95NsRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95NsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95NsRecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95NsRecordType string

const (
	DNSRecordResultMpBuiT95NsRecordTypeNs DNSRecordResultMpBuiT95NsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95NsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                  `json:"source"`
	JSON   dnsRecordResultMpBuiT95NsRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95NsRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95NsRecordMeta]
type dnsRecordResultMpBuiT95NsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95NsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95NsRecordTtlNumber].
type DNSRecordResultMpBuiT95NsRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95NsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95NsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95NsRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95NsRecordTtlNumber1 DNSRecordResultMpBuiT95NsRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95PtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95PtrRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95PtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95PtrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95PtrRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95PtrRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95PtrRecord]
type dnsRecordResultMpBuiT95PtrRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95PtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95PtrRecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95PtrRecordType string

const (
	DNSRecordResultMpBuiT95PtrRecordTypePtr DNSRecordResultMpBuiT95PtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95PtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsRecordResultMpBuiT95PtrRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95PtrRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95PtrRecordMeta]
type dnsRecordResultMpBuiT95PtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95PtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95PtrRecordTtlNumber].
type DNSRecordResultMpBuiT95PtrRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95PtrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95PtrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95PtrRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95PtrRecordTtlNumber1 DNSRecordResultMpBuiT95PtrRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95SmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordResultMpBuiT95SmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95SmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95SmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95SmimeaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95SmimeaRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SmimeaRecordJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SmimeaRecord]
type dnsRecordResultMpBuiT95SmimeaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95SmimeaRecord) implementsDNSRecordResult() {}

// Components of a SMIMEA record.
type DNSRecordResultMpBuiT95SmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                     `json:"usage"`
	JSON  dnsRecordResultMpBuiT95SmimeaRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SmimeaRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SmimeaRecordData]
type dnsRecordResultMpBuiT95SmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95SmimeaRecordType string

const (
	DNSRecordResultMpBuiT95SmimeaRecordTypeSmimea DNSRecordResultMpBuiT95SmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95SmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsRecordResultMpBuiT95SmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SmimeaRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SmimeaRecordMeta]
type dnsRecordResultMpBuiT95SmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95SmimeaRecordTtlNumber].
type DNSRecordResultMpBuiT95SmimeaRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95SmimeaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95SmimeaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95SmimeaRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95SmimeaRecordTtlNumber1 DNSRecordResultMpBuiT95SmimeaRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95SrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordResultMpBuiT95SrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95SrvRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95SrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95SrvRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95SrvRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SrvRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95SrvRecord]
type dnsRecordResultMpBuiT95SrvRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95SrvRecord) implementsDNSRecordResult() {}

// Components of a SRV record.
type DNSRecordResultMpBuiT95SrvRecordData struct {
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
	Weight float64                                  `json:"weight"`
	JSON   dnsRecordResultMpBuiT95SrvRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SrvRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SrvRecordData]
type dnsRecordResultMpBuiT95SrvRecordDataJSON struct {
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

func (r *DNSRecordResultMpBuiT95SrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95SrvRecordType string

const (
	DNSRecordResultMpBuiT95SrvRecordTypeSrv DNSRecordResultMpBuiT95SrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95SrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsRecordResultMpBuiT95SrvRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SrvRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SrvRecordMeta]
type dnsRecordResultMpBuiT95SrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95SrvRecordTtlNumber].
type DNSRecordResultMpBuiT95SrvRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95SrvRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95SrvRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95SrvRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95SrvRecordTtlNumber1 DNSRecordResultMpBuiT95SrvRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95SshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordResultMpBuiT95SshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95SshfpRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95SshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95SshfpRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95SshfpRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SshfpRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95SshfpRecord]
type dnsRecordResultMpBuiT95SshfpRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95SshfpRecord) implementsDNSRecordResult() {}

// Components of a SSHFP record.
type DNSRecordResultMpBuiT95SshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                    `json:"type"`
	JSON dnsRecordResultMpBuiT95SshfpRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SshfpRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SshfpRecordData]
type dnsRecordResultMpBuiT95SshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95SshfpRecordType string

const (
	DNSRecordResultMpBuiT95SshfpRecordTypeSshfp DNSRecordResultMpBuiT95SshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95SshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsRecordResultMpBuiT95SshfpRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SshfpRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SshfpRecordMeta]
type dnsRecordResultMpBuiT95SshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95SshfpRecordTtlNumber].
type DNSRecordResultMpBuiT95SshfpRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95SshfpRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95SshfpRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95SshfpRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95SshfpRecordTtlNumber1 DNSRecordResultMpBuiT95SshfpRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95SvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordResultMpBuiT95SvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95SvcbRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95SvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95SvcbRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95SvcbRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SvcbRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95SvcbRecord]
type dnsRecordResultMpBuiT95SvcbRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95SvcbRecord) implementsDNSRecordResult() {}

// Components of a SVCB record.
type DNSRecordResultMpBuiT95SvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                    `json:"value"`
	JSON  dnsRecordResultMpBuiT95SvcbRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SvcbRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SvcbRecordData]
type dnsRecordResultMpBuiT95SvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95SvcbRecordType string

const (
	DNSRecordResultMpBuiT95SvcbRecordTypeSvcb DNSRecordResultMpBuiT95SvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95SvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsRecordResultMpBuiT95SvcbRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95SvcbRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95SvcbRecordMeta]
type dnsRecordResultMpBuiT95SvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95SvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95SvcbRecordTtlNumber].
type DNSRecordResultMpBuiT95SvcbRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95SvcbRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95SvcbRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95SvcbRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95SvcbRecordTtlNumber1 DNSRecordResultMpBuiT95SvcbRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95TlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordResultMpBuiT95TlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95TlsaRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95TlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95TlsaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95TlsaRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95TlsaRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95TlsaRecord]
type dnsRecordResultMpBuiT95TlsaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95TlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95TlsaRecord) implementsDNSRecordResult() {}

// Components of a TLSA record.
type DNSRecordResultMpBuiT95TlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                   `json:"usage"`
	JSON  dnsRecordResultMpBuiT95TlsaRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95TlsaRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95TlsaRecordData]
type dnsRecordResultMpBuiT95TlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95TlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95TlsaRecordType string

const (
	DNSRecordResultMpBuiT95TlsaRecordTypeTlsa DNSRecordResultMpBuiT95TlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95TlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsRecordResultMpBuiT95TlsaRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95TlsaRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95TlsaRecordMeta]
type dnsRecordResultMpBuiT95TlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95TlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95TlsaRecordTtlNumber].
type DNSRecordResultMpBuiT95TlsaRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95TlsaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95TlsaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95TlsaRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95TlsaRecordTtlNumber1 DNSRecordResultMpBuiT95TlsaRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95TxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95TxtRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95TxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95TxtRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95TxtRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95TxtRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95TxtRecord]
type dnsRecordResultMpBuiT95TxtRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95TxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95TxtRecord) implementsDNSRecordResult() {}

// Record type.
type DNSRecordResultMpBuiT95TxtRecordType string

const (
	DNSRecordResultMpBuiT95TxtRecordTypeTxt DNSRecordResultMpBuiT95TxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95TxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsRecordResultMpBuiT95TxtRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95TxtRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95TxtRecordMeta]
type dnsRecordResultMpBuiT95TxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95TxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95TxtRecordTtlNumber].
type DNSRecordResultMpBuiT95TxtRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95TxtRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95TxtRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95TxtRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95TxtRecordTtlNumber1 DNSRecordResultMpBuiT95TxtRecordTtlNumber = 1
)

type DNSRecordResultMpBuiT95UriRecord struct {
	// Components of a URI record.
	Data DNSRecordResultMpBuiT95UriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordResultMpBuiT95UriRecordType `json:"type,required"`
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
	Meta DNSRecordResultMpBuiT95UriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSRecordResultMpBuiT95UriRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordResultMpBuiT95UriRecordJSON `json:"-"`
}

// dnsRecordResultMpBuiT95UriRecordJSON contains the JSON metadata for the struct
// [DNSRecordResultMpBuiT95UriRecord]
type dnsRecordResultMpBuiT95UriRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95UriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordResultMpBuiT95UriRecord) implementsDNSRecordResult() {}

// Components of a URI record.
type DNSRecordResultMpBuiT95UriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                  `json:"weight"`
	JSON   dnsRecordResultMpBuiT95UriRecordDataJSON `json:"-"`
}

// dnsRecordResultMpBuiT95UriRecordDataJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95UriRecordData]
type dnsRecordResultMpBuiT95UriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95UriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordResultMpBuiT95UriRecordType string

const (
	DNSRecordResultMpBuiT95UriRecordTypeUri DNSRecordResultMpBuiT95UriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordResultMpBuiT95UriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsRecordResultMpBuiT95UriRecordMetaJSON `json:"-"`
}

// dnsRecordResultMpBuiT95UriRecordMetaJSON contains the JSON metadata for the
// struct [DNSRecordResultMpBuiT95UriRecordMeta]
type dnsRecordResultMpBuiT95UriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordResultMpBuiT95UriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordResultMpBuiT95UriRecordTtlNumber].
type DNSRecordResultMpBuiT95UriRecordTtl interface {
	ImplementsDNSRecordResultMpBuiT95UriRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordResultMpBuiT95UriRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordResultMpBuiT95UriRecordTtlNumber float64

const (
	DNSRecordResultMpBuiT95UriRecordTtlNumber1 DNSRecordResultMpBuiT95UriRecordTtlNumber = 1
)

// Whether the API call was successful
type DNSRecordSuccess bool

const (
	DNSRecordSuccessTrue DNSRecordSuccess = true
)

type ZoneDNSRecordDeleteResponse struct {
	Result ZoneDNSRecordDeleteResponseResult `json:"result"`
	JSON   zoneDNSRecordDeleteResponseJSON   `json:"-"`
}

// zoneDNSRecordDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneDNSRecordDeleteResponse]
type zoneDNSRecordDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSRecordDeleteResponseResult struct {
	// Identifier
	ID   string                                `json:"id"`
	JSON zoneDNSRecordDeleteResponseResultJSON `json:"-"`
}

// zoneDNSRecordDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneDNSRecordDeleteResponseResult]
type zoneDNSRecordDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecord] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecord].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse interface {
	implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse)(nil)).Elem(), "")
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMeta `json:"meta"`
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
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTypeA ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95ARecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMeta `json:"meta"`
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
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTypeAaaa ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95AaaaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecord struct {
	// Components of a CAA record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a CAA record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                                         `json:"value"`
	JSON  zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTypeCaa ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CaaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecord struct {
	// Components of a CERT record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a CERT record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                                         `json:"type"`
	JSON zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTypeCert ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CertRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMeta `json:"meta"`
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
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTypeCname ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95CnameRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecord struct {
	// Components of a DNSKEY record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a DNSKEY record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                                            `json:"public_key"`
	JSON      zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTypeDnskey ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DnskeyRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecord struct {
	// Components of a DS record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a DS record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                                                       `json:"key_tag"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTypeDs ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95DsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecord struct {
	// Components of a HTTPS record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a HTTPS record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                           `json:"value"`
	JSON  zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTypeHTTPs ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95HTTPsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecord struct {
	// Components of a LOC record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a LOC record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                                                        `json:"size"`
	JSON zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataJSON struct {
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

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLatDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLatDirectionN ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLatDirection = "N"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLatDirectionS ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLatDirection = "S"
)

// Longitude direction.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLongDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLongDirectionE ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLongDirection = "E"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLongDirectionW ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordDataLongDirection = "W"
)

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTypeLoc ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95LocRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTypeMx ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95MxRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecord struct {
	// Components of a NAPTR record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a NAPTR record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordData struct {
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
	Service string                                                                           `json:"service"`
	JSON    zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTypeNaptr ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NaptrRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTypeNs ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95NsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTypePtr ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95PtrRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecord struct {
	// Components of a SMIMEA record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SMIMEA record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                           `json:"usage"`
	JSON  zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTypeSmimea ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SmimeaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecord struct {
	// Components of a SRV record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SRV record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordData struct {
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
	Weight float64                                                                        `json:"weight"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordDataJSON struct {
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

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTypeSrv ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SrvRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecord struct {
	// Components of a SSHFP record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SSHFP record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                                          `json:"type"`
	JSON zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTypeSshfp ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SshfpRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecord struct {
	// Components of a SVCB record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SVCB record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                          `json:"value"`
	JSON  zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTypeSvcb ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95SvcbRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecord struct {
	// Components of a TLSA record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a TLSA record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                         `json:"usage"`
	JSON  zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTypeTlsa ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TlsaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTypeTxt ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95TxtRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecord struct {
	// Components of a URI record.
	Data ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordType `json:"type,required"`
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
	Meta ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecord]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordJSON struct {
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
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecord) implementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a URI record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                                                        `json:"weight"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordDataJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordDataJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordData]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTypeUri ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMetaJSON `json:"-"`
}

// zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMetaJSON
// contains the JSON metadata for the struct
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMeta]
type zoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsResponseMpBuiT95UriRecordTtlNumber = 1
)

// This interface is a union satisfied by one of the following:
// [ZoneDNSRecordUpdateParamsMpBuiT95ARecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95CaaRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95CertRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95CnameRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95DsRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95LocRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95MxRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95NsRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95PtrRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95SrvRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95TxtRecord],
// [ZoneDNSRecordUpdateParamsMpBuiT95UriRecord].
type ZoneDNSRecordUpdateParams interface {
	ImplementsZoneDNSRecordUpdateParams()
}

type ZoneDNSRecordUpdateParamsMpBuiT95ARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95ARecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95ARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95ARecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95ARecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95ARecordTypeA ZoneDNSRecordUpdateParamsMpBuiT95ARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95ARecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95ARecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTypeAaaa ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95AaaaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95CaaRecord struct {
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95CaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95CaaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a CAA record.
type ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTypeCaa ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95CaaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95CertRecord struct {
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95CertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95CertRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a CERT record.
type ZoneDNSRecordUpdateParamsMpBuiT95CertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95CertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95CertRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTypeCert ZoneDNSRecordUpdateParamsMpBuiT95CertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95CertRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95CnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95CnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95CnameRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTypeCname ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95CnameRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a DNSKEY record.
type ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTypeDnskey ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95DnskeyRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95DsRecord struct {
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95DsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95DsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95DsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95DsRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a DS record.
type ZoneDNSRecordUpdateParamsMpBuiT95DsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95DsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95DsRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTypeDs ZoneDNSRecordUpdateParamsMpBuiT95DsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95DsRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a HTTPS record.
type ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTypeHTTPs ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95HTTPsRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95LocRecord struct {
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95LocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95LocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95LocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95LocRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a LOC record.
type ZoneDNSRecordUpdateParamsMpBuiT95LocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLongDirection] `json:"long_direction"`
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

func (r ZoneDNSRecordUpdateParamsMpBuiT95LocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLatDirection string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLatDirectionN ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLatDirection = "N"
	ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLatDirectionS ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLatDirection = "S"
)

// Longitude direction.
type ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLongDirection string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLongDirectionE ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLongDirection = "E"
	ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLongDirectionW ZoneDNSRecordUpdateParamsMpBuiT95LocRecordDataLongDirection = "W"
)

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95LocRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTypeLoc ZoneDNSRecordUpdateParamsMpBuiT95LocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95LocRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95MxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95MxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95MxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95MxRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95MxRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTypeMx ZoneDNSRecordUpdateParamsMpBuiT95MxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95MxRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a NAPTR record.
type ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordData struct {
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

func (r ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTypeNaptr ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95NaptrRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95NsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95NsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95NsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95NsRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95NsRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTypeNs ZoneDNSRecordUpdateParamsMpBuiT95NsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95NsRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95PtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95PtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95PtrRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTypePtr ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95PtrRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SMIMEA record.
type ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTypeSmimea ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95SmimeaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95SrvRecord struct {
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95SrvRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SRV record.
type ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordData struct {
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

func (r ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTypeSrv ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95SrvRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SSHFP record.
type ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTypeSshfp ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95SshfpRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SVCB record.
type ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTypeSvcb ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95SvcbRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a TLSA record.
type ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTypeTlsa ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95TlsaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95TxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95TxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95TxtRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTypeTxt ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95TxtRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMpBuiT95UriRecord struct {
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordUpdateParamsMpBuiT95UriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMpBuiT95UriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95UriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMpBuiT95UriRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a URI record.
type ZoneDNSRecordUpdateParamsMpBuiT95UriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordUpdateParamsMpBuiT95UriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsMpBuiT95UriRecordType string

const (
	ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTypeUri ZoneDNSRecordUpdateParamsMpBuiT95UriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtl()
}

type ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtlNumber1 ZoneDNSRecordUpdateParamsMpBuiT95UriRecordTtlNumber = 1
)

// This interface is a union satisfied by one of the following:
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecord].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTypeA ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95ARecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTypeAaaa ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95AaaaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecord struct {
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a CAA record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTypeCaa ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CaaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecord struct {
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a CERT record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTypeCert ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CertRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTypeCname ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95CnameRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a DNSKEY record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTypeDnskey ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DnskeyRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecord struct {
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a DS record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTypeDs ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95DsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a HTTPS record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTypeHTTPs ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95HTTPsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecord struct {
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a LOC record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLongDirection] `json:"long_direction"`
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

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLatDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLatDirectionN ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLatDirection = "N"
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLatDirectionS ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLatDirection = "S"
)

// Longitude direction.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLongDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLongDirectionE ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLongDirection = "E"
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLongDirectionW ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordDataLongDirection = "W"
)

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTypeLoc ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95LocRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTypeMx ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95MxRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a NAPTR record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordData struct {
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

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTypeNaptr ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NaptrRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTypeNs ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95NsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTypePtr ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95PtrRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SMIMEA record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTypeSmimea ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SmimeaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecord struct {
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SRV record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordData struct {
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

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTypeSrv ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SrvRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SSHFP record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTypeSshfp ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SshfpRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SVCB record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTypeSvcb ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95SvcbRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a TLSA record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTypeTlsa ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TlsaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTypeTxt ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95TxtRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecord struct {
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a URI record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTypeUri ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMpBuiT95UriRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParams struct {
	Comment param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment] `query:"comment"`
	// DNS record content.
	Content param.Field[string] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch] `query:"match"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `query:"name"`
	// Field to order DNS records by.
	Order param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder] `query:"order"`
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
	Search param.Field[string]                                                 `query:"search"`
	Tag    param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag] `query:"tag"`
	// Whether to match all tag search requirements or at least one (any). If set to
	// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
	// logical OR instead. Note that the regular `match` parameter is still used to
	// combine the resulting condition with other filters that aren't related to tags.
	TagMatch param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch] `query:"tag_match"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType] `query:"type"`
}

// URLQuery serializes [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParams]'s
// query parameters as `url.Values`.
func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment struct {
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

// URLQuery serializes
// [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment]'s query parameters
// as `url.Values`.
func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order DNS records in.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirectionAsc  ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection = "asc"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirectionDesc ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatchAny ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch = "any"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatchAll ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch = "all"
)

// Field to order DNS records by.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderType    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "type"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderName    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "name"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderContent ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "content"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderTtl     ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "ttl"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderProxied ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "proxied"
)

type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag struct {
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

// URLQuery serializes [ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag]'s
// query parameters as `url.Values`.
func (r ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all tag search requirements or at least one (any). If set to
// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
// logical OR instead. Note that the regular `match` parameter is still used to
// combine the resulting condition with other filters that aren't related to tags.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatchAny ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch = "any"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatchAll ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch = "all"
)

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType string

const (
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeA      ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "A"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeAaaa   ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "AAAA"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeCaa    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "CAA"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeCert   ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "CERT"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeCname  ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "CNAME"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeDnskey ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "DNSKEY"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeDs     ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "DS"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeHTTPs  ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "HTTPS"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeLoc    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "LOC"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeMx     ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "MX"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeNaptr  ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "NAPTR"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeNs     ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "NS"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypePtr    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "PTR"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSmimea ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SMIMEA"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSrv    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SRV"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSshfp  ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SSHFP"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSvcb   ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SVCB"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeTlsa   ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "TLSA"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeTxt    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "TXT"
	ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeUri    ZoneDNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "URI"
)

// This interface is a union satisfied by one of the following:
// [ZoneDNSRecordPatchParamsMpBuiT95ARecord],
// [ZoneDNSRecordPatchParamsMpBuiT95AaaaRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95CaaRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95CertRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95CnameRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95DsRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95LocRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95MxRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95NaptrRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95NsRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95PtrRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95SrvRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95SshfpRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95SvcbRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95TlsaRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95TxtRecord],
// [ZoneDNSRecordPatchParamsMpBuiT95UriRecord].
type ZoneDNSRecordPatchParams interface {
	ImplementsZoneDNSRecordPatchParams()
}

type ZoneDNSRecordPatchParamsMpBuiT95ARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95ARecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95ARecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95ARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95ARecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95ARecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95ARecordTypeA ZoneDNSRecordPatchParamsMpBuiT95ARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95ARecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95ARecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95ARecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95ARecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95ARecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95ARecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95AaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95AaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95AaaaRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTypeAaaa ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95AaaaRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95CaaRecord struct {
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95CaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95CaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95CaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95CaaRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a CAA record.
type ZoneDNSRecordPatchParamsMpBuiT95CaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95CaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95CaaRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTypeCaa ZoneDNSRecordPatchParamsMpBuiT95CaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95CaaRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95CertRecord struct {
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95CertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95CertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95CertRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95CertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95CertRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a CERT record.
type ZoneDNSRecordPatchParamsMpBuiT95CertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95CertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95CertRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95CertRecordTypeCert ZoneDNSRecordPatchParamsMpBuiT95CertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95CertRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95CertRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95CertRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95CertRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95CertRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95CertRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95CnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95CnameRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95CnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95CnameRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95CnameRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTypeCname ZoneDNSRecordPatchParamsMpBuiT95CnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95CnameRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a DNSKEY record.
type ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTypeDnskey ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95DnskeyRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95DsRecord struct {
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95DsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95DsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95DsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95DsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95DsRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a DS record.
type ZoneDNSRecordPatchParamsMpBuiT95DsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95DsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95DsRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95DsRecordTypeDs ZoneDNSRecordPatchParamsMpBuiT95DsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95DsRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95DsRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95DsRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95DsRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95DsRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95DsRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a HTTPS record.
type ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTypeHTTPs ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95HTTPsRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95LocRecord struct {
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95LocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95LocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95LocRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95LocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95LocRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a LOC record.
type ZoneDNSRecordPatchParamsMpBuiT95LocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLongDirection] `json:"long_direction"`
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

func (r ZoneDNSRecordPatchParamsMpBuiT95LocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLatDirection string

const (
	ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLatDirectionN ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLatDirection = "N"
	ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLatDirectionS ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLatDirection = "S"
)

// Longitude direction.
type ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLongDirection string

const (
	ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLongDirectionE ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLongDirection = "E"
	ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLongDirectionW ZoneDNSRecordPatchParamsMpBuiT95LocRecordDataLongDirection = "W"
)

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95LocRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95LocRecordTypeLoc ZoneDNSRecordPatchParamsMpBuiT95LocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95LocRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95LocRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95LocRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95LocRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95LocRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95LocRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95MxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95MxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95MxRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95MxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95MxRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95MxRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95MxRecordTypeMx ZoneDNSRecordPatchParamsMpBuiT95MxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95MxRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95MxRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95MxRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95MxRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95MxRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95MxRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95NaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95NaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95NaptrRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a NAPTR record.
type ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordData struct {
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

func (r ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTypeNaptr ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95NaptrRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95NsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95NsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95NsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95NsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95NsRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95NsRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95NsRecordTypeNs ZoneDNSRecordPatchParamsMpBuiT95NsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95NsRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95NsRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95NsRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95NsRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95NsRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95NsRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95PtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95PtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95PtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95PtrRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95PtrRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTypePtr ZoneDNSRecordPatchParamsMpBuiT95PtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95PtrRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a SMIMEA record.
type ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTypeSmimea ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95SmimeaRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95SrvRecord struct {
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95SrvRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95SrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95SrvRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a SRV record.
type ZoneDNSRecordPatchParamsMpBuiT95SrvRecordData struct {
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

func (r ZoneDNSRecordPatchParamsMpBuiT95SrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95SrvRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTypeSrv ZoneDNSRecordPatchParamsMpBuiT95SrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95SrvRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95SshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95SshfpRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a SSHFP record.
type ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTypeSshfp ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95SshfpRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95SvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95SvcbRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a SVCB record.
type ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTypeSvcb ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95SvcbRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95TlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95TlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95TlsaRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a TLSA record.
type ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTypeTlsa ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95TlsaRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95TxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95TxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95TxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95TxtRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95TxtRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTypeTxt ZoneDNSRecordPatchParamsMpBuiT95TxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95TxtRecordTtlNumber = 1
)

type ZoneDNSRecordPatchParamsMpBuiT95UriRecord struct {
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordPatchParamsMpBuiT95UriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordPatchParamsMpBuiT95UriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordPatchParamsMpBuiT95UriRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95UriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordPatchParamsMpBuiT95UriRecord) ImplementsZoneDNSRecordPatchParams() {

}

// Components of a URI record.
type ZoneDNSRecordPatchParamsMpBuiT95UriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordPatchParamsMpBuiT95UriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordPatchParamsMpBuiT95UriRecordType string

const (
	ZoneDNSRecordPatchParamsMpBuiT95UriRecordTypeUri ZoneDNSRecordPatchParamsMpBuiT95UriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordPatchParamsMpBuiT95UriRecordTtlNumber].
type ZoneDNSRecordPatchParamsMpBuiT95UriRecordTtl interface {
	ImplementsZoneDNSRecordPatchParamsMpBuiT95UriRecordTtl()
}

type ZoneDNSRecordPatchParamsMpBuiT95UriRecordTtlNumber float64

const (
	ZoneDNSRecordPatchParamsMpBuiT95UriRecordTtlNumber1 ZoneDNSRecordPatchParamsMpBuiT95UriRecordTtlNumber = 1
)
