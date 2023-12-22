// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *ZoneDNSRecordService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *DNSResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update DNS Record
func (r *ZoneDNSRecordService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneDNSRecordUpdateParams, opts ...option.RequestOption) (res *DNSResponseSingle, err error) {
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

// Create a new DNS record for a zone. See the record object definitions for
// required attributes for each record type.
func (r *ZoneDNSRecordService) DNSRecordsForAZoneNewDNSRecord(ctx context.Context, zoneIdentifier string, body ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams, opts ...option.RequestOption) (res *DNSResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *ZoneDNSRecordService) DNSRecordsForAZoneListDNSRecords(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *DNSResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DNSResponseCollection struct {
	Errors     []DNSResponseCollectionError    `json:"errors"`
	Messages   []DNSResponseCollectionMessage  `json:"messages"`
	Result     []DNSResponseCollectionResult   `json:"result"`
	ResultInfo DNSResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DNSResponseCollectionSuccess `json:"success"`
	JSON    dnsResponseCollectionJSON    `json:"-"`
}

// dnsResponseCollectionJSON contains the JSON metadata for the struct
// [DNSResponseCollection]
type dnsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseCollectionError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dnsResponseCollectionErrorJSON `json:"-"`
}

// dnsResponseCollectionErrorJSON contains the JSON metadata for the struct
// [DNSResponseCollectionError]
type dnsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseCollectionMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dnsResponseCollectionMessageJSON `json:"-"`
}

// dnsResponseCollectionMessageJSON contains the JSON metadata for the struct
// [DNSResponseCollectionMessage]
type dnsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DNSResponseCollectionResultARecord],
// [DNSResponseCollectionResultAaaaRecord], [DNSResponseCollectionResultCaaRecord],
// [DNSResponseCollectionResultCertRecord],
// [DNSResponseCollectionResultCnameRecord],
// [DNSResponseCollectionResultDnskeyRecord],
// [DNSResponseCollectionResultDsRecord], [DNSResponseCollectionResultHTTPsRecord],
// [DNSResponseCollectionResultLocRecord], [DNSResponseCollectionResultMxRecord],
// [DNSResponseCollectionResultNaptrRecord], [DNSResponseCollectionResultNsRecord],
// [DNSResponseCollectionResultPtrRecord],
// [DNSResponseCollectionResultSmimeaRecord],
// [DNSResponseCollectionResultSrvRecord],
// [DNSResponseCollectionResultSshfpRecord],
// [DNSResponseCollectionResultSvcbRecord],
// [DNSResponseCollectionResultTlsaRecord], [DNSResponseCollectionResultTxtRecord]
// or [DNSResponseCollectionResultUriRecord].
type DNSResponseCollectionResult interface {
	implementsDNSResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSResponseCollectionResult)(nil)).Elem(), "")
}

type DNSResponseCollectionResultARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultARecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultARecordMeta `json:"meta"`
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
	Ttl DNSResponseCollectionResultARecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultARecordJSON `json:"-"`
}

// dnsResponseCollectionResultARecordJSON contains the JSON metadata for the struct
// [DNSResponseCollectionResultARecord]
type dnsResponseCollectionResultARecordJSON struct {
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

func (r *DNSResponseCollectionResultARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultARecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultARecordType string

const (
	DNSResponseCollectionResultARecordTypeA DNSResponseCollectionResultARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsResponseCollectionResultARecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultARecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultARecordMeta]
type dnsResponseCollectionResultARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultARecordTtlNumber].
type DNSResponseCollectionResultARecordTtl interface {
	ImplementsDNSResponseCollectionResultARecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultARecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultARecordTtlNumber float64

const (
	DNSResponseCollectionResultARecordTtlNumber1 DNSResponseCollectionResultARecordTtlNumber = 1
)

type DNSResponseCollectionResultAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultAaaaRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultAaaaRecordMeta `json:"meta"`
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
	Ttl DNSResponseCollectionResultAaaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultAaaaRecordJSON `json:"-"`
}

// dnsResponseCollectionResultAaaaRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultAaaaRecord]
type dnsResponseCollectionResultAaaaRecordJSON struct {
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

func (r *DNSResponseCollectionResultAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultAaaaRecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultAaaaRecordType string

const (
	DNSResponseCollectionResultAaaaRecordTypeAaaa DNSResponseCollectionResultAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsResponseCollectionResultAaaaRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultAaaaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultAaaaRecordMeta]
type dnsResponseCollectionResultAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultAaaaRecordTtlNumber].
type DNSResponseCollectionResultAaaaRecordTtl interface {
	ImplementsDNSResponseCollectionResultAaaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultAaaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultAaaaRecordTtlNumber float64

const (
	DNSResponseCollectionResultAaaaRecordTtlNumber1 DNSResponseCollectionResultAaaaRecordTtlNumber = 1
)

type DNSResponseCollectionResultCaaRecord struct {
	// Components of a CAA record.
	Data DNSResponseCollectionResultCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultCaaRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultCaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultCaaRecordJSON `json:"-"`
}

// dnsResponseCollectionResultCaaRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCaaRecord]
type dnsResponseCollectionResultCaaRecordJSON struct {
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

func (r *DNSResponseCollectionResultCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultCaaRecord) implementsDNSResponseCollectionResult() {}

// Components of a CAA record.
type DNSResponseCollectionResultCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                       `json:"value"`
	JSON  dnsResponseCollectionResultCaaRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultCaaRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCaaRecordData]
type dnsResponseCollectionResultCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultCaaRecordType string

const (
	DNSResponseCollectionResultCaaRecordTypeCaa DNSResponseCollectionResultCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsResponseCollectionResultCaaRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultCaaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCaaRecordMeta]
type dnsResponseCollectionResultCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultCaaRecordTtlNumber].
type DNSResponseCollectionResultCaaRecordTtl interface {
	ImplementsDNSResponseCollectionResultCaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultCaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultCaaRecordTtlNumber float64

const (
	DNSResponseCollectionResultCaaRecordTtlNumber1 DNSResponseCollectionResultCaaRecordTtlNumber = 1
)

type DNSResponseCollectionResultCertRecord struct {
	// Components of a CERT record.
	Data DNSResponseCollectionResultCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultCertRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultCertRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultCertRecordJSON `json:"-"`
}

// dnsResponseCollectionResultCertRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCertRecord]
type dnsResponseCollectionResultCertRecordJSON struct {
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

func (r *DNSResponseCollectionResultCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultCertRecord) implementsDNSResponseCollectionResult() {}

// Components of a CERT record.
type DNSResponseCollectionResultCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                       `json:"type"`
	JSON dnsResponseCollectionResultCertRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultCertRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCertRecordData]
type dnsResponseCollectionResultCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultCertRecordType string

const (
	DNSResponseCollectionResultCertRecordTypeCert DNSResponseCollectionResultCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsResponseCollectionResultCertRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultCertRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCertRecordMeta]
type dnsResponseCollectionResultCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultCertRecordTtlNumber].
type DNSResponseCollectionResultCertRecordTtl interface {
	ImplementsDNSResponseCollectionResultCertRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultCertRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultCertRecordTtlNumber float64

const (
	DNSResponseCollectionResultCertRecordTtlNumber1 DNSResponseCollectionResultCertRecordTtlNumber = 1
)

type DNSResponseCollectionResultCnameRecord struct {
	// A valid hostname.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultCnameRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultCnameRecordMeta `json:"meta"`
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
	Ttl DNSResponseCollectionResultCnameRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultCnameRecordJSON `json:"-"`
}

// dnsResponseCollectionResultCnameRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultCnameRecord]
type dnsResponseCollectionResultCnameRecordJSON struct {
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

func (r *DNSResponseCollectionResultCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultCnameRecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultCnameRecordType string

const (
	DNSResponseCollectionResultCnameRecordTypeCname DNSResponseCollectionResultCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsResponseCollectionResultCnameRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultCnameRecordMetaJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultCnameRecordMeta]
type dnsResponseCollectionResultCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultCnameRecordTtlNumber].
type DNSResponseCollectionResultCnameRecordTtl interface {
	ImplementsDNSResponseCollectionResultCnameRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultCnameRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultCnameRecordTtlNumber float64

const (
	DNSResponseCollectionResultCnameRecordTtlNumber1 DNSResponseCollectionResultCnameRecordTtlNumber = 1
)

type DNSResponseCollectionResultDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSResponseCollectionResultDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultDnskeyRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultDnskeyRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultDnskeyRecordJSON `json:"-"`
}

// dnsResponseCollectionResultDnskeyRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultDnskeyRecord]
type dnsResponseCollectionResultDnskeyRecordJSON struct {
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

func (r *DNSResponseCollectionResultDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultDnskeyRecord) implementsDNSResponseCollectionResult() {}

// Components of a DNSKEY record.
type DNSResponseCollectionResultDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                          `json:"public_key"`
	JSON      dnsResponseCollectionResultDnskeyRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultDnskeyRecordDataJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultDnskeyRecordData]
type dnsResponseCollectionResultDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultDnskeyRecordType string

const (
	DNSResponseCollectionResultDnskeyRecordTypeDnskey DNSResponseCollectionResultDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsResponseCollectionResultDnskeyRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultDnskeyRecordMetaJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultDnskeyRecordMeta]
type dnsResponseCollectionResultDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultDnskeyRecordTtlNumber].
type DNSResponseCollectionResultDnskeyRecordTtl interface {
	ImplementsDNSResponseCollectionResultDnskeyRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultDnskeyRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultDnskeyRecordTtlNumber float64

const (
	DNSResponseCollectionResultDnskeyRecordTtlNumber1 DNSResponseCollectionResultDnskeyRecordTtlNumber = 1
)

type DNSResponseCollectionResultDsRecord struct {
	// Components of a DS record.
	Data DNSResponseCollectionResultDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultDsRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultDsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultDsRecordJSON `json:"-"`
}

// dnsResponseCollectionResultDsRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultDsRecord]
type dnsResponseCollectionResultDsRecordJSON struct {
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

func (r *DNSResponseCollectionResultDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultDsRecord) implementsDNSResponseCollectionResult() {}

// Components of a DS record.
type DNSResponseCollectionResultDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                     `json:"key_tag"`
	JSON   dnsResponseCollectionResultDsRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultDsRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultDsRecordData]
type dnsResponseCollectionResultDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultDsRecordType string

const (
	DNSResponseCollectionResultDsRecordTypeDs DNSResponseCollectionResultDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsResponseCollectionResultDsRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultDsRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultDsRecordMeta]
type dnsResponseCollectionResultDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultDsRecordTtlNumber].
type DNSResponseCollectionResultDsRecordTtl interface {
	ImplementsDNSResponseCollectionResultDsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultDsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultDsRecordTtlNumber float64

const (
	DNSResponseCollectionResultDsRecordTtlNumber1 DNSResponseCollectionResultDsRecordTtlNumber = 1
)

type DNSResponseCollectionResultHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSResponseCollectionResultHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultHTTPsRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultHTTPsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultHTTPsRecordJSON `json:"-"`
}

// dnsResponseCollectionResultHTTPsRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultHTTPsRecord]
type dnsResponseCollectionResultHTTPsRecordJSON struct {
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

func (r *DNSResponseCollectionResultHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultHTTPsRecord) implementsDNSResponseCollectionResult() {}

// Components of a HTTPS record.
type DNSResponseCollectionResultHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                         `json:"value"`
	JSON  dnsResponseCollectionResultHTTPsRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultHTTPsRecordDataJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultHTTPsRecordData]
type dnsResponseCollectionResultHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultHTTPsRecordType string

const (
	DNSResponseCollectionResultHTTPsRecordTypeHTTPs DNSResponseCollectionResultHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsResponseCollectionResultHTTPsRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultHTTPsRecordMetaJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultHTTPsRecordMeta]
type dnsResponseCollectionResultHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultHTTPsRecordTtlNumber].
type DNSResponseCollectionResultHTTPsRecordTtl interface {
	ImplementsDNSResponseCollectionResultHTTPsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultHTTPsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultHTTPsRecordTtlNumber float64

const (
	DNSResponseCollectionResultHTTPsRecordTtlNumber1 DNSResponseCollectionResultHTTPsRecordTtlNumber = 1
)

type DNSResponseCollectionResultLocRecord struct {
	// Components of a LOC record.
	Data DNSResponseCollectionResultLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultLocRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultLocRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultLocRecordJSON `json:"-"`
}

// dnsResponseCollectionResultLocRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultLocRecord]
type dnsResponseCollectionResultLocRecordJSON struct {
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

func (r *DNSResponseCollectionResultLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultLocRecord) implementsDNSResponseCollectionResult() {}

// Components of a LOC record.
type DNSResponseCollectionResultLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSResponseCollectionResultLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSResponseCollectionResultLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                      `json:"size"`
	JSON dnsResponseCollectionResultLocRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultLocRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultLocRecordData]
type dnsResponseCollectionResultLocRecordDataJSON struct {
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

func (r *DNSResponseCollectionResultLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSResponseCollectionResultLocRecordDataLatDirection string

const (
	DNSResponseCollectionResultLocRecordDataLatDirectionN DNSResponseCollectionResultLocRecordDataLatDirection = "N"
	DNSResponseCollectionResultLocRecordDataLatDirectionS DNSResponseCollectionResultLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSResponseCollectionResultLocRecordDataLongDirection string

const (
	DNSResponseCollectionResultLocRecordDataLongDirectionE DNSResponseCollectionResultLocRecordDataLongDirection = "E"
	DNSResponseCollectionResultLocRecordDataLongDirectionW DNSResponseCollectionResultLocRecordDataLongDirection = "W"
)

// Record type.
type DNSResponseCollectionResultLocRecordType string

const (
	DNSResponseCollectionResultLocRecordTypeLoc DNSResponseCollectionResultLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsResponseCollectionResultLocRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultLocRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultLocRecordMeta]
type dnsResponseCollectionResultLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultLocRecordTtlNumber].
type DNSResponseCollectionResultLocRecordTtl interface {
	ImplementsDNSResponseCollectionResultLocRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultLocRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultLocRecordTtlNumber float64

const (
	DNSResponseCollectionResultLocRecordTtlNumber1 DNSResponseCollectionResultLocRecordTtlNumber = 1
)

type DNSResponseCollectionResultMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSResponseCollectionResultMxRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultMxRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultMxRecordJSON `json:"-"`
}

// dnsResponseCollectionResultMxRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultMxRecord]
type dnsResponseCollectionResultMxRecordJSON struct {
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

func (r *DNSResponseCollectionResultMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultMxRecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultMxRecordType string

const (
	DNSResponseCollectionResultMxRecordTypeMx DNSResponseCollectionResultMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsResponseCollectionResultMxRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultMxRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultMxRecordMeta]
type dnsResponseCollectionResultMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultMxRecordTtlNumber].
type DNSResponseCollectionResultMxRecordTtl interface {
	ImplementsDNSResponseCollectionResultMxRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultMxRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultMxRecordTtlNumber float64

const (
	DNSResponseCollectionResultMxRecordTtlNumber1 DNSResponseCollectionResultMxRecordTtlNumber = 1
)

type DNSResponseCollectionResultNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSResponseCollectionResultNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultNaptrRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultNaptrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultNaptrRecordJSON `json:"-"`
}

// dnsResponseCollectionResultNaptrRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultNaptrRecord]
type dnsResponseCollectionResultNaptrRecordJSON struct {
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

func (r *DNSResponseCollectionResultNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultNaptrRecord) implementsDNSResponseCollectionResult() {}

// Components of a NAPTR record.
type DNSResponseCollectionResultNaptrRecordData struct {
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
	Service string                                         `json:"service"`
	JSON    dnsResponseCollectionResultNaptrRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultNaptrRecordDataJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultNaptrRecordData]
type dnsResponseCollectionResultNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultNaptrRecordType string

const (
	DNSResponseCollectionResultNaptrRecordTypeNaptr DNSResponseCollectionResultNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsResponseCollectionResultNaptrRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultNaptrRecordMetaJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultNaptrRecordMeta]
type dnsResponseCollectionResultNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultNaptrRecordTtlNumber].
type DNSResponseCollectionResultNaptrRecordTtl interface {
	ImplementsDNSResponseCollectionResultNaptrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultNaptrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultNaptrRecordTtlNumber float64

const (
	DNSResponseCollectionResultNaptrRecordTtlNumber1 DNSResponseCollectionResultNaptrRecordTtlNumber = 1
)

type DNSResponseCollectionResultNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultNsRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultNsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultNsRecordJSON `json:"-"`
}

// dnsResponseCollectionResultNsRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultNsRecord]
type dnsResponseCollectionResultNsRecordJSON struct {
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

func (r *DNSResponseCollectionResultNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultNsRecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultNsRecordType string

const (
	DNSResponseCollectionResultNsRecordTypeNs DNSResponseCollectionResultNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsResponseCollectionResultNsRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultNsRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultNsRecordMeta]
type dnsResponseCollectionResultNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultNsRecordTtlNumber].
type DNSResponseCollectionResultNsRecordTtl interface {
	ImplementsDNSResponseCollectionResultNsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultNsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultNsRecordTtlNumber float64

const (
	DNSResponseCollectionResultNsRecordTtlNumber1 DNSResponseCollectionResultNsRecordTtlNumber = 1
)

type DNSResponseCollectionResultPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultPtrRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultPtrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultPtrRecordJSON `json:"-"`
}

// dnsResponseCollectionResultPtrRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultPtrRecord]
type dnsResponseCollectionResultPtrRecordJSON struct {
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

func (r *DNSResponseCollectionResultPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultPtrRecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultPtrRecordType string

const (
	DNSResponseCollectionResultPtrRecordTypePtr DNSResponseCollectionResultPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsResponseCollectionResultPtrRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultPtrRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultPtrRecordMeta]
type dnsResponseCollectionResultPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultPtrRecordTtlNumber].
type DNSResponseCollectionResultPtrRecordTtl interface {
	ImplementsDNSResponseCollectionResultPtrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultPtrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultPtrRecordTtlNumber float64

const (
	DNSResponseCollectionResultPtrRecordTtlNumber1 DNSResponseCollectionResultPtrRecordTtlNumber = 1
)

type DNSResponseCollectionResultSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSResponseCollectionResultSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultSmimeaRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultSmimeaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultSmimeaRecordJSON `json:"-"`
}

// dnsResponseCollectionResultSmimeaRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSmimeaRecord]
type dnsResponseCollectionResultSmimeaRecordJSON struct {
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

func (r *DNSResponseCollectionResultSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultSmimeaRecord) implementsDNSResponseCollectionResult() {}

// Components of a SMIMEA record.
type DNSResponseCollectionResultSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                         `json:"usage"`
	JSON  dnsResponseCollectionResultSmimeaRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultSmimeaRecordDataJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultSmimeaRecordData]
type dnsResponseCollectionResultSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultSmimeaRecordType string

const (
	DNSResponseCollectionResultSmimeaRecordTypeSmimea DNSResponseCollectionResultSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsResponseCollectionResultSmimeaRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultSmimeaRecordMetaJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultSmimeaRecordMeta]
type dnsResponseCollectionResultSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultSmimeaRecordTtlNumber].
type DNSResponseCollectionResultSmimeaRecordTtl interface {
	ImplementsDNSResponseCollectionResultSmimeaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultSmimeaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultSmimeaRecordTtlNumber float64

const (
	DNSResponseCollectionResultSmimeaRecordTtlNumber1 DNSResponseCollectionResultSmimeaRecordTtlNumber = 1
)

type DNSResponseCollectionResultSrvRecord struct {
	// Components of a SRV record.
	Data DNSResponseCollectionResultSrvRecordData `json:"data,required"`
	// Record type.
	Type DNSResponseCollectionResultSrvRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Service, protocol, and SRV name content. See 'data' for setting the individual
	// component values.
	Name string `json:"name"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultSrvRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultSrvRecordJSON `json:"-"`
}

// dnsResponseCollectionResultSrvRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSrvRecord]
type dnsResponseCollectionResultSrvRecordJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultSrvRecord) implementsDNSResponseCollectionResult() {}

// Components of a SRV record.
type DNSResponseCollectionResultSrvRecordData struct {
	// A valid hostname.
	Name string `json:"name" format:"hostname"`
	// The port of the service.
	Port float64 `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// A valid protocol.
	Proto string `json:"proto"`
	// A service type, prefixed with an underscore.
	Service string `json:"service"`
	// A valid hostname.
	Target string `json:"target" format:"hostname"`
	// The record weight.
	Weight float64                                      `json:"weight"`
	JSON   dnsResponseCollectionResultSrvRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultSrvRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSrvRecordData]
type dnsResponseCollectionResultSrvRecordDataJSON struct {
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

func (r *DNSResponseCollectionResultSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultSrvRecordType string

const (
	DNSResponseCollectionResultSrvRecordTypeSrv DNSResponseCollectionResultSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsResponseCollectionResultSrvRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultSrvRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSrvRecordMeta]
type dnsResponseCollectionResultSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultSrvRecordTtlNumber].
type DNSResponseCollectionResultSrvRecordTtl interface {
	ImplementsDNSResponseCollectionResultSrvRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultSrvRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultSrvRecordTtlNumber float64

const (
	DNSResponseCollectionResultSrvRecordTtlNumber1 DNSResponseCollectionResultSrvRecordTtlNumber = 1
)

type DNSResponseCollectionResultSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSResponseCollectionResultSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultSshfpRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultSshfpRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultSshfpRecordJSON `json:"-"`
}

// dnsResponseCollectionResultSshfpRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSshfpRecord]
type dnsResponseCollectionResultSshfpRecordJSON struct {
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

func (r *DNSResponseCollectionResultSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultSshfpRecord) implementsDNSResponseCollectionResult() {}

// Components of a SSHFP record.
type DNSResponseCollectionResultSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                        `json:"type"`
	JSON dnsResponseCollectionResultSshfpRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultSshfpRecordDataJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultSshfpRecordData]
type dnsResponseCollectionResultSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultSshfpRecordType string

const (
	DNSResponseCollectionResultSshfpRecordTypeSshfp DNSResponseCollectionResultSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsResponseCollectionResultSshfpRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultSshfpRecordMetaJSON contains the JSON metadata for
// the struct [DNSResponseCollectionResultSshfpRecordMeta]
type dnsResponseCollectionResultSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultSshfpRecordTtlNumber].
type DNSResponseCollectionResultSshfpRecordTtl interface {
	ImplementsDNSResponseCollectionResultSshfpRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultSshfpRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultSshfpRecordTtlNumber float64

const (
	DNSResponseCollectionResultSshfpRecordTtlNumber1 DNSResponseCollectionResultSshfpRecordTtlNumber = 1
)

type DNSResponseCollectionResultSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSResponseCollectionResultSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultSvcbRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultSvcbRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultSvcbRecordJSON `json:"-"`
}

// dnsResponseCollectionResultSvcbRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSvcbRecord]
type dnsResponseCollectionResultSvcbRecordJSON struct {
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

func (r *DNSResponseCollectionResultSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultSvcbRecord) implementsDNSResponseCollectionResult() {}

// Components of a SVCB record.
type DNSResponseCollectionResultSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                        `json:"value"`
	JSON  dnsResponseCollectionResultSvcbRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultSvcbRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSvcbRecordData]
type dnsResponseCollectionResultSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultSvcbRecordType string

const (
	DNSResponseCollectionResultSvcbRecordTypeSvcb DNSResponseCollectionResultSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsResponseCollectionResultSvcbRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultSvcbRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultSvcbRecordMeta]
type dnsResponseCollectionResultSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultSvcbRecordTtlNumber].
type DNSResponseCollectionResultSvcbRecordTtl interface {
	ImplementsDNSResponseCollectionResultSvcbRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultSvcbRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultSvcbRecordTtlNumber float64

const (
	DNSResponseCollectionResultSvcbRecordTtlNumber1 DNSResponseCollectionResultSvcbRecordTtlNumber = 1
)

type DNSResponseCollectionResultTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSResponseCollectionResultTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultTlsaRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultTlsaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                    `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultTlsaRecordJSON `json:"-"`
}

// dnsResponseCollectionResultTlsaRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultTlsaRecord]
type dnsResponseCollectionResultTlsaRecordJSON struct {
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

func (r *DNSResponseCollectionResultTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultTlsaRecord) implementsDNSResponseCollectionResult() {}

// Components of a TLSA record.
type DNSResponseCollectionResultTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                       `json:"usage"`
	JSON  dnsResponseCollectionResultTlsaRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultTlsaRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultTlsaRecordData]
type dnsResponseCollectionResultTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSResponseCollectionResultTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultTlsaRecordType string

const (
	DNSResponseCollectionResultTlsaRecordTypeTlsa DNSResponseCollectionResultTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                        `json:"source"`
	JSON   dnsResponseCollectionResultTlsaRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultTlsaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultTlsaRecordMeta]
type dnsResponseCollectionResultTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultTlsaRecordTtlNumber].
type DNSResponseCollectionResultTlsaRecordTtl interface {
	ImplementsDNSResponseCollectionResultTlsaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultTlsaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultTlsaRecordTtlNumber float64

const (
	DNSResponseCollectionResultTlsaRecordTtlNumber1 DNSResponseCollectionResultTlsaRecordTtlNumber = 1
)

type DNSResponseCollectionResultTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseCollectionResultTxtRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultTxtRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultTxtRecordJSON `json:"-"`
}

// dnsResponseCollectionResultTxtRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultTxtRecord]
type dnsResponseCollectionResultTxtRecordJSON struct {
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

func (r *DNSResponseCollectionResultTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultTxtRecord) implementsDNSResponseCollectionResult() {}

// Record type.
type DNSResponseCollectionResultTxtRecordType string

const (
	DNSResponseCollectionResultTxtRecordTypeTxt DNSResponseCollectionResultTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsResponseCollectionResultTxtRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultTxtRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultTxtRecordMeta]
type dnsResponseCollectionResultTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultTxtRecordTtlNumber].
type DNSResponseCollectionResultTxtRecordTtl interface {
	ImplementsDNSResponseCollectionResultTxtRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultTxtRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultTxtRecordTtlNumber float64

const (
	DNSResponseCollectionResultTxtRecordTtlNumber1 DNSResponseCollectionResultTxtRecordTtlNumber = 1
)

type DNSResponseCollectionResultUriRecord struct {
	// Components of a URI record.
	Data DNSResponseCollectionResultUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSResponseCollectionResultUriRecordType `json:"type,required"`
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
	Meta DNSResponseCollectionResultUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseCollectionResultUriRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                   `json:"zone_name" format:"hostname"`
	JSON     dnsResponseCollectionResultUriRecordJSON `json:"-"`
}

// dnsResponseCollectionResultUriRecordJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultUriRecord]
type dnsResponseCollectionResultUriRecordJSON struct {
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

func (r *DNSResponseCollectionResultUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseCollectionResultUriRecord) implementsDNSResponseCollectionResult() {}

// Components of a URI record.
type DNSResponseCollectionResultUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                      `json:"weight"`
	JSON   dnsResponseCollectionResultUriRecordDataJSON `json:"-"`
}

// dnsResponseCollectionResultUriRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultUriRecordData]
type dnsResponseCollectionResultUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseCollectionResultUriRecordType string

const (
	DNSResponseCollectionResultUriRecordTypeUri DNSResponseCollectionResultUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseCollectionResultUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                       `json:"source"`
	JSON   dnsResponseCollectionResultUriRecordMetaJSON `json:"-"`
}

// dnsResponseCollectionResultUriRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseCollectionResultUriRecordMeta]
type dnsResponseCollectionResultUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseCollectionResultUriRecordTtlNumber].
type DNSResponseCollectionResultUriRecordTtl interface {
	ImplementsDNSResponseCollectionResultUriRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseCollectionResultUriRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseCollectionResultUriRecordTtlNumber float64

const (
	DNSResponseCollectionResultUriRecordTtlNumber1 DNSResponseCollectionResultUriRecordTtlNumber = 1
)

type DNSResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                             `json:"total_count"`
	JSON       dnsResponseCollectionResultInfoJSON `json:"-"`
}

// dnsResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [DNSResponseCollectionResultInfo]
type dnsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSResponseCollectionSuccess bool

const (
	DNSResponseCollectionSuccessTrue DNSResponseCollectionSuccess = true
)

type DNSResponseSingle struct {
	Errors   []DNSResponseSingleError   `json:"errors"`
	Messages []DNSResponseSingleMessage `json:"messages"`
	Result   DNSResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success DNSResponseSingleSuccess `json:"success"`
	JSON    dnsResponseSingleJSON    `json:"-"`
}

// dnsResponseSingleJSON contains the JSON metadata for the struct
// [DNSResponseSingle]
type dnsResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseSingleError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    dnsResponseSingleErrorJSON `json:"-"`
}

// dnsResponseSingleErrorJSON contains the JSON metadata for the struct
// [DNSResponseSingleError]
type dnsResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseSingleMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    dnsResponseSingleMessageJSON `json:"-"`
}

// dnsResponseSingleMessageJSON contains the JSON metadata for the struct
// [DNSResponseSingleMessage]
type dnsResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DNSResponseSingleResultARecord],
// [DNSResponseSingleResultAaaaRecord], [DNSResponseSingleResultCaaRecord],
// [DNSResponseSingleResultCertRecord], [DNSResponseSingleResultCnameRecord],
// [DNSResponseSingleResultDnskeyRecord], [DNSResponseSingleResultDsRecord],
// [DNSResponseSingleResultHTTPsRecord], [DNSResponseSingleResultLocRecord],
// [DNSResponseSingleResultMxRecord], [DNSResponseSingleResultNaptrRecord],
// [DNSResponseSingleResultNsRecord], [DNSResponseSingleResultPtrRecord],
// [DNSResponseSingleResultSmimeaRecord], [DNSResponseSingleResultSrvRecord],
// [DNSResponseSingleResultSshfpRecord], [DNSResponseSingleResultSvcbRecord],
// [DNSResponseSingleResultTlsaRecord], [DNSResponseSingleResultTxtRecord] or
// [DNSResponseSingleResultUriRecord].
type DNSResponseSingleResult interface {
	implementsDNSResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSResponseSingleResult)(nil)).Elem(), "")
}

type DNSResponseSingleResultARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultARecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultARecordMeta `json:"meta"`
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
	Ttl DNSResponseSingleResultARecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                             `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultARecordJSON `json:"-"`
}

// dnsResponseSingleResultARecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultARecord]
type dnsResponseSingleResultARecordJSON struct {
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

func (r *DNSResponseSingleResultARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultARecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultARecordType string

const (
	DNSResponseSingleResultARecordTypeA DNSResponseSingleResultARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                 `json:"source"`
	JSON   dnsResponseSingleResultARecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultARecordMetaJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultARecordMeta]
type dnsResponseSingleResultARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultARecordTtlNumber].
type DNSResponseSingleResultARecordTtl interface {
	ImplementsDNSResponseSingleResultARecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultARecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultARecordTtlNumber float64

const (
	DNSResponseSingleResultARecordTtlNumber1 DNSResponseSingleResultARecordTtlNumber = 1
)

type DNSResponseSingleResultAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultAaaaRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultAaaaRecordMeta `json:"meta"`
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
	Ttl DNSResponseSingleResultAaaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultAaaaRecordJSON `json:"-"`
}

// dnsResponseSingleResultAaaaRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultAaaaRecord]
type dnsResponseSingleResultAaaaRecordJSON struct {
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

func (r *DNSResponseSingleResultAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultAaaaRecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultAaaaRecordType string

const (
	DNSResponseSingleResultAaaaRecordTypeAaaa DNSResponseSingleResultAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsResponseSingleResultAaaaRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultAaaaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultAaaaRecordMeta]
type dnsResponseSingleResultAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultAaaaRecordTtlNumber].
type DNSResponseSingleResultAaaaRecordTtl interface {
	ImplementsDNSResponseSingleResultAaaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultAaaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultAaaaRecordTtlNumber float64

const (
	DNSResponseSingleResultAaaaRecordTtlNumber1 DNSResponseSingleResultAaaaRecordTtlNumber = 1
)

type DNSResponseSingleResultCaaRecord struct {
	// Components of a CAA record.
	Data DNSResponseSingleResultCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultCaaRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultCaaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultCaaRecordJSON `json:"-"`
}

// dnsResponseSingleResultCaaRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultCaaRecord]
type dnsResponseSingleResultCaaRecordJSON struct {
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

func (r *DNSResponseSingleResultCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultCaaRecord) implementsDNSResponseSingleResult() {}

// Components of a CAA record.
type DNSResponseSingleResultCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                   `json:"value"`
	JSON  dnsResponseSingleResultCaaRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultCaaRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultCaaRecordData]
type dnsResponseSingleResultCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultCaaRecordType string

const (
	DNSResponseSingleResultCaaRecordTypeCaa DNSResponseSingleResultCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsResponseSingleResultCaaRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultCaaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultCaaRecordMeta]
type dnsResponseSingleResultCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultCaaRecordTtlNumber].
type DNSResponseSingleResultCaaRecordTtl interface {
	ImplementsDNSResponseSingleResultCaaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultCaaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultCaaRecordTtlNumber float64

const (
	DNSResponseSingleResultCaaRecordTtlNumber1 DNSResponseSingleResultCaaRecordTtlNumber = 1
)

type DNSResponseSingleResultCertRecord struct {
	// Components of a CERT record.
	Data DNSResponseSingleResultCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultCertRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultCertRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultCertRecordJSON `json:"-"`
}

// dnsResponseSingleResultCertRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultCertRecord]
type dnsResponseSingleResultCertRecordJSON struct {
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

func (r *DNSResponseSingleResultCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultCertRecord) implementsDNSResponseSingleResult() {}

// Components of a CERT record.
type DNSResponseSingleResultCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                   `json:"type"`
	JSON dnsResponseSingleResultCertRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultCertRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultCertRecordData]
type dnsResponseSingleResultCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultCertRecordType string

const (
	DNSResponseSingleResultCertRecordTypeCert DNSResponseSingleResultCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsResponseSingleResultCertRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultCertRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultCertRecordMeta]
type dnsResponseSingleResultCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultCertRecordTtlNumber].
type DNSResponseSingleResultCertRecordTtl interface {
	ImplementsDNSResponseSingleResultCertRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultCertRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultCertRecordTtlNumber float64

const (
	DNSResponseSingleResultCertRecordTtlNumber1 DNSResponseSingleResultCertRecordTtlNumber = 1
)

type DNSResponseSingleResultCnameRecord struct {
	// A valid hostname.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultCnameRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultCnameRecordMeta `json:"meta"`
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
	Ttl DNSResponseSingleResultCnameRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultCnameRecordJSON `json:"-"`
}

// dnsResponseSingleResultCnameRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultCnameRecord]
type dnsResponseSingleResultCnameRecordJSON struct {
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

func (r *DNSResponseSingleResultCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultCnameRecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultCnameRecordType string

const (
	DNSResponseSingleResultCnameRecordTypeCname DNSResponseSingleResultCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsResponseSingleResultCnameRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultCnameRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultCnameRecordMeta]
type dnsResponseSingleResultCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultCnameRecordTtlNumber].
type DNSResponseSingleResultCnameRecordTtl interface {
	ImplementsDNSResponseSingleResultCnameRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultCnameRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultCnameRecordTtlNumber float64

const (
	DNSResponseSingleResultCnameRecordTtlNumber1 DNSResponseSingleResultCnameRecordTtlNumber = 1
)

type DNSResponseSingleResultDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSResponseSingleResultDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultDnskeyRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultDnskeyRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultDnskeyRecordJSON `json:"-"`
}

// dnsResponseSingleResultDnskeyRecordJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultDnskeyRecord]
type dnsResponseSingleResultDnskeyRecordJSON struct {
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

func (r *DNSResponseSingleResultDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultDnskeyRecord) implementsDNSResponseSingleResult() {}

// Components of a DNSKEY record.
type DNSResponseSingleResultDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                      `json:"public_key"`
	JSON      dnsResponseSingleResultDnskeyRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultDnskeyRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultDnskeyRecordData]
type dnsResponseSingleResultDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultDnskeyRecordType string

const (
	DNSResponseSingleResultDnskeyRecordTypeDnskey DNSResponseSingleResultDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsResponseSingleResultDnskeyRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultDnskeyRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultDnskeyRecordMeta]
type dnsResponseSingleResultDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultDnskeyRecordTtlNumber].
type DNSResponseSingleResultDnskeyRecordTtl interface {
	ImplementsDNSResponseSingleResultDnskeyRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultDnskeyRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultDnskeyRecordTtlNumber float64

const (
	DNSResponseSingleResultDnskeyRecordTtlNumber1 DNSResponseSingleResultDnskeyRecordTtlNumber = 1
)

type DNSResponseSingleResultDsRecord struct {
	// Components of a DS record.
	Data DNSResponseSingleResultDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultDsRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultDsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultDsRecordJSON `json:"-"`
}

// dnsResponseSingleResultDsRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultDsRecord]
type dnsResponseSingleResultDsRecordJSON struct {
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

func (r *DNSResponseSingleResultDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultDsRecord) implementsDNSResponseSingleResult() {}

// Components of a DS record.
type DNSResponseSingleResultDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                 `json:"key_tag"`
	JSON   dnsResponseSingleResultDsRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultDsRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultDsRecordData]
type dnsResponseSingleResultDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultDsRecordType string

const (
	DNSResponseSingleResultDsRecordTypeDs DNSResponseSingleResultDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                  `json:"source"`
	JSON   dnsResponseSingleResultDsRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultDsRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultDsRecordMeta]
type dnsResponseSingleResultDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultDsRecordTtlNumber].
type DNSResponseSingleResultDsRecordTtl interface {
	ImplementsDNSResponseSingleResultDsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultDsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultDsRecordTtlNumber float64

const (
	DNSResponseSingleResultDsRecordTtlNumber1 DNSResponseSingleResultDsRecordTtlNumber = 1
)

type DNSResponseSingleResultHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSResponseSingleResultHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultHTTPsRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultHTTPsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultHTTPsRecordJSON `json:"-"`
}

// dnsResponseSingleResultHTTPsRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultHTTPsRecord]
type dnsResponseSingleResultHTTPsRecordJSON struct {
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

func (r *DNSResponseSingleResultHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultHTTPsRecord) implementsDNSResponseSingleResult() {}

// Components of a HTTPS record.
type DNSResponseSingleResultHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                     `json:"value"`
	JSON  dnsResponseSingleResultHTTPsRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultHTTPsRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultHTTPsRecordData]
type dnsResponseSingleResultHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultHTTPsRecordType string

const (
	DNSResponseSingleResultHTTPsRecordTypeHTTPs DNSResponseSingleResultHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsResponseSingleResultHTTPsRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultHTTPsRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultHTTPsRecordMeta]
type dnsResponseSingleResultHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultHTTPsRecordTtlNumber].
type DNSResponseSingleResultHTTPsRecordTtl interface {
	ImplementsDNSResponseSingleResultHTTPsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultHTTPsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultHTTPsRecordTtlNumber float64

const (
	DNSResponseSingleResultHTTPsRecordTtlNumber1 DNSResponseSingleResultHTTPsRecordTtlNumber = 1
)

type DNSResponseSingleResultLocRecord struct {
	// Components of a LOC record.
	Data DNSResponseSingleResultLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultLocRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultLocRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultLocRecordJSON `json:"-"`
}

// dnsResponseSingleResultLocRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultLocRecord]
type dnsResponseSingleResultLocRecordJSON struct {
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

func (r *DNSResponseSingleResultLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultLocRecord) implementsDNSResponseSingleResult() {}

// Components of a LOC record.
type DNSResponseSingleResultLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSResponseSingleResultLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSResponseSingleResultLocRecordDataLongDirection `json:"long_direction"`
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
	JSON dnsResponseSingleResultLocRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultLocRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultLocRecordData]
type dnsResponseSingleResultLocRecordDataJSON struct {
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

func (r *DNSResponseSingleResultLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSResponseSingleResultLocRecordDataLatDirection string

const (
	DNSResponseSingleResultLocRecordDataLatDirectionN DNSResponseSingleResultLocRecordDataLatDirection = "N"
	DNSResponseSingleResultLocRecordDataLatDirectionS DNSResponseSingleResultLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSResponseSingleResultLocRecordDataLongDirection string

const (
	DNSResponseSingleResultLocRecordDataLongDirectionE DNSResponseSingleResultLocRecordDataLongDirection = "E"
	DNSResponseSingleResultLocRecordDataLongDirectionW DNSResponseSingleResultLocRecordDataLongDirection = "W"
)

// Record type.
type DNSResponseSingleResultLocRecordType string

const (
	DNSResponseSingleResultLocRecordTypeLoc DNSResponseSingleResultLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsResponseSingleResultLocRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultLocRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultLocRecordMeta]
type dnsResponseSingleResultLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultLocRecordTtlNumber].
type DNSResponseSingleResultLocRecordTtl interface {
	ImplementsDNSResponseSingleResultLocRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultLocRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultLocRecordTtlNumber float64

const (
	DNSResponseSingleResultLocRecordTtlNumber1 DNSResponseSingleResultLocRecordTtlNumber = 1
)

type DNSResponseSingleResultMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSResponseSingleResultMxRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultMxRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultMxRecordJSON `json:"-"`
}

// dnsResponseSingleResultMxRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultMxRecord]
type dnsResponseSingleResultMxRecordJSON struct {
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

func (r *DNSResponseSingleResultMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultMxRecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultMxRecordType string

const (
	DNSResponseSingleResultMxRecordTypeMx DNSResponseSingleResultMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                  `json:"source"`
	JSON   dnsResponseSingleResultMxRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultMxRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultMxRecordMeta]
type dnsResponseSingleResultMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultMxRecordTtlNumber].
type DNSResponseSingleResultMxRecordTtl interface {
	ImplementsDNSResponseSingleResultMxRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultMxRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultMxRecordTtlNumber float64

const (
	DNSResponseSingleResultMxRecordTtlNumber1 DNSResponseSingleResultMxRecordTtlNumber = 1
)

type DNSResponseSingleResultNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSResponseSingleResultNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultNaptrRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultNaptrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultNaptrRecordJSON `json:"-"`
}

// dnsResponseSingleResultNaptrRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultNaptrRecord]
type dnsResponseSingleResultNaptrRecordJSON struct {
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

func (r *DNSResponseSingleResultNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultNaptrRecord) implementsDNSResponseSingleResult() {}

// Components of a NAPTR record.
type DNSResponseSingleResultNaptrRecordData struct {
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
	JSON    dnsResponseSingleResultNaptrRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultNaptrRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultNaptrRecordData]
type dnsResponseSingleResultNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultNaptrRecordType string

const (
	DNSResponseSingleResultNaptrRecordTypeNaptr DNSResponseSingleResultNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsResponseSingleResultNaptrRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultNaptrRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultNaptrRecordMeta]
type dnsResponseSingleResultNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultNaptrRecordTtlNumber].
type DNSResponseSingleResultNaptrRecordTtl interface {
	ImplementsDNSResponseSingleResultNaptrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultNaptrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultNaptrRecordTtlNumber float64

const (
	DNSResponseSingleResultNaptrRecordTtlNumber1 DNSResponseSingleResultNaptrRecordTtlNumber = 1
)

type DNSResponseSingleResultNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultNsRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultNsRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                              `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultNsRecordJSON `json:"-"`
}

// dnsResponseSingleResultNsRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultNsRecord]
type dnsResponseSingleResultNsRecordJSON struct {
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

func (r *DNSResponseSingleResultNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultNsRecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultNsRecordType string

const (
	DNSResponseSingleResultNsRecordTypeNs DNSResponseSingleResultNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                  `json:"source"`
	JSON   dnsResponseSingleResultNsRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultNsRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultNsRecordMeta]
type dnsResponseSingleResultNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultNsRecordTtlNumber].
type DNSResponseSingleResultNsRecordTtl interface {
	ImplementsDNSResponseSingleResultNsRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultNsRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultNsRecordTtlNumber float64

const (
	DNSResponseSingleResultNsRecordTtlNumber1 DNSResponseSingleResultNsRecordTtlNumber = 1
)

type DNSResponseSingleResultPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultPtrRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultPtrRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultPtrRecordJSON `json:"-"`
}

// dnsResponseSingleResultPtrRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultPtrRecord]
type dnsResponseSingleResultPtrRecordJSON struct {
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

func (r *DNSResponseSingleResultPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultPtrRecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultPtrRecordType string

const (
	DNSResponseSingleResultPtrRecordTypePtr DNSResponseSingleResultPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsResponseSingleResultPtrRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultPtrRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultPtrRecordMeta]
type dnsResponseSingleResultPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultPtrRecordTtlNumber].
type DNSResponseSingleResultPtrRecordTtl interface {
	ImplementsDNSResponseSingleResultPtrRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultPtrRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultPtrRecordTtlNumber float64

const (
	DNSResponseSingleResultPtrRecordTtlNumber1 DNSResponseSingleResultPtrRecordTtlNumber = 1
)

type DNSResponseSingleResultSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSResponseSingleResultSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultSmimeaRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultSmimeaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                  `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultSmimeaRecordJSON `json:"-"`
}

// dnsResponseSingleResultSmimeaRecordJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSmimeaRecord]
type dnsResponseSingleResultSmimeaRecordJSON struct {
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

func (r *DNSResponseSingleResultSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultSmimeaRecord) implementsDNSResponseSingleResult() {}

// Components of a SMIMEA record.
type DNSResponseSingleResultSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                     `json:"usage"`
	JSON  dnsResponseSingleResultSmimeaRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultSmimeaRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSmimeaRecordData]
type dnsResponseSingleResultSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSResponseSingleResultSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultSmimeaRecordType string

const (
	DNSResponseSingleResultSmimeaRecordTypeSmimea DNSResponseSingleResultSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                      `json:"source"`
	JSON   dnsResponseSingleResultSmimeaRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultSmimeaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSmimeaRecordMeta]
type dnsResponseSingleResultSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultSmimeaRecordTtlNumber].
type DNSResponseSingleResultSmimeaRecordTtl interface {
	ImplementsDNSResponseSingleResultSmimeaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultSmimeaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultSmimeaRecordTtlNumber float64

const (
	DNSResponseSingleResultSmimeaRecordTtlNumber1 DNSResponseSingleResultSmimeaRecordTtlNumber = 1
)

type DNSResponseSingleResultSrvRecord struct {
	// Components of a SRV record.
	Data DNSResponseSingleResultSrvRecordData `json:"data,required"`
	// Record type.
	Type DNSResponseSingleResultSrvRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Service, protocol, and SRV name content. See 'data' for setting the individual
	// component values.
	Name string `json:"name"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultSrvRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultSrvRecordJSON `json:"-"`
}

// dnsResponseSingleResultSrvRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultSrvRecord]
type dnsResponseSingleResultSrvRecordJSON struct {
	Data        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	Ttl         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultSrvRecord) implementsDNSResponseSingleResult() {}

// Components of a SRV record.
type DNSResponseSingleResultSrvRecordData struct {
	// A valid hostname.
	Name string `json:"name" format:"hostname"`
	// The port of the service.
	Port float64 `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// A valid protocol.
	Proto string `json:"proto"`
	// A service type, prefixed with an underscore.
	Service string `json:"service"`
	// A valid hostname.
	Target string `json:"target" format:"hostname"`
	// The record weight.
	Weight float64                                  `json:"weight"`
	JSON   dnsResponseSingleResultSrvRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultSrvRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSrvRecordData]
type dnsResponseSingleResultSrvRecordDataJSON struct {
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

func (r *DNSResponseSingleResultSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultSrvRecordType string

const (
	DNSResponseSingleResultSrvRecordTypeSrv DNSResponseSingleResultSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsResponseSingleResultSrvRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultSrvRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSrvRecordMeta]
type dnsResponseSingleResultSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultSrvRecordTtlNumber].
type DNSResponseSingleResultSrvRecordTtl interface {
	ImplementsDNSResponseSingleResultSrvRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultSrvRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultSrvRecordTtlNumber float64

const (
	DNSResponseSingleResultSrvRecordTtlNumber1 DNSResponseSingleResultSrvRecordTtlNumber = 1
)

type DNSResponseSingleResultSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSResponseSingleResultSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultSshfpRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultSshfpRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                 `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultSshfpRecordJSON `json:"-"`
}

// dnsResponseSingleResultSshfpRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultSshfpRecord]
type dnsResponseSingleResultSshfpRecordJSON struct {
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

func (r *DNSResponseSingleResultSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultSshfpRecord) implementsDNSResponseSingleResult() {}

// Components of a SSHFP record.
type DNSResponseSingleResultSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                    `json:"type"`
	JSON dnsResponseSingleResultSshfpRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultSshfpRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSshfpRecordData]
type dnsResponseSingleResultSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultSshfpRecordType string

const (
	DNSResponseSingleResultSshfpRecordTypeSshfp DNSResponseSingleResultSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                     `json:"source"`
	JSON   dnsResponseSingleResultSshfpRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultSshfpRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSshfpRecordMeta]
type dnsResponseSingleResultSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultSshfpRecordTtlNumber].
type DNSResponseSingleResultSshfpRecordTtl interface {
	ImplementsDNSResponseSingleResultSshfpRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultSshfpRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultSshfpRecordTtlNumber float64

const (
	DNSResponseSingleResultSshfpRecordTtlNumber1 DNSResponseSingleResultSshfpRecordTtlNumber = 1
)

type DNSResponseSingleResultSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSResponseSingleResultSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultSvcbRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultSvcbRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultSvcbRecordJSON `json:"-"`
}

// dnsResponseSingleResultSvcbRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultSvcbRecord]
type dnsResponseSingleResultSvcbRecordJSON struct {
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

func (r *DNSResponseSingleResultSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultSvcbRecord) implementsDNSResponseSingleResult() {}

// Components of a SVCB record.
type DNSResponseSingleResultSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                    `json:"value"`
	JSON  dnsResponseSingleResultSvcbRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultSvcbRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSvcbRecordData]
type dnsResponseSingleResultSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultSvcbRecordType string

const (
	DNSResponseSingleResultSvcbRecordTypeSvcb DNSResponseSingleResultSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsResponseSingleResultSvcbRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultSvcbRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultSvcbRecordMeta]
type dnsResponseSingleResultSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultSvcbRecordTtlNumber].
type DNSResponseSingleResultSvcbRecordTtl interface {
	ImplementsDNSResponseSingleResultSvcbRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultSvcbRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultSvcbRecordTtlNumber float64

const (
	DNSResponseSingleResultSvcbRecordTtlNumber1 DNSResponseSingleResultSvcbRecordTtlNumber = 1
)

type DNSResponseSingleResultTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSResponseSingleResultTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultTlsaRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultTlsaRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultTlsaRecordJSON `json:"-"`
}

// dnsResponseSingleResultTlsaRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultTlsaRecord]
type dnsResponseSingleResultTlsaRecordJSON struct {
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

func (r *DNSResponseSingleResultTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultTlsaRecord) implementsDNSResponseSingleResult() {}

// Components of a TLSA record.
type DNSResponseSingleResultTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                   `json:"usage"`
	JSON  dnsResponseSingleResultTlsaRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultTlsaRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultTlsaRecordData]
type dnsResponseSingleResultTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSResponseSingleResultTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultTlsaRecordType string

const (
	DNSResponseSingleResultTlsaRecordTypeTlsa DNSResponseSingleResultTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                    `json:"source"`
	JSON   dnsResponseSingleResultTlsaRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultTlsaRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultTlsaRecordMeta]
type dnsResponseSingleResultTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultTlsaRecordTtlNumber].
type DNSResponseSingleResultTlsaRecordTtl interface {
	ImplementsDNSResponseSingleResultTlsaRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultTlsaRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultTlsaRecordTtlNumber float64

const (
	DNSResponseSingleResultTlsaRecordTtlNumber1 DNSResponseSingleResultTlsaRecordTtlNumber = 1
)

type DNSResponseSingleResultTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Record type.
	Type DNSResponseSingleResultTxtRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultTxtRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultTxtRecordJSON `json:"-"`
}

// dnsResponseSingleResultTxtRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultTxtRecord]
type dnsResponseSingleResultTxtRecordJSON struct {
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

func (r *DNSResponseSingleResultTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultTxtRecord) implementsDNSResponseSingleResult() {}

// Record type.
type DNSResponseSingleResultTxtRecordType string

const (
	DNSResponseSingleResultTxtRecordTypeTxt DNSResponseSingleResultTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsResponseSingleResultTxtRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultTxtRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultTxtRecordMeta]
type dnsResponseSingleResultTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultTxtRecordTtlNumber].
type DNSResponseSingleResultTxtRecordTtl interface {
	ImplementsDNSResponseSingleResultTxtRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultTxtRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultTxtRecordTtlNumber float64

const (
	DNSResponseSingleResultTxtRecordTtlNumber1 DNSResponseSingleResultTxtRecordTtlNumber = 1
)

type DNSResponseSingleResultUriRecord struct {
	// Components of a URI record.
	Data DNSResponseSingleResultUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSResponseSingleResultUriRecordType `json:"type,required"`
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
	Meta DNSResponseSingleResultUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl DNSResponseSingleResultUriRecordTtl `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                               `json:"zone_name" format:"hostname"`
	JSON     dnsResponseSingleResultUriRecordJSON `json:"-"`
}

// dnsResponseSingleResultUriRecordJSON contains the JSON metadata for the struct
// [DNSResponseSingleResultUriRecord]
type dnsResponseSingleResultUriRecordJSON struct {
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

func (r *DNSResponseSingleResultUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSResponseSingleResultUriRecord) implementsDNSResponseSingleResult() {}

// Components of a URI record.
type DNSResponseSingleResultUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                  `json:"weight"`
	JSON   dnsResponseSingleResultUriRecordDataJSON `json:"-"`
}

// dnsResponseSingleResultUriRecordDataJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultUriRecordData]
type dnsResponseSingleResultUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSResponseSingleResultUriRecordType string

const (
	DNSResponseSingleResultUriRecordTypeUri DNSResponseSingleResultUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSResponseSingleResultUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                   `json:"source"`
	JSON   dnsResponseSingleResultUriRecordMetaJSON `json:"-"`
}

// dnsResponseSingleResultUriRecordMetaJSON contains the JSON metadata for the
// struct [DNSResponseSingleResultUriRecordMeta]
type dnsResponseSingleResultUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingleResultUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSResponseSingleResultUriRecordTtlNumber].
type DNSResponseSingleResultUriRecordTtl interface {
	ImplementsDNSResponseSingleResultUriRecordTtl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSResponseSingleResultUriRecordTtl)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSResponseSingleResultUriRecordTtlNumber float64

const (
	DNSResponseSingleResultUriRecordTtlNumber1 DNSResponseSingleResultUriRecordTtlNumber = 1
)

// Whether the API call was successful
type DNSResponseSingleSuccess bool

const (
	DNSResponseSingleSuccessTrue DNSResponseSingleSuccess = true
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

// This interface is a union satisfied by one of the following:
// [ZoneDNSRecordUpdateParamsARecord], [ZoneDNSRecordUpdateParamsAaaaRecord],
// [ZoneDNSRecordUpdateParamsCaaRecord], [ZoneDNSRecordUpdateParamsCertRecord],
// [ZoneDNSRecordUpdateParamsCnameRecord], [ZoneDNSRecordUpdateParamsDnskeyRecord],
// [ZoneDNSRecordUpdateParamsDsRecord], [ZoneDNSRecordUpdateParamsHTTPsRecord],
// [ZoneDNSRecordUpdateParamsLocRecord], [ZoneDNSRecordUpdateParamsMxRecord],
// [ZoneDNSRecordUpdateParamsNaptrRecord], [ZoneDNSRecordUpdateParamsNsRecord],
// [ZoneDNSRecordUpdateParamsPtrRecord], [ZoneDNSRecordUpdateParamsSmimeaRecord],
// [ZoneDNSRecordUpdateParamsSrvRecord], [ZoneDNSRecordUpdateParamsSshfpRecord],
// [ZoneDNSRecordUpdateParamsSvcbRecord], [ZoneDNSRecordUpdateParamsTlsaRecord],
// [ZoneDNSRecordUpdateParamsTxtRecord], [ZoneDNSRecordUpdateParamsUriRecord].
type ZoneDNSRecordUpdateParams interface {
	ImplementsZoneDNSRecordUpdateParams()
}

type ZoneDNSRecordUpdateParamsARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsARecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordUpdateParamsARecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsARecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsARecordType string

const (
	ZoneDNSRecordUpdateParamsARecordTypeA ZoneDNSRecordUpdateParamsARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsARecordTtlNumber].
type ZoneDNSRecordUpdateParamsARecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsARecordTtl()
}

type ZoneDNSRecordUpdateParamsARecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsARecordTtlNumber1 ZoneDNSRecordUpdateParamsARecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsAaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsAaaaRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordUpdateParamsAaaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsAaaaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsAaaaRecordType string

const (
	ZoneDNSRecordUpdateParamsAaaaRecordTypeAaaa ZoneDNSRecordUpdateParamsAaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsAaaaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsAaaaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsAaaaRecordTtl()
}

type ZoneDNSRecordUpdateParamsAaaaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsAaaaRecordTtlNumber1 ZoneDNSRecordUpdateParamsAaaaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsCaaRecord struct {
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordUpdateParamsCaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsCaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsCaaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a CAA record.
type ZoneDNSRecordUpdateParamsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordUpdateParamsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsCaaRecordType string

const (
	ZoneDNSRecordUpdateParamsCaaRecordTypeCaa ZoneDNSRecordUpdateParamsCaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsCaaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsCaaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsCaaRecordTtl()
}

type ZoneDNSRecordUpdateParamsCaaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsCaaRecordTtlNumber1 ZoneDNSRecordUpdateParamsCaaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsCertRecord struct {
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordUpdateParamsCertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsCertRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsCertRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a CERT record.
type ZoneDNSRecordUpdateParamsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordUpdateParamsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsCertRecordType string

const (
	ZoneDNSRecordUpdateParamsCertRecordTypeCert ZoneDNSRecordUpdateParamsCertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsCertRecordTtlNumber].
type ZoneDNSRecordUpdateParamsCertRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsCertRecordTtl()
}

type ZoneDNSRecordUpdateParamsCertRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsCertRecordTtlNumber1 ZoneDNSRecordUpdateParamsCertRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsCnameRecord struct {
	// A valid hostname.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsCnameRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordUpdateParamsCnameRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsCnameRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsCnameRecordType string

const (
	ZoneDNSRecordUpdateParamsCnameRecordTypeCname ZoneDNSRecordUpdateParamsCnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsCnameRecordTtlNumber].
type ZoneDNSRecordUpdateParamsCnameRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsCnameRecordTtl()
}

type ZoneDNSRecordUpdateParamsCnameRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsCnameRecordTtlNumber1 ZoneDNSRecordUpdateParamsCnameRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordUpdateParamsDnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsDnskeyRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsDnskeyRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a DNSKEY record.
type ZoneDNSRecordUpdateParamsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordUpdateParamsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsDnskeyRecordType string

const (
	ZoneDNSRecordUpdateParamsDnskeyRecordTypeDnskey ZoneDNSRecordUpdateParamsDnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsDnskeyRecordTtlNumber].
type ZoneDNSRecordUpdateParamsDnskeyRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsDnskeyRecordTtl()
}

type ZoneDNSRecordUpdateParamsDnskeyRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsDnskeyRecordTtlNumber1 ZoneDNSRecordUpdateParamsDnskeyRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsDsRecord struct {
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordUpdateParamsDsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsDsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsDsRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a DS record.
type ZoneDNSRecordUpdateParamsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordUpdateParamsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsDsRecordType string

const (
	ZoneDNSRecordUpdateParamsDsRecordTypeDs ZoneDNSRecordUpdateParamsDsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsDsRecordTtlNumber].
type ZoneDNSRecordUpdateParamsDsRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsDsRecordTtl()
}

type ZoneDNSRecordUpdateParamsDsRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsDsRecordTtlNumber1 ZoneDNSRecordUpdateParamsDsRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordUpdateParamsHTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsHTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsHTTPsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsHTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsHTTPsRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a HTTPS record.
type ZoneDNSRecordUpdateParamsHTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordUpdateParamsHTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsHTTPsRecordType string

const (
	ZoneDNSRecordUpdateParamsHTTPsRecordTypeHTTPs ZoneDNSRecordUpdateParamsHTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsHTTPsRecordTtlNumber].
type ZoneDNSRecordUpdateParamsHTTPsRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsHTTPsRecordTtl()
}

type ZoneDNSRecordUpdateParamsHTTPsRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsHTTPsRecordTtlNumber1 ZoneDNSRecordUpdateParamsHTTPsRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsLocRecord struct {
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordUpdateParamsLocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsLocRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsLocRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a LOC record.
type ZoneDNSRecordUpdateParamsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordUpdateParamsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordUpdateParamsLocRecordDataLongDirection] `json:"long_direction"`
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

func (r ZoneDNSRecordUpdateParamsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordUpdateParamsLocRecordDataLatDirection string

const (
	ZoneDNSRecordUpdateParamsLocRecordDataLatDirectionN ZoneDNSRecordUpdateParamsLocRecordDataLatDirection = "N"
	ZoneDNSRecordUpdateParamsLocRecordDataLatDirectionS ZoneDNSRecordUpdateParamsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type ZoneDNSRecordUpdateParamsLocRecordDataLongDirection string

const (
	ZoneDNSRecordUpdateParamsLocRecordDataLongDirectionE ZoneDNSRecordUpdateParamsLocRecordDataLongDirection = "E"
	ZoneDNSRecordUpdateParamsLocRecordDataLongDirectionW ZoneDNSRecordUpdateParamsLocRecordDataLongDirection = "W"
)

// Record type.
type ZoneDNSRecordUpdateParamsLocRecordType string

const (
	ZoneDNSRecordUpdateParamsLocRecordTypeLoc ZoneDNSRecordUpdateParamsLocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsLocRecordTtlNumber].
type ZoneDNSRecordUpdateParamsLocRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsLocRecordTtl()
}

type ZoneDNSRecordUpdateParamsLocRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsLocRecordTtlNumber1 ZoneDNSRecordUpdateParamsLocRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsMxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsMxRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsMxRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsMxRecordType string

const (
	ZoneDNSRecordUpdateParamsMxRecordTypeMx ZoneDNSRecordUpdateParamsMxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsMxRecordTtlNumber].
type ZoneDNSRecordUpdateParamsMxRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsMxRecordTtl()
}

type ZoneDNSRecordUpdateParamsMxRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsMxRecordTtlNumber1 ZoneDNSRecordUpdateParamsMxRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsNaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordUpdateParamsNaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsNaptrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsNaptrRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a NAPTR record.
type ZoneDNSRecordUpdateParamsNaptrRecordData struct {
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

func (r ZoneDNSRecordUpdateParamsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsNaptrRecordType string

const (
	ZoneDNSRecordUpdateParamsNaptrRecordTypeNaptr ZoneDNSRecordUpdateParamsNaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsNaptrRecordTtlNumber].
type ZoneDNSRecordUpdateParamsNaptrRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsNaptrRecordTtl()
}

type ZoneDNSRecordUpdateParamsNaptrRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsNaptrRecordTtlNumber1 ZoneDNSRecordUpdateParamsNaptrRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsNsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsNsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsNsRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsNsRecordType string

const (
	ZoneDNSRecordUpdateParamsNsRecordTypeNs ZoneDNSRecordUpdateParamsNsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsNsRecordTtlNumber].
type ZoneDNSRecordUpdateParamsNsRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsNsRecordTtl()
}

type ZoneDNSRecordUpdateParamsNsRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsNsRecordTtlNumber1 ZoneDNSRecordUpdateParamsNsRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsPtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsPtrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsPtrRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsPtrRecordType string

const (
	ZoneDNSRecordUpdateParamsPtrRecordTypePtr ZoneDNSRecordUpdateParamsPtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsPtrRecordTtlNumber].
type ZoneDNSRecordUpdateParamsPtrRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsPtrRecordTtl()
}

type ZoneDNSRecordUpdateParamsPtrRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsPtrRecordTtlNumber1 ZoneDNSRecordUpdateParamsPtrRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordUpdateParamsSmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsSmimeaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsSmimeaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SMIMEA record.
type ZoneDNSRecordUpdateParamsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordUpdateParamsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsSmimeaRecordType string

const (
	ZoneDNSRecordUpdateParamsSmimeaRecordTypeSmimea ZoneDNSRecordUpdateParamsSmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsSmimeaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsSmimeaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsSmimeaRecordTtl()
}

type ZoneDNSRecordUpdateParamsSmimeaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsSmimeaRecordTtlNumber1 ZoneDNSRecordUpdateParamsSmimeaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsSrvRecord struct {
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordUpdateParamsSrvRecordData] `json:"data,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsSrvRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsSrvRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SRV record.
type ZoneDNSRecordUpdateParamsSrvRecordData struct {
	// A valid hostname.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol.
	Proto param.Field[string] `json:"proto"`
	// A service type, prefixed with an underscore.
	Service param.Field[string] `json:"service"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordUpdateParamsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsSrvRecordType string

const (
	ZoneDNSRecordUpdateParamsSrvRecordTypeSrv ZoneDNSRecordUpdateParamsSrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsSrvRecordTtlNumber].
type ZoneDNSRecordUpdateParamsSrvRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsSrvRecordTtl()
}

type ZoneDNSRecordUpdateParamsSrvRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsSrvRecordTtlNumber1 ZoneDNSRecordUpdateParamsSrvRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsSshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordUpdateParamsSshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsSshfpRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsSshfpRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SSHFP record.
type ZoneDNSRecordUpdateParamsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordUpdateParamsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsSshfpRecordType string

const (
	ZoneDNSRecordUpdateParamsSshfpRecordTypeSshfp ZoneDNSRecordUpdateParamsSshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsSshfpRecordTtlNumber].
type ZoneDNSRecordUpdateParamsSshfpRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsSshfpRecordTtl()
}

type ZoneDNSRecordUpdateParamsSshfpRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsSshfpRecordTtlNumber1 ZoneDNSRecordUpdateParamsSshfpRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsSvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordUpdateParamsSvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsSvcbRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsSvcbRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a SVCB record.
type ZoneDNSRecordUpdateParamsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordUpdateParamsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsSvcbRecordType string

const (
	ZoneDNSRecordUpdateParamsSvcbRecordTypeSvcb ZoneDNSRecordUpdateParamsSvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsSvcbRecordTtlNumber].
type ZoneDNSRecordUpdateParamsSvcbRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsSvcbRecordTtl()
}

type ZoneDNSRecordUpdateParamsSvcbRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsSvcbRecordTtlNumber1 ZoneDNSRecordUpdateParamsSvcbRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsTlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordUpdateParamsTlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsTlsaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsTlsaRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a TLSA record.
type ZoneDNSRecordUpdateParamsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordUpdateParamsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsTlsaRecordType string

const (
	ZoneDNSRecordUpdateParamsTlsaRecordTypeTlsa ZoneDNSRecordUpdateParamsTlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordUpdateParamsTlsaRecordTtlNumber].
type ZoneDNSRecordUpdateParamsTlsaRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsTlsaRecordTtl()
}

type ZoneDNSRecordUpdateParamsTlsaRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsTlsaRecordTtlNumber1 ZoneDNSRecordUpdateParamsTlsaRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsTxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsTxtRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsTxtRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Record type.
type ZoneDNSRecordUpdateParamsTxtRecordType string

const (
	ZoneDNSRecordUpdateParamsTxtRecordTypeTxt ZoneDNSRecordUpdateParamsTxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsTxtRecordTtlNumber].
type ZoneDNSRecordUpdateParamsTxtRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsTxtRecordTtl()
}

type ZoneDNSRecordUpdateParamsTxtRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsTxtRecordTtlNumber1 ZoneDNSRecordUpdateParamsTxtRecordTtlNumber = 1
)

type ZoneDNSRecordUpdateParamsUriRecord struct {
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordUpdateParamsUriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordUpdateParamsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordUpdateParamsUriRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordUpdateParamsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordUpdateParamsUriRecord) ImplementsZoneDNSRecordUpdateParams() {

}

// Components of a URI record.
type ZoneDNSRecordUpdateParamsUriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordUpdateParamsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordUpdateParamsUriRecordType string

const (
	ZoneDNSRecordUpdateParamsUriRecordTypeUri ZoneDNSRecordUpdateParamsUriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat], [ZoneDNSRecordUpdateParamsUriRecordTtlNumber].
type ZoneDNSRecordUpdateParamsUriRecordTtl interface {
	ImplementsZoneDNSRecordUpdateParamsUriRecordTtl()
}

type ZoneDNSRecordUpdateParamsUriRecordTtlNumber float64

const (
	ZoneDNSRecordUpdateParamsUriRecordTtlNumber1 ZoneDNSRecordUpdateParamsUriRecordTtlNumber = 1
)

// This interface is a union satisfied by one of the following:
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecord],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecord].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTypeA ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsARecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTypeAaaa ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsAaaaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecord struct {
	// Components of a CAA record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a CAA record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTypeCaa ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCaaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecord struct {
	// Components of a CERT record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a CERT record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTypeCert ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCertRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecord struct {
	// A valid hostname.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordType] `json:"type,required"`
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
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTypeCname ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsCnameRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a DNSKEY record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTypeDnskey ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDnskeyRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecord struct {
	// Components of a DS record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a DS record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTypeDs ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a HTTPS record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTypeHTTPs ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsHTTPsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecord struct {
	// Components of a LOC record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a LOC record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLongDirection] `json:"long_direction"`
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

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLatDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLatDirectionN ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLatDirection = "N"
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLatDirectionS ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLongDirection string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLongDirectionE ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLongDirection = "E"
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLongDirectionW ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordDataLongDirection = "W"
)

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTypeLoc ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsLocRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTypeMx ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsMxRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a NAPTR record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordData struct {
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

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTypeNaptr ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNaptrRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTypeNs ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsNsRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTypePtr ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsPtrRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SMIMEA record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTypeSmimea ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSmimeaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecord struct {
	// Components of a SRV record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordData] `json:"data,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SRV record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordData struct {
	// A valid hostname.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid protocol.
	Proto param.Field[string] `json:"proto"`
	// A service type, prefixed with an underscore.
	Service param.Field[string] `json:"service"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTypeSrv ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSrvRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SSHFP record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTypeSshfp ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSshfpRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SVCB record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTypeSvcb ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsSvcbRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a TLSA record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTypeTlsa ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTlsaRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTypeTxt ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsTxtRecordTtlNumber = 1
)

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecord struct {
	// Components of a URI record.
	Data param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex).
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	Ttl param.Field[ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtl] `json:"ttl"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecord) ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a URI record.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordType string

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTypeUri ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtlNumber].
type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtl interface {
	ImplementsZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtl()
}

type ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtlNumber float64

const (
	ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtlNumber1 ZoneDNSRecordDNSRecordsForAZoneNewDNSRecordParamsUriRecordTtlNumber = 1
)
