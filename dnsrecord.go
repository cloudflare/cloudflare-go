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
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
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
func (r *DNSRecordService) Update(ctx context.Context, zoneID string, dnsRecordID string, body DNSRecordUpdateParams, opts ...option.RequestOption) (res *DNSRecordUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/%s", zoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
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
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *DNSRecordService) DNSRecordsForAZoneListDNSRecords(ctx context.Context, zoneID string, query DNSRecordDNSRecordsForAZoneListDNSRecordsParams, opts ...option.RequestOption) (res *DNSRecordDNSRecordsForAZoneListDNSRecordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DNSRecordGetResponse struct {
	Errors   []DNSRecordGetResponseError   `json:"errors"`
	Messages []DNSRecordGetResponseMessage `json:"messages"`
	Result   DNSRecordGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DNSRecordGetResponseSuccess `json:"success"`
	JSON    dnsRecordGetResponseJSON    `json:"-"`
}

// dnsRecordGetResponseJSON contains the JSON metadata for the struct
// [DNSRecordGetResponse]
type dnsRecordGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordGetResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    dnsRecordGetResponseErrorJSON `json:"-"`
}

// dnsRecordGetResponseErrorJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseError]
type dnsRecordGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordGetResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    dnsRecordGetResponseMessageJSON `json:"-"`
}

// dnsRecordGetResponseMessageJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseMessage]
type dnsRecordGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DNSRecordGetResponseResultDNSRecordsARecord],
// [DNSRecordGetResponseResultDNSRecordsAaaaRecord],
// [DNSRecordGetResponseResultDNSRecordsCaaRecord],
// [DNSRecordGetResponseResultDNSRecordsCertRecord],
// [DNSRecordGetResponseResultDNSRecordsCnameRecord],
// [DNSRecordGetResponseResultDNSRecordsDnskeyRecord],
// [DNSRecordGetResponseResultDNSRecordsDsRecord],
// [DNSRecordGetResponseResultDNSRecordsHTTPsRecord],
// [DNSRecordGetResponseResultDNSRecordsLocRecord],
// [DNSRecordGetResponseResultDNSRecordsMxRecord],
// [DNSRecordGetResponseResultDNSRecordsNaptrRecord],
// [DNSRecordGetResponseResultDNSRecordsNsRecord],
// [DNSRecordGetResponseResultDNSRecordsPtrRecord],
// [DNSRecordGetResponseResultDNSRecordsSmimeaRecord],
// [DNSRecordGetResponseResultDNSRecordsSrvRecord],
// [DNSRecordGetResponseResultDNSRecordsSshfpRecord],
// [DNSRecordGetResponseResultDNSRecordsSvcbRecord],
// [DNSRecordGetResponseResultDNSRecordsTlsaRecord],
// [DNSRecordGetResponseResultDNSRecordsTxtRecord] or
// [DNSRecordGetResponseResultDNSRecordsUriRecord].
type DNSRecordGetResponseResult interface {
	implementsDNSRecordGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordGetResponseResult)(nil)).Elem(), "")
}

type DNSRecordGetResponseResultDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseResultDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsARecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsARecord]
type dnsRecordGetResponseResultDNSRecordsARecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsARecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsARecordType string

const (
	DNSRecordGetResponseResultDNSRecordsARecordTypeA DNSRecordGetResponseResultDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsARecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsARecordMeta]
type dnsRecordGetResponseResultDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsARecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsARecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsARecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsARecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsARecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseResultDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsAaaaRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsAaaaRecord]
type dnsRecordGetResponseResultDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsAaaaRecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsAaaaRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsAaaaRecordTypeAaaa DNSRecordGetResponseResultDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                 `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsAaaaRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsAaaaRecordMeta]
type dnsRecordGetResponseResultDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsAaaaRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsAaaaRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordGetResponseResultDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCaaRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsCaaRecord]
type dnsRecordGetResponseResultDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsCaaRecord) implementsDNSRecordGetResponseResult() {}

// Components of a CAA record.
type DNSRecordGetResponseResultDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                `json:"value"`
	JSON  dnsRecordGetResponseResultDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCaaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsCaaRecordData]
type dnsRecordGetResponseResultDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsCaaRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsCaaRecordTypeCaa DNSRecordGetResponseResultDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCaaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsCaaRecordMeta]
type dnsRecordGetResponseResultDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsCaaRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsCaaRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordGetResponseResultDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCertRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsCertRecord]
type dnsRecordGetResponseResultDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsCertRecord) implementsDNSRecordGetResponseResult() {}

// Components of a CERT record.
type DNSRecordGetResponseResultDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                `json:"type"`
	JSON dnsRecordGetResponseResultDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCertRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsCertRecordData]
type dnsRecordGetResponseResultDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsCertRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsCertRecordTypeCert DNSRecordGetResponseResultDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                 `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCertRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsCertRecordMeta]
type dnsRecordGetResponseResultDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsCertRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsCertRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseResultDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCnameRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsCnameRecord]
type dnsRecordGetResponseResultDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsCnameRecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsCnameRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsCnameRecordTypeCname DNSRecordGetResponseResultDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsCnameRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsCnameRecordMeta]
type dnsRecordGetResponseResultDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsCnameRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsCnameRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordGetResponseResultDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsDnskeyRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsDnskeyRecord]
type dnsRecordGetResponseResultDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsDnskeyRecord) implementsDNSRecordGetResponseResult() {}

// Components of a DNSKEY record.
type DNSRecordGetResponseResultDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                   `json:"public_key"`
	JSON      dnsRecordGetResponseResultDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsDnskeyRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsDnskeyRecordData]
type dnsRecordGetResponseResultDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsDnskeyRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsDnskeyRecordTypeDnskey DNSRecordGetResponseResultDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsDnskeyRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsDnskeyRecordMeta]
type dnsRecordGetResponseResultDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordGetResponseResultDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsDsRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsDsRecord]
type dnsRecordGetResponseResultDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsDsRecord) implementsDNSRecordGetResponseResult() {}

// Components of a DS record.
type DNSRecordGetResponseResultDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                              `json:"key_tag"`
	JSON   dnsRecordGetResponseResultDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsDsRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsDsRecordData]
type dnsRecordGetResponseResultDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsDsRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsDsRecordTypeDs DNSRecordGetResponseResultDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsDsRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsDsRecordMeta]
type dnsRecordGetResponseResultDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsDsRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsDsRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordGetResponseResultDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsHTTPsRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsHTTPsRecord]
type dnsRecordGetResponseResultDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsHTTPsRecord) implementsDNSRecordGetResponseResult() {}

// Components of a HTTPS record.
type DNSRecordGetResponseResultDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                  `json:"value"`
	JSON  dnsRecordGetResponseResultDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsHTTPsRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsHTTPsRecordData]
type dnsRecordGetResponseResultDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsHTTPsRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsHTTPsRecordTypeHTTPs DNSRecordGetResponseResultDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsHTTPsRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsHTTPsRecordMeta]
type dnsRecordGetResponseResultDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordGetResponseResultDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsLocRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsLocRecord]
type dnsRecordGetResponseResultDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsLocRecord) implementsDNSRecordGetResponseResult() {}

// Components of a LOC record.
type DNSRecordGetResponseResultDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordGetResponseResultDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordGetResponseResultDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                               `json:"size"`
	JSON dnsRecordGetResponseResultDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsLocRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsLocRecordData]
type dnsRecordGetResponseResultDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordGetResponseResultDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordGetResponseResultDNSRecordsLocRecordDataLatDirectionN DNSRecordGetResponseResultDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordGetResponseResultDNSRecordsLocRecordDataLatDirectionS DNSRecordGetResponseResultDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordGetResponseResultDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordGetResponseResultDNSRecordsLocRecordDataLongDirectionE DNSRecordGetResponseResultDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordGetResponseResultDNSRecordsLocRecordDataLongDirectionW DNSRecordGetResponseResultDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordGetResponseResultDNSRecordsLocRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsLocRecordTypeLoc DNSRecordGetResponseResultDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsLocRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsLocRecordMeta]
type dnsRecordGetResponseResultDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsLocRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsLocRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsMxRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsMxRecord]
type dnsRecordGetResponseResultDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsMxRecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsMxRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsMxRecordTypeMx DNSRecordGetResponseResultDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsMxRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsMxRecordMeta]
type dnsRecordGetResponseResultDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsMxRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsMxRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordGetResponseResultDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsNaptrRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsNaptrRecord]
type dnsRecordGetResponseResultDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsNaptrRecord) implementsDNSRecordGetResponseResult() {}

// Components of a NAPTR record.
type DNSRecordGetResponseResultDNSRecordsNaptrRecordData struct {
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
	Service string                                                  `json:"service"`
	JSON    dnsRecordGetResponseResultDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsNaptrRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsNaptrRecordData]
type dnsRecordGetResponseResultDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsNaptrRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsNaptrRecordTypeNaptr DNSRecordGetResponseResultDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsNaptrRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsNaptrRecordMeta]
type dnsRecordGetResponseResultDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsNaptrRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsNaptrRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsNsRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsNsRecord]
type dnsRecordGetResponseResultDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsNsRecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsNsRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsNsRecordTypeNs DNSRecordGetResponseResultDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsNsRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsNsRecordMeta]
type dnsRecordGetResponseResultDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsNsRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsNsRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsPtrRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsPtrRecord]
type dnsRecordGetResponseResultDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsPtrRecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsPtrRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsPtrRecordTypePtr DNSRecordGetResponseResultDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsPtrRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsPtrRecordMeta]
type dnsRecordGetResponseResultDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsPtrRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsPtrRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordGetResponseResultDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSmimeaRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsSmimeaRecord]
type dnsRecordGetResponseResultDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsSmimeaRecord) implementsDNSRecordGetResponseResult() {}

// Components of a SMIMEA record.
type DNSRecordGetResponseResultDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                  `json:"usage"`
	JSON  dnsRecordGetResponseResultDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSmimeaRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsSmimeaRecordData]
type dnsRecordGetResponseResultDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsSmimeaRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsSmimeaRecordTypeSmimea DNSRecordGetResponseResultDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSmimeaRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsSmimeaRecordMeta]
type dnsRecordGetResponseResultDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordGetResponseResultDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSrvRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsSrvRecord]
type dnsRecordGetResponseResultDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsSrvRecord) implementsDNSRecordGetResponseResult() {}

// Components of a SRV record.
type DNSRecordGetResponseResultDNSRecordsSrvRecordData struct {
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
	Weight float64                                               `json:"weight"`
	JSON   dnsRecordGetResponseResultDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSrvRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsSrvRecordData]
type dnsRecordGetResponseResultDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsSrvRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsSrvRecordTypeSrv DNSRecordGetResponseResultDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSrvRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsSrvRecordMeta]
type dnsRecordGetResponseResultDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsSrvRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsSrvRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordGetResponseResultDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSshfpRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsSshfpRecord]
type dnsRecordGetResponseResultDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsSshfpRecord) implementsDNSRecordGetResponseResult() {}

// Components of a SSHFP record.
type DNSRecordGetResponseResultDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                 `json:"type"`
	JSON dnsRecordGetResponseResultDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSshfpRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsSshfpRecordData]
type dnsRecordGetResponseResultDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsSshfpRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsSshfpRecordTypeSshfp DNSRecordGetResponseResultDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSshfpRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsSshfpRecordMeta]
type dnsRecordGetResponseResultDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsSshfpRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsSshfpRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordGetResponseResultDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSvcbRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsSvcbRecord]
type dnsRecordGetResponseResultDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsSvcbRecord) implementsDNSRecordGetResponseResult() {}

// Components of a SVCB record.
type DNSRecordGetResponseResultDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                 `json:"value"`
	JSON  dnsRecordGetResponseResultDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSvcbRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsSvcbRecordData]
type dnsRecordGetResponseResultDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsSvcbRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsSvcbRecordTypeSvcb DNSRecordGetResponseResultDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                 `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsSvcbRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsSvcbRecordMeta]
type dnsRecordGetResponseResultDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsSvcbRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsSvcbRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordGetResponseResultDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsTlsaRecordJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsTlsaRecord]
type dnsRecordGetResponseResultDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsTlsaRecord) implementsDNSRecordGetResponseResult() {}

// Components of a TLSA record.
type DNSRecordGetResponseResultDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                `json:"usage"`
	JSON  dnsRecordGetResponseResultDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsTlsaRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsTlsaRecordData]
type dnsRecordGetResponseResultDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsTlsaRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsTlsaRecordTypeTlsa DNSRecordGetResponseResultDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                 `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsTlsaRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordGetResponseResultDNSRecordsTlsaRecordMeta]
type dnsRecordGetResponseResultDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsTlsaRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsTlsaRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsTxtRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsTxtRecord]
type dnsRecordGetResponseResultDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsTxtRecord) implementsDNSRecordGetResponseResult() {}

// Record type.
type DNSRecordGetResponseResultDNSRecordsTxtRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsTxtRecordTypeTxt DNSRecordGetResponseResultDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsTxtRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsTxtRecordMeta]
type dnsRecordGetResponseResultDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsTxtRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsTxtRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordGetResponseResultDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordGetResponseResultDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseResultDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseResultDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseResultDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseResultDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsUriRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseResultDNSRecordsUriRecord]
type dnsRecordGetResponseResultDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordGetResponseResultDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseResultDNSRecordsUriRecord) implementsDNSRecordGetResponseResult() {}

// Components of a URI record.
type DNSRecordGetResponseResultDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                               `json:"weight"`
	JSON   dnsRecordGetResponseResultDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsUriRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsUriRecordData]
type dnsRecordGetResponseResultDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseResultDNSRecordsUriRecordType string

const (
	DNSRecordGetResponseResultDNSRecordsUriRecordTypeUri DNSRecordGetResponseResultDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseResultDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordGetResponseResultDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseResultDNSRecordsUriRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseResultDNSRecordsUriRecordMeta]
type dnsRecordGetResponseResultDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseResultDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseResultDNSRecordsUriRecordTTLNumber].
type DNSRecordGetResponseResultDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordGetResponseResultDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseResultDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseResultDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordGetResponseResultDNSRecordsUriRecordTTLNumber1 DNSRecordGetResponseResultDNSRecordsUriRecordTTLNumber = 1
)

// Whether the API call was successful
type DNSRecordGetResponseSuccess bool

const (
	DNSRecordGetResponseSuccessTrue DNSRecordGetResponseSuccess = true
)

type DNSRecordUpdateResponse struct {
	Errors   []DNSRecordUpdateResponseError   `json:"errors"`
	Messages []DNSRecordUpdateResponseMessage `json:"messages"`
	Result   DNSRecordUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DNSRecordUpdateResponseSuccess `json:"success"`
	JSON    dnsRecordUpdateResponseJSON    `json:"-"`
}

// dnsRecordUpdateResponseJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponse]
type dnsRecordUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordUpdateResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dnsRecordUpdateResponseErrorJSON `json:"-"`
}

// dnsRecordUpdateResponseErrorJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseError]
type dnsRecordUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordUpdateResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    dnsRecordUpdateResponseMessageJSON `json:"-"`
}

// dnsRecordUpdateResponseMessageJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseMessage]
type dnsRecordUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [DNSRecordUpdateResponseResultDNSRecordsARecord],
// [DNSRecordUpdateResponseResultDNSRecordsAaaaRecord],
// [DNSRecordUpdateResponseResultDNSRecordsCaaRecord],
// [DNSRecordUpdateResponseResultDNSRecordsCertRecord],
// [DNSRecordUpdateResponseResultDNSRecordsCnameRecord],
// [DNSRecordUpdateResponseResultDNSRecordsDnskeyRecord],
// [DNSRecordUpdateResponseResultDNSRecordsDsRecord],
// [DNSRecordUpdateResponseResultDNSRecordsHTTPsRecord],
// [DNSRecordUpdateResponseResultDNSRecordsLocRecord],
// [DNSRecordUpdateResponseResultDNSRecordsMxRecord],
// [DNSRecordUpdateResponseResultDNSRecordsNaptrRecord],
// [DNSRecordUpdateResponseResultDNSRecordsNsRecord],
// [DNSRecordUpdateResponseResultDNSRecordsPtrRecord],
// [DNSRecordUpdateResponseResultDNSRecordsSmimeaRecord],
// [DNSRecordUpdateResponseResultDNSRecordsSrvRecord],
// [DNSRecordUpdateResponseResultDNSRecordsSshfpRecord],
// [DNSRecordUpdateResponseResultDNSRecordsSvcbRecord],
// [DNSRecordUpdateResponseResultDNSRecordsTlsaRecord],
// [DNSRecordUpdateResponseResultDNSRecordsTxtRecord] or
// [DNSRecordUpdateResponseResultDNSRecordsUriRecord].
type DNSRecordUpdateResponseResult interface {
	implementsDNSRecordUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordUpdateResponseResult)(nil)).Elem(), "")
}

type DNSRecordUpdateResponseResultDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseResultDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                             `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsARecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsARecord]
type dnsRecordUpdateResponseResultDNSRecordsARecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsARecord) implementsDNSRecordUpdateResponseResult() {}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsARecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsARecordTypeA DNSRecordUpdateResponseResultDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                 `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsARecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsARecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsARecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsARecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsARecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsARecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsARecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsAaaaRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsAaaaRecord]
type dnsRecordUpdateResponseResultDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsAaaaRecord) implementsDNSRecordUpdateResponseResult() {
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsAaaaRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTypeAaaa DNSRecordUpdateResponseResultDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                    `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsAaaaRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsAaaaRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordUpdateResponseResultDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCaaRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsCaaRecord]
type dnsRecordUpdateResponseResultDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsCaaRecord) implementsDNSRecordUpdateResponseResult() {}

// Components of a CAA record.
type DNSRecordUpdateResponseResultDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                   `json:"value"`
	JSON  dnsRecordUpdateResponseResultDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCaaRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsCaaRecordData]
type dnsRecordUpdateResponseResultDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsCaaRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsCaaRecordTypeCaa DNSRecordUpdateResponseResultDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCaaRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsCaaRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordUpdateResponseResultDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCertRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsCertRecord]
type dnsRecordUpdateResponseResultDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsCertRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a CERT record.
type DNSRecordUpdateResponseResultDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                   `json:"type"`
	JSON dnsRecordUpdateResponseResultDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCertRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsCertRecordData]
type dnsRecordUpdateResponseResultDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsCertRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsCertRecordTypeCert DNSRecordUpdateResponseResultDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                    `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCertRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsCertRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsCertRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsCertRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCnameRecordJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsCnameRecord]
type dnsRecordUpdateResponseResultDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsCnameRecord) implementsDNSRecordUpdateResponseResult() {
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsCnameRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsCnameRecordTypeCname DNSRecordUpdateResponseResultDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                     `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsCnameRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsCnameRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsDnskeyRecord]
type dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsDnskeyRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a DNSKEY record.
type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                      `json:"public_key"`
	JSON      dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordDataJSON contains the JSON
// metadata for the struct
// [DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordData]
type dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTypeDnskey DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                      `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordMetaJSON contains the JSON
// metadata for the struct
// [DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordUpdateResponseResultDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsDsRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsDsRecord]
type dnsRecordUpdateResponseResultDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsDsRecord) implementsDNSRecordUpdateResponseResult() {}

// Components of a DS record.
type DNSRecordUpdateResponseResultDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                                 `json:"key_tag"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsDsRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsDsRecordData]
type dnsRecordUpdateResponseResultDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsDsRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsDsRecordTypeDs DNSRecordUpdateResponseResultDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsDsRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsDsRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsDsRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsDsRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsHTTPsRecord]
type dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsHTTPsRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a HTTPS record.
type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                     `json:"value"`
	JSON  dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordData]
type dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTypeHTTPs DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                     `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordUpdateResponseResultDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsLocRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsLocRecord]
type dnsRecordUpdateResponseResultDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsLocRecord) implementsDNSRecordUpdateResponseResult() {}

// Components of a LOC record.
type DNSRecordUpdateResponseResultDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                                  `json:"size"`
	JSON dnsRecordUpdateResponseResultDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsLocRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsLocRecordData]
type dnsRecordUpdateResponseResultDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLatDirectionN DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLatDirectionS DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLongDirectionE DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLongDirectionW DNSRecordUpdateResponseResultDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsLocRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsLocRecordTypeLoc DNSRecordUpdateResponseResultDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsLocRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsLocRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsLocRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsLocRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsMxRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsMxRecord]
type dnsRecordUpdateResponseResultDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsMxRecord) implementsDNSRecordUpdateResponseResult() {}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsMxRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsMxRecordTypeMx DNSRecordUpdateResponseResultDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsMxRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsMxRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsMxRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsMxRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordUpdateResponseResultDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsNaptrRecordJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsNaptrRecord]
type dnsRecordUpdateResponseResultDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsNaptrRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a NAPTR record.
type DNSRecordUpdateResponseResultDNSRecordsNaptrRecordData struct {
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
	Service string                                                     `json:"service"`
	JSON    dnsRecordUpdateResponseResultDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsNaptrRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsNaptrRecordData]
type dnsRecordUpdateResponseResultDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsNaptrRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTypeNaptr DNSRecordUpdateResponseResultDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                     `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsNaptrRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsNaptrRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                              `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsNsRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsNsRecord]
type dnsRecordUpdateResponseResultDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsNsRecord) implementsDNSRecordUpdateResponseResult() {}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsNsRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsNsRecordTypeNs DNSRecordUpdateResponseResultDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                  `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsNsRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsNsRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsNsRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsNsRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsPtrRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsPtrRecord]
type dnsRecordUpdateResponseResultDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsPtrRecord) implementsDNSRecordUpdateResponseResult() {}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsPtrRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsPtrRecordTypePtr DNSRecordUpdateResponseResultDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsPtrRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsPtrRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                  `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSmimeaRecord]
type dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsSmimeaRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a SMIMEA record.
type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                     `json:"usage"`
	JSON  dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordDataJSON contains the JSON
// metadata for the struct
// [DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordData]
type dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTypeSmimea DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                      `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordMetaJSON contains the JSON
// metadata for the struct
// [DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordUpdateResponseResultDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSrvRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsSrvRecord]
type dnsRecordUpdateResponseResultDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsSrvRecord) implementsDNSRecordUpdateResponseResult() {}

// Components of a SRV record.
type DNSRecordUpdateResponseResultDNSRecordsSrvRecordData struct {
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
	Weight float64                                                  `json:"weight"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSrvRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSrvRecordData]
type dnsRecordUpdateResponseResultDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsSrvRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsSrvRecordTypeSrv DNSRecordUpdateResponseResultDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSrvRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSrvRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordUpdateResponseResultDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                 `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSshfpRecordJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSshfpRecord]
type dnsRecordUpdateResponseResultDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsSshfpRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a SSHFP record.
type DNSRecordUpdateResponseResultDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                    `json:"type"`
	JSON dnsRecordUpdateResponseResultDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSshfpRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSshfpRecordData]
type dnsRecordUpdateResponseResultDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsSshfpRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTypeSshfp DNSRecordUpdateResponseResultDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                     `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSshfpRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSshfpRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordUpdateResponseResultDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSvcbRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsSvcbRecord]
type dnsRecordUpdateResponseResultDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsSvcbRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a SVCB record.
type DNSRecordUpdateResponseResultDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                    `json:"value"`
	JSON  dnsRecordUpdateResponseResultDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSvcbRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSvcbRecordData]
type dnsRecordUpdateResponseResultDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsSvcbRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTypeSvcb DNSRecordUpdateResponseResultDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                    `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsSvcbRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsSvcbRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordUpdateResponseResultDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsTlsaRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsTlsaRecord]
type dnsRecordUpdateResponseResultDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsTlsaRecord) implementsDNSRecordUpdateResponseResult() {
}

// Components of a TLSA record.
type DNSRecordUpdateResponseResultDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                   `json:"usage"`
	JSON  dnsRecordUpdateResponseResultDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsTlsaRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsTlsaRecordData]
type dnsRecordUpdateResponseResultDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsTlsaRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTypeTlsa DNSRecordUpdateResponseResultDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                    `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsTlsaRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsTlsaRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsTxtRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsTxtRecord]
type dnsRecordUpdateResponseResultDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsTxtRecord) implementsDNSRecordUpdateResponseResult() {}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsTxtRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsTxtRecordTypeTxt DNSRecordUpdateResponseResultDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsTxtRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsTxtRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordUpdateResponseResultDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordUpdateResponseResultDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseResultDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseResultDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseResultDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                               `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseResultDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsUriRecordJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseResultDNSRecordsUriRecord]
type dnsRecordUpdateResponseResultDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordUpdateResponseResultDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseResultDNSRecordsUriRecord) implementsDNSRecordUpdateResponseResult() {}

// Components of a URI record.
type DNSRecordUpdateResponseResultDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                                  `json:"weight"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsUriRecordDataJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsUriRecordData]
type dnsRecordUpdateResponseResultDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseResultDNSRecordsUriRecordType string

const (
	DNSRecordUpdateResponseResultDNSRecordsUriRecordTypeUri DNSRecordUpdateResponseResultDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseResultDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                   `json:"source"`
	JSON   dnsRecordUpdateResponseResultDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseResultDNSRecordsUriRecordMetaJSON contains the JSON
// metadata for the struct [DNSRecordUpdateResponseResultDNSRecordsUriRecordMeta]
type dnsRecordUpdateResponseResultDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseResultDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseResultDNSRecordsUriRecordTTLNumber].
type DNSRecordUpdateResponseResultDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordUpdateResponseResultDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseResultDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseResultDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordUpdateResponseResultDNSRecordsUriRecordTTLNumber1 DNSRecordUpdateResponseResultDNSRecordsUriRecordTTLNumber = 1
)

// Whether the API call was successful
type DNSRecordUpdateResponseSuccess bool

const (
	DNSRecordUpdateResponseSuccessTrue DNSRecordUpdateResponseSuccess = true
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

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponse struct {
	Errors   []DNSRecordDNSRecordsForAZoneNewDNSRecordResponseError   `json:"errors"`
	Messages []DNSRecordDNSRecordsForAZoneNewDNSRecordResponseMessage `json:"messages"`
	Result   DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DNSRecordDNSRecordsForAZoneNewDNSRecordResponseSuccess `json:"success"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseJSON    `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseJSON contains the JSON metadata
// for the struct [DNSRecordDNSRecordsForAZoneNewDNSRecordResponse]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseErrorJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseErrorJSON contains the JSON
// metadata for the struct [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseError]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseMessageJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseMessageJSON contains the JSON
// metadata for the struct [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseMessage]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecord],
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecord] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecord].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult interface {
	implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult)(nil)).Elem(), "")
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTypeA DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                         `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsARecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTypeAaaa DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a CAA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                                           `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTypeCaa DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a CERT record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                                           `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTypeCert DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTypeCname DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a DNSKEY record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                                              `json:"public_key"`
	JSON      dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTypeDnskey DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                              `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a DS record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                                                         `json:"key_tag"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTypeDs DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a HTTPS record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                             `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTypeHTTPs DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a LOC record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                                                          `json:"size"`
	JSON dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLatDirectionN DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLatDirectionS DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLongDirectionE DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLongDirectionW DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTypeLoc DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTypeMx DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a NAPTR record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordData struct {
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
	Service string                                                                             `json:"service"`
	JSON    dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTypeNaptr DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTypeNs DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                          `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTypePtr DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a SMIMEA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                             `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTypeSmimea DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                              `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a SRV record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordData struct {
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
	Weight float64                                                                          `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTypeSrv DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a SSHFP record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                                            `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTypeSshfp DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a SVCB record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                            `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTypeSvcb DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a TLSA record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                           `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTypeTlsa DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTypeTxt DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecord]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecord) implementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResult() {
}

// Components of a URI record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                                                          `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordData]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordType string

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTypeUri DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMeta]
type dnsRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTLNumber1 DNSRecordDNSRecordsForAZoneNewDNSRecordResponseResultDNSRecordsUriRecordTTLNumber = 1
)

// Whether the API call was successful
type DNSRecordDNSRecordsForAZoneNewDNSRecordResponseSuccess bool

const (
	DNSRecordDNSRecordsForAZoneNewDNSRecordResponseSuccessTrue DNSRecordDNSRecordsForAZoneNewDNSRecordResponseSuccess = true
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponse struct {
	Errors     []DNSRecordDNSRecordsForAZoneListDNSRecordsResponseError    `json:"errors"`
	Messages   []DNSRecordDNSRecordsForAZoneListDNSRecordsResponseMessage  `json:"messages"`
	Result     []DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult   `json:"result"`
	ResultInfo DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DNSRecordDNSRecordsForAZoneListDNSRecordsResponseSuccess `json:"success"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseJSON    `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseJSON contains the JSON metadata
// for the struct [DNSRecordDNSRecordsForAZoneListDNSRecordsResponse]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseErrorJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseErrorJSON contains the JSON
// metadata for the struct [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseError]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseMessageJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseMessageJSON contains the JSON
// metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseMessage]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecord],
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecord] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecord].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult interface {
	implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult)(nil)).Elem(), "")
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTypeA DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                           `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsARecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTypeAaaa DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                              `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsAaaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecord struct {
	// Components of a CAA record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a CAA record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                                                             `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTypeCaa DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCaaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a CERT record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                                                             `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTypeCert DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                              `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMeta `json:"meta"`
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
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTypeCname DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                               `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsCnameRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a DNSKEY record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                                                `json:"public_key"`
	JSON      dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTypeDnskey DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                                `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDnskeyRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecord struct {
	// Components of a DS record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a DS record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                                                           `json:"key_tag"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTypeDs DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsDsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a HTTPS record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                               `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTypeHTTPs DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                               `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsHTTPsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecord struct {
	// Components of a LOC record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a LOC record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                                                                            `json:"size"`
	JSON dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLatDirection string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLatDirectionN DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLatDirection = "N"
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLatDirectionS DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLongDirection string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLongDirectionE DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLongDirection = "E"
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLongDirectionW DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTypeLoc DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsLocRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTypeMx DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsMxRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a NAPTR record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordData struct {
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
	Service string                                                                               `json:"service"`
	JSON    dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTypeNaptr DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                               `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNaptrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTypeNs DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                            `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsNsRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTypePtr DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsPtrRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a SMIMEA record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                               `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTypeSmimea DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                                `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecord struct {
	// Components of a SRV record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a SRV record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordData struct {
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
	Weight float64                                                                            `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordDataJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTypeSrv DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSrvRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a SSHFP record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                                                              `json:"type"`
	JSON dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTypeSshfp DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                               `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSshfpRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecord struct {
	// Components of a SVCB record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a SVCB record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                                                              `json:"value"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTypeSvcb DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                              `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsSvcbRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecord struct {
	// Components of a TLSA record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a TLSA record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                                                             `json:"usage"`
	JSON  dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTypeTlsa DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                              `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTlsaRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTypeTxt DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsTxtRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecord struct {
	// Components of a URI record.
	Data DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordType `json:"type,required"`
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
	Meta DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecord]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordJSON struct {
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

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecord) implementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResult() {
}

// Components of a URI record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                                                            `json:"weight"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordDataJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordDataJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordData]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordType string

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTypeUri DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                                             `json:"source"`
	JSON   dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMetaJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMetaJSON
// contains the JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMeta]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTLNumber].
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTL interface {
	ImplementsDNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTLNumber float64

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTLNumber1 DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultDNSRecordsUriRecordTTLNumber = 1
)

type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfoJSON `json:"-"`
}

// dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfoJSON contains the
// JSON metadata for the struct
// [DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfo]
type dnsRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordDNSRecordsForAZoneListDNSRecordsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordDNSRecordsForAZoneListDNSRecordsResponseSuccess bool

const (
	DNSRecordDNSRecordsForAZoneListDNSRecordsResponseSuccessTrue DNSRecordDNSRecordsForAZoneListDNSRecordsResponseSuccess = true
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
