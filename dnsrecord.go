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
	Exports *DNSRecordExportService
	Imports *DNSRecordImportService
	Scans   *DNSRecordScanService
}

// NewDNSRecordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSRecordService(opts ...option.RequestOption) (r *DNSRecordService) {
	r = &DNSRecordService{}
	r.Options = opts
	r.Exports = NewDNSRecordExportService(opts...)
	r.Imports = NewDNSRecordImportService(opts...)
	r.Scans = NewDNSRecordScanService(opts...)
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

// Update an existing DNS record.
//
// Notes:
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

// Create a new DNS record for a zone.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *DNSRecordService) DNSRecordsForAZoneNewDNSRecord(ctx context.Context, zoneID string, body DNSRecordDNSRecordsForAZoneNewDNSRecordParams, opts ...option.RequestOption) (res *DNSRecordDNSRecordsForAZoneNewDNSRecordResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *DNSRecordService) DNSRecordsForAZoneListDNSRecords(ctx context.Context, zoneID string, query DNSRecordDNSRecordsForAZoneListDNSRecordsParams, opts ...option.RequestOption) (res *[]DNSRecordDNSRecordsForAZoneListDNSRecordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DNSRecordGetResponseDNSRecordsARecord],
// [DNSRecordGetResponseDNSRecordsAaaaRecord],
// [DNSRecordGetResponseDNSRecordsCaaRecord],
// [DNSRecordGetResponseDNSRecordsCertRecord],
// [DNSRecordGetResponseDNSRecordsCnameRecord],
// [DNSRecordGetResponseDNSRecordsDnskeyRecord],
// [DNSRecordGetResponseDNSRecordsDsRecord],
// [DNSRecordGetResponseDNSRecordsHTTPsRecord],
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

type DNSRecordGetResponseDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordGetResponseDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsHTTPsRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsHTTPsRecord]
type dnsRecordGetResponseDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsHTTPsRecord) implementsDNSRecordGetResponse() {}

// Components of a HTTPS record.
type DNSRecordGetResponseDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                            `json:"value"`
	JSON  dnsRecordGetResponseDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsHTTPsRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsHTTPsRecordData]
type dnsRecordGetResponseDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsHTTPsRecordType string

const (
	DNSRecordGetResponseDNSRecordsHTTPsRecordTypeHTTPs DNSRecordGetResponseDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsHTTPsRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsHTTPsRecordMeta]
type dnsRecordGetResponseDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsHTTPsRecordTTLNumber1 DNSRecordGetResponseDNSRecordsHTTPsRecordTTLNumber = 1
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

// Union satisfied by [DNSRecordUpdateResponseDNSRecordsARecord],
// [DNSRecordUpdateResponseDNSRecordsAaaaRecord],
// [DNSRecordUpdateResponseDNSRecordsCaaRecord],
// [DNSRecordUpdateResponseDNSRecordsCertRecord],
// [DNSRecordUpdateResponseDNSRecordsCnameRecord],
// [DNSRecordUpdateResponseDNSRecordsDnskeyRecord],
// [DNSRecordUpdateResponseDNSRecordsDsRecord],
// [DNSRecordUpdateResponseDNSRecordsHTTPsRecord],
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

type DNSRecordUpdateResponseDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordUpdateResponseDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsHTTPsRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsHTTPsRecord]
type dnsRecordUpdateResponseDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsHTTPsRecord) implementsDNSRecordUpdateResponse() {}

// Components of a HTTPS record.
type DNSRecordUpdateResponseDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                               `json:"value"`
	JSON  dnsRecordUpdateResponseDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsHTTPsRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsHTTPsRecordData]
type dnsRecordUpdateResponseDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsHTTPsRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsHTTPsRecordTypeHTTPs DNSRecordUpdateResponseDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsHTTPsRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsHTTPsRecordMeta]
type dnsRecordUpdateResponseDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsHTTPsRecordTTLNumber = 1
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

// Union satisfied by
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecord] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecord].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponse interface {
	implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponse)(nil)).Elem(), "")
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTypeA DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                   `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a CAA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                                     `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTypeCaa DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a CERT record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                                     `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTypeCert DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTypeCname DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a DNSKEY record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                                        `json:"public_key"`
	JSON      dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a DS record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                                                   `json:"key_tag"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTypeDs DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                    `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a HTTPS record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                       `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTypeHTTPs DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a LOC record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                                                    `json:"size"`
	JSON dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTypeLoc DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTypeMx DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                    `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a NAPTR record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordData struct {
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
	Service string                                                                       `json:"service"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTypeNs DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                    `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTypePtr DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a SMIMEA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                       `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a SRV record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordData struct {
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
	Weight float64                                                                    `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTypeSrv DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a SSHFP record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                                      `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a SVCB record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                      `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a TLSA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                     `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTypeTxt DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponse() {
}

// Components of a URI record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                                                    `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTypeUri DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseDNSRecordsUriRecordTTLNumber = 1
)

// Union satisfied by
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecord] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecord].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponse interface {
	implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponse)(nil)).Elem(), "")
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTypeA DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                     `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTypeAaaa DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a CAA record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                                       `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTypeCaa DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a CERT record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                                       `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTypeCert DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTypeCname DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a DNSKEY record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                                          `json:"public_key"`
	JSON      dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTypeDnskey DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a DS record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                                                     `json:"key_tag"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTypeDs DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a HTTPS record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                         `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTypeHTTPs DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a LOC record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                                                      `json:"size"`
	JSON dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLatDirectionN DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLatDirectionS DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLongDirectionE DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLongDirectionW DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTypeLoc DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTypeMx DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a NAPTR record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordData struct {
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
	Service string                                                                         `json:"service"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTypeNaptr DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTypeNs DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                      `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTypePtr DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SMIMEA record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                         `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SRV record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordData struct {
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
	Weight float64                                                                      `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTypeSrv DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SSHFP record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                                        `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTypeSshfp DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a SVCB record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                        `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTypeSvcb DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a TLSA record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                       `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTypeTlsa DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                        `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTypeTxt DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                   `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponse() {
}

// Components of a URI record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                                                      `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTypeUri DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                       `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseDNSRecordsUriRecordTTLNumber = 1
)

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

// This interface is a union satisfied by one of the following:
// [DNSRecordUpdateParamsDNSRecordsARecord],
// [DNSRecordUpdateParamsDNSRecordsAaaaRecord],
// [DNSRecordUpdateParamsDNSRecordsCaaRecord],
// [DNSRecordUpdateParamsDNSRecordsCertRecord],
// [DNSRecordUpdateParamsDNSRecordsCnameRecord],
// [DNSRecordUpdateParamsDNSRecordsDnskeyRecord],
// [DNSRecordUpdateParamsDNSRecordsDsRecord],
// [DNSRecordUpdateParamsDNSRecordsHTTPsRecord],
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

type DNSRecordUpdateParamsDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[DNSRecordUpdateParamsDNSRecordsHTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordUpdateParamsDNSRecordsHTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordUpdateParamsDNSRecordsHTTPsRecordTTL] `json:"ttl"`
}

func (r DNSRecordUpdateParamsDNSRecordsHTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordUpdateParamsDNSRecordsHTTPsRecord) ImplementsDNSRecordUpdateParams() {

}

// Components of a HTTPS record.
type DNSRecordUpdateParamsDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordUpdateParamsDNSRecordsHTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordUpdateParamsDNSRecordsHTTPsRecordType string

const (
	DNSRecordUpdateParamsDNSRecordsHTTPsRecordTypeHTTPs DNSRecordUpdateParamsDNSRecordsHTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordUpdateParamsDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordUpdateParamsDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordUpdateParamsDNSRecordsHTTPsRecordTTL()
}

type DNSRecordUpdateParamsDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordUpdateParamsDNSRecordsHTTPsRecordTTLNumber1 DNSRecordUpdateParamsDNSRecordsHTTPsRecordTTLNumber = 1
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

// This interface is a union satisfied by one of the following:
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecord].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParams interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content param.Field[string] `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTypeA DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordType = "A"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsARecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content param.Field[string] `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTypeAaaa DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordType = "AAAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a CAA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTypeCaa DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordType = "CAA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a CERT record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTypeCert DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordType = "CERT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordType] `json:"type,required"`
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
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTypeCname DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordType = "CNAME"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a DNSKEY record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTypeDnskey DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a DS record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTypeDs DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordType = "DS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a HTTPS record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTypeHTTPs DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordType = "HTTPS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a LOC record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLongDirection] `json:"long_direction"`
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

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Latitude direction.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLatDirectionN DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLatDirectionS DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLongDirectionE DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLongDirectionW DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTypeLoc DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordType = "LOC"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content param.Field[string] `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTypeMx DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordType = "MX"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a NAPTR record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordData struct {
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

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTypeNaptr DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordType = "NAPTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content param.Field[interface{}] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTypeNs DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordType = "NS"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTypePtr DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordType = "PTR"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SMIMEA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTypeSmimea DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SRV record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordData struct {
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

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTypeSrv DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordType = "SRV"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SSHFP record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTypeSshfp DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordType = "SSHFP"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a SVCB record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordData struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTypeSvcb DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordType = "SVCB"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a TLSA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTypeTlsa DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordType = "TLSA"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTypeTxt DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordType = "TXT"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordData] `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority,required"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]string] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTL] `json:"ttl"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecord) ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParams() {

}

// Components of a URI record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordData struct {
	// The record content.
	Content param.Field[string] `json:"content"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTypeUri DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordType = "URI"
)

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Satisfied by [shared.UnionFloat],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTL()
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordParamsDNSRecordsUriRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelope struct {
	Errors   []DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordDNSRecordsForAZoneNewDNSRecordResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelope]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrors]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessages]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeSuccess bool

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeSuccessTrue DNSRecordDNSRecordsForAZoneNewDNSRecordResponseEnvelopeSuccess = true
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsParams struct {
	Comment param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment] `query:"comment"`
	// DNS record content.
	Content param.Field[string] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch] `query:"match"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `query:"name"`
	// Field to order DNS records by.
	Order param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder] `query:"order"`
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
	Search param.Field[string]                                             `query:"search"`
	Tag    param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag] `query:"tag"`
	// Whether to match all tag search requirements or at least one (any). If set to
	// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
	// logical OR instead. Note that the regular `match` parameter is still used to
	// combine the resulting condition with other filters that aren't related to tags.
	TagMatch param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch] `query:"tag_match"`
	// Record type.
	Type param.Field[DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType] `query:"type"`
}

// URLQuery serializes [DNSRecordDNSRecordsForAZoneListDNSRecordsParams]'s query
// parameters as `url.Values`.
func (r DNSRecordDNSRecordsForAZoneListDNSRecordsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment struct {
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

// URLQuery serializes [DNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment]'s
// query parameters as `url.Values`.
func (r DNSRecordDNSRecordsForAZoneListDNSRecordsParamsComment) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order DNS records in.
type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirectionAsc  DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection = "asc"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirectionDesc DNSRecordDNSRecordsForAZoneListDNSRecordsParamsDirection = "desc"
)

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatchAny DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch = "any"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatchAll DNSRecordDNSRecordsForAZoneListDNSRecordsParamsMatch = "all"
)

// Field to order DNS records by.
type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderType    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "type"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderName    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "name"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderContent DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "content"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderTTL     DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "ttl"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrderProxied DNSRecordDNSRecordsForAZoneListDNSRecordsParamsOrder = "proxied"
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag struct {
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

// URLQuery serializes [DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag]'s query
// parameters as `url.Values`.
func (r DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTag) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to match all tag search requirements or at least one (any). If set to
// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
// logical OR instead. Note that the regular `match` parameter is still used to
// combine the resulting condition with other filters that aren't related to tags.
type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatchAny DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch = "any"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatchAll DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTagMatch = "all"
)

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeA      DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "A"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeAaaa   DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "AAAA"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeCaa    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "CAA"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeCert   DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "CERT"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeCname  DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "CNAME"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeDnskey DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "DNSKEY"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeDs     DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "DS"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeHTTPs  DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "HTTPS"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeLoc    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "LOC"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeMx     DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "MX"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeNaptr  DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "NAPTR"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeNs     DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "NS"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypePtr    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "PTR"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSmimea DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SMIMEA"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSrv    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SRV"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSshfp  DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SSHFP"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeSvcb   DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "SVCB"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeTlsa   DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "TLSA"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeTxt    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "TXT"
	DNSRecordDNSRecordsForAZoneListDNSRecordsParamsTypeUri    DNSRecordDNSRecordsForAZoneListDNSRecordsParamsType = "URI"
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelope struct {
	Errors   []DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessages `json:"messages,required"`
	Result   []DNSRecordDNSRecordsForAZoneListDNSRecordsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeJSON       `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelope]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrors]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessages]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeSuccess bool

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeSuccessTrue DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeSuccess = true
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfoJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfo]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
