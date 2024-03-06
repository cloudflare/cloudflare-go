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
func (r *DNSRecordService) New(ctx context.Context, params DNSRecordNewParams, opts ...option.RequestOption) (res *DNSRecordNewResponse, err error) {
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
func (r *DNSRecordService) Update(ctx context.Context, dnsRecordID string, params DNSRecordUpdateParams, opts ...option.RequestOption) (res *DNSRecordUpdateResponse, err error) {
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
func (r *DNSRecordService) List(ctx context.Context, params DNSRecordListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[DNSRecordListResponse], err error) {
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
func (r *DNSRecordService) ListAutoPaging(ctx context.Context, params DNSRecordListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[DNSRecordListResponse] {
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
func (r *DNSRecordService) Edit(ctx context.Context, dnsRecordID string, params DNSRecordEditParams, opts ...option.RequestOption) (res *DNSRecordEditResponse, err error) {
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
func (r *DNSRecordService) Get(ctx context.Context, dnsRecordID string, query DNSRecordGetParams, opts ...option.RequestOption) (res *DNSRecordGetResponse, err error) {
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

// Union satisfied by [DNSRecordNewResponseDNSRecordsARecord],
// [DNSRecordNewResponseDNSRecordsAAAARecord],
// [DNSRecordNewResponseDNSRecordsCAARecord],
// [DNSRecordNewResponseDNSRecordsCertRecord],
// [DNSRecordNewResponseDNSRecordsCNAMERecord],
// [DNSRecordNewResponseDNSRecordsDNSKEYRecord],
// [DNSRecordNewResponseDNSRecordsDSRecord],
// [DNSRecordNewResponseDNSRecordsHTTPSRecord],
// [DNSRecordNewResponseDNSRecordsLOCRecord],
// [DNSRecordNewResponseDNSRecordsMXRecord],
// [DNSRecordNewResponseDNSRecordsNAPTRRecord],
// [DNSRecordNewResponseDNSRecordsNSRecord],
// [DNSRecordNewResponseDNSRecordsPTRRecord],
// [DNSRecordNewResponseDNSRecordsSmimeaRecord],
// [DNSRecordNewResponseDNSRecordsSRVRecord],
// [DNSRecordNewResponseDNSRecordsSSHFPRecord],
// [DNSRecordNewResponseDNSRecordsSVCBRecord],
// [DNSRecordNewResponseDNSRecordsTLSARecord],
// [DNSRecordNewResponseDNSRecordsTXTRecord] or
// [DNSRecordNewResponseDNSRecordsURIRecord].
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

type DNSRecordNewResponseDNSRecordsAAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsAAAARecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsAAAARecordMeta `json:"meta"`
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
	TTL DNSRecordNewResponseDNSRecordsAAAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsAAAARecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsAAAARecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsAAAARecord]
type dnsRecordNewResponseDNSRecordsAAAARecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsAAAARecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsAAAARecordType string

const (
	DNSRecordNewResponseDNSRecordsAAAARecordTypeAAAA DNSRecordNewResponseDNSRecordsAAAARecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsAAAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsAAAARecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsAAAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsAAAARecordMeta]
type dnsRecordNewResponseDNSRecordsAAAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsAAAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsAAAARecordTTLNumber].
type DNSRecordNewResponseDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsAAAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsAAAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsAAAARecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsAAAARecordTTLNumber1 DNSRecordNewResponseDNSRecordsAAAARecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsCAARecord struct {
	// Components of a CAA record.
	Data DNSRecordNewResponseDNSRecordsCAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsCAARecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsCAARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsCAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsCAARecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCAARecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsCAARecord]
type dnsRecordNewResponseDNSRecordsCAARecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsCAARecord) implementsDNSRecordNewResponse() {}

// Components of a CAA record.
type DNSRecordNewResponseDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                          `json:"value"`
	JSON  dnsRecordNewResponseDNSRecordsCAARecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCAARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCAARecordData]
type dnsRecordNewResponseDNSRecordsCAARecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsCAARecordType string

const (
	DNSRecordNewResponseDNSRecordsCAARecordTypeCAA DNSRecordNewResponseDNSRecordsCAARecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsCAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsCAARecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCAARecordMeta]
type dnsRecordNewResponseDNSRecordsCAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsCAARecordTTLNumber].
type DNSRecordNewResponseDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsCAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsCAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsCAARecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsCAARecordTTLNumber1 DNSRecordNewResponseDNSRecordsCAARecordTTLNumber = 1
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

type DNSRecordNewResponseDNSRecordsCNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsCNAMERecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsCNAMERecordMeta `json:"meta"`
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
	TTL DNSRecordNewResponseDNSRecordsCNAMERecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsCNAMERecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCNAMERecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsCNAMERecord]
type dnsRecordNewResponseDNSRecordsCNAMERecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsCNAMERecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsCNAMERecordType string

const (
	DNSRecordNewResponseDNSRecordsCNAMERecordTypeCNAME DNSRecordNewResponseDNSRecordsCNAMERecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsCNAMERecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsCNAMERecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsCNAMERecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsCNAMERecordMeta]
type dnsRecordNewResponseDNSRecordsCNAMERecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsCNAMERecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsCNAMERecordTTLNumber].
type DNSRecordNewResponseDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsCNAMERecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsCNAMERecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsCNAMERecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsCNAMERecordTTLNumber1 DNSRecordNewResponseDNSRecordsCNAMERecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsDNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordNewResponseDNSRecordsDNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsDNSKEYRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsDNSKEYRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsDNSKEYRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsDNSKEYRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDNSKEYRecordJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsDNSKEYRecord]
type dnsRecordNewResponseDNSRecordsDNSKEYRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsDNSKEYRecord) implementsDNSRecordNewResponse() {}

// Components of a DNSKEY record.
type DNSRecordNewResponseDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                             `json:"public_key"`
	JSON      dnsRecordNewResponseDNSRecordsDNSKEYRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDNSKEYRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordNewResponseDNSRecordsDNSKEYRecordData]
type dnsRecordNewResponseDNSRecordsDNSKEYRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsDNSKEYRecordType string

const (
	DNSRecordNewResponseDNSRecordsDNSKEYRecordTypeDNSKEY DNSRecordNewResponseDNSRecordsDNSKEYRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsDNSKEYRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsDNSKEYRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDNSKEYRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordNewResponseDNSRecordsDNSKEYRecordMeta]
type dnsRecordNewResponseDNSRecordsDNSKEYRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDNSKEYRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsDNSKEYRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsDNSKEYRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsDNSKEYRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsDNSKEYRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsDNSKEYRecordTTLNumber1 DNSRecordNewResponseDNSRecordsDNSKEYRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsDSRecord struct {
	// Components of a DS record.
	Data DNSRecordNewResponseDNSRecordsDSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsDSRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsDSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsDSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsDSRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDSRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsDSRecord]
type dnsRecordNewResponseDNSRecordsDSRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsDSRecord) implementsDNSRecordNewResponse() {}

// Components of a DS record.
type DNSRecordNewResponseDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                        `json:"key_tag"`
	JSON   dnsRecordNewResponseDNSRecordsDSRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsDSRecordData]
type dnsRecordNewResponseDNSRecordsDSRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsDSRecordType string

const (
	DNSRecordNewResponseDNSRecordsDSRecordTypeDS DNSRecordNewResponseDNSRecordsDSRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsDSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsDSRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsDSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsDSRecordMeta]
type dnsRecordNewResponseDNSRecordsDSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsDSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsDSRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsDSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsDSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsDSRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsDSRecordTTLNumber1 DNSRecordNewResponseDNSRecordsDSRecordTTLNumber = 1
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

type DNSRecordNewResponseDNSRecordsLOCRecord struct {
	// Components of a LOC record.
	Data DNSRecordNewResponseDNSRecordsLOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsLOCRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsLOCRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsLOCRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsLOCRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsLOCRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsLOCRecord]
type dnsRecordNewResponseDNSRecordsLOCRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsLOCRecord) implementsDNSRecordNewResponse() {}

// Components of a LOC record.
type DNSRecordNewResponseDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordNewResponseDNSRecordsLOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordNewResponseDNSRecordsLOCRecordDataLongDirection `json:"long_direction"`
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
	JSON dnsRecordNewResponseDNSRecordsLOCRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsLOCRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsLOCRecordData]
type dnsRecordNewResponseDNSRecordsLOCRecordDataJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsLOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordNewResponseDNSRecordsLOCRecordDataLatDirection string

const (
	DNSRecordNewResponseDNSRecordsLOCRecordDataLatDirectionN DNSRecordNewResponseDNSRecordsLOCRecordDataLatDirection = "N"
	DNSRecordNewResponseDNSRecordsLOCRecordDataLatDirectionS DNSRecordNewResponseDNSRecordsLOCRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordNewResponseDNSRecordsLOCRecordDataLongDirection string

const (
	DNSRecordNewResponseDNSRecordsLOCRecordDataLongDirectionE DNSRecordNewResponseDNSRecordsLOCRecordDataLongDirection = "E"
	DNSRecordNewResponseDNSRecordsLOCRecordDataLongDirectionW DNSRecordNewResponseDNSRecordsLOCRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordNewResponseDNSRecordsLOCRecordType string

const (
	DNSRecordNewResponseDNSRecordsLOCRecordTypeLOC DNSRecordNewResponseDNSRecordsLOCRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsLOCRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsLOCRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsLOCRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsLOCRecordMeta]
type dnsRecordNewResponseDNSRecordsLOCRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsLOCRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsLOCRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsLOCRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsLOCRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsLOCRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsLOCRecordTTLNumber1 DNSRecordNewResponseDNSRecordsLOCRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsMXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsMXRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsMXRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsMXRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsMXRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsMXRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsMXRecord]
type dnsRecordNewResponseDNSRecordsMXRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsMXRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsMXRecordType string

const (
	DNSRecordNewResponseDNSRecordsMXRecordTypeMX DNSRecordNewResponseDNSRecordsMXRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsMXRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsMXRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsMXRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsMXRecordMeta]
type dnsRecordNewResponseDNSRecordsMXRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsMXRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsMXRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsMXRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsMXRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsMXRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsMXRecordTTLNumber1 DNSRecordNewResponseDNSRecordsMXRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsNAPTRRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordNewResponseDNSRecordsNAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsNAPTRRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsNAPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsNAPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsNAPTRRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNAPTRRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsNAPTRRecord]
type dnsRecordNewResponseDNSRecordsNAPTRRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsNAPTRRecord) implementsDNSRecordNewResponse() {}

// Components of a NAPTR record.
type DNSRecordNewResponseDNSRecordsNAPTRRecordData struct {
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
	JSON    dnsRecordNewResponseDNSRecordsNAPTRRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNAPTRRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsNAPTRRecordData]
type dnsRecordNewResponseDNSRecordsNAPTRRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsNAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsNAPTRRecordType string

const (
	DNSRecordNewResponseDNSRecordsNAPTRRecordTypeNAPTR DNSRecordNewResponseDNSRecordsNAPTRRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsNAPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsNAPTRRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNAPTRRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsNAPTRRecordMeta]
type dnsRecordNewResponseDNSRecordsNAPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsNAPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsNAPTRRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsNAPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsNAPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsNAPTRRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsNAPTRRecordTTLNumber1 DNSRecordNewResponseDNSRecordsNAPTRRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsNSRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsNSRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsNSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsNSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsNSRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNSRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsNSRecord]
type dnsRecordNewResponseDNSRecordsNSRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsNSRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsNSRecordType string

const (
	DNSRecordNewResponseDNSRecordsNSRecordTypeNS DNSRecordNewResponseDNSRecordsNSRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsNSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsNSRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsNSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsNSRecordMeta]
type dnsRecordNewResponseDNSRecordsNSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsNSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsNSRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsNSRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsNSRecordTTLNumber1 DNSRecordNewResponseDNSRecordsNSRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsPTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsPTRRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsPTRRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsPTRRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsPTRRecord]
type dnsRecordNewResponseDNSRecordsPTRRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsPTRRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsPTRRecordType string

const (
	DNSRecordNewResponseDNSRecordsPTRRecordTypePTR DNSRecordNewResponseDNSRecordsPTRRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsPTRRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsPTRRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsPTRRecordMeta]
type dnsRecordNewResponseDNSRecordsPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsPTRRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsPTRRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsPTRRecordTTLNumber1 DNSRecordNewResponseDNSRecordsPTRRecordTTLNumber = 1
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

type DNSRecordNewResponseDNSRecordsSRVRecord struct {
	// Components of a SRV record.
	Data DNSRecordNewResponseDNSRecordsSRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSRVRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSRVRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSRVRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSRVRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSRVRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsSRVRecord]
type dnsRecordNewResponseDNSRecordsSRVRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSRVRecord) implementsDNSRecordNewResponse() {}

// Components of a SRV record.
type DNSRecordNewResponseDNSRecordsSRVRecordData struct {
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
	JSON   dnsRecordNewResponseDNSRecordsSRVRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSRVRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSRVRecordData]
type dnsRecordNewResponseDNSRecordsSRVRecordDataJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSRVRecordType string

const (
	DNSRecordNewResponseDNSRecordsSRVRecordTypeSRV DNSRecordNewResponseDNSRecordsSRVRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSRVRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSRVRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSRVRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSRVRecordMeta]
type dnsRecordNewResponseDNSRecordsSRVRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSRVRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSRVRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSRVRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSRVRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSRVRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSRVRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSRVRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsSSHFPRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordNewResponseDNSRecordsSSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSSHFPRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSSHFPRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSSHFPRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSSHFPRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSSHFPRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsSSHFPRecord]
type dnsRecordNewResponseDNSRecordsSSHFPRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSSHFPRecord) implementsDNSRecordNewResponse() {}

// Components of a SSHFP record.
type DNSRecordNewResponseDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                           `json:"type"`
	JSON dnsRecordNewResponseDNSRecordsSSHFPRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSSHFPRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSSHFPRecordData]
type dnsRecordNewResponseDNSRecordsSSHFPRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSSHFPRecordType string

const (
	DNSRecordNewResponseDNSRecordsSSHFPRecordTypeSSHFP DNSRecordNewResponseDNSRecordsSSHFPRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSSHFPRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSSHFPRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSSHFPRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSSHFPRecordMeta]
type dnsRecordNewResponseDNSRecordsSSHFPRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSSHFPRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSSHFPRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSSHFPRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSSHFPRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSSHFPRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSSHFPRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSSHFPRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsSVCBRecord struct {
	// Components of a SVCB record.
	Data DNSRecordNewResponseDNSRecordsSVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsSVCBRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsSVCBRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsSVCBRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsSVCBRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSVCBRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsSVCBRecord]
type dnsRecordNewResponseDNSRecordsSVCBRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsSVCBRecord) implementsDNSRecordNewResponse() {}

// Components of a SVCB record.
type DNSRecordNewResponseDNSRecordsSVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                           `json:"value"`
	JSON  dnsRecordNewResponseDNSRecordsSVCBRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSVCBRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSVCBRecordData]
type dnsRecordNewResponseDNSRecordsSVCBRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsSVCBRecordType string

const (
	DNSRecordNewResponseDNSRecordsSVCBRecordTypeSVCB DNSRecordNewResponseDNSRecordsSVCBRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsSVCBRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsSVCBRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsSVCBRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsSVCBRecordMeta]
type dnsRecordNewResponseDNSRecordsSVCBRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsSVCBRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsSVCBRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsSVCBRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsSVCBRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsSVCBRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsSVCBRecordTTLNumber1 DNSRecordNewResponseDNSRecordsSVCBRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsTLSARecord struct {
	// Components of a TLSA record.
	Data DNSRecordNewResponseDNSRecordsTLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsTLSARecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsTLSARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsTLSARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsTLSARecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTLSARecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsTLSARecord]
type dnsRecordNewResponseDNSRecordsTLSARecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsTLSARecord) implementsDNSRecordNewResponse() {}

// Components of a TLSA record.
type DNSRecordNewResponseDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                          `json:"usage"`
	JSON  dnsRecordNewResponseDNSRecordsTLSARecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTLSARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsTLSARecordData]
type dnsRecordNewResponseDNSRecordsTLSARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsTLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsTLSARecordType string

const (
	DNSRecordNewResponseDNSRecordsTLSARecordTypeTLSA DNSRecordNewResponseDNSRecordsTLSARecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsTLSARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsTLSARecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTLSARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsTLSARecordMeta]
type dnsRecordNewResponseDNSRecordsTLSARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsTLSARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsTLSARecordTTLNumber].
type DNSRecordNewResponseDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsTLSARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsTLSARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsTLSARecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsTLSARecordTTLNumber1 DNSRecordNewResponseDNSRecordsTLSARecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsTXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsTXTRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsTXTRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsTXTRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsTXTRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTXTRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsTXTRecord]
type dnsRecordNewResponseDNSRecordsTXTRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsTXTRecord) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseDNSRecordsTXTRecordType string

const (
	DNSRecordNewResponseDNSRecordsTXTRecordTypeTXT DNSRecordNewResponseDNSRecordsTXTRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsTXTRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsTXTRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsTXTRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsTXTRecordMeta]
type dnsRecordNewResponseDNSRecordsTXTRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsTXTRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsTXTRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsTXTRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsTXTRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsTXTRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsTXTRecordTTLNumber1 DNSRecordNewResponseDNSRecordsTXTRecordTTLNumber = 1
)

type DNSRecordNewResponseDNSRecordsURIRecord struct {
	// Components of a URI record.
	Data DNSRecordNewResponseDNSRecordsURIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordNewResponseDNSRecordsURIRecordType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSRecordsURIRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSRecordsURIRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSRecordsURIRecordJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsURIRecordJSON contains the JSON metadata for the
// struct [DNSRecordNewResponseDNSRecordsURIRecord]
type dnsRecordNewResponseDNSRecordsURIRecordJSON struct {
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

func (r *DNSRecordNewResponseDNSRecordsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordNewResponseDNSRecordsURIRecord) implementsDNSRecordNewResponse() {}

// Components of a URI record.
type DNSRecordNewResponseDNSRecordsURIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                         `json:"weight"`
	JSON   dnsRecordNewResponseDNSRecordsURIRecordDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsURIRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsURIRecordData]
type dnsRecordNewResponseDNSRecordsURIRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsURIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordNewResponseDNSRecordsURIRecordType string

const (
	DNSRecordNewResponseDNSRecordsURIRecordTypeURI DNSRecordNewResponseDNSRecordsURIRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSRecordsURIRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordNewResponseDNSRecordsURIRecordMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSRecordsURIRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordNewResponseDNSRecordsURIRecordMeta]
type dnsRecordNewResponseDNSRecordsURIRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSRecordsURIRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordNewResponseDNSRecordsURIRecordTTLNumber].
type DNSRecordNewResponseDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordNewResponseDNSRecordsURIRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSRecordsURIRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordNewResponseDNSRecordsURIRecordTTLNumber float64

const (
	DNSRecordNewResponseDNSRecordsURIRecordTTLNumber1 DNSRecordNewResponseDNSRecordsURIRecordTTLNumber = 1
)

// Union satisfied by [DNSRecordUpdateResponseDNSRecordsARecord],
// [DNSRecordUpdateResponseDNSRecordsAAAARecord],
// [DNSRecordUpdateResponseDNSRecordsCAARecord],
// [DNSRecordUpdateResponseDNSRecordsCertRecord],
// [DNSRecordUpdateResponseDNSRecordsCNAMERecord],
// [DNSRecordUpdateResponseDNSRecordsDNSKEYRecord],
// [DNSRecordUpdateResponseDNSRecordsDSRecord],
// [DNSRecordUpdateResponseDNSRecordsHTTPSRecord],
// [DNSRecordUpdateResponseDNSRecordsLOCRecord],
// [DNSRecordUpdateResponseDNSRecordsMXRecord],
// [DNSRecordUpdateResponseDNSRecordsNAPTRRecord],
// [DNSRecordUpdateResponseDNSRecordsNSRecord],
// [DNSRecordUpdateResponseDNSRecordsPTRRecord],
// [DNSRecordUpdateResponseDNSRecordsSmimeaRecord],
// [DNSRecordUpdateResponseDNSRecordsSRVRecord],
// [DNSRecordUpdateResponseDNSRecordsSSHFPRecord],
// [DNSRecordUpdateResponseDNSRecordsSVCBRecord],
// [DNSRecordUpdateResponseDNSRecordsTLSARecord],
// [DNSRecordUpdateResponseDNSRecordsTXTRecord] or
// [DNSRecordUpdateResponseDNSRecordsURIRecord].
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

type DNSRecordUpdateResponseDNSRecordsAAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsAAAARecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsAAAARecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseDNSRecordsAAAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsAAAARecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsAAAARecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsAAAARecord]
type dnsRecordUpdateResponseDNSRecordsAAAARecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsAAAARecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsAAAARecordType string

const (
	DNSRecordUpdateResponseDNSRecordsAAAARecordTypeAAAA DNSRecordUpdateResponseDNSRecordsAAAARecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsAAAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsAAAARecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsAAAARecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsAAAARecordMeta]
type dnsRecordUpdateResponseDNSRecordsAAAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsAAAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsAAAARecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsAAAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsAAAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsAAAARecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsAAAARecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsAAAARecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsCAARecord struct {
	// Components of a CAA record.
	Data DNSRecordUpdateResponseDNSRecordsCAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsCAARecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsCAARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsCAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsCAARecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCAARecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsCAARecord]
type dnsRecordUpdateResponseDNSRecordsCAARecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsCAARecord) implementsDNSRecordUpdateResponse() {}

// Components of a CAA record.
type DNSRecordUpdateResponseDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                             `json:"value"`
	JSON  dnsRecordUpdateResponseDNSRecordsCAARecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCAARecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCAARecordData]
type dnsRecordUpdateResponseDNSRecordsCAARecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsCAARecordType string

const (
	DNSRecordUpdateResponseDNSRecordsCAARecordTypeCAA DNSRecordUpdateResponseDNSRecordsCAARecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsCAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsCAARecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCAARecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCAARecordMeta]
type dnsRecordUpdateResponseDNSRecordsCAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsCAARecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsCAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsCAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsCAARecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsCAARecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsCAARecordTTLNumber = 1
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

type DNSRecordUpdateResponseDNSRecordsCNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsCNAMERecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsCNAMERecordMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseDNSRecordsCNAMERecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsCNAMERecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCNAMERecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsCNAMERecord]
type dnsRecordUpdateResponseDNSRecordsCNAMERecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsCNAMERecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsCNAMERecordType string

const (
	DNSRecordUpdateResponseDNSRecordsCNAMERecordTypeCNAME DNSRecordUpdateResponseDNSRecordsCNAMERecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsCNAMERecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsCNAMERecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsCNAMERecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsCNAMERecordMeta]
type dnsRecordUpdateResponseDNSRecordsCNAMERecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsCNAMERecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsCNAMERecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsCNAMERecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsCNAMERecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsCNAMERecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsCNAMERecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsCNAMERecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsDNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordUpdateResponseDNSRecordsDNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsDNSKEYRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsDNSKEYRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsDNSKEYRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDNSKEYRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsDNSKEYRecord]
type dnsRecordUpdateResponseDNSRecordsDNSKEYRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsDNSKEYRecord) implementsDNSRecordUpdateResponse() {}

// Components of a DNSKEY record.
type DNSRecordUpdateResponseDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                                `json:"public_key"`
	JSON      dnsRecordUpdateResponseDNSRecordsDNSKEYRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDNSKEYRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsDNSKEYRecordData]
type dnsRecordUpdateResponseDNSRecordsDNSKEYRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsDNSKEYRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTypeDNSKEY DNSRecordUpdateResponseDNSRecordsDNSKEYRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsDNSKEYRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                                `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsDNSKEYRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDNSKEYRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsDNSKEYRecordMeta]
type dnsRecordUpdateResponseDNSRecordsDNSKEYRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDNSKEYRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsDNSKEYRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsDSRecord struct {
	// Components of a DS record.
	Data DNSRecordUpdateResponseDNSRecordsDSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsDSRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsDSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsDSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsDSRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDSRecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsDSRecord]
type dnsRecordUpdateResponseDNSRecordsDSRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsDSRecord) implementsDNSRecordUpdateResponse() {}

// Components of a DS record.
type DNSRecordUpdateResponseDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                           `json:"key_tag"`
	JSON   dnsRecordUpdateResponseDNSRecordsDSRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsDSRecordData]
type dnsRecordUpdateResponseDNSRecordsDSRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsDSRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsDSRecordTypeDS DNSRecordUpdateResponseDNSRecordsDSRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsDSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsDSRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsDSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsDSRecordMeta]
type dnsRecordUpdateResponseDNSRecordsDSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsDSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsDSRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsDSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsDSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsDSRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsDSRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsDSRecordTTLNumber = 1
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

type DNSRecordUpdateResponseDNSRecordsLOCRecord struct {
	// Components of a LOC record.
	Data DNSRecordUpdateResponseDNSRecordsLOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsLOCRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsLOCRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsLOCRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsLOCRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsLOCRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsLOCRecord]
type dnsRecordUpdateResponseDNSRecordsLOCRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsLOCRecord) implementsDNSRecordUpdateResponse() {}

// Components of a LOC record.
type DNSRecordUpdateResponseDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordUpdateResponseDNSRecordsLOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordUpdateResponseDNSRecordsLOCRecordDataLongDirection `json:"long_direction"`
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
	JSON dnsRecordUpdateResponseDNSRecordsLOCRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsLOCRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsLOCRecordData]
type dnsRecordUpdateResponseDNSRecordsLOCRecordDataJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsLOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordUpdateResponseDNSRecordsLOCRecordDataLatDirection string

const (
	DNSRecordUpdateResponseDNSRecordsLOCRecordDataLatDirectionN DNSRecordUpdateResponseDNSRecordsLOCRecordDataLatDirection = "N"
	DNSRecordUpdateResponseDNSRecordsLOCRecordDataLatDirectionS DNSRecordUpdateResponseDNSRecordsLOCRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordUpdateResponseDNSRecordsLOCRecordDataLongDirection string

const (
	DNSRecordUpdateResponseDNSRecordsLOCRecordDataLongDirectionE DNSRecordUpdateResponseDNSRecordsLOCRecordDataLongDirection = "E"
	DNSRecordUpdateResponseDNSRecordsLOCRecordDataLongDirectionW DNSRecordUpdateResponseDNSRecordsLOCRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordUpdateResponseDNSRecordsLOCRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsLOCRecordTypeLOC DNSRecordUpdateResponseDNSRecordsLOCRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsLOCRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsLOCRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsLOCRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsLOCRecordMeta]
type dnsRecordUpdateResponseDNSRecordsLOCRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsLOCRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsLOCRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsLOCRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsLOCRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsLOCRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsLOCRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsLOCRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsMXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsMXRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsMXRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsMXRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsMXRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsMXRecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsMXRecord]
type dnsRecordUpdateResponseDNSRecordsMXRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsMXRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsMXRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsMXRecordTypeMX DNSRecordUpdateResponseDNSRecordsMXRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsMXRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsMXRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsMXRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsMXRecordMeta]
type dnsRecordUpdateResponseDNSRecordsMXRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsMXRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsMXRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsMXRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsMXRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsMXRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsMXRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsMXRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsNAPTRRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordUpdateResponseDNSRecordsNAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsNAPTRRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsNAPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsNAPTRRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNAPTRRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsNAPTRRecord]
type dnsRecordUpdateResponseDNSRecordsNAPTRRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsNAPTRRecord) implementsDNSRecordUpdateResponse() {}

// Components of a NAPTR record.
type DNSRecordUpdateResponseDNSRecordsNAPTRRecordData struct {
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
	JSON    dnsRecordUpdateResponseDNSRecordsNAPTRRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNAPTRRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsNAPTRRecordData]
type dnsRecordUpdateResponseDNSRecordsNAPTRRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsNAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsNAPTRRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsNAPTRRecordTypeNAPTR DNSRecordUpdateResponseDNSRecordsNAPTRRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsNAPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsNAPTRRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNAPTRRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsNAPTRRecordMeta]
type dnsRecordUpdateResponseDNSRecordsNAPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsNAPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsNAPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsNAPTRRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsNSRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsNSRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsNSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsNSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsNSRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNSRecordJSON contains the JSON metadata for the
// struct [DNSRecordUpdateResponseDNSRecordsNSRecord]
type dnsRecordUpdateResponseDNSRecordsNSRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsNSRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsNSRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsNSRecordTypeNS DNSRecordUpdateResponseDNSRecordsNSRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsNSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsNSRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsNSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsNSRecordMeta]
type dnsRecordUpdateResponseDNSRecordsNSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsNSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsNSRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsNSRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsNSRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsNSRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsPTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsPTRRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsPTRRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsPTRRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsPTRRecord]
type dnsRecordUpdateResponseDNSRecordsPTRRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsPTRRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsPTRRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsPTRRecordTypePTR DNSRecordUpdateResponseDNSRecordsPTRRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsPTRRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsPTRRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsPTRRecordMeta]
type dnsRecordUpdateResponseDNSRecordsPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsPTRRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsPTRRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsPTRRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsPTRRecordTTLNumber = 1
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

type DNSRecordUpdateResponseDNSRecordsSRVRecord struct {
	// Components of a SRV record.
	Data DNSRecordUpdateResponseDNSRecordsSRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSRVRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSRVRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSRVRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSRVRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSRVRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSRVRecord]
type dnsRecordUpdateResponseDNSRecordsSRVRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSRVRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SRV record.
type DNSRecordUpdateResponseDNSRecordsSRVRecordData struct {
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
	JSON   dnsRecordUpdateResponseDNSRecordsSRVRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSRVRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSRVRecordData]
type dnsRecordUpdateResponseDNSRecordsSRVRecordDataJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSRVRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSRVRecordTypeSRV DNSRecordUpdateResponseDNSRecordsSRVRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSRVRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSRVRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSRVRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSRVRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSRVRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSRVRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSRVRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSRVRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSRVRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSRVRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSRVRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSRVRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsSSHFPRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordUpdateResponseDNSRecordsSSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSSHFPRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSSHFPRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSSHFPRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSSHFPRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSSHFPRecord]
type dnsRecordUpdateResponseDNSRecordsSSHFPRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSSHFPRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SSHFP record.
type DNSRecordUpdateResponseDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                              `json:"type"`
	JSON dnsRecordUpdateResponseDNSRecordsSSHFPRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSSHFPRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSSHFPRecordData]
type dnsRecordUpdateResponseDNSRecordsSSHFPRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSSHFPRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSSHFPRecordTypeSSHFP DNSRecordUpdateResponseDNSRecordsSSHFPRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSSHFPRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                               `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSSHFPRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSSHFPRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSSHFPRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSSHFPRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSSHFPRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSSHFPRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSSHFPRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsSVCBRecord struct {
	// Components of a SVCB record.
	Data DNSRecordUpdateResponseDNSRecordsSVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsSVCBRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsSVCBRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsSVCBRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsSVCBRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSVCBRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsSVCBRecord]
type dnsRecordUpdateResponseDNSRecordsSVCBRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsSVCBRecord) implementsDNSRecordUpdateResponse() {}

// Components of a SVCB record.
type DNSRecordUpdateResponseDNSRecordsSVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                              `json:"value"`
	JSON  dnsRecordUpdateResponseDNSRecordsSVCBRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSVCBRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSVCBRecordData]
type dnsRecordUpdateResponseDNSRecordsSVCBRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsSVCBRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsSVCBRecordTypeSVCB DNSRecordUpdateResponseDNSRecordsSVCBRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsSVCBRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsSVCBRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsSVCBRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsSVCBRecordMeta]
type dnsRecordUpdateResponseDNSRecordsSVCBRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsSVCBRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsSVCBRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsSVCBRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsSVCBRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsSVCBRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsSVCBRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsSVCBRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsTLSARecord struct {
	// Components of a TLSA record.
	Data DNSRecordUpdateResponseDNSRecordsTLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsTLSARecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsTLSARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsTLSARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsTLSARecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTLSARecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsTLSARecord]
type dnsRecordUpdateResponseDNSRecordsTLSARecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsTLSARecord) implementsDNSRecordUpdateResponse() {}

// Components of a TLSA record.
type DNSRecordUpdateResponseDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                             `json:"usage"`
	JSON  dnsRecordUpdateResponseDNSRecordsTLSARecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTLSARecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsTLSARecordData]
type dnsRecordUpdateResponseDNSRecordsTLSARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsTLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsTLSARecordType string

const (
	DNSRecordUpdateResponseDNSRecordsTLSARecordTypeTLSA DNSRecordUpdateResponseDNSRecordsTLSARecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsTLSARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsTLSARecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTLSARecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsTLSARecordMeta]
type dnsRecordUpdateResponseDNSRecordsTLSARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsTLSARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsTLSARecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsTLSARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsTLSARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsTLSARecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsTLSARecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsTLSARecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsTXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsTXTRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsTXTRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsTXTRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsTXTRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTXTRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsTXTRecord]
type dnsRecordUpdateResponseDNSRecordsTXTRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsTXTRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseDNSRecordsTXTRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsTXTRecordTypeTXT DNSRecordUpdateResponseDNSRecordsTXTRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsTXTRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsTXTRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsTXTRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsTXTRecordMeta]
type dnsRecordUpdateResponseDNSRecordsTXTRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsTXTRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsTXTRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsTXTRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsTXTRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsTXTRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsTXTRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsTXTRecordTTLNumber = 1
)

type DNSRecordUpdateResponseDNSRecordsURIRecord struct {
	// Components of a URI record.
	Data DNSRecordUpdateResponseDNSRecordsURIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSRecordsURIRecordType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSRecordsURIRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSRecordsURIRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSRecordsURIRecordJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsURIRecordJSON contains the JSON metadata for
// the struct [DNSRecordUpdateResponseDNSRecordsURIRecord]
type dnsRecordUpdateResponseDNSRecordsURIRecordJSON struct {
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

func (r *DNSRecordUpdateResponseDNSRecordsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordUpdateResponseDNSRecordsURIRecord) implementsDNSRecordUpdateResponse() {}

// Components of a URI record.
type DNSRecordUpdateResponseDNSRecordsURIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                            `json:"weight"`
	JSON   dnsRecordUpdateResponseDNSRecordsURIRecordDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsURIRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsURIRecordData]
type dnsRecordUpdateResponseDNSRecordsURIRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsURIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordUpdateResponseDNSRecordsURIRecordType string

const (
	DNSRecordUpdateResponseDNSRecordsURIRecordTypeURI DNSRecordUpdateResponseDNSRecordsURIRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSRecordsURIRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordUpdateResponseDNSRecordsURIRecordMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSRecordsURIRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordUpdateResponseDNSRecordsURIRecordMeta]
type dnsRecordUpdateResponseDNSRecordsURIRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSRecordsURIRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSRecordsURIRecordTTLNumber].
type DNSRecordUpdateResponseDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordUpdateResponseDNSRecordsURIRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSRecordsURIRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSRecordsURIRecordTTLNumber float64

const (
	DNSRecordUpdateResponseDNSRecordsURIRecordTTLNumber1 DNSRecordUpdateResponseDNSRecordsURIRecordTTLNumber = 1
)

// Union satisfied by [DNSRecordListResponseDNSRecordsARecord],
// [DNSRecordListResponseDNSRecordsAAAARecord],
// [DNSRecordListResponseDNSRecordsCAARecord],
// [DNSRecordListResponseDNSRecordsCertRecord],
// [DNSRecordListResponseDNSRecordsCNAMERecord],
// [DNSRecordListResponseDNSRecordsDNSKEYRecord],
// [DNSRecordListResponseDNSRecordsDSRecord],
// [DNSRecordListResponseDNSRecordsHTTPSRecord],
// [DNSRecordListResponseDNSRecordsLOCRecord],
// [DNSRecordListResponseDNSRecordsMXRecord],
// [DNSRecordListResponseDNSRecordsNAPTRRecord],
// [DNSRecordListResponseDNSRecordsNSRecord],
// [DNSRecordListResponseDNSRecordsPTRRecord],
// [DNSRecordListResponseDNSRecordsSmimeaRecord],
// [DNSRecordListResponseDNSRecordsSRVRecord],
// [DNSRecordListResponseDNSRecordsSSHFPRecord],
// [DNSRecordListResponseDNSRecordsSVCBRecord],
// [DNSRecordListResponseDNSRecordsTLSARecord],
// [DNSRecordListResponseDNSRecordsTXTRecord] or
// [DNSRecordListResponseDNSRecordsURIRecord].
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

type DNSRecordListResponseDNSRecordsAAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsAAAARecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsAAAARecordMeta `json:"meta"`
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
	TTL DNSRecordListResponseDNSRecordsAAAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsAAAARecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsAAAARecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsAAAARecord]
type dnsRecordListResponseDNSRecordsAAAARecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsAAAARecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsAAAARecordType string

const (
	DNSRecordListResponseDNSRecordsAAAARecordTypeAAAA DNSRecordListResponseDNSRecordsAAAARecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsAAAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsAAAARecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsAAAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsAAAARecordMeta]
type dnsRecordListResponseDNSRecordsAAAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsAAAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsAAAARecordTTLNumber].
type DNSRecordListResponseDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsAAAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsAAAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsAAAARecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsAAAARecordTTLNumber1 DNSRecordListResponseDNSRecordsAAAARecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsCAARecord struct {
	// Components of a CAA record.
	Data DNSRecordListResponseDNSRecordsCAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsCAARecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsCAARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsCAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsCAARecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCAARecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsCAARecord]
type dnsRecordListResponseDNSRecordsCAARecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsCAARecord) implementsDNSRecordListResponse() {}

// Components of a CAA record.
type DNSRecordListResponseDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                           `json:"value"`
	JSON  dnsRecordListResponseDNSRecordsCAARecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCAARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCAARecordData]
type dnsRecordListResponseDNSRecordsCAARecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsCAARecordType string

const (
	DNSRecordListResponseDNSRecordsCAARecordTypeCAA DNSRecordListResponseDNSRecordsCAARecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsCAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsCAARecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCAARecordMeta]
type dnsRecordListResponseDNSRecordsCAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsCAARecordTTLNumber].
type DNSRecordListResponseDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsCAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsCAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsCAARecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsCAARecordTTLNumber1 DNSRecordListResponseDNSRecordsCAARecordTTLNumber = 1
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

type DNSRecordListResponseDNSRecordsCNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsCNAMERecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsCNAMERecordMeta `json:"meta"`
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
	TTL DNSRecordListResponseDNSRecordsCNAMERecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsCNAMERecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCNAMERecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsCNAMERecord]
type dnsRecordListResponseDNSRecordsCNAMERecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsCNAMERecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsCNAMERecordType string

const (
	DNSRecordListResponseDNSRecordsCNAMERecordTypeCNAME DNSRecordListResponseDNSRecordsCNAMERecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsCNAMERecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsCNAMERecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsCNAMERecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsCNAMERecordMeta]
type dnsRecordListResponseDNSRecordsCNAMERecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsCNAMERecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsCNAMERecordTTLNumber].
type DNSRecordListResponseDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsCNAMERecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsCNAMERecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsCNAMERecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsCNAMERecordTTLNumber1 DNSRecordListResponseDNSRecordsCNAMERecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsDNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordListResponseDNSRecordsDNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsDNSKEYRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsDNSKEYRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsDNSKEYRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsDNSKEYRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDNSKEYRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsDNSKEYRecord]
type dnsRecordListResponseDNSRecordsDNSKEYRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsDNSKEYRecord) implementsDNSRecordListResponse() {}

// Components of a DNSKEY record.
type DNSRecordListResponseDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                              `json:"public_key"`
	JSON      dnsRecordListResponseDNSRecordsDNSKEYRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDNSKEYRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsDNSKEYRecordData]
type dnsRecordListResponseDNSRecordsDNSKEYRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsDNSKEYRecordType string

const (
	DNSRecordListResponseDNSRecordsDNSKEYRecordTypeDNSKEY DNSRecordListResponseDNSRecordsDNSKEYRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsDNSKEYRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsDNSKEYRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDNSKEYRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsDNSKEYRecordMeta]
type dnsRecordListResponseDNSRecordsDNSKEYRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDNSKEYRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsDNSKEYRecordTTLNumber].
type DNSRecordListResponseDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsDNSKEYRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsDNSKEYRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsDNSKEYRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsDNSKEYRecordTTLNumber1 DNSRecordListResponseDNSRecordsDNSKEYRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsDSRecord struct {
	// Components of a DS record.
	Data DNSRecordListResponseDNSRecordsDSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsDSRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsDSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsDSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsDSRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDSRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsDSRecord]
type dnsRecordListResponseDNSRecordsDSRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsDSRecord) implementsDNSRecordListResponse() {}

// Components of a DS record.
type DNSRecordListResponseDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                         `json:"key_tag"`
	JSON   dnsRecordListResponseDNSRecordsDSRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsDSRecordData]
type dnsRecordListResponseDNSRecordsDSRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsDSRecordType string

const (
	DNSRecordListResponseDNSRecordsDSRecordTypeDS DNSRecordListResponseDNSRecordsDSRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsDSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsDSRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsDSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsDSRecordMeta]
type dnsRecordListResponseDNSRecordsDSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsDSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsDSRecordTTLNumber].
type DNSRecordListResponseDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsDSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsDSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsDSRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsDSRecordTTLNumber1 DNSRecordListResponseDNSRecordsDSRecordTTLNumber = 1
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

type DNSRecordListResponseDNSRecordsLOCRecord struct {
	// Components of a LOC record.
	Data DNSRecordListResponseDNSRecordsLOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsLOCRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsLOCRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsLOCRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsLOCRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsLOCRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsLOCRecord]
type dnsRecordListResponseDNSRecordsLOCRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsLOCRecord) implementsDNSRecordListResponse() {}

// Components of a LOC record.
type DNSRecordListResponseDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordListResponseDNSRecordsLOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordListResponseDNSRecordsLOCRecordDataLongDirection `json:"long_direction"`
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
	JSON dnsRecordListResponseDNSRecordsLOCRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsLOCRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsLOCRecordData]
type dnsRecordListResponseDNSRecordsLOCRecordDataJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsLOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordListResponseDNSRecordsLOCRecordDataLatDirection string

const (
	DNSRecordListResponseDNSRecordsLOCRecordDataLatDirectionN DNSRecordListResponseDNSRecordsLOCRecordDataLatDirection = "N"
	DNSRecordListResponseDNSRecordsLOCRecordDataLatDirectionS DNSRecordListResponseDNSRecordsLOCRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordListResponseDNSRecordsLOCRecordDataLongDirection string

const (
	DNSRecordListResponseDNSRecordsLOCRecordDataLongDirectionE DNSRecordListResponseDNSRecordsLOCRecordDataLongDirection = "E"
	DNSRecordListResponseDNSRecordsLOCRecordDataLongDirectionW DNSRecordListResponseDNSRecordsLOCRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordListResponseDNSRecordsLOCRecordType string

const (
	DNSRecordListResponseDNSRecordsLOCRecordTypeLOC DNSRecordListResponseDNSRecordsLOCRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsLOCRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsLOCRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsLOCRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsLOCRecordMeta]
type dnsRecordListResponseDNSRecordsLOCRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsLOCRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsLOCRecordTTLNumber].
type DNSRecordListResponseDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsLOCRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsLOCRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsLOCRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsLOCRecordTTLNumber1 DNSRecordListResponseDNSRecordsLOCRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsMXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsMXRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsMXRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsMXRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsMXRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsMXRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsMXRecord]
type dnsRecordListResponseDNSRecordsMXRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsMXRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsMXRecordType string

const (
	DNSRecordListResponseDNSRecordsMXRecordTypeMX DNSRecordListResponseDNSRecordsMXRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsMXRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsMXRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsMXRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsMXRecordMeta]
type dnsRecordListResponseDNSRecordsMXRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsMXRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsMXRecordTTLNumber].
type DNSRecordListResponseDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsMXRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsMXRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsMXRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsMXRecordTTLNumber1 DNSRecordListResponseDNSRecordsMXRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsNAPTRRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordListResponseDNSRecordsNAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsNAPTRRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsNAPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsNAPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsNAPTRRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNAPTRRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsNAPTRRecord]
type dnsRecordListResponseDNSRecordsNAPTRRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsNAPTRRecord) implementsDNSRecordListResponse() {}

// Components of a NAPTR record.
type DNSRecordListResponseDNSRecordsNAPTRRecordData struct {
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
	JSON    dnsRecordListResponseDNSRecordsNAPTRRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNAPTRRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsNAPTRRecordData]
type dnsRecordListResponseDNSRecordsNAPTRRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsNAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsNAPTRRecordType string

const (
	DNSRecordListResponseDNSRecordsNAPTRRecordTypeNAPTR DNSRecordListResponseDNSRecordsNAPTRRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsNAPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsNAPTRRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNAPTRRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsNAPTRRecordMeta]
type dnsRecordListResponseDNSRecordsNAPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsNAPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsNAPTRRecordTTLNumber].
type DNSRecordListResponseDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsNAPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsNAPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsNAPTRRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsNAPTRRecordTTLNumber1 DNSRecordListResponseDNSRecordsNAPTRRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsNSRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsNSRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsNSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsNSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsNSRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNSRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsNSRecord]
type dnsRecordListResponseDNSRecordsNSRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsNSRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsNSRecordType string

const (
	DNSRecordListResponseDNSRecordsNSRecordTypeNS DNSRecordListResponseDNSRecordsNSRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsNSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsNSRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsNSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsNSRecordMeta]
type dnsRecordListResponseDNSRecordsNSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsNSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsNSRecordTTLNumber].
type DNSRecordListResponseDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsNSRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsNSRecordTTLNumber1 DNSRecordListResponseDNSRecordsNSRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsPTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsPTRRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsPTRRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsPTRRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsPTRRecord]
type dnsRecordListResponseDNSRecordsPTRRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsPTRRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsPTRRecordType string

const (
	DNSRecordListResponseDNSRecordsPTRRecordTypePTR DNSRecordListResponseDNSRecordsPTRRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsPTRRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsPTRRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsPTRRecordMeta]
type dnsRecordListResponseDNSRecordsPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsPTRRecordTTLNumber].
type DNSRecordListResponseDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsPTRRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsPTRRecordTTLNumber1 DNSRecordListResponseDNSRecordsPTRRecordTTLNumber = 1
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

type DNSRecordListResponseDNSRecordsSRVRecord struct {
	// Components of a SRV record.
	Data DNSRecordListResponseDNSRecordsSRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSRVRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSRVRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSRVRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSRVRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSRVRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsSRVRecord]
type dnsRecordListResponseDNSRecordsSRVRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSRVRecord) implementsDNSRecordListResponse() {}

// Components of a SRV record.
type DNSRecordListResponseDNSRecordsSRVRecordData struct {
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
	JSON   dnsRecordListResponseDNSRecordsSRVRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSRVRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSRVRecordData]
type dnsRecordListResponseDNSRecordsSRVRecordDataJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSRVRecordType string

const (
	DNSRecordListResponseDNSRecordsSRVRecordTypeSRV DNSRecordListResponseDNSRecordsSRVRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSRVRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSRVRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSRVRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSRVRecordMeta]
type dnsRecordListResponseDNSRecordsSRVRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSRVRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSRVRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSRVRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSRVRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSRVRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSRVRecordTTLNumber1 DNSRecordListResponseDNSRecordsSRVRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsSSHFPRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordListResponseDNSRecordsSSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSSHFPRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSSHFPRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSSHFPRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSSHFPRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSSHFPRecordJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSSHFPRecord]
type dnsRecordListResponseDNSRecordsSSHFPRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSSHFPRecord) implementsDNSRecordListResponse() {}

// Components of a SSHFP record.
type DNSRecordListResponseDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                            `json:"type"`
	JSON dnsRecordListResponseDNSRecordsSSHFPRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSSHFPRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsSSHFPRecordData]
type dnsRecordListResponseDNSRecordsSSHFPRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSSHFPRecordType string

const (
	DNSRecordListResponseDNSRecordsSSHFPRecordTypeSSHFP DNSRecordListResponseDNSRecordsSSHFPRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSSHFPRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSSHFPRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSSHFPRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordListResponseDNSRecordsSSHFPRecordMeta]
type dnsRecordListResponseDNSRecordsSSHFPRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSSHFPRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSSHFPRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSSHFPRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSSHFPRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSSHFPRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSSHFPRecordTTLNumber1 DNSRecordListResponseDNSRecordsSSHFPRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsSVCBRecord struct {
	// Components of a SVCB record.
	Data DNSRecordListResponseDNSRecordsSVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsSVCBRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsSVCBRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsSVCBRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsSVCBRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSVCBRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsSVCBRecord]
type dnsRecordListResponseDNSRecordsSVCBRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsSVCBRecord) implementsDNSRecordListResponse() {}

// Components of a SVCB record.
type DNSRecordListResponseDNSRecordsSVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                            `json:"value"`
	JSON  dnsRecordListResponseDNSRecordsSVCBRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSVCBRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSVCBRecordData]
type dnsRecordListResponseDNSRecordsSVCBRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsSVCBRecordType string

const (
	DNSRecordListResponseDNSRecordsSVCBRecordTypeSVCB DNSRecordListResponseDNSRecordsSVCBRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsSVCBRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsSVCBRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsSVCBRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsSVCBRecordMeta]
type dnsRecordListResponseDNSRecordsSVCBRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsSVCBRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsSVCBRecordTTLNumber].
type DNSRecordListResponseDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsSVCBRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsSVCBRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsSVCBRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsSVCBRecordTTLNumber1 DNSRecordListResponseDNSRecordsSVCBRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsTLSARecord struct {
	// Components of a TLSA record.
	Data DNSRecordListResponseDNSRecordsTLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsTLSARecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsTLSARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsTLSARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsTLSARecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTLSARecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsTLSARecord]
type dnsRecordListResponseDNSRecordsTLSARecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsTLSARecord) implementsDNSRecordListResponse() {}

// Components of a TLSA record.
type DNSRecordListResponseDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                           `json:"usage"`
	JSON  dnsRecordListResponseDNSRecordsTLSARecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTLSARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsTLSARecordData]
type dnsRecordListResponseDNSRecordsTLSARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsTLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsTLSARecordType string

const (
	DNSRecordListResponseDNSRecordsTLSARecordTypeTLSA DNSRecordListResponseDNSRecordsTLSARecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsTLSARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsTLSARecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTLSARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsTLSARecordMeta]
type dnsRecordListResponseDNSRecordsTLSARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsTLSARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsTLSARecordTTLNumber].
type DNSRecordListResponseDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsTLSARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsTLSARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsTLSARecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsTLSARecordTTLNumber1 DNSRecordListResponseDNSRecordsTLSARecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsTXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsTXTRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsTXTRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsTXTRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsTXTRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTXTRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsTXTRecord]
type dnsRecordListResponseDNSRecordsTXTRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsTXTRecord) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseDNSRecordsTXTRecordType string

const (
	DNSRecordListResponseDNSRecordsTXTRecordTypeTXT DNSRecordListResponseDNSRecordsTXTRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsTXTRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsTXTRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsTXTRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsTXTRecordMeta]
type dnsRecordListResponseDNSRecordsTXTRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsTXTRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsTXTRecordTTLNumber].
type DNSRecordListResponseDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsTXTRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsTXTRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsTXTRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsTXTRecordTTLNumber1 DNSRecordListResponseDNSRecordsTXTRecordTTLNumber = 1
)

type DNSRecordListResponseDNSRecordsURIRecord struct {
	// Components of a URI record.
	Data DNSRecordListResponseDNSRecordsURIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordListResponseDNSRecordsURIRecordType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSRecordsURIRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSRecordsURIRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSRecordsURIRecordJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsURIRecordJSON contains the JSON metadata for the
// struct [DNSRecordListResponseDNSRecordsURIRecord]
type dnsRecordListResponseDNSRecordsURIRecordJSON struct {
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

func (r *DNSRecordListResponseDNSRecordsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordListResponseDNSRecordsURIRecord) implementsDNSRecordListResponse() {}

// Components of a URI record.
type DNSRecordListResponseDNSRecordsURIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                          `json:"weight"`
	JSON   dnsRecordListResponseDNSRecordsURIRecordDataJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsURIRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsURIRecordData]
type dnsRecordListResponseDNSRecordsURIRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsURIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordListResponseDNSRecordsURIRecordType string

const (
	DNSRecordListResponseDNSRecordsURIRecordTypeURI DNSRecordListResponseDNSRecordsURIRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSRecordsURIRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordListResponseDNSRecordsURIRecordMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSRecordsURIRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordListResponseDNSRecordsURIRecordMeta]
type dnsRecordListResponseDNSRecordsURIRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSRecordsURIRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSRecordsURIRecordTTLNumber].
type DNSRecordListResponseDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordListResponseDNSRecordsURIRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSRecordsURIRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordListResponseDNSRecordsURIRecordTTLNumber float64

const (
	DNSRecordListResponseDNSRecordsURIRecordTTLNumber1 DNSRecordListResponseDNSRecordsURIRecordTTLNumber = 1
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

// Union satisfied by [DNSRecordEditResponseDNSRecordsARecord],
// [DNSRecordEditResponseDNSRecordsAAAARecord],
// [DNSRecordEditResponseDNSRecordsCAARecord],
// [DNSRecordEditResponseDNSRecordsCertRecord],
// [DNSRecordEditResponseDNSRecordsCNAMERecord],
// [DNSRecordEditResponseDNSRecordsDNSKEYRecord],
// [DNSRecordEditResponseDNSRecordsDSRecord],
// [DNSRecordEditResponseDNSRecordsHTTPSRecord],
// [DNSRecordEditResponseDNSRecordsLOCRecord],
// [DNSRecordEditResponseDNSRecordsMXRecord],
// [DNSRecordEditResponseDNSRecordsNAPTRRecord],
// [DNSRecordEditResponseDNSRecordsNSRecord],
// [DNSRecordEditResponseDNSRecordsPTRRecord],
// [DNSRecordEditResponseDNSRecordsSmimeaRecord],
// [DNSRecordEditResponseDNSRecordsSRVRecord],
// [DNSRecordEditResponseDNSRecordsSSHFPRecord],
// [DNSRecordEditResponseDNSRecordsSVCBRecord],
// [DNSRecordEditResponseDNSRecordsTLSARecord],
// [DNSRecordEditResponseDNSRecordsTXTRecord] or
// [DNSRecordEditResponseDNSRecordsURIRecord].
type DNSRecordEditResponse interface {
	implementsDNSRecordEditResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DNSRecordEditResponse)(nil)).Elem(), "")
}

type DNSRecordEditResponseDNSRecordsARecord struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsARecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsARecordMeta `json:"meta"`
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
	TTL DNSRecordEditResponseDNSRecordsARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsARecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsARecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsARecord]
type dnsRecordEditResponseDNSRecordsARecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsARecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsARecordType string

const (
	DNSRecordEditResponseDNSRecordsARecordTypeA DNSRecordEditResponseDNSRecordsARecordType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsARecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsARecordMeta]
type dnsRecordEditResponseDNSRecordsARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsARecordTTLNumber].
type DNSRecordEditResponseDNSRecordsARecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsARecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsARecordTTLNumber1 DNSRecordEditResponseDNSRecordsARecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsAAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsAAAARecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsAAAARecordMeta `json:"meta"`
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
	TTL DNSRecordEditResponseDNSRecordsAAAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsAAAARecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsAAAARecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsAAAARecord]
type dnsRecordEditResponseDNSRecordsAAAARecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsAAAARecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsAAAARecordType string

const (
	DNSRecordEditResponseDNSRecordsAAAARecordTypeAAAA DNSRecordEditResponseDNSRecordsAAAARecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsAAAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsAAAARecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsAAAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsAAAARecordMeta]
type dnsRecordEditResponseDNSRecordsAAAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsAAAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsAAAARecordTTLNumber].
type DNSRecordEditResponseDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsAAAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsAAAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsAAAARecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsAAAARecordTTLNumber1 DNSRecordEditResponseDNSRecordsAAAARecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsCAARecord struct {
	// Components of a CAA record.
	Data DNSRecordEditResponseDNSRecordsCAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsCAARecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsCAARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsCAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsCAARecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCAARecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsCAARecord]
type dnsRecordEditResponseDNSRecordsCAARecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsCAARecord) implementsDNSRecordEditResponse() {}

// Components of a CAA record.
type DNSRecordEditResponseDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                           `json:"value"`
	JSON  dnsRecordEditResponseDNSRecordsCAARecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCAARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsCAARecordData]
type dnsRecordEditResponseDNSRecordsCAARecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsCAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsCAARecordType string

const (
	DNSRecordEditResponseDNSRecordsCAARecordTypeCAA DNSRecordEditResponseDNSRecordsCAARecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsCAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsCAARecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsCAARecordMeta]
type dnsRecordEditResponseDNSRecordsCAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsCAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsCAARecordTTLNumber].
type DNSRecordEditResponseDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsCAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsCAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsCAARecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsCAARecordTTLNumber1 DNSRecordEditResponseDNSRecordsCAARecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsCertRecord struct {
	// Components of a CERT record.
	Data DNSRecordEditResponseDNSRecordsCertRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsCertRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsCertRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsCertRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsCertRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCertRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsCertRecord]
type dnsRecordEditResponseDNSRecordsCertRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsCertRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsCertRecord) implementsDNSRecordEditResponse() {}

// Components of a CERT record.
type DNSRecordEditResponseDNSRecordsCertRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                                           `json:"type"`
	JSON dnsRecordEditResponseDNSRecordsCertRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCertRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsCertRecordData]
type dnsRecordEditResponseDNSRecordsCertRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsCertRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsCertRecordType string

const (
	DNSRecordEditResponseDNSRecordsCertRecordTypeCert DNSRecordEditResponseDNSRecordsCertRecordType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsCertRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsCertRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCertRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsCertRecordMeta]
type dnsRecordEditResponseDNSRecordsCertRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsCertRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsCertRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsCertRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsCertRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsCertRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsCertRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsCertRecordTTLNumber1 DNSRecordEditResponseDNSRecordsCertRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsCNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsCNAMERecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsCNAMERecordMeta `json:"meta"`
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
	TTL DNSRecordEditResponseDNSRecordsCNAMERecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsCNAMERecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCNAMERecordJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsCNAMERecord]
type dnsRecordEditResponseDNSRecordsCNAMERecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsCNAMERecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsCNAMERecordType string

const (
	DNSRecordEditResponseDNSRecordsCNAMERecordTypeCNAME DNSRecordEditResponseDNSRecordsCNAMERecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsCNAMERecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsCNAMERecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsCNAMERecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsCNAMERecordMeta]
type dnsRecordEditResponseDNSRecordsCNAMERecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsCNAMERecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsCNAMERecordTTLNumber].
type DNSRecordEditResponseDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsCNAMERecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsCNAMERecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsCNAMERecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsCNAMERecordTTLNumber1 DNSRecordEditResponseDNSRecordsCNAMERecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsDNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordEditResponseDNSRecordsDNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsDNSKEYRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsDNSKEYRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsDNSKEYRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsDNSKEYRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsDNSKEYRecordJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsDNSKEYRecord]
type dnsRecordEditResponseDNSRecordsDNSKEYRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsDNSKEYRecord) implementsDNSRecordEditResponse() {}

// Components of a DNSKEY record.
type DNSRecordEditResponseDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                              `json:"public_key"`
	JSON      dnsRecordEditResponseDNSRecordsDNSKEYRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsDNSKEYRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsDNSKEYRecordData]
type dnsRecordEditResponseDNSRecordsDNSKEYRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsDNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsDNSKEYRecordType string

const (
	DNSRecordEditResponseDNSRecordsDNSKEYRecordTypeDNSKEY DNSRecordEditResponseDNSRecordsDNSKEYRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsDNSKEYRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsDNSKEYRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsDNSKEYRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsDNSKEYRecordMeta]
type dnsRecordEditResponseDNSRecordsDNSKEYRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsDNSKEYRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsDNSKEYRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsDNSKEYRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsDNSKEYRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsDNSKEYRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsDNSKEYRecordTTLNumber1 DNSRecordEditResponseDNSRecordsDNSKEYRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsDSRecord struct {
	// Components of a DS record.
	Data DNSRecordEditResponseDNSRecordsDSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsDSRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsDSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsDSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsDSRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsDSRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsDSRecord]
type dnsRecordEditResponseDNSRecordsDSRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsDSRecord) implementsDNSRecordEditResponse() {}

// Components of a DS record.
type DNSRecordEditResponseDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                         `json:"key_tag"`
	JSON   dnsRecordEditResponseDNSRecordsDSRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsDSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsDSRecordData]
type dnsRecordEditResponseDNSRecordsDSRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsDSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsDSRecordType string

const (
	DNSRecordEditResponseDNSRecordsDSRecordTypeDS DNSRecordEditResponseDNSRecordsDSRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsDSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsDSRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsDSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsDSRecordMeta]
type dnsRecordEditResponseDNSRecordsDSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsDSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsDSRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsDSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsDSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsDSRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsDSRecordTTLNumber1 DNSRecordEditResponseDNSRecordsDSRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsHTTPSRecord struct {
	// Components of a HTTPS record.
	Data DNSRecordEditResponseDNSRecordsHTTPSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsHTTPSRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsHTTPSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsHTTPSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsHTTPSRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsHTTPSRecordJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsHTTPSRecord]
type dnsRecordEditResponseDNSRecordsHTTPSRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsHTTPSRecord) implementsDNSRecordEditResponse() {}

// Components of a HTTPS record.
type DNSRecordEditResponseDNSRecordsHTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                             `json:"value"`
	JSON  dnsRecordEditResponseDNSRecordsHTTPSRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsHTTPSRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsHTTPSRecordData]
type dnsRecordEditResponseDNSRecordsHTTPSRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsHTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsHTTPSRecordType string

const (
	DNSRecordEditResponseDNSRecordsHTTPSRecordTypeHTTPS DNSRecordEditResponseDNSRecordsHTTPSRecordType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsHTTPSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsHTTPSRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsHTTPSRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsHTTPSRecordMeta]
type dnsRecordEditResponseDNSRecordsHTTPSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsHTTPSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsHTTPSRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsHTTPSRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsHTTPSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsHTTPSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsHTTPSRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsHTTPSRecordTTLNumber1 DNSRecordEditResponseDNSRecordsHTTPSRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsLOCRecord struct {
	// Components of a LOC record.
	Data DNSRecordEditResponseDNSRecordsLOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsLOCRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsLOCRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsLOCRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsLOCRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsLOCRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsLOCRecord]
type dnsRecordEditResponseDNSRecordsLOCRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsLOCRecord) implementsDNSRecordEditResponse() {}

// Components of a LOC record.
type DNSRecordEditResponseDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordEditResponseDNSRecordsLOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordEditResponseDNSRecordsLOCRecordDataLongDirection `json:"long_direction"`
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
	JSON dnsRecordEditResponseDNSRecordsLOCRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsLOCRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsLOCRecordData]
type dnsRecordEditResponseDNSRecordsLOCRecordDataJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsLOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordEditResponseDNSRecordsLOCRecordDataLatDirection string

const (
	DNSRecordEditResponseDNSRecordsLOCRecordDataLatDirectionN DNSRecordEditResponseDNSRecordsLOCRecordDataLatDirection = "N"
	DNSRecordEditResponseDNSRecordsLOCRecordDataLatDirectionS DNSRecordEditResponseDNSRecordsLOCRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordEditResponseDNSRecordsLOCRecordDataLongDirection string

const (
	DNSRecordEditResponseDNSRecordsLOCRecordDataLongDirectionE DNSRecordEditResponseDNSRecordsLOCRecordDataLongDirection = "E"
	DNSRecordEditResponseDNSRecordsLOCRecordDataLongDirectionW DNSRecordEditResponseDNSRecordsLOCRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordEditResponseDNSRecordsLOCRecordType string

const (
	DNSRecordEditResponseDNSRecordsLOCRecordTypeLOC DNSRecordEditResponseDNSRecordsLOCRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsLOCRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsLOCRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsLOCRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsLOCRecordMeta]
type dnsRecordEditResponseDNSRecordsLOCRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsLOCRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsLOCRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsLOCRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsLOCRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsLOCRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsLOCRecordTTLNumber1 DNSRecordEditResponseDNSRecordsLOCRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsMXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsMXRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsMXRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsMXRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsMXRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsMXRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsMXRecord]
type dnsRecordEditResponseDNSRecordsMXRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsMXRecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsMXRecordType string

const (
	DNSRecordEditResponseDNSRecordsMXRecordTypeMX DNSRecordEditResponseDNSRecordsMXRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsMXRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsMXRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsMXRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsMXRecordMeta]
type dnsRecordEditResponseDNSRecordsMXRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsMXRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsMXRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsMXRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsMXRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsMXRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsMXRecordTTLNumber1 DNSRecordEditResponseDNSRecordsMXRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsNAPTRRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordEditResponseDNSRecordsNAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsNAPTRRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsNAPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsNAPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsNAPTRRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsNAPTRRecordJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsNAPTRRecord]
type dnsRecordEditResponseDNSRecordsNAPTRRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsNAPTRRecord) implementsDNSRecordEditResponse() {}

// Components of a NAPTR record.
type DNSRecordEditResponseDNSRecordsNAPTRRecordData struct {
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
	JSON    dnsRecordEditResponseDNSRecordsNAPTRRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsNAPTRRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsNAPTRRecordData]
type dnsRecordEditResponseDNSRecordsNAPTRRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsNAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsNAPTRRecordType string

const (
	DNSRecordEditResponseDNSRecordsNAPTRRecordTypeNAPTR DNSRecordEditResponseDNSRecordsNAPTRRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsNAPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsNAPTRRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsNAPTRRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsNAPTRRecordMeta]
type dnsRecordEditResponseDNSRecordsNAPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsNAPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsNAPTRRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsNAPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsNAPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsNAPTRRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsNAPTRRecordTTLNumber1 DNSRecordEditResponseDNSRecordsNAPTRRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsNSRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsNSRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsNSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsNSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsNSRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsNSRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsNSRecord]
type dnsRecordEditResponseDNSRecordsNSRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsNSRecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsNSRecordType string

const (
	DNSRecordEditResponseDNSRecordsNSRecordTypeNS DNSRecordEditResponseDNSRecordsNSRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsNSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsNSRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsNSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsNSRecordMeta]
type dnsRecordEditResponseDNSRecordsNSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsNSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsNSRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsNSRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsNSRecordTTLNumber1 DNSRecordEditResponseDNSRecordsNSRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsPTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsPTRRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsPTRRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsPTRRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsPTRRecord]
type dnsRecordEditResponseDNSRecordsPTRRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsPTRRecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsPTRRecordType string

const (
	DNSRecordEditResponseDNSRecordsPTRRecordTypePTR DNSRecordEditResponseDNSRecordsPTRRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsPTRRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsPTRRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsPTRRecordMeta]
type dnsRecordEditResponseDNSRecordsPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsPTRRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsPTRRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsPTRRecordTTLNumber1 DNSRecordEditResponseDNSRecordsPTRRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsSmimeaRecord struct {
	// Components of a SMIMEA record.
	Data DNSRecordEditResponseDNSRecordsSmimeaRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsSmimeaRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsSmimeaRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsSmimeaRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsSmimeaRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSmimeaRecordJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsSmimeaRecord]
type dnsRecordEditResponseDNSRecordsSmimeaRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsSmimeaRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsSmimeaRecord) implementsDNSRecordEditResponse() {}

// Components of a SMIMEA record.
type DNSRecordEditResponseDNSRecordsSmimeaRecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                             `json:"usage"`
	JSON  dnsRecordEditResponseDNSRecordsSmimeaRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSmimeaRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsSmimeaRecordData]
type dnsRecordEditResponseDNSRecordsSmimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSmimeaRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsSmimeaRecordType string

const (
	DNSRecordEditResponseDNSRecordsSmimeaRecordTypeSmimea DNSRecordEditResponseDNSRecordsSmimeaRecordType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsSmimeaRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                              `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsSmimeaRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSmimeaRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsSmimeaRecordMeta]
type dnsRecordEditResponseDNSRecordsSmimeaRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSmimeaRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsSmimeaRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsSmimeaRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsSmimeaRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsSmimeaRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsSmimeaRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsSmimeaRecordTTLNumber1 DNSRecordEditResponseDNSRecordsSmimeaRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsSRVRecord struct {
	// Components of a SRV record.
	Data DNSRecordEditResponseDNSRecordsSRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsSRVRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsSRVRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsSRVRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsSRVRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSRVRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsSRVRecord]
type dnsRecordEditResponseDNSRecordsSRVRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsSRVRecord) implementsDNSRecordEditResponse() {}

// Components of a SRV record.
type DNSRecordEditResponseDNSRecordsSRVRecordData struct {
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
	JSON   dnsRecordEditResponseDNSRecordsSRVRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSRVRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsSRVRecordData]
type dnsRecordEditResponseDNSRecordsSRVRecordDataJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsSRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsSRVRecordType string

const (
	DNSRecordEditResponseDNSRecordsSRVRecordTypeSRV DNSRecordEditResponseDNSRecordsSRVRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsSRVRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsSRVRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSRVRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsSRVRecordMeta]
type dnsRecordEditResponseDNSRecordsSRVRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSRVRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsSRVRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsSRVRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsSRVRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsSRVRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsSRVRecordTTLNumber1 DNSRecordEditResponseDNSRecordsSRVRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsSSHFPRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordEditResponseDNSRecordsSSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsSSHFPRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsSSHFPRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsSSHFPRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsSSHFPRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSSHFPRecordJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsSSHFPRecord]
type dnsRecordEditResponseDNSRecordsSSHFPRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsSSHFPRecord) implementsDNSRecordEditResponse() {}

// Components of a SSHFP record.
type DNSRecordEditResponseDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                            `json:"type"`
	JSON dnsRecordEditResponseDNSRecordsSSHFPRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSSHFPRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsSSHFPRecordData]
type dnsRecordEditResponseDNSRecordsSSHFPRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsSSHFPRecordType string

const (
	DNSRecordEditResponseDNSRecordsSSHFPRecordTypeSSHFP DNSRecordEditResponseDNSRecordsSSHFPRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsSSHFPRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsSSHFPRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSSHFPRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordEditResponseDNSRecordsSSHFPRecordMeta]
type dnsRecordEditResponseDNSRecordsSSHFPRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSSHFPRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsSSHFPRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsSSHFPRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsSSHFPRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsSSHFPRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsSSHFPRecordTTLNumber1 DNSRecordEditResponseDNSRecordsSSHFPRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsSVCBRecord struct {
	// Components of a SVCB record.
	Data DNSRecordEditResponseDNSRecordsSVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsSVCBRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsSVCBRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsSVCBRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsSVCBRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSVCBRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsSVCBRecord]
type dnsRecordEditResponseDNSRecordsSVCBRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsSVCBRecord) implementsDNSRecordEditResponse() {}

// Components of a SVCB record.
type DNSRecordEditResponseDNSRecordsSVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                            `json:"value"`
	JSON  dnsRecordEditResponseDNSRecordsSVCBRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSVCBRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsSVCBRecordData]
type dnsRecordEditResponseDNSRecordsSVCBRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsSVCBRecordType string

const (
	DNSRecordEditResponseDNSRecordsSVCBRecordTypeSVCB DNSRecordEditResponseDNSRecordsSVCBRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsSVCBRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsSVCBRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsSVCBRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsSVCBRecordMeta]
type dnsRecordEditResponseDNSRecordsSVCBRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsSVCBRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsSVCBRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsSVCBRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsSVCBRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsSVCBRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsSVCBRecordTTLNumber1 DNSRecordEditResponseDNSRecordsSVCBRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsTLSARecord struct {
	// Components of a TLSA record.
	Data DNSRecordEditResponseDNSRecordsTLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsTLSARecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsTLSARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsTLSARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsTLSARecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsTLSARecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsTLSARecord]
type dnsRecordEditResponseDNSRecordsTLSARecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsTLSARecord) implementsDNSRecordEditResponse() {}

// Components of a TLSA record.
type DNSRecordEditResponseDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                           `json:"usage"`
	JSON  dnsRecordEditResponseDNSRecordsTLSARecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsTLSARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsTLSARecordData]
type dnsRecordEditResponseDNSRecordsTLSARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsTLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsTLSARecordType string

const (
	DNSRecordEditResponseDNSRecordsTLSARecordTypeTLSA DNSRecordEditResponseDNSRecordsTLSARecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsTLSARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsTLSARecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsTLSARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsTLSARecordMeta]
type dnsRecordEditResponseDNSRecordsTLSARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsTLSARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsTLSARecordTTLNumber].
type DNSRecordEditResponseDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsTLSARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsTLSARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsTLSARecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsTLSARecordTTLNumber1 DNSRecordEditResponseDNSRecordsTLSARecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsTXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsTXTRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsTXTRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsTXTRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsTXTRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsTXTRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsTXTRecord]
type dnsRecordEditResponseDNSRecordsTXTRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsTXTRecord) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseDNSRecordsTXTRecordType string

const (
	DNSRecordEditResponseDNSRecordsTXTRecordTypeTXT DNSRecordEditResponseDNSRecordsTXTRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsTXTRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsTXTRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsTXTRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsTXTRecordMeta]
type dnsRecordEditResponseDNSRecordsTXTRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsTXTRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsTXTRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsTXTRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsTXTRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsTXTRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsTXTRecordTTLNumber1 DNSRecordEditResponseDNSRecordsTXTRecordTTLNumber = 1
)

type DNSRecordEditResponseDNSRecordsURIRecord struct {
	// Components of a URI record.
	Data DNSRecordEditResponseDNSRecordsURIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordEditResponseDNSRecordsURIRecordType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSRecordsURIRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSRecordsURIRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSRecordsURIRecordJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsURIRecordJSON contains the JSON metadata for the
// struct [DNSRecordEditResponseDNSRecordsURIRecord]
type dnsRecordEditResponseDNSRecordsURIRecordJSON struct {
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

func (r *DNSRecordEditResponseDNSRecordsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordEditResponseDNSRecordsURIRecord) implementsDNSRecordEditResponse() {}

// Components of a URI record.
type DNSRecordEditResponseDNSRecordsURIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                          `json:"weight"`
	JSON   dnsRecordEditResponseDNSRecordsURIRecordDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsURIRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsURIRecordData]
type dnsRecordEditResponseDNSRecordsURIRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsURIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordEditResponseDNSRecordsURIRecordType string

const (
	DNSRecordEditResponseDNSRecordsURIRecordTypeURI DNSRecordEditResponseDNSRecordsURIRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSRecordsURIRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordEditResponseDNSRecordsURIRecordMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSRecordsURIRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordEditResponseDNSRecordsURIRecordMeta]
type dnsRecordEditResponseDNSRecordsURIRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSRecordsURIRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSRecordsURIRecordTTLNumber].
type DNSRecordEditResponseDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordEditResponseDNSRecordsURIRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSRecordsURIRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordEditResponseDNSRecordsURIRecordTTLNumber float64

const (
	DNSRecordEditResponseDNSRecordsURIRecordTTLNumber1 DNSRecordEditResponseDNSRecordsURIRecordTTLNumber = 1
)

// Union satisfied by [DNSRecordGetResponseDNSRecordsARecord],
// [DNSRecordGetResponseDNSRecordsAAAARecord],
// [DNSRecordGetResponseDNSRecordsCAARecord],
// [DNSRecordGetResponseDNSRecordsCertRecord],
// [DNSRecordGetResponseDNSRecordsCNAMERecord],
// [DNSRecordGetResponseDNSRecordsDNSKEYRecord],
// [DNSRecordGetResponseDNSRecordsDSRecord],
// [DNSRecordGetResponseDNSRecordsHTTPSRecord],
// [DNSRecordGetResponseDNSRecordsLOCRecord],
// [DNSRecordGetResponseDNSRecordsMXRecord],
// [DNSRecordGetResponseDNSRecordsNAPTRRecord],
// [DNSRecordGetResponseDNSRecordsNSRecord],
// [DNSRecordGetResponseDNSRecordsPTRRecord],
// [DNSRecordGetResponseDNSRecordsSmimeaRecord],
// [DNSRecordGetResponseDNSRecordsSRVRecord],
// [DNSRecordGetResponseDNSRecordsSSHFPRecord],
// [DNSRecordGetResponseDNSRecordsSVCBRecord],
// [DNSRecordGetResponseDNSRecordsTLSARecord],
// [DNSRecordGetResponseDNSRecordsTXTRecord] or
// [DNSRecordGetResponseDNSRecordsURIRecord].
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

type DNSRecordGetResponseDNSRecordsAAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsAAAARecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsAAAARecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseDNSRecordsAAAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsAAAARecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsAAAARecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsAAAARecord]
type dnsRecordGetResponseDNSRecordsAAAARecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsAAAARecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsAAAARecordType string

const (
	DNSRecordGetResponseDNSRecordsAAAARecordTypeAAAA DNSRecordGetResponseDNSRecordsAAAARecordType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsAAAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsAAAARecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsAAAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsAAAARecordMeta]
type dnsRecordGetResponseDNSRecordsAAAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsAAAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsAAAARecordTTLNumber].
type DNSRecordGetResponseDNSRecordsAAAARecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsAAAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsAAAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsAAAARecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsAAAARecordTTLNumber1 DNSRecordGetResponseDNSRecordsAAAARecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsCAARecord struct {
	// Components of a CAA record.
	Data DNSRecordGetResponseDNSRecordsCAARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsCAARecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsCAARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsCAARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsCAARecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCAARecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsCAARecord]
type dnsRecordGetResponseDNSRecordsCAARecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsCAARecord) implementsDNSRecordGetResponse() {}

// Components of a CAA record.
type DNSRecordGetResponseDNSRecordsCAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                                          `json:"value"`
	JSON  dnsRecordGetResponseDNSRecordsCAARecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCAARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCAARecordData]
type dnsRecordGetResponseDNSRecordsCAARecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsCAARecordType string

const (
	DNSRecordGetResponseDNSRecordsCAARecordTypeCAA DNSRecordGetResponseDNSRecordsCAARecordType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsCAARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsCAARecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCAARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCAARecordMeta]
type dnsRecordGetResponseDNSRecordsCAARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCAARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsCAARecordTTLNumber].
type DNSRecordGetResponseDNSRecordsCAARecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsCAARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsCAARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsCAARecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsCAARecordTTLNumber1 DNSRecordGetResponseDNSRecordsCAARecordTTLNumber = 1
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

type DNSRecordGetResponseDNSRecordsCNAMERecord struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsCNAMERecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsCNAMERecordMeta `json:"meta"`
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
	TTL DNSRecordGetResponseDNSRecordsCNAMERecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsCNAMERecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCNAMERecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsCNAMERecord]
type dnsRecordGetResponseDNSRecordsCNAMERecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsCNAMERecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsCNAMERecordType string

const (
	DNSRecordGetResponseDNSRecordsCNAMERecordTypeCNAME DNSRecordGetResponseDNSRecordsCNAMERecordType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsCNAMERecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsCNAMERecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsCNAMERecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsCNAMERecordMeta]
type dnsRecordGetResponseDNSRecordsCNAMERecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsCNAMERecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsCNAMERecordTTLNumber].
type DNSRecordGetResponseDNSRecordsCNAMERecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsCNAMERecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsCNAMERecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsCNAMERecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsCNAMERecordTTLNumber1 DNSRecordGetResponseDNSRecordsCNAMERecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsDNSKEYRecord struct {
	// Components of a DNSKEY record.
	Data DNSRecordGetResponseDNSRecordsDNSKEYRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsDNSKEYRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsDNSKEYRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsDNSKEYRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsDNSKEYRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDNSKEYRecordJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsDNSKEYRecord]
type dnsRecordGetResponseDNSRecordsDNSKEYRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsDNSKEYRecord) implementsDNSRecordGetResponse() {}

// Components of a DNSKEY record.
type DNSRecordGetResponseDNSRecordsDNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                             `json:"public_key"`
	JSON      dnsRecordGetResponseDNSRecordsDNSKEYRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDNSKEYRecordDataJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseDNSRecordsDNSKEYRecordData]
type dnsRecordGetResponseDNSRecordsDNSKEYRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsDNSKEYRecordType string

const (
	DNSRecordGetResponseDNSRecordsDNSKEYRecordTypeDNSKEY DNSRecordGetResponseDNSRecordsDNSKEYRecordType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsDNSKEYRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                             `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsDNSKEYRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDNSKEYRecordMetaJSON contains the JSON metadata
// for the struct [DNSRecordGetResponseDNSRecordsDNSKEYRecordMeta]
type dnsRecordGetResponseDNSRecordsDNSKEYRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDNSKEYRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsDNSKEYRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsDNSKEYRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsDNSKEYRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsDNSKEYRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsDNSKEYRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsDNSKEYRecordTTLNumber1 DNSRecordGetResponseDNSRecordsDNSKEYRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsDSRecord struct {
	// Components of a DS record.
	Data DNSRecordGetResponseDNSRecordsDSRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsDSRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsDSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsDSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsDSRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDSRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsDSRecord]
type dnsRecordGetResponseDNSRecordsDSRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsDSRecord) implementsDNSRecordGetResponse() {}

// Components of a DS record.
type DNSRecordGetResponseDNSRecordsDSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                                        `json:"key_tag"`
	JSON   dnsRecordGetResponseDNSRecordsDSRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDSRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsDSRecordData]
type dnsRecordGetResponseDNSRecordsDSRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsDSRecordType string

const (
	DNSRecordGetResponseDNSRecordsDSRecordTypeDS DNSRecordGetResponseDNSRecordsDSRecordType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsDSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsDSRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsDSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsDSRecordMeta]
type dnsRecordGetResponseDNSRecordsDSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsDSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsDSRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsDSRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsDSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsDSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsDSRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsDSRecordTTLNumber1 DNSRecordGetResponseDNSRecordsDSRecordTTLNumber = 1
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

type DNSRecordGetResponseDNSRecordsLOCRecord struct {
	// Components of a LOC record.
	Data DNSRecordGetResponseDNSRecordsLOCRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsLOCRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsLOCRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsLOCRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsLOCRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsLOCRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsLOCRecord]
type dnsRecordGetResponseDNSRecordsLOCRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsLOCRecord) implementsDNSRecordGetResponse() {}

// Components of a LOC record.
type DNSRecordGetResponseDNSRecordsLOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordGetResponseDNSRecordsLOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordGetResponseDNSRecordsLOCRecordDataLongDirection `json:"long_direction"`
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
	JSON dnsRecordGetResponseDNSRecordsLOCRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsLOCRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsLOCRecordData]
type dnsRecordGetResponseDNSRecordsLOCRecordDataJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsLOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Latitude direction.
type DNSRecordGetResponseDNSRecordsLOCRecordDataLatDirection string

const (
	DNSRecordGetResponseDNSRecordsLOCRecordDataLatDirectionN DNSRecordGetResponseDNSRecordsLOCRecordDataLatDirection = "N"
	DNSRecordGetResponseDNSRecordsLOCRecordDataLatDirectionS DNSRecordGetResponseDNSRecordsLOCRecordDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordGetResponseDNSRecordsLOCRecordDataLongDirection string

const (
	DNSRecordGetResponseDNSRecordsLOCRecordDataLongDirectionE DNSRecordGetResponseDNSRecordsLOCRecordDataLongDirection = "E"
	DNSRecordGetResponseDNSRecordsLOCRecordDataLongDirectionW DNSRecordGetResponseDNSRecordsLOCRecordDataLongDirection = "W"
)

// Record type.
type DNSRecordGetResponseDNSRecordsLOCRecordType string

const (
	DNSRecordGetResponseDNSRecordsLOCRecordTypeLOC DNSRecordGetResponseDNSRecordsLOCRecordType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsLOCRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsLOCRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsLOCRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsLOCRecordMeta]
type dnsRecordGetResponseDNSRecordsLOCRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsLOCRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsLOCRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsLOCRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsLOCRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsLOCRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsLOCRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsLOCRecordTTLNumber1 DNSRecordGetResponseDNSRecordsLOCRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsMXRecord struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsMXRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsMXRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsMXRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsMXRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsMXRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsMXRecord]
type dnsRecordGetResponseDNSRecordsMXRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsMXRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsMXRecordType string

const (
	DNSRecordGetResponseDNSRecordsMXRecordTypeMX DNSRecordGetResponseDNSRecordsMXRecordType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsMXRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsMXRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsMXRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsMXRecordMeta]
type dnsRecordGetResponseDNSRecordsMXRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsMXRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsMXRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsMXRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsMXRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsMXRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsMXRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsMXRecordTTLNumber1 DNSRecordGetResponseDNSRecordsMXRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsNAPTRRecord struct {
	// Components of a NAPTR record.
	Data DNSRecordGetResponseDNSRecordsNAPTRRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsNAPTRRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsNAPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsNAPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsNAPTRRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNAPTRRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsNAPTRRecord]
type dnsRecordGetResponseDNSRecordsNAPTRRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsNAPTRRecord) implementsDNSRecordGetResponse() {}

// Components of a NAPTR record.
type DNSRecordGetResponseDNSRecordsNAPTRRecordData struct {
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
	JSON    dnsRecordGetResponseDNSRecordsNAPTRRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNAPTRRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsNAPTRRecordData]
type dnsRecordGetResponseDNSRecordsNAPTRRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsNAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsNAPTRRecordType string

const (
	DNSRecordGetResponseDNSRecordsNAPTRRecordTypeNAPTR DNSRecordGetResponseDNSRecordsNAPTRRecordType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsNAPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsNAPTRRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNAPTRRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsNAPTRRecordMeta]
type dnsRecordGetResponseDNSRecordsNAPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsNAPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsNAPTRRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsNAPTRRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsNAPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsNAPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsNAPTRRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsNAPTRRecordTTLNumber1 DNSRecordGetResponseDNSRecordsNAPTRRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsNSRecord struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsNSRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsNSRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsNSRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsNSRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNSRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsNSRecord]
type dnsRecordGetResponseDNSRecordsNSRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsNSRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsNSRecordType string

const (
	DNSRecordGetResponseDNSRecordsNSRecordTypeNS DNSRecordGetResponseDNSRecordsNSRecordType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsNSRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                         `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsNSRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsNSRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsNSRecordMeta]
type dnsRecordGetResponseDNSRecordsNSRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsNSRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsNSRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsNSRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsNSRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsNSRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsNSRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsNSRecordTTLNumber1 DNSRecordGetResponseDNSRecordsNSRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsPTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsPTRRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsPTRRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsPTRRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsPTRRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsPTRRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsPTRRecord]
type dnsRecordGetResponseDNSRecordsPTRRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsPTRRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsPTRRecordType string

const (
	DNSRecordGetResponseDNSRecordsPTRRecordTypePTR DNSRecordGetResponseDNSRecordsPTRRecordType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsPTRRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsPTRRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsPTRRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsPTRRecordMeta]
type dnsRecordGetResponseDNSRecordsPTRRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsPTRRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsPTRRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsPTRRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsPTRRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsPTRRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsPTRRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsPTRRecordTTLNumber1 DNSRecordGetResponseDNSRecordsPTRRecordTTLNumber = 1
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

type DNSRecordGetResponseDNSRecordsSRVRecord struct {
	// Components of a SRV record.
	Data DNSRecordGetResponseDNSRecordsSRVRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSRVRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSRVRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSRVRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSRVRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSRVRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsSRVRecord]
type dnsRecordGetResponseDNSRecordsSRVRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSRVRecord) implementsDNSRecordGetResponse() {}

// Components of a SRV record.
type DNSRecordGetResponseDNSRecordsSRVRecordData struct {
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
	JSON   dnsRecordGetResponseDNSRecordsSRVRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSRVRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSRVRecordData]
type dnsRecordGetResponseDNSRecordsSRVRecordDataJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSRVRecordType string

const (
	DNSRecordGetResponseDNSRecordsSRVRecordTypeSRV DNSRecordGetResponseDNSRecordsSRVRecordType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSRVRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSRVRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSRVRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSRVRecordMeta]
type dnsRecordGetResponseDNSRecordsSRVRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSRVRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSRVRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSRVRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSRVRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSRVRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSRVRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSRVRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSRVRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsSSHFPRecord struct {
	// Components of a SSHFP record.
	Data DNSRecordGetResponseDNSRecordsSSHFPRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSSHFPRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSSHFPRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSSHFPRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSSHFPRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSSHFPRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsSSHFPRecord]
type dnsRecordGetResponseDNSRecordsSSHFPRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSSHFPRecord) implementsDNSRecordGetResponse() {}

// Components of a SSHFP record.
type DNSRecordGetResponseDNSRecordsSSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                                           `json:"type"`
	JSON dnsRecordGetResponseDNSRecordsSSHFPRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSSHFPRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSSHFPRecordData]
type dnsRecordGetResponseDNSRecordsSSHFPRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSSHFPRecordType string

const (
	DNSRecordGetResponseDNSRecordsSSHFPRecordTypeSSHFP DNSRecordGetResponseDNSRecordsSSHFPRecordType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSSHFPRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                            `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSSHFPRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSSHFPRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSSHFPRecordMeta]
type dnsRecordGetResponseDNSRecordsSSHFPRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSSHFPRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSSHFPRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSSHFPRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSSHFPRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSSHFPRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSSHFPRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSSHFPRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSSHFPRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsSVCBRecord struct {
	// Components of a SVCB record.
	Data DNSRecordGetResponseDNSRecordsSVCBRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsSVCBRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsSVCBRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsSVCBRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsSVCBRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSVCBRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsSVCBRecord]
type dnsRecordGetResponseDNSRecordsSVCBRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsSVCBRecord) implementsDNSRecordGetResponse() {}

// Components of a SVCB record.
type DNSRecordGetResponseDNSRecordsSVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                                           `json:"value"`
	JSON  dnsRecordGetResponseDNSRecordsSVCBRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSVCBRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSVCBRecordData]
type dnsRecordGetResponseDNSRecordsSVCBRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsSVCBRecordType string

const (
	DNSRecordGetResponseDNSRecordsSVCBRecordTypeSVCB DNSRecordGetResponseDNSRecordsSVCBRecordType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsSVCBRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsSVCBRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsSVCBRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsSVCBRecordMeta]
type dnsRecordGetResponseDNSRecordsSVCBRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsSVCBRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsSVCBRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsSVCBRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsSVCBRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsSVCBRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsSVCBRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsSVCBRecordTTLNumber1 DNSRecordGetResponseDNSRecordsSVCBRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsTLSARecord struct {
	// Components of a TLSA record.
	Data DNSRecordGetResponseDNSRecordsTLSARecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsTLSARecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsTLSARecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsTLSARecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsTLSARecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTLSARecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsTLSARecord]
type dnsRecordGetResponseDNSRecordsTLSARecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsTLSARecord) implementsDNSRecordGetResponse() {}

// Components of a TLSA record.
type DNSRecordGetResponseDNSRecordsTLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                                          `json:"usage"`
	JSON  dnsRecordGetResponseDNSRecordsTLSARecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTLSARecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsTLSARecordData]
type dnsRecordGetResponseDNSRecordsTLSARecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsTLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsTLSARecordType string

const (
	DNSRecordGetResponseDNSRecordsTLSARecordTypeTLSA DNSRecordGetResponseDNSRecordsTLSARecordType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsTLSARecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                           `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsTLSARecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTLSARecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsTLSARecordMeta]
type dnsRecordGetResponseDNSRecordsTLSARecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsTLSARecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsTLSARecordTTLNumber].
type DNSRecordGetResponseDNSRecordsTLSARecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsTLSARecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsTLSARecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsTLSARecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsTLSARecordTTLNumber1 DNSRecordGetResponseDNSRecordsTLSARecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsTXTRecord struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsTXTRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsTXTRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsTXTRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsTXTRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTXTRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsTXTRecord]
type dnsRecordGetResponseDNSRecordsTXTRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsTXTRecord) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseDNSRecordsTXTRecordType string

const (
	DNSRecordGetResponseDNSRecordsTXTRecordTypeTXT DNSRecordGetResponseDNSRecordsTXTRecordType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsTXTRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsTXTRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsTXTRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsTXTRecordMeta]
type dnsRecordGetResponseDNSRecordsTXTRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsTXTRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsTXTRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsTXTRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsTXTRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsTXTRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsTXTRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsTXTRecordTTLNumber1 DNSRecordGetResponseDNSRecordsTXTRecordTTLNumber = 1
)

type DNSRecordGetResponseDNSRecordsURIRecord struct {
	// Components of a URI record.
	Data DNSRecordGetResponseDNSRecordsURIRecordData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseDNSRecordsURIRecordType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSRecordsURIRecordMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSRecordsURIRecordTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSRecordsURIRecordJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsURIRecordJSON contains the JSON metadata for the
// struct [DNSRecordGetResponseDNSRecordsURIRecord]
type dnsRecordGetResponseDNSRecordsURIRecordJSON struct {
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

func (r *DNSRecordGetResponseDNSRecordsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DNSRecordGetResponseDNSRecordsURIRecord) implementsDNSRecordGetResponse() {}

// Components of a URI record.
type DNSRecordGetResponseDNSRecordsURIRecordData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                                         `json:"weight"`
	JSON   dnsRecordGetResponseDNSRecordsURIRecordDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsURIRecordDataJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsURIRecordData]
type dnsRecordGetResponseDNSRecordsURIRecordDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsURIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Record type.
type DNSRecordGetResponseDNSRecordsURIRecordType string

const (
	DNSRecordGetResponseDNSRecordsURIRecordTypeURI DNSRecordGetResponseDNSRecordsURIRecordType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSRecordsURIRecordMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                          `json:"source"`
	JSON   dnsRecordGetResponseDNSRecordsURIRecordMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSRecordsURIRecordMetaJSON contains the JSON metadata for
// the struct [DNSRecordGetResponseDNSRecordsURIRecordMeta]
type dnsRecordGetResponseDNSRecordsURIRecordMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSRecordsURIRecordMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordGetResponseDNSRecordsURIRecordTTLNumber].
type DNSRecordGetResponseDNSRecordsURIRecordTTL interface {
	ImplementsDNSRecordGetResponseDNSRecordsURIRecordTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSRecordsURIRecordTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type DNSRecordGetResponseDNSRecordsURIRecordTTLNumber float64

const (
	DNSRecordGetResponseDNSRecordsURIRecordTTLNumber1 DNSRecordGetResponseDNSRecordsURIRecordTTLNumber = 1
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
	DNSRecordNewParamsTypeURI DNSRecordNewParamsType = "URI"
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
	DNSRecordUpdateParamsTypeURI DNSRecordUpdateParamsType = "URI"
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
	DNSRecordListParamsTypeAAAA   DNSRecordListParamsType = "AAAA"
	DNSRecordListParamsTypeCAA    DNSRecordListParamsType = "CAA"
	DNSRecordListParamsTypeCert   DNSRecordListParamsType = "CERT"
	DNSRecordListParamsTypeCNAME  DNSRecordListParamsType = "CNAME"
	DNSRecordListParamsTypeDNSKEY DNSRecordListParamsType = "DNSKEY"
	DNSRecordListParamsTypeDS     DNSRecordListParamsType = "DS"
	DNSRecordListParamsTypeHTTPS  DNSRecordListParamsType = "HTTPS"
	DNSRecordListParamsTypeLOC    DNSRecordListParamsType = "LOC"
	DNSRecordListParamsTypeMX     DNSRecordListParamsType = "MX"
	DNSRecordListParamsTypeNAPTR  DNSRecordListParamsType = "NAPTR"
	DNSRecordListParamsTypeNS     DNSRecordListParamsType = "NS"
	DNSRecordListParamsTypePTR    DNSRecordListParamsType = "PTR"
	DNSRecordListParamsTypeSmimea DNSRecordListParamsType = "SMIMEA"
	DNSRecordListParamsTypeSRV    DNSRecordListParamsType = "SRV"
	DNSRecordListParamsTypeSSHFP  DNSRecordListParamsType = "SSHFP"
	DNSRecordListParamsTypeSVCB   DNSRecordListParamsType = "SVCB"
	DNSRecordListParamsTypeTLSA   DNSRecordListParamsType = "TLSA"
	DNSRecordListParamsTypeTXT    DNSRecordListParamsType = "TXT"
	DNSRecordListParamsTypeURI    DNSRecordListParamsType = "URI"
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
	DNSRecordEditParamsTypeURI DNSRecordEditParamsType = "URI"
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
	Result   DNSRecordEditResponse                   `json:"result,required"`
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
