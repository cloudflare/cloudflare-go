// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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

// Union satisfied by [DNSRecordNewResponseA], [DNSRecordNewResponseAAAA],
// [DNSRecordNewResponseCAA], [DNSRecordNewResponseCert],
// [DNSRecordNewResponseCNAME], [DNSRecordNewResponseDNSKEY],
// [DNSRecordNewResponseDS], [DNSRecordNewResponseHTTPS],
// [DNSRecordNewResponseLOC], [DNSRecordNewResponseMX],
// [DNSRecordNewResponseNAPTR], [DNSRecordNewResponseNS],
// [DNSRecordNewResponsePTR], [DNSRecordNewResponseSmimea],
// [DNSRecordNewResponseSRV], [DNSRecordNewResponseSSHFP],
// [DNSRecordNewResponseSVCB], [DNSRecordNewResponseTLSA],
// [DNSRecordNewResponseTXT] or [DNSRecordNewResponseURI].
type DNSRecordNewResponse interface {
	implementsDNSRecordNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordNewResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordNewResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseAType `json:"type,required"`
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
	Meta DNSRecordNewResponseAMeta `json:"meta"`
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
	TTL DNSRecordNewResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseAJSON `json:"-"`
}

// dnsRecordNewResponseAJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseA]
type dnsRecordNewResponseAJSON struct {
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

func (r *DNSRecordNewResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseA) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseAType string

const (
	DNSRecordNewResponseATypeA DNSRecordNewResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   dnsRecordNewResponseAMetaJSON `json:"-"`
}

// dnsRecordNewResponseAMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseAMeta]
type dnsRecordNewResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseATTLNumber].
type DNSRecordNewResponseATTL interface {
	ImplementsDNSRecordNewResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseATTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseATTLNumber float64

const (
	DNSRecordNewResponseATTLNumber1 DNSRecordNewResponseATTLNumber = 1
)

type DNSRecordNewResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseAAAAType `json:"type,required"`
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
	Meta DNSRecordNewResponseAAAAMeta `json:"meta"`
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
	TTL DNSRecordNewResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseAAAAJSON `json:"-"`
}

// dnsRecordNewResponseAAAAJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseAAAA]
type dnsRecordNewResponseAAAAJSON struct {
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

func (r *DNSRecordNewResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseAAAA) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseAAAAType string

const (
	DNSRecordNewResponseAAAATypeAAAA DNSRecordNewResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordNewResponseAAAAMetaJSON `json:"-"`
}

// dnsRecordNewResponseAAAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseAAAAMeta]
type dnsRecordNewResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseAAAATTLNumber].
type DNSRecordNewResponseAAAATTL interface {
	ImplementsDNSRecordNewResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseAAAATTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseAAAATTLNumber float64

const (
	DNSRecordNewResponseAAAATTLNumber1 DNSRecordNewResponseAAAATTLNumber = 1
)

type DNSRecordNewResponseCAA struct {
	// Components of a CAA record.
	Data DNSRecordNewResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseCAAType `json:"type,required"`
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
	Meta DNSRecordNewResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseCAAJSON `json:"-"`
}

// dnsRecordNewResponseCAAJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCAA]
type dnsRecordNewResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseCAA) implementsDNSRecordNewResponse() {}

// Components of a CAA record.
type DNSRecordNewResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                          `json:"value"`
	JSON  dnsRecordNewResponseCAADataJSON `json:"-"`
}

// dnsRecordNewResponseCAADataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCAAData]
type dnsRecordNewResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseCAAType string

const (
	DNSRecordNewResponseCAATypeCAA DNSRecordNewResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordNewResponseCAAMetaJSON `json:"-"`
}

// dnsRecordNewResponseCAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCAAMeta]
type dnsRecordNewResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseCAATTLNumber].
type DNSRecordNewResponseCAATTL interface {
	ImplementsDNSRecordNewResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseCAATTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseCAATTLNumber float64

const (
	DNSRecordNewResponseCAATTLNumber1 DNSRecordNewResponseCAATTLNumber = 1
)

type DNSRecordNewResponseCert struct {
	// Components of a CERT record.
	Data DNSRecordNewResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseCertType `json:"type,required"`
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
	Meta DNSRecordNewResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseCertJSON `json:"-"`
}

// dnsRecordNewResponseCertJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCert]
type dnsRecordNewResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseCert) implementsDNSRecordNewResponse() {}

// Components of a CERT record.
type DNSRecordNewResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                          `json:"type"`
	JSON dnsRecordNewResponseCertDataJSON `json:"-"`
}

// dnsRecordNewResponseCertDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCertData]
type dnsRecordNewResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseCertType string

const (
	DNSRecordNewResponseCertTypeCert DNSRecordNewResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordNewResponseCertMetaJSON `json:"-"`
}

// dnsRecordNewResponseCertMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCertMeta]
type dnsRecordNewResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseCertTTLNumber].
type DNSRecordNewResponseCertTTL interface {
	ImplementsDNSRecordNewResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseCertTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseCertTTLNumber float64

const (
	DNSRecordNewResponseCertTTLNumber1 DNSRecordNewResponseCertTTLNumber = 1
)

type DNSRecordNewResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseCNAMEType `json:"type,required"`
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
	Meta DNSRecordNewResponseCNAMEMeta `json:"meta"`
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
	TTL DNSRecordNewResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseCNAMEJSON `json:"-"`
}

// dnsRecordNewResponseCNAMEJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCNAME]
type dnsRecordNewResponseCNAMEJSON struct {
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

func (r *DNSRecordNewResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseCNAME) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseCNAMEType string

const (
	DNSRecordNewResponseCNAMETypeCNAME DNSRecordNewResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordNewResponseCNAMEMetaJSON `json:"-"`
}

// dnsRecordNewResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseCNAMEMeta]
type dnsRecordNewResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseCNAMETTLNumber].
type DNSRecordNewResponseCNAMETTL interface {
	ImplementsDNSRecordNewResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseCNAMETTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseCNAMETTLNumber float64

const (
	DNSRecordNewResponseCNAMETTLNumber1 DNSRecordNewResponseCNAMETTLNumber = 1
)

type DNSRecordNewResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data DNSRecordNewResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDNSKEYType `json:"type,required"`
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
	Meta DNSRecordNewResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDNSKEYJSON `json:"-"`
}

// dnsRecordNewResponseDNSKEYJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseDNSKEY]
type dnsRecordNewResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseDNSKEY) implementsDNSRecordNewResponse() {}

// Components of a DNSKEY record.
type DNSRecordNewResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                             `json:"public_key"`
	JSON      dnsRecordNewResponseDNSKEYDataJSON `json:"-"`
}

// dnsRecordNewResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseDNSKEYData]
type dnsRecordNewResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseDNSKEYType string

const (
	DNSRecordNewResponseDNSKEYTypeDNSKEY DNSRecordNewResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordNewResponseDNSKEYMetaJSON `json:"-"`
}

// dnsRecordNewResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseDNSKEYMeta]
type dnsRecordNewResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseDNSKEYTTLNumber].
type DNSRecordNewResponseDNSKEYTTL interface {
	ImplementsDNSRecordNewResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseDNSKEYTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseDNSKEYTTLNumber float64

const (
	DNSRecordNewResponseDNSKEYTTLNumber1 DNSRecordNewResponseDNSKEYTTLNumber = 1
)

type DNSRecordNewResponseDS struct {
	// Components of a DS record.
	Data DNSRecordNewResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseDSType `json:"type,required"`
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
	Meta DNSRecordNewResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseDSJSON `json:"-"`
}

// dnsRecordNewResponseDSJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseDS]
type dnsRecordNewResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseDS) implementsDNSRecordNewResponse() {}

// Components of a DS record.
type DNSRecordNewResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                        `json:"key_tag"`
	JSON   dnsRecordNewResponseDSDataJSON `json:"-"`
}

// dnsRecordNewResponseDSDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseDSData]
type dnsRecordNewResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseDSType string

const (
	DNSRecordNewResponseDSTypeDS DNSRecordNewResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordNewResponseDSMetaJSON `json:"-"`
}

// dnsRecordNewResponseDSMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseDSMeta]
type dnsRecordNewResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseDSTTLNumber].
type DNSRecordNewResponseDSTTL interface {
	ImplementsDNSRecordNewResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseDSTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseDSTTLNumber float64

const (
	DNSRecordNewResponseDSTTLNumber1 DNSRecordNewResponseDSTTLNumber = 1
)

type DNSRecordNewResponseHTTPS struct {
	// Components of a HTTPS record.
	Data DNSRecordNewResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseHTTPSType `json:"type,required"`
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
	Meta DNSRecordNewResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseHTTPSJSON `json:"-"`
}

// dnsRecordNewResponseHTTPSJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseHTTPS]
type dnsRecordNewResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseHTTPS) implementsDNSRecordNewResponse() {}

// Components of a HTTPS record.
type DNSRecordNewResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                            `json:"value"`
	JSON  dnsRecordNewResponseHTTPSDataJSON `json:"-"`
}

// dnsRecordNewResponseHTTPSDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseHTTPSData]
type dnsRecordNewResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseHTTPSType string

const (
	DNSRecordNewResponseHTTPSTypeHTTPS DNSRecordNewResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordNewResponseHTTPSMetaJSON `json:"-"`
}

// dnsRecordNewResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseHTTPSMeta]
type dnsRecordNewResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseHTTPSTTLNumber].
type DNSRecordNewResponseHTTPSTTL interface {
	ImplementsDNSRecordNewResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseHTTPSTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseHTTPSTTLNumber float64

const (
	DNSRecordNewResponseHTTPSTTLNumber1 DNSRecordNewResponseHTTPSTTLNumber = 1
)

type DNSRecordNewResponseLOC struct {
	// Components of a LOC record.
	Data DNSRecordNewResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseLOCType `json:"type,required"`
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
	Meta DNSRecordNewResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseLOCJSON `json:"-"`
}

// dnsRecordNewResponseLOCJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseLOC]
type dnsRecordNewResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseLOC) implementsDNSRecordNewResponse() {}

// Components of a LOC record.
type DNSRecordNewResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordNewResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordNewResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                         `json:"size"`
	JSON dnsRecordNewResponseLOCDataJSON `json:"-"`
}

// dnsRecordNewResponseLOCDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseLOCData]
type dnsRecordNewResponseLOCDataJSON struct {
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

func (r *DNSRecordNewResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordNewResponseLOCDataLatDirection string

const (
	DNSRecordNewResponseLOCDataLatDirectionN DNSRecordNewResponseLOCDataLatDirection = "N"
	DNSRecordNewResponseLOCDataLatDirectionS DNSRecordNewResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordNewResponseLOCDataLongDirection string

const (
	DNSRecordNewResponseLOCDataLongDirectionE DNSRecordNewResponseLOCDataLongDirection = "E"
	DNSRecordNewResponseLOCDataLongDirectionW DNSRecordNewResponseLOCDataLongDirection = "W"
)

// Record type.
type DNSRecordNewResponseLOCType string

const (
	DNSRecordNewResponseLOCTypeLOC DNSRecordNewResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordNewResponseLOCMetaJSON `json:"-"`
}

// dnsRecordNewResponseLOCMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseLOCMeta]
type dnsRecordNewResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseLOCTTLNumber].
type DNSRecordNewResponseLOCTTL interface {
	ImplementsDNSRecordNewResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseLOCTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseLOCTTLNumber float64

const (
	DNSRecordNewResponseLOCTTLNumber1 DNSRecordNewResponseLOCTTLNumber = 1
)

type DNSRecordNewResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordNewResponseMXType `json:"type,required"`
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
	Meta DNSRecordNewResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseMXJSON `json:"-"`
}

// dnsRecordNewResponseMXJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseMX]
type dnsRecordNewResponseMXJSON struct {
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

func (r *DNSRecordNewResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseMX) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseMXType string

const (
	DNSRecordNewResponseMXTypeMX DNSRecordNewResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordNewResponseMXMetaJSON `json:"-"`
}

// dnsRecordNewResponseMXMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseMXMeta]
type dnsRecordNewResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseMXTTLNumber].
type DNSRecordNewResponseMXTTL interface {
	ImplementsDNSRecordNewResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseMXTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseMXTTLNumber float64

const (
	DNSRecordNewResponseMXTTLNumber1 DNSRecordNewResponseMXTTLNumber = 1
)

type DNSRecordNewResponseNAPTR struct {
	// Components of a NAPTR record.
	Data DNSRecordNewResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseNAPTRType `json:"type,required"`
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
	Meta DNSRecordNewResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseNAPTRJSON `json:"-"`
}

// dnsRecordNewResponseNAPTRJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseNAPTR]
type dnsRecordNewResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseNAPTR) implementsDNSRecordNewResponse() {}

// Components of a NAPTR record.
type DNSRecordNewResponseNAPTRData struct {
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
	Service string                            `json:"service"`
	JSON    dnsRecordNewResponseNAPTRDataJSON `json:"-"`
}

// dnsRecordNewResponseNAPTRDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseNAPTRData]
type dnsRecordNewResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseNAPTRType string

const (
	DNSRecordNewResponseNAPTRTypeNAPTR DNSRecordNewResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordNewResponseNAPTRMetaJSON `json:"-"`
}

// dnsRecordNewResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseNAPTRMeta]
type dnsRecordNewResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseNAPTRTTLNumber].
type DNSRecordNewResponseNAPTRTTL interface {
	ImplementsDNSRecordNewResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseNAPTRTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseNAPTRTTLNumber float64

const (
	DNSRecordNewResponseNAPTRTTLNumber1 DNSRecordNewResponseNAPTRTTLNumber = 1
)

type DNSRecordNewResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseNSType `json:"type,required"`
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
	Meta DNSRecordNewResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseNSJSON `json:"-"`
}

// dnsRecordNewResponseNSJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseNS]
type dnsRecordNewResponseNSJSON struct {
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

func (r *DNSRecordNewResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseNS) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseNSType string

const (
	DNSRecordNewResponseNSTypeNS DNSRecordNewResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordNewResponseNSMetaJSON `json:"-"`
}

// dnsRecordNewResponseNSMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseNSMeta]
type dnsRecordNewResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseNSTTLNumber].
type DNSRecordNewResponseNSTTL interface {
	ImplementsDNSRecordNewResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseNSTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseNSTTLNumber float64

const (
	DNSRecordNewResponseNSTTLNumber1 DNSRecordNewResponseNSTTLNumber = 1
)

type DNSRecordNewResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponsePTRType `json:"type,required"`
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
	Meta DNSRecordNewResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponsePTRJSON `json:"-"`
}

// dnsRecordNewResponsePTRJSON contains the JSON metadata for the struct
// [DNSRecordNewResponsePTR]
type dnsRecordNewResponsePTRJSON struct {
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

func (r *DNSRecordNewResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponsePTR) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponsePTRType string

const (
	DNSRecordNewResponsePTRTypePTR DNSRecordNewResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordNewResponsePTRMetaJSON `json:"-"`
}

// dnsRecordNewResponsePTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponsePTRMeta]
type dnsRecordNewResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponsePTRTTLNumber].
type DNSRecordNewResponsePTRTTL interface {
	ImplementsDNSRecordNewResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponsePTRTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponsePTRTTLNumber float64

const (
	DNSRecordNewResponsePTRTTLNumber1 DNSRecordNewResponsePTRTTLNumber = 1
)

type DNSRecordNewResponseSmimea struct {
	// Components of a SMIMEA record.
	Data DNSRecordNewResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseSmimeaType `json:"type,required"`
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
	Meta DNSRecordNewResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseSmimeaJSON `json:"-"`
}

// dnsRecordNewResponseSmimeaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSmimea]
type dnsRecordNewResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseSmimea) implementsDNSRecordNewResponse() {}

// Components of a SMIMEA record.
type DNSRecordNewResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                            `json:"usage"`
	JSON  dnsRecordNewResponseSmimeaDataJSON `json:"-"`
}

// dnsRecordNewResponseSmimeaDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSmimeaData]
type dnsRecordNewResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordNewResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseSmimeaType string

const (
	DNSRecordNewResponseSmimeaTypeSmimea DNSRecordNewResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordNewResponseSmimeaMetaJSON `json:"-"`
}

// dnsRecordNewResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSmimeaMeta]
type dnsRecordNewResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseSmimeaTTLNumber].
type DNSRecordNewResponseSmimeaTTL interface {
	ImplementsDNSRecordNewResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseSmimeaTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseSmimeaTTLNumber float64

const (
	DNSRecordNewResponseSmimeaTTLNumber1 DNSRecordNewResponseSmimeaTTLNumber = 1
)

type DNSRecordNewResponseSRV struct {
	// Components of a SRV record.
	Data DNSRecordNewResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseSRVType `json:"type,required"`
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
	Meta DNSRecordNewResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseSRVJSON `json:"-"`
}

// dnsRecordNewResponseSRVJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSRV]
type dnsRecordNewResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseSRV) implementsDNSRecordNewResponse() {}

// Components of a SRV record.
type DNSRecordNewResponseSRVData struct {
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
	Weight float64                         `json:"weight"`
	JSON   dnsRecordNewResponseSRVDataJSON `json:"-"`
}

// dnsRecordNewResponseSRVDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSRVData]
type dnsRecordNewResponseSRVDataJSON struct {
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

func (r *DNSRecordNewResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseSRVType string

const (
	DNSRecordNewResponseSRVTypeSRV DNSRecordNewResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordNewResponseSRVMetaJSON `json:"-"`
}

// dnsRecordNewResponseSRVMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSRVMeta]
type dnsRecordNewResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseSRVTTLNumber].
type DNSRecordNewResponseSRVTTL interface {
	ImplementsDNSRecordNewResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseSRVTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseSRVTTLNumber float64

const (
	DNSRecordNewResponseSRVTTLNumber1 DNSRecordNewResponseSRVTTLNumber = 1
)

type DNSRecordNewResponseSSHFP struct {
	// Components of a SSHFP record.
	Data DNSRecordNewResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseSSHFPType `json:"type,required"`
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
	Meta DNSRecordNewResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseSSHFPJSON `json:"-"`
}

// dnsRecordNewResponseSSHFPJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSSHFP]
type dnsRecordNewResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseSSHFP) implementsDNSRecordNewResponse() {}

// Components of a SSHFP record.
type DNSRecordNewResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                           `json:"type"`
	JSON dnsRecordNewResponseSSHFPDataJSON `json:"-"`
}

// dnsRecordNewResponseSSHFPDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSSHFPData]
type dnsRecordNewResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseSSHFPType string

const (
	DNSRecordNewResponseSSHFPTypeSSHFP DNSRecordNewResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordNewResponseSSHFPMetaJSON `json:"-"`
}

// dnsRecordNewResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSSHFPMeta]
type dnsRecordNewResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseSSHFPTTLNumber].
type DNSRecordNewResponseSSHFPTTL interface {
	ImplementsDNSRecordNewResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseSSHFPTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseSSHFPTTLNumber float64

const (
	DNSRecordNewResponseSSHFPTTLNumber1 DNSRecordNewResponseSSHFPTTLNumber = 1
)

type DNSRecordNewResponseSVCB struct {
	// Components of a SVCB record.
	Data DNSRecordNewResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseSVCBType `json:"type,required"`
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
	Meta DNSRecordNewResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseSVCBJSON `json:"-"`
}

// dnsRecordNewResponseSVCBJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSVCB]
type dnsRecordNewResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseSVCB) implementsDNSRecordNewResponse() {}

// Components of a SVCB record.
type DNSRecordNewResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                           `json:"value"`
	JSON  dnsRecordNewResponseSVCBDataJSON `json:"-"`
}

// dnsRecordNewResponseSVCBDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSVCBData]
type dnsRecordNewResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseSVCBType string

const (
	DNSRecordNewResponseSVCBTypeSVCB DNSRecordNewResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordNewResponseSVCBMetaJSON `json:"-"`
}

// dnsRecordNewResponseSVCBMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseSVCBMeta]
type dnsRecordNewResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseSVCBTTLNumber].
type DNSRecordNewResponseSVCBTTL interface {
	ImplementsDNSRecordNewResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseSVCBTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseSVCBTTLNumber float64

const (
	DNSRecordNewResponseSVCBTTLNumber1 DNSRecordNewResponseSVCBTTLNumber = 1
)

type DNSRecordNewResponseTLSA struct {
	// Components of a TLSA record.
	Data DNSRecordNewResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseTLSAType `json:"type,required"`
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
	Meta DNSRecordNewResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseTLSAJSON `json:"-"`
}

// dnsRecordNewResponseTLSAJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseTLSA]
type dnsRecordNewResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseTLSA) implementsDNSRecordNewResponse() {}

// Components of a TLSA record.
type DNSRecordNewResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                          `json:"usage"`
	JSON  dnsRecordNewResponseTLSADataJSON `json:"-"`
}

// dnsRecordNewResponseTLSADataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseTLSAData]
type dnsRecordNewResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordNewResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseTLSAType string

const (
	DNSRecordNewResponseTLSATypeTLSA DNSRecordNewResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordNewResponseTLSAMetaJSON `json:"-"`
}

// dnsRecordNewResponseTLSAMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseTLSAMeta]
type dnsRecordNewResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseTLSATTLNumber].
type DNSRecordNewResponseTLSATTL interface {
	ImplementsDNSRecordNewResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseTLSATTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseTLSATTLNumber float64

const (
	DNSRecordNewResponseTLSATTLNumber1 DNSRecordNewResponseTLSATTLNumber = 1
)

type DNSRecordNewResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordNewResponseTXTType `json:"type,required"`
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
	Meta DNSRecordNewResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseTXTJSON `json:"-"`
}

// dnsRecordNewResponseTXTJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseTXT]
type dnsRecordNewResponseTXTJSON struct {
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

func (r *DNSRecordNewResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseTXT) implementsDNSRecordNewResponse() {}

// Record type.
type DNSRecordNewResponseTXTType string

const (
	DNSRecordNewResponseTXTTypeTXT DNSRecordNewResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordNewResponseTXTMetaJSON `json:"-"`
}

// dnsRecordNewResponseTXTMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseTXTMeta]
type dnsRecordNewResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseTXTTTLNumber].
type DNSRecordNewResponseTXTTTL interface {
	ImplementsDNSRecordNewResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseTXTTTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseTXTTTLNumber float64

const (
	DNSRecordNewResponseTXTTTLNumber1 DNSRecordNewResponseTXTTTLNumber = 1
)

type DNSRecordNewResponseURI struct {
	// Components of a URI record.
	Data DNSRecordNewResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordNewResponseURIType `json:"type,required"`
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
	Meta DNSRecordNewResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordNewResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordNewResponseURIJSON `json:"-"`
}

// dnsRecordNewResponseURIJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseURI]
type dnsRecordNewResponseURIJSON struct {
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

func (r *DNSRecordNewResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordNewResponseURI) implementsDNSRecordNewResponse() {}

// Components of a URI record.
type DNSRecordNewResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                         `json:"weight"`
	JSON   dnsRecordNewResponseURIDataJSON `json:"-"`
}

// dnsRecordNewResponseURIDataJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseURIData]
type dnsRecordNewResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordNewResponseURIType string

const (
	DNSRecordNewResponseURITypeURI DNSRecordNewResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordNewResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordNewResponseURIMetaJSON `json:"-"`
}

// dnsRecordNewResponseURIMetaJSON contains the JSON metadata for the struct
// [DNSRecordNewResponseURIMeta]
type dnsRecordNewResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordNewResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordNewResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordNewResponseURITTLNumber].
type DNSRecordNewResponseURITTL interface {
	ImplementsDNSRecordNewResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordNewResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordNewResponseURITTLNumber(0)),
		},
	)
}

type DNSRecordNewResponseURITTLNumber float64

const (
	DNSRecordNewResponseURITTLNumber1 DNSRecordNewResponseURITTLNumber = 1
)

// Union satisfied by [DNSRecordUpdateResponseA], [DNSRecordUpdateResponseAAAA],
// [DNSRecordUpdateResponseCAA], [DNSRecordUpdateResponseCert],
// [DNSRecordUpdateResponseCNAME], [DNSRecordUpdateResponseDNSKEY],
// [DNSRecordUpdateResponseDS], [DNSRecordUpdateResponseHTTPS],
// [DNSRecordUpdateResponseLOC], [DNSRecordUpdateResponseMX],
// [DNSRecordUpdateResponseNAPTR], [DNSRecordUpdateResponseNS],
// [DNSRecordUpdateResponsePTR], [DNSRecordUpdateResponseSmimea],
// [DNSRecordUpdateResponseSRV], [DNSRecordUpdateResponseSSHFP],
// [DNSRecordUpdateResponseSVCB], [DNSRecordUpdateResponseTLSA],
// [DNSRecordUpdateResponseTXT] or [DNSRecordUpdateResponseURI].
type DNSRecordUpdateResponse interface {
	implementsDNSRecordUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordUpdateResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordUpdateResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseAType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseAMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseAJSON `json:"-"`
}

// dnsRecordUpdateResponseAJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseA]
type dnsRecordUpdateResponseAJSON struct {
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

func (r *DNSRecordUpdateResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseA) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseAType string

const (
	DNSRecordUpdateResponseATypeA DNSRecordUpdateResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordUpdateResponseAMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseAMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseAMeta]
type dnsRecordUpdateResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseATTLNumber].
type DNSRecordUpdateResponseATTL interface {
	ImplementsDNSRecordUpdateResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseATTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseATTLNumber float64

const (
	DNSRecordUpdateResponseATTLNumber1 DNSRecordUpdateResponseATTLNumber = 1
)

type DNSRecordUpdateResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseAAAAType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseAAAAMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseAAAAJSON `json:"-"`
}

// dnsRecordUpdateResponseAAAAJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseAAAA]
type dnsRecordUpdateResponseAAAAJSON struct {
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

func (r *DNSRecordUpdateResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseAAAA) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseAAAAType string

const (
	DNSRecordUpdateResponseAAAATypeAAAA DNSRecordUpdateResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordUpdateResponseAAAAMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseAAAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseAAAAMeta]
type dnsRecordUpdateResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseAAAATTLNumber].
type DNSRecordUpdateResponseAAAATTL interface {
	ImplementsDNSRecordUpdateResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseAAAATTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseAAAATTLNumber float64

const (
	DNSRecordUpdateResponseAAAATTLNumber1 DNSRecordUpdateResponseAAAATTLNumber = 1
)

type DNSRecordUpdateResponseCAA struct {
	// Components of a CAA record.
	Data DNSRecordUpdateResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseCAAType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseCAAJSON `json:"-"`
}

// dnsRecordUpdateResponseCAAJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCAA]
type dnsRecordUpdateResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseCAA) implementsDNSRecordUpdateResponse() {}

// Components of a CAA record.
type DNSRecordUpdateResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                             `json:"value"`
	JSON  dnsRecordUpdateResponseCAADataJSON `json:"-"`
}

// dnsRecordUpdateResponseCAADataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCAAData]
type dnsRecordUpdateResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseCAAType string

const (
	DNSRecordUpdateResponseCAATypeCAA DNSRecordUpdateResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordUpdateResponseCAAMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseCAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCAAMeta]
type dnsRecordUpdateResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseCAATTLNumber].
type DNSRecordUpdateResponseCAATTL interface {
	ImplementsDNSRecordUpdateResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseCAATTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseCAATTLNumber float64

const (
	DNSRecordUpdateResponseCAATTLNumber1 DNSRecordUpdateResponseCAATTLNumber = 1
)

type DNSRecordUpdateResponseCert struct {
	// Components of a CERT record.
	Data DNSRecordUpdateResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseCertType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseCertJSON `json:"-"`
}

// dnsRecordUpdateResponseCertJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCert]
type dnsRecordUpdateResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseCert) implementsDNSRecordUpdateResponse() {}

// Components of a CERT record.
type DNSRecordUpdateResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                             `json:"type"`
	JSON dnsRecordUpdateResponseCertDataJSON `json:"-"`
}

// dnsRecordUpdateResponseCertDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCertData]
type dnsRecordUpdateResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseCertType string

const (
	DNSRecordUpdateResponseCertTypeCert DNSRecordUpdateResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordUpdateResponseCertMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseCertMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCertMeta]
type dnsRecordUpdateResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseCertTTLNumber].
type DNSRecordUpdateResponseCertTTL interface {
	ImplementsDNSRecordUpdateResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseCertTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseCertTTLNumber float64

const (
	DNSRecordUpdateResponseCertTTLNumber1 DNSRecordUpdateResponseCertTTLNumber = 1
)

type DNSRecordUpdateResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseCNAMEType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseCNAMEMeta `json:"meta"`
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
	TTL DNSRecordUpdateResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseCNAMEJSON `json:"-"`
}

// dnsRecordUpdateResponseCNAMEJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCNAME]
type dnsRecordUpdateResponseCNAMEJSON struct {
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

func (r *DNSRecordUpdateResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseCNAME) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseCNAMEType string

const (
	DNSRecordUpdateResponseCNAMETypeCNAME DNSRecordUpdateResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                               `json:"source"`
	JSON   dnsRecordUpdateResponseCNAMEMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseCNAMEMeta]
type dnsRecordUpdateResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseCNAMETTLNumber].
type DNSRecordUpdateResponseCNAMETTL interface {
	ImplementsDNSRecordUpdateResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseCNAMETTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseCNAMETTLNumber float64

const (
	DNSRecordUpdateResponseCNAMETTLNumber1 DNSRecordUpdateResponseCNAMETTLNumber = 1
)

type DNSRecordUpdateResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data DNSRecordUpdateResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDNSKEYType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDNSKEYJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSKEYJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseDNSKEY]
type dnsRecordUpdateResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseDNSKEY) implementsDNSRecordUpdateResponse() {}

// Components of a DNSKEY record.
type DNSRecordUpdateResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                                `json:"public_key"`
	JSON      dnsRecordUpdateResponseDNSKEYDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseDNSKEYData]
type dnsRecordUpdateResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseDNSKEYType string

const (
	DNSRecordUpdateResponseDNSKEYTypeDNSKEY DNSRecordUpdateResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                `json:"source"`
	JSON   dnsRecordUpdateResponseDNSKEYMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseDNSKEYMeta]
type dnsRecordUpdateResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseDNSKEYTTLNumber].
type DNSRecordUpdateResponseDNSKEYTTL interface {
	ImplementsDNSRecordUpdateResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseDNSKEYTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseDNSKEYTTLNumber float64

const (
	DNSRecordUpdateResponseDNSKEYTTLNumber1 DNSRecordUpdateResponseDNSKEYTTLNumber = 1
)

type DNSRecordUpdateResponseDS struct {
	// Components of a DS record.
	Data DNSRecordUpdateResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseDSType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseDSJSON `json:"-"`
}

// dnsRecordUpdateResponseDSJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseDS]
type dnsRecordUpdateResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseDS) implementsDNSRecordUpdateResponse() {}

// Components of a DS record.
type DNSRecordUpdateResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                           `json:"key_tag"`
	JSON   dnsRecordUpdateResponseDSDataJSON `json:"-"`
}

// dnsRecordUpdateResponseDSDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseDSData]
type dnsRecordUpdateResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseDSType string

const (
	DNSRecordUpdateResponseDSTypeDS DNSRecordUpdateResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordUpdateResponseDSMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseDSMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseDSMeta]
type dnsRecordUpdateResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseDSTTLNumber].
type DNSRecordUpdateResponseDSTTL interface {
	ImplementsDNSRecordUpdateResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseDSTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseDSTTLNumber float64

const (
	DNSRecordUpdateResponseDSTTLNumber1 DNSRecordUpdateResponseDSTTLNumber = 1
)

type DNSRecordUpdateResponseHTTPS struct {
	// Components of a HTTPS record.
	Data DNSRecordUpdateResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseHTTPSType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseHTTPSJSON `json:"-"`
}

// dnsRecordUpdateResponseHTTPSJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseHTTPS]
type dnsRecordUpdateResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseHTTPS) implementsDNSRecordUpdateResponse() {}

// Components of a HTTPS record.
type DNSRecordUpdateResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                               `json:"value"`
	JSON  dnsRecordUpdateResponseHTTPSDataJSON `json:"-"`
}

// dnsRecordUpdateResponseHTTPSDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseHTTPSData]
type dnsRecordUpdateResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseHTTPSType string

const (
	DNSRecordUpdateResponseHTTPSTypeHTTPS DNSRecordUpdateResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                               `json:"source"`
	JSON   dnsRecordUpdateResponseHTTPSMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseHTTPSMeta]
type dnsRecordUpdateResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseHTTPSTTLNumber].
type DNSRecordUpdateResponseHTTPSTTL interface {
	ImplementsDNSRecordUpdateResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseHTTPSTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseHTTPSTTLNumber float64

const (
	DNSRecordUpdateResponseHTTPSTTLNumber1 DNSRecordUpdateResponseHTTPSTTLNumber = 1
)

type DNSRecordUpdateResponseLOC struct {
	// Components of a LOC record.
	Data DNSRecordUpdateResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseLOCType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseLOCJSON `json:"-"`
}

// dnsRecordUpdateResponseLOCJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseLOC]
type dnsRecordUpdateResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseLOC) implementsDNSRecordUpdateResponse() {}

// Components of a LOC record.
type DNSRecordUpdateResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordUpdateResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordUpdateResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                            `json:"size"`
	JSON dnsRecordUpdateResponseLOCDataJSON `json:"-"`
}

// dnsRecordUpdateResponseLOCDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseLOCData]
type dnsRecordUpdateResponseLOCDataJSON struct {
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

func (r *DNSRecordUpdateResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordUpdateResponseLOCDataLatDirection string

const (
	DNSRecordUpdateResponseLOCDataLatDirectionN DNSRecordUpdateResponseLOCDataLatDirection = "N"
	DNSRecordUpdateResponseLOCDataLatDirectionS DNSRecordUpdateResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordUpdateResponseLOCDataLongDirection string

const (
	DNSRecordUpdateResponseLOCDataLongDirectionE DNSRecordUpdateResponseLOCDataLongDirection = "E"
	DNSRecordUpdateResponseLOCDataLongDirectionW DNSRecordUpdateResponseLOCDataLongDirection = "W"
)

// Record type.
type DNSRecordUpdateResponseLOCType string

const (
	DNSRecordUpdateResponseLOCTypeLOC DNSRecordUpdateResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordUpdateResponseLOCMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseLOCMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseLOCMeta]
type dnsRecordUpdateResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseLOCTTLNumber].
type DNSRecordUpdateResponseLOCTTL interface {
	ImplementsDNSRecordUpdateResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseLOCTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseLOCTTLNumber float64

const (
	DNSRecordUpdateResponseLOCTTLNumber1 DNSRecordUpdateResponseLOCTTLNumber = 1
)

type DNSRecordUpdateResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseMXType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseMXJSON `json:"-"`
}

// dnsRecordUpdateResponseMXJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseMX]
type dnsRecordUpdateResponseMXJSON struct {
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

func (r *DNSRecordUpdateResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseMX) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseMXType string

const (
	DNSRecordUpdateResponseMXTypeMX DNSRecordUpdateResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordUpdateResponseMXMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseMXMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseMXMeta]
type dnsRecordUpdateResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseMXTTLNumber].
type DNSRecordUpdateResponseMXTTL interface {
	ImplementsDNSRecordUpdateResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseMXTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseMXTTLNumber float64

const (
	DNSRecordUpdateResponseMXTTLNumber1 DNSRecordUpdateResponseMXTTLNumber = 1
)

type DNSRecordUpdateResponseNAPTR struct {
	// Components of a NAPTR record.
	Data DNSRecordUpdateResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseNAPTRType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseNAPTRJSON `json:"-"`
}

// dnsRecordUpdateResponseNAPTRJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseNAPTR]
type dnsRecordUpdateResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseNAPTR) implementsDNSRecordUpdateResponse() {}

// Components of a NAPTR record.
type DNSRecordUpdateResponseNAPTRData struct {
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
	Service string                               `json:"service"`
	JSON    dnsRecordUpdateResponseNAPTRDataJSON `json:"-"`
}

// dnsRecordUpdateResponseNAPTRDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseNAPTRData]
type dnsRecordUpdateResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseNAPTRType string

const (
	DNSRecordUpdateResponseNAPTRTypeNAPTR DNSRecordUpdateResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                               `json:"source"`
	JSON   dnsRecordUpdateResponseNAPTRMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseNAPTRMeta]
type dnsRecordUpdateResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseNAPTRTTLNumber].
type DNSRecordUpdateResponseNAPTRTTL interface {
	ImplementsDNSRecordUpdateResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseNAPTRTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseNAPTRTTLNumber float64

const (
	DNSRecordUpdateResponseNAPTRTTLNumber1 DNSRecordUpdateResponseNAPTRTTLNumber = 1
)

type DNSRecordUpdateResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseNSType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseNSJSON `json:"-"`
}

// dnsRecordUpdateResponseNSJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseNS]
type dnsRecordUpdateResponseNSJSON struct {
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

func (r *DNSRecordUpdateResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseNS) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseNSType string

const (
	DNSRecordUpdateResponseNSTypeNS DNSRecordUpdateResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordUpdateResponseNSMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseNSMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseNSMeta]
type dnsRecordUpdateResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseNSTTLNumber].
type DNSRecordUpdateResponseNSTTL interface {
	ImplementsDNSRecordUpdateResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseNSTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseNSTTLNumber float64

const (
	DNSRecordUpdateResponseNSTTLNumber1 DNSRecordUpdateResponseNSTTLNumber = 1
)

type DNSRecordUpdateResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponsePTRType `json:"type,required"`
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
	Meta DNSRecordUpdateResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponsePTRJSON `json:"-"`
}

// dnsRecordUpdateResponsePTRJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponsePTR]
type dnsRecordUpdateResponsePTRJSON struct {
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

func (r *DNSRecordUpdateResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponsePTR) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponsePTRType string

const (
	DNSRecordUpdateResponsePTRTypePTR DNSRecordUpdateResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordUpdateResponsePTRMetaJSON `json:"-"`
}

// dnsRecordUpdateResponsePTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponsePTRMeta]
type dnsRecordUpdateResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponsePTRTTLNumber].
type DNSRecordUpdateResponsePTRTTL interface {
	ImplementsDNSRecordUpdateResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponsePTRTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponsePTRTTLNumber float64

const (
	DNSRecordUpdateResponsePTRTTLNumber1 DNSRecordUpdateResponsePTRTTLNumber = 1
)

type DNSRecordUpdateResponseSmimea struct {
	// Components of a SMIMEA record.
	Data DNSRecordUpdateResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseSmimeaType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                            `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseSmimeaJSON `json:"-"`
}

// dnsRecordUpdateResponseSmimeaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSmimea]
type dnsRecordUpdateResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseSmimea) implementsDNSRecordUpdateResponse() {}

// Components of a SMIMEA record.
type DNSRecordUpdateResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                               `json:"usage"`
	JSON  dnsRecordUpdateResponseSmimeaDataJSON `json:"-"`
}

// dnsRecordUpdateResponseSmimeaDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSmimeaData]
type dnsRecordUpdateResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseSmimeaType string

const (
	DNSRecordUpdateResponseSmimeaTypeSmimea DNSRecordUpdateResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                                `json:"source"`
	JSON   dnsRecordUpdateResponseSmimeaMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSmimeaMeta]
type dnsRecordUpdateResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseSmimeaTTLNumber].
type DNSRecordUpdateResponseSmimeaTTL interface {
	ImplementsDNSRecordUpdateResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseSmimeaTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseSmimeaTTLNumber float64

const (
	DNSRecordUpdateResponseSmimeaTTLNumber1 DNSRecordUpdateResponseSmimeaTTLNumber = 1
)

type DNSRecordUpdateResponseSRV struct {
	// Components of a SRV record.
	Data DNSRecordUpdateResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseSRVType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseSRVJSON `json:"-"`
}

// dnsRecordUpdateResponseSRVJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSRV]
type dnsRecordUpdateResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseSRV) implementsDNSRecordUpdateResponse() {}

// Components of a SRV record.
type DNSRecordUpdateResponseSRVData struct {
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
	Weight float64                            `json:"weight"`
	JSON   dnsRecordUpdateResponseSRVDataJSON `json:"-"`
}

// dnsRecordUpdateResponseSRVDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSRVData]
type dnsRecordUpdateResponseSRVDataJSON struct {
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

func (r *DNSRecordUpdateResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseSRVType string

const (
	DNSRecordUpdateResponseSRVTypeSRV DNSRecordUpdateResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordUpdateResponseSRVMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseSRVMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSRVMeta]
type dnsRecordUpdateResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseSRVTTLNumber].
type DNSRecordUpdateResponseSRVTTL interface {
	ImplementsDNSRecordUpdateResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseSRVTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseSRVTTLNumber float64

const (
	DNSRecordUpdateResponseSRVTTLNumber1 DNSRecordUpdateResponseSRVTTLNumber = 1
)

type DNSRecordUpdateResponseSSHFP struct {
	// Components of a SSHFP record.
	Data DNSRecordUpdateResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseSSHFPType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                           `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseSSHFPJSON `json:"-"`
}

// dnsRecordUpdateResponseSSHFPJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSSHFP]
type dnsRecordUpdateResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseSSHFP) implementsDNSRecordUpdateResponse() {}

// Components of a SSHFP record.
type DNSRecordUpdateResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                              `json:"type"`
	JSON dnsRecordUpdateResponseSSHFPDataJSON `json:"-"`
}

// dnsRecordUpdateResponseSSHFPDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSSHFPData]
type dnsRecordUpdateResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseSSHFPType string

const (
	DNSRecordUpdateResponseSSHFPTypeSSHFP DNSRecordUpdateResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                               `json:"source"`
	JSON   dnsRecordUpdateResponseSSHFPMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSSHFPMeta]
type dnsRecordUpdateResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseSSHFPTTLNumber].
type DNSRecordUpdateResponseSSHFPTTL interface {
	ImplementsDNSRecordUpdateResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseSSHFPTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseSSHFPTTLNumber float64

const (
	DNSRecordUpdateResponseSSHFPTTLNumber1 DNSRecordUpdateResponseSSHFPTTLNumber = 1
)

type DNSRecordUpdateResponseSVCB struct {
	// Components of a SVCB record.
	Data DNSRecordUpdateResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseSVCBType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseSVCBJSON `json:"-"`
}

// dnsRecordUpdateResponseSVCBJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSVCB]
type dnsRecordUpdateResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseSVCB) implementsDNSRecordUpdateResponse() {}

// Components of a SVCB record.
type DNSRecordUpdateResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                              `json:"value"`
	JSON  dnsRecordUpdateResponseSVCBDataJSON `json:"-"`
}

// dnsRecordUpdateResponseSVCBDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSVCBData]
type dnsRecordUpdateResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseSVCBType string

const (
	DNSRecordUpdateResponseSVCBTypeSVCB DNSRecordUpdateResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordUpdateResponseSVCBMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseSVCBMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseSVCBMeta]
type dnsRecordUpdateResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseSVCBTTLNumber].
type DNSRecordUpdateResponseSVCBTTL interface {
	ImplementsDNSRecordUpdateResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseSVCBTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseSVCBTTLNumber float64

const (
	DNSRecordUpdateResponseSVCBTTLNumber1 DNSRecordUpdateResponseSVCBTTLNumber = 1
)

type DNSRecordUpdateResponseTLSA struct {
	// Components of a TLSA record.
	Data DNSRecordUpdateResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseTLSAType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseTLSAJSON `json:"-"`
}

// dnsRecordUpdateResponseTLSAJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseTLSA]
type dnsRecordUpdateResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseTLSA) implementsDNSRecordUpdateResponse() {}

// Components of a TLSA record.
type DNSRecordUpdateResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                             `json:"usage"`
	JSON  dnsRecordUpdateResponseTLSADataJSON `json:"-"`
}

// dnsRecordUpdateResponseTLSADataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseTLSAData]
type dnsRecordUpdateResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseTLSAType string

const (
	DNSRecordUpdateResponseTLSATypeTLSA DNSRecordUpdateResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordUpdateResponseTLSAMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseTLSAMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseTLSAMeta]
type dnsRecordUpdateResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordUpdateResponseTLSATTLNumber].
type DNSRecordUpdateResponseTLSATTL interface {
	ImplementsDNSRecordUpdateResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseTLSATTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseTLSATTLNumber float64

const (
	DNSRecordUpdateResponseTLSATTLNumber1 DNSRecordUpdateResponseTLSATTLNumber = 1
)

type DNSRecordUpdateResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordUpdateResponseTXTType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseTXTJSON `json:"-"`
}

// dnsRecordUpdateResponseTXTJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseTXT]
type dnsRecordUpdateResponseTXTJSON struct {
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

func (r *DNSRecordUpdateResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseTXT) implementsDNSRecordUpdateResponse() {}

// Record type.
type DNSRecordUpdateResponseTXTType string

const (
	DNSRecordUpdateResponseTXTTypeTXT DNSRecordUpdateResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordUpdateResponseTXTMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseTXTMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseTXTMeta]
type dnsRecordUpdateResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseTXTTTLNumber].
type DNSRecordUpdateResponseTXTTTL interface {
	ImplementsDNSRecordUpdateResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseTXTTTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseTXTTTLNumber float64

const (
	DNSRecordUpdateResponseTXTTTLNumber1 DNSRecordUpdateResponseTXTTTLNumber = 1
)

type DNSRecordUpdateResponseURI struct {
	// Components of a URI record.
	Data DNSRecordUpdateResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordUpdateResponseURIType `json:"type,required"`
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
	Meta DNSRecordUpdateResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordUpdateResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordUpdateResponseURIJSON `json:"-"`
}

// dnsRecordUpdateResponseURIJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseURI]
type dnsRecordUpdateResponseURIJSON struct {
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

func (r *DNSRecordUpdateResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordUpdateResponseURI) implementsDNSRecordUpdateResponse() {}

// Components of a URI record.
type DNSRecordUpdateResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                            `json:"weight"`
	JSON   dnsRecordUpdateResponseURIDataJSON `json:"-"`
}

// dnsRecordUpdateResponseURIDataJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseURIData]
type dnsRecordUpdateResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordUpdateResponseURIType string

const (
	DNSRecordUpdateResponseURITypeURI DNSRecordUpdateResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordUpdateResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordUpdateResponseURIMetaJSON `json:"-"`
}

// dnsRecordUpdateResponseURIMetaJSON contains the JSON metadata for the struct
// [DNSRecordUpdateResponseURIMeta]
type dnsRecordUpdateResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordUpdateResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordUpdateResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordUpdateResponseURITTLNumber].
type DNSRecordUpdateResponseURITTL interface {
	ImplementsDNSRecordUpdateResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordUpdateResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordUpdateResponseURITTLNumber(0)),
		},
	)
}

type DNSRecordUpdateResponseURITTLNumber float64

const (
	DNSRecordUpdateResponseURITTLNumber1 DNSRecordUpdateResponseURITTLNumber = 1
)

// Union satisfied by [DNSRecordListResponseA], [DNSRecordListResponseAAAA],
// [DNSRecordListResponseCAA], [DNSRecordListResponseCert],
// [DNSRecordListResponseCNAME], [DNSRecordListResponseDNSKEY],
// [DNSRecordListResponseDS], [DNSRecordListResponseHTTPS],
// [DNSRecordListResponseLOC], [DNSRecordListResponseMX],
// [DNSRecordListResponseNAPTR], [DNSRecordListResponseNS],
// [DNSRecordListResponsePTR], [DNSRecordListResponseSmimea],
// [DNSRecordListResponseSRV], [DNSRecordListResponseSSHFP],
// [DNSRecordListResponseSVCB], [DNSRecordListResponseTLSA],
// [DNSRecordListResponseTXT] or [DNSRecordListResponseURI].
type DNSRecordListResponse interface {
	implementsDNSRecordListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordListResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordListResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseAType `json:"type,required"`
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
	Meta DNSRecordListResponseAMeta `json:"meta"`
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
	TTL DNSRecordListResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseAJSON `json:"-"`
}

// dnsRecordListResponseAJSON contains the JSON metadata for the struct
// [DNSRecordListResponseA]
type dnsRecordListResponseAJSON struct {
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

func (r *DNSRecordListResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseA) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseAType string

const (
	DNSRecordListResponseATypeA DNSRecordListResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordListResponseAMetaJSON `json:"-"`
}

// dnsRecordListResponseAMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseAMeta]
type dnsRecordListResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseATTLNumber].
type DNSRecordListResponseATTL interface {
	ImplementsDNSRecordListResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseATTLNumber(0)),
		},
	)
}

type DNSRecordListResponseATTLNumber float64

const (
	DNSRecordListResponseATTLNumber1 DNSRecordListResponseATTLNumber = 1
)

type DNSRecordListResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseAAAAType `json:"type,required"`
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
	Meta DNSRecordListResponseAAAAMeta `json:"meta"`
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
	TTL DNSRecordListResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseAAAAJSON `json:"-"`
}

// dnsRecordListResponseAAAAJSON contains the JSON metadata for the struct
// [DNSRecordListResponseAAAA]
type dnsRecordListResponseAAAAJSON struct {
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

func (r *DNSRecordListResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseAAAA) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseAAAAType string

const (
	DNSRecordListResponseAAAATypeAAAA DNSRecordListResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordListResponseAAAAMetaJSON `json:"-"`
}

// dnsRecordListResponseAAAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseAAAAMeta]
type dnsRecordListResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseAAAATTLNumber].
type DNSRecordListResponseAAAATTL interface {
	ImplementsDNSRecordListResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseAAAATTLNumber(0)),
		},
	)
}

type DNSRecordListResponseAAAATTLNumber float64

const (
	DNSRecordListResponseAAAATTLNumber1 DNSRecordListResponseAAAATTLNumber = 1
)

type DNSRecordListResponseCAA struct {
	// Components of a CAA record.
	Data DNSRecordListResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseCAAType `json:"type,required"`
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
	Meta DNSRecordListResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseCAAJSON `json:"-"`
}

// dnsRecordListResponseCAAJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCAA]
type dnsRecordListResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseCAA) implementsDNSRecordListResponse() {}

// Components of a CAA record.
type DNSRecordListResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                           `json:"value"`
	JSON  dnsRecordListResponseCAADataJSON `json:"-"`
}

// dnsRecordListResponseCAADataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCAAData]
type dnsRecordListResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseCAAType string

const (
	DNSRecordListResponseCAATypeCAA DNSRecordListResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordListResponseCAAMetaJSON `json:"-"`
}

// dnsRecordListResponseCAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCAAMeta]
type dnsRecordListResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseCAATTLNumber].
type DNSRecordListResponseCAATTL interface {
	ImplementsDNSRecordListResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseCAATTLNumber(0)),
		},
	)
}

type DNSRecordListResponseCAATTLNumber float64

const (
	DNSRecordListResponseCAATTLNumber1 DNSRecordListResponseCAATTLNumber = 1
)

type DNSRecordListResponseCert struct {
	// Components of a CERT record.
	Data DNSRecordListResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseCertType `json:"type,required"`
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
	Meta DNSRecordListResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseCertJSON `json:"-"`
}

// dnsRecordListResponseCertJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCert]
type dnsRecordListResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseCert) implementsDNSRecordListResponse() {}

// Components of a CERT record.
type DNSRecordListResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                           `json:"type"`
	JSON dnsRecordListResponseCertDataJSON `json:"-"`
}

// dnsRecordListResponseCertDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCertData]
type dnsRecordListResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseCertType string

const (
	DNSRecordListResponseCertTypeCert DNSRecordListResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordListResponseCertMetaJSON `json:"-"`
}

// dnsRecordListResponseCertMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCertMeta]
type dnsRecordListResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseCertTTLNumber].
type DNSRecordListResponseCertTTL interface {
	ImplementsDNSRecordListResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseCertTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseCertTTLNumber float64

const (
	DNSRecordListResponseCertTTLNumber1 DNSRecordListResponseCertTTLNumber = 1
)

type DNSRecordListResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseCNAMEType `json:"type,required"`
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
	Meta DNSRecordListResponseCNAMEMeta `json:"meta"`
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
	TTL DNSRecordListResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseCNAMEJSON `json:"-"`
}

// dnsRecordListResponseCNAMEJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCNAME]
type dnsRecordListResponseCNAMEJSON struct {
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

func (r *DNSRecordListResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseCNAME) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseCNAMEType string

const (
	DNSRecordListResponseCNAMETypeCNAME DNSRecordListResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordListResponseCNAMEMetaJSON `json:"-"`
}

// dnsRecordListResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseCNAMEMeta]
type dnsRecordListResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseCNAMETTLNumber].
type DNSRecordListResponseCNAMETTL interface {
	ImplementsDNSRecordListResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseCNAMETTLNumber(0)),
		},
	)
}

type DNSRecordListResponseCNAMETTLNumber float64

const (
	DNSRecordListResponseCNAMETTLNumber1 DNSRecordListResponseCNAMETTLNumber = 1
)

type DNSRecordListResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data DNSRecordListResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDNSKEYType `json:"type,required"`
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
	Meta DNSRecordListResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDNSKEYJSON `json:"-"`
}

// dnsRecordListResponseDNSKEYJSON contains the JSON metadata for the struct
// [DNSRecordListResponseDNSKEY]
type dnsRecordListResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseDNSKEY) implementsDNSRecordListResponse() {}

// Components of a DNSKEY record.
type DNSRecordListResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                              `json:"public_key"`
	JSON      dnsRecordListResponseDNSKEYDataJSON `json:"-"`
}

// dnsRecordListResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseDNSKEYData]
type dnsRecordListResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseDNSKEYType string

const (
	DNSRecordListResponseDNSKEYTypeDNSKEY DNSRecordListResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordListResponseDNSKEYMetaJSON `json:"-"`
}

// dnsRecordListResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseDNSKEYMeta]
type dnsRecordListResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseDNSKEYTTLNumber].
type DNSRecordListResponseDNSKEYTTL interface {
	ImplementsDNSRecordListResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseDNSKEYTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseDNSKEYTTLNumber float64

const (
	DNSRecordListResponseDNSKEYTTLNumber1 DNSRecordListResponseDNSKEYTTLNumber = 1
)

type DNSRecordListResponseDS struct {
	// Components of a DS record.
	Data DNSRecordListResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseDSType `json:"type,required"`
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
	Meta DNSRecordListResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseDSJSON `json:"-"`
}

// dnsRecordListResponseDSJSON contains the JSON metadata for the struct
// [DNSRecordListResponseDS]
type dnsRecordListResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseDS) implementsDNSRecordListResponse() {}

// Components of a DS record.
type DNSRecordListResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                         `json:"key_tag"`
	JSON   dnsRecordListResponseDSDataJSON `json:"-"`
}

// dnsRecordListResponseDSDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseDSData]
type dnsRecordListResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseDSType string

const (
	DNSRecordListResponseDSTypeDS DNSRecordListResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordListResponseDSMetaJSON `json:"-"`
}

// dnsRecordListResponseDSMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseDSMeta]
type dnsRecordListResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseDSTTLNumber].
type DNSRecordListResponseDSTTL interface {
	ImplementsDNSRecordListResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseDSTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseDSTTLNumber float64

const (
	DNSRecordListResponseDSTTLNumber1 DNSRecordListResponseDSTTLNumber = 1
)

type DNSRecordListResponseHTTPS struct {
	// Components of a HTTPS record.
	Data DNSRecordListResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseHTTPSType `json:"type,required"`
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
	Meta DNSRecordListResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseHTTPSJSON `json:"-"`
}

// dnsRecordListResponseHTTPSJSON contains the JSON metadata for the struct
// [DNSRecordListResponseHTTPS]
type dnsRecordListResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseHTTPS) implementsDNSRecordListResponse() {}

// Components of a HTTPS record.
type DNSRecordListResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                             `json:"value"`
	JSON  dnsRecordListResponseHTTPSDataJSON `json:"-"`
}

// dnsRecordListResponseHTTPSDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseHTTPSData]
type dnsRecordListResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseHTTPSType string

const (
	DNSRecordListResponseHTTPSTypeHTTPS DNSRecordListResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordListResponseHTTPSMetaJSON `json:"-"`
}

// dnsRecordListResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseHTTPSMeta]
type dnsRecordListResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseHTTPSTTLNumber].
type DNSRecordListResponseHTTPSTTL interface {
	ImplementsDNSRecordListResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseHTTPSTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseHTTPSTTLNumber float64

const (
	DNSRecordListResponseHTTPSTTLNumber1 DNSRecordListResponseHTTPSTTLNumber = 1
)

type DNSRecordListResponseLOC struct {
	// Components of a LOC record.
	Data DNSRecordListResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseLOCType `json:"type,required"`
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
	Meta DNSRecordListResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseLOCJSON `json:"-"`
}

// dnsRecordListResponseLOCJSON contains the JSON metadata for the struct
// [DNSRecordListResponseLOC]
type dnsRecordListResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseLOC) implementsDNSRecordListResponse() {}

// Components of a LOC record.
type DNSRecordListResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordListResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordListResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                          `json:"size"`
	JSON dnsRecordListResponseLOCDataJSON `json:"-"`
}

// dnsRecordListResponseLOCDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseLOCData]
type dnsRecordListResponseLOCDataJSON struct {
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

func (r *DNSRecordListResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordListResponseLOCDataLatDirection string

const (
	DNSRecordListResponseLOCDataLatDirectionN DNSRecordListResponseLOCDataLatDirection = "N"
	DNSRecordListResponseLOCDataLatDirectionS DNSRecordListResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordListResponseLOCDataLongDirection string

const (
	DNSRecordListResponseLOCDataLongDirectionE DNSRecordListResponseLOCDataLongDirection = "E"
	DNSRecordListResponseLOCDataLongDirectionW DNSRecordListResponseLOCDataLongDirection = "W"
)

// Record type.
type DNSRecordListResponseLOCType string

const (
	DNSRecordListResponseLOCTypeLOC DNSRecordListResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordListResponseLOCMetaJSON `json:"-"`
}

// dnsRecordListResponseLOCMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseLOCMeta]
type dnsRecordListResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseLOCTTLNumber].
type DNSRecordListResponseLOCTTL interface {
	ImplementsDNSRecordListResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseLOCTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseLOCTTLNumber float64

const (
	DNSRecordListResponseLOCTTLNumber1 DNSRecordListResponseLOCTTLNumber = 1
)

type DNSRecordListResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordListResponseMXType `json:"type,required"`
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
	Meta DNSRecordListResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseMXJSON `json:"-"`
}

// dnsRecordListResponseMXJSON contains the JSON metadata for the struct
// [DNSRecordListResponseMX]
type dnsRecordListResponseMXJSON struct {
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

func (r *DNSRecordListResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseMX) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseMXType string

const (
	DNSRecordListResponseMXTypeMX DNSRecordListResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordListResponseMXMetaJSON `json:"-"`
}

// dnsRecordListResponseMXMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseMXMeta]
type dnsRecordListResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseMXTTLNumber].
type DNSRecordListResponseMXTTL interface {
	ImplementsDNSRecordListResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseMXTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseMXTTLNumber float64

const (
	DNSRecordListResponseMXTTLNumber1 DNSRecordListResponseMXTTLNumber = 1
)

type DNSRecordListResponseNAPTR struct {
	// Components of a NAPTR record.
	Data DNSRecordListResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseNAPTRType `json:"type,required"`
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
	Meta DNSRecordListResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseNAPTRJSON `json:"-"`
}

// dnsRecordListResponseNAPTRJSON contains the JSON metadata for the struct
// [DNSRecordListResponseNAPTR]
type dnsRecordListResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseNAPTR) implementsDNSRecordListResponse() {}

// Components of a NAPTR record.
type DNSRecordListResponseNAPTRData struct {
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
	Service string                             `json:"service"`
	JSON    dnsRecordListResponseNAPTRDataJSON `json:"-"`
}

// dnsRecordListResponseNAPTRDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseNAPTRData]
type dnsRecordListResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseNAPTRType string

const (
	DNSRecordListResponseNAPTRTypeNAPTR DNSRecordListResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordListResponseNAPTRMetaJSON `json:"-"`
}

// dnsRecordListResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseNAPTRMeta]
type dnsRecordListResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseNAPTRTTLNumber].
type DNSRecordListResponseNAPTRTTL interface {
	ImplementsDNSRecordListResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseNAPTRTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseNAPTRTTLNumber float64

const (
	DNSRecordListResponseNAPTRTTLNumber1 DNSRecordListResponseNAPTRTTLNumber = 1
)

type DNSRecordListResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseNSType `json:"type,required"`
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
	Meta DNSRecordListResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseNSJSON `json:"-"`
}

// dnsRecordListResponseNSJSON contains the JSON metadata for the struct
// [DNSRecordListResponseNS]
type dnsRecordListResponseNSJSON struct {
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

func (r *DNSRecordListResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseNS) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseNSType string

const (
	DNSRecordListResponseNSTypeNS DNSRecordListResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordListResponseNSMetaJSON `json:"-"`
}

// dnsRecordListResponseNSMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseNSMeta]
type dnsRecordListResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseNSTTLNumber].
type DNSRecordListResponseNSTTL interface {
	ImplementsDNSRecordListResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseNSTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseNSTTLNumber float64

const (
	DNSRecordListResponseNSTTLNumber1 DNSRecordListResponseNSTTLNumber = 1
)

type DNSRecordListResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponsePTRType `json:"type,required"`
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
	Meta DNSRecordListResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponsePTRJSON `json:"-"`
}

// dnsRecordListResponsePTRJSON contains the JSON metadata for the struct
// [DNSRecordListResponsePTR]
type dnsRecordListResponsePTRJSON struct {
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

func (r *DNSRecordListResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponsePTR) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponsePTRType string

const (
	DNSRecordListResponsePTRTypePTR DNSRecordListResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordListResponsePTRMetaJSON `json:"-"`
}

// dnsRecordListResponsePTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponsePTRMeta]
type dnsRecordListResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponsePTRTTLNumber].
type DNSRecordListResponsePTRTTL interface {
	ImplementsDNSRecordListResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponsePTRTTLNumber(0)),
		},
	)
}

type DNSRecordListResponsePTRTTLNumber float64

const (
	DNSRecordListResponsePTRTTLNumber1 DNSRecordListResponsePTRTTLNumber = 1
)

type DNSRecordListResponseSmimea struct {
	// Components of a SMIMEA record.
	Data DNSRecordListResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseSmimeaType `json:"type,required"`
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
	Meta DNSRecordListResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseSmimeaJSON `json:"-"`
}

// dnsRecordListResponseSmimeaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSmimea]
type dnsRecordListResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseSmimea) implementsDNSRecordListResponse() {}

// Components of a SMIMEA record.
type DNSRecordListResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                             `json:"usage"`
	JSON  dnsRecordListResponseSmimeaDataJSON `json:"-"`
}

// dnsRecordListResponseSmimeaDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSmimeaData]
type dnsRecordListResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordListResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseSmimeaType string

const (
	DNSRecordListResponseSmimeaTypeSmimea DNSRecordListResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordListResponseSmimeaMetaJSON `json:"-"`
}

// dnsRecordListResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSmimeaMeta]
type dnsRecordListResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordListResponseSmimeaTTLNumber].
type DNSRecordListResponseSmimeaTTL interface {
	ImplementsDNSRecordListResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseSmimeaTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseSmimeaTTLNumber float64

const (
	DNSRecordListResponseSmimeaTTLNumber1 DNSRecordListResponseSmimeaTTLNumber = 1
)

type DNSRecordListResponseSRV struct {
	// Components of a SRV record.
	Data DNSRecordListResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseSRVType `json:"type,required"`
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
	Meta DNSRecordListResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseSRVJSON `json:"-"`
}

// dnsRecordListResponseSRVJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSRV]
type dnsRecordListResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseSRV) implementsDNSRecordListResponse() {}

// Components of a SRV record.
type DNSRecordListResponseSRVData struct {
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
	Weight float64                          `json:"weight"`
	JSON   dnsRecordListResponseSRVDataJSON `json:"-"`
}

// dnsRecordListResponseSRVDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSRVData]
type dnsRecordListResponseSRVDataJSON struct {
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

func (r *DNSRecordListResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseSRVType string

const (
	DNSRecordListResponseSRVTypeSRV DNSRecordListResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordListResponseSRVMetaJSON `json:"-"`
}

// dnsRecordListResponseSRVMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSRVMeta]
type dnsRecordListResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseSRVTTLNumber].
type DNSRecordListResponseSRVTTL interface {
	ImplementsDNSRecordListResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseSRVTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseSRVTTLNumber float64

const (
	DNSRecordListResponseSRVTTLNumber1 DNSRecordListResponseSRVTTLNumber = 1
)

type DNSRecordListResponseSSHFP struct {
	// Components of a SSHFP record.
	Data DNSRecordListResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseSSHFPType `json:"type,required"`
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
	Meta DNSRecordListResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseSSHFPJSON `json:"-"`
}

// dnsRecordListResponseSSHFPJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSSHFP]
type dnsRecordListResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseSSHFP) implementsDNSRecordListResponse() {}

// Components of a SSHFP record.
type DNSRecordListResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                            `json:"type"`
	JSON dnsRecordListResponseSSHFPDataJSON `json:"-"`
}

// dnsRecordListResponseSSHFPDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSSHFPData]
type dnsRecordListResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseSSHFPType string

const (
	DNSRecordListResponseSSHFPTypeSSHFP DNSRecordListResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordListResponseSSHFPMetaJSON `json:"-"`
}

// dnsRecordListResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSSHFPMeta]
type dnsRecordListResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseSSHFPTTLNumber].
type DNSRecordListResponseSSHFPTTL interface {
	ImplementsDNSRecordListResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseSSHFPTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseSSHFPTTLNumber float64

const (
	DNSRecordListResponseSSHFPTTLNumber1 DNSRecordListResponseSSHFPTTLNumber = 1
)

type DNSRecordListResponseSVCB struct {
	// Components of a SVCB record.
	Data DNSRecordListResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseSVCBType `json:"type,required"`
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
	Meta DNSRecordListResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseSVCBJSON `json:"-"`
}

// dnsRecordListResponseSVCBJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSVCB]
type dnsRecordListResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseSVCB) implementsDNSRecordListResponse() {}

// Components of a SVCB record.
type DNSRecordListResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                            `json:"value"`
	JSON  dnsRecordListResponseSVCBDataJSON `json:"-"`
}

// dnsRecordListResponseSVCBDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSVCBData]
type dnsRecordListResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseSVCBType string

const (
	DNSRecordListResponseSVCBTypeSVCB DNSRecordListResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordListResponseSVCBMetaJSON `json:"-"`
}

// dnsRecordListResponseSVCBMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseSVCBMeta]
type dnsRecordListResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseSVCBTTLNumber].
type DNSRecordListResponseSVCBTTL interface {
	ImplementsDNSRecordListResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseSVCBTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseSVCBTTLNumber float64

const (
	DNSRecordListResponseSVCBTTLNumber1 DNSRecordListResponseSVCBTTLNumber = 1
)

type DNSRecordListResponseTLSA struct {
	// Components of a TLSA record.
	Data DNSRecordListResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseTLSAType `json:"type,required"`
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
	Meta DNSRecordListResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseTLSAJSON `json:"-"`
}

// dnsRecordListResponseTLSAJSON contains the JSON metadata for the struct
// [DNSRecordListResponseTLSA]
type dnsRecordListResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseTLSA) implementsDNSRecordListResponse() {}

// Components of a TLSA record.
type DNSRecordListResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                           `json:"usage"`
	JSON  dnsRecordListResponseTLSADataJSON `json:"-"`
}

// dnsRecordListResponseTLSADataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseTLSAData]
type dnsRecordListResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordListResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseTLSAType string

const (
	DNSRecordListResponseTLSATypeTLSA DNSRecordListResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordListResponseTLSAMetaJSON `json:"-"`
}

// dnsRecordListResponseTLSAMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseTLSAMeta]
type dnsRecordListResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseTLSATTLNumber].
type DNSRecordListResponseTLSATTL interface {
	ImplementsDNSRecordListResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseTLSATTLNumber(0)),
		},
	)
}

type DNSRecordListResponseTLSATTLNumber float64

const (
	DNSRecordListResponseTLSATTLNumber1 DNSRecordListResponseTLSATTLNumber = 1
)

type DNSRecordListResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordListResponseTXTType `json:"type,required"`
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
	Meta DNSRecordListResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseTXTJSON `json:"-"`
}

// dnsRecordListResponseTXTJSON contains the JSON metadata for the struct
// [DNSRecordListResponseTXT]
type dnsRecordListResponseTXTJSON struct {
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

func (r *DNSRecordListResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseTXT) implementsDNSRecordListResponse() {}

// Record type.
type DNSRecordListResponseTXTType string

const (
	DNSRecordListResponseTXTTypeTXT DNSRecordListResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordListResponseTXTMetaJSON `json:"-"`
}

// dnsRecordListResponseTXTMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseTXTMeta]
type dnsRecordListResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseTXTTTLNumber].
type DNSRecordListResponseTXTTTL interface {
	ImplementsDNSRecordListResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseTXTTTLNumber(0)),
		},
	)
}

type DNSRecordListResponseTXTTTLNumber float64

const (
	DNSRecordListResponseTXTTTLNumber1 DNSRecordListResponseTXTTTLNumber = 1
)

type DNSRecordListResponseURI struct {
	// Components of a URI record.
	Data DNSRecordListResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordListResponseURIType `json:"type,required"`
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
	Meta DNSRecordListResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordListResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordListResponseURIJSON `json:"-"`
}

// dnsRecordListResponseURIJSON contains the JSON metadata for the struct
// [DNSRecordListResponseURI]
type dnsRecordListResponseURIJSON struct {
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

func (r *DNSRecordListResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordListResponseURI) implementsDNSRecordListResponse() {}

// Components of a URI record.
type DNSRecordListResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                          `json:"weight"`
	JSON   dnsRecordListResponseURIDataJSON `json:"-"`
}

// dnsRecordListResponseURIDataJSON contains the JSON metadata for the struct
// [DNSRecordListResponseURIData]
type dnsRecordListResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordListResponseURIType string

const (
	DNSRecordListResponseURITypeURI DNSRecordListResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordListResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordListResponseURIMetaJSON `json:"-"`
}

// dnsRecordListResponseURIMetaJSON contains the JSON metadata for the struct
// [DNSRecordListResponseURIMeta]
type dnsRecordListResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordListResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordListResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordListResponseURITTLNumber].
type DNSRecordListResponseURITTL interface {
	ImplementsDNSRecordListResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordListResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordListResponseURITTLNumber(0)),
		},
	)
}

type DNSRecordListResponseURITTLNumber float64

const (
	DNSRecordListResponseURITTLNumber1 DNSRecordListResponseURITTLNumber = 1
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

func (r dnsRecordDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [DNSRecordEditResponseA], [DNSRecordEditResponseAAAA],
// [DNSRecordEditResponseCAA], [DNSRecordEditResponseCert],
// [DNSRecordEditResponseCNAME], [DNSRecordEditResponseDNSKEY],
// [DNSRecordEditResponseDS], [DNSRecordEditResponseHTTPS],
// [DNSRecordEditResponseLOC], [DNSRecordEditResponseMX],
// [DNSRecordEditResponseNAPTR], [DNSRecordEditResponseNS],
// [DNSRecordEditResponsePTR], [DNSRecordEditResponseSmimea],
// [DNSRecordEditResponseSRV], [DNSRecordEditResponseSSHFP],
// [DNSRecordEditResponseSVCB], [DNSRecordEditResponseTLSA],
// [DNSRecordEditResponseTXT] or [DNSRecordEditResponseURI].
type DNSRecordEditResponse interface {
	implementsDNSRecordEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordEditResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordEditResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseAType `json:"type,required"`
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
	Meta DNSRecordEditResponseAMeta `json:"meta"`
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
	TTL DNSRecordEditResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseAJSON `json:"-"`
}

// dnsRecordEditResponseAJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseA]
type dnsRecordEditResponseAJSON struct {
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

func (r *DNSRecordEditResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseA) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseAType string

const (
	DNSRecordEditResponseATypeA DNSRecordEditResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordEditResponseAMetaJSON `json:"-"`
}

// dnsRecordEditResponseAMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseAMeta]
type dnsRecordEditResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseATTLNumber].
type DNSRecordEditResponseATTL interface {
	ImplementsDNSRecordEditResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseATTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseATTLNumber float64

const (
	DNSRecordEditResponseATTLNumber1 DNSRecordEditResponseATTLNumber = 1
)

type DNSRecordEditResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseAAAAType `json:"type,required"`
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
	Meta DNSRecordEditResponseAAAAMeta `json:"meta"`
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
	TTL DNSRecordEditResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseAAAAJSON `json:"-"`
}

// dnsRecordEditResponseAAAAJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseAAAA]
type dnsRecordEditResponseAAAAJSON struct {
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

func (r *DNSRecordEditResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseAAAA) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseAAAAType string

const (
	DNSRecordEditResponseAAAATypeAAAA DNSRecordEditResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordEditResponseAAAAMetaJSON `json:"-"`
}

// dnsRecordEditResponseAAAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseAAAAMeta]
type dnsRecordEditResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseAAAATTLNumber].
type DNSRecordEditResponseAAAATTL interface {
	ImplementsDNSRecordEditResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseAAAATTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseAAAATTLNumber float64

const (
	DNSRecordEditResponseAAAATTLNumber1 DNSRecordEditResponseAAAATTLNumber = 1
)

type DNSRecordEditResponseCAA struct {
	// Components of a CAA record.
	Data DNSRecordEditResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseCAAType `json:"type,required"`
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
	Meta DNSRecordEditResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseCAAJSON `json:"-"`
}

// dnsRecordEditResponseCAAJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCAA]
type dnsRecordEditResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseCAA) implementsDNSRecordEditResponse() {}

// Components of a CAA record.
type DNSRecordEditResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                           `json:"value"`
	JSON  dnsRecordEditResponseCAADataJSON `json:"-"`
}

// dnsRecordEditResponseCAADataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCAAData]
type dnsRecordEditResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseCAAType string

const (
	DNSRecordEditResponseCAATypeCAA DNSRecordEditResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordEditResponseCAAMetaJSON `json:"-"`
}

// dnsRecordEditResponseCAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCAAMeta]
type dnsRecordEditResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseCAATTLNumber].
type DNSRecordEditResponseCAATTL interface {
	ImplementsDNSRecordEditResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseCAATTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseCAATTLNumber float64

const (
	DNSRecordEditResponseCAATTLNumber1 DNSRecordEditResponseCAATTLNumber = 1
)

type DNSRecordEditResponseCert struct {
	// Components of a CERT record.
	Data DNSRecordEditResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseCertType `json:"type,required"`
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
	Meta DNSRecordEditResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseCertJSON `json:"-"`
}

// dnsRecordEditResponseCertJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCert]
type dnsRecordEditResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseCert) implementsDNSRecordEditResponse() {}

// Components of a CERT record.
type DNSRecordEditResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                           `json:"type"`
	JSON dnsRecordEditResponseCertDataJSON `json:"-"`
}

// dnsRecordEditResponseCertDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCertData]
type dnsRecordEditResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseCertType string

const (
	DNSRecordEditResponseCertTypeCert DNSRecordEditResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordEditResponseCertMetaJSON `json:"-"`
}

// dnsRecordEditResponseCertMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCertMeta]
type dnsRecordEditResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseCertTTLNumber].
type DNSRecordEditResponseCertTTL interface {
	ImplementsDNSRecordEditResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseCertTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseCertTTLNumber float64

const (
	DNSRecordEditResponseCertTTLNumber1 DNSRecordEditResponseCertTTLNumber = 1
)

type DNSRecordEditResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseCNAMEType `json:"type,required"`
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
	Meta DNSRecordEditResponseCNAMEMeta `json:"meta"`
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
	TTL DNSRecordEditResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseCNAMEJSON `json:"-"`
}

// dnsRecordEditResponseCNAMEJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCNAME]
type dnsRecordEditResponseCNAMEJSON struct {
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

func (r *DNSRecordEditResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseCNAME) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseCNAMEType string

const (
	DNSRecordEditResponseCNAMETypeCNAME DNSRecordEditResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordEditResponseCNAMEMetaJSON `json:"-"`
}

// dnsRecordEditResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseCNAMEMeta]
type dnsRecordEditResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseCNAMETTLNumber].
type DNSRecordEditResponseCNAMETTL interface {
	ImplementsDNSRecordEditResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseCNAMETTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseCNAMETTLNumber float64

const (
	DNSRecordEditResponseCNAMETTLNumber1 DNSRecordEditResponseCNAMETTLNumber = 1
)

type DNSRecordEditResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data DNSRecordEditResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDNSKEYType `json:"type,required"`
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
	Meta DNSRecordEditResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDNSKEYJSON `json:"-"`
}

// dnsRecordEditResponseDNSKEYJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseDNSKEY]
type dnsRecordEditResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseDNSKEY) implementsDNSRecordEditResponse() {}

// Components of a DNSKEY record.
type DNSRecordEditResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                              `json:"public_key"`
	JSON      dnsRecordEditResponseDNSKEYDataJSON `json:"-"`
}

// dnsRecordEditResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseDNSKEYData]
type dnsRecordEditResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseDNSKEYType string

const (
	DNSRecordEditResponseDNSKEYTypeDNSKEY DNSRecordEditResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordEditResponseDNSKEYMetaJSON `json:"-"`
}

// dnsRecordEditResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseDNSKEYMeta]
type dnsRecordEditResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseDNSKEYTTLNumber].
type DNSRecordEditResponseDNSKEYTTL interface {
	ImplementsDNSRecordEditResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseDNSKEYTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseDNSKEYTTLNumber float64

const (
	DNSRecordEditResponseDNSKEYTTLNumber1 DNSRecordEditResponseDNSKEYTTLNumber = 1
)

type DNSRecordEditResponseDS struct {
	// Components of a DS record.
	Data DNSRecordEditResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseDSType `json:"type,required"`
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
	Meta DNSRecordEditResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseDSJSON `json:"-"`
}

// dnsRecordEditResponseDSJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseDS]
type dnsRecordEditResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseDS) implementsDNSRecordEditResponse() {}

// Components of a DS record.
type DNSRecordEditResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                         `json:"key_tag"`
	JSON   dnsRecordEditResponseDSDataJSON `json:"-"`
}

// dnsRecordEditResponseDSDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseDSData]
type dnsRecordEditResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseDSType string

const (
	DNSRecordEditResponseDSTypeDS DNSRecordEditResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordEditResponseDSMetaJSON `json:"-"`
}

// dnsRecordEditResponseDSMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseDSMeta]
type dnsRecordEditResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseDSTTLNumber].
type DNSRecordEditResponseDSTTL interface {
	ImplementsDNSRecordEditResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseDSTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseDSTTLNumber float64

const (
	DNSRecordEditResponseDSTTLNumber1 DNSRecordEditResponseDSTTLNumber = 1
)

type DNSRecordEditResponseHTTPS struct {
	// Components of a HTTPS record.
	Data DNSRecordEditResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseHTTPSType `json:"type,required"`
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
	Meta DNSRecordEditResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseHTTPSJSON `json:"-"`
}

// dnsRecordEditResponseHTTPSJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseHTTPS]
type dnsRecordEditResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseHTTPS) implementsDNSRecordEditResponse() {}

// Components of a HTTPS record.
type DNSRecordEditResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                             `json:"value"`
	JSON  dnsRecordEditResponseHTTPSDataJSON `json:"-"`
}

// dnsRecordEditResponseHTTPSDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseHTTPSData]
type dnsRecordEditResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseHTTPSType string

const (
	DNSRecordEditResponseHTTPSTypeHTTPS DNSRecordEditResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordEditResponseHTTPSMetaJSON `json:"-"`
}

// dnsRecordEditResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseHTTPSMeta]
type dnsRecordEditResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseHTTPSTTLNumber].
type DNSRecordEditResponseHTTPSTTL interface {
	ImplementsDNSRecordEditResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseHTTPSTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseHTTPSTTLNumber float64

const (
	DNSRecordEditResponseHTTPSTTLNumber1 DNSRecordEditResponseHTTPSTTLNumber = 1
)

type DNSRecordEditResponseLOC struct {
	// Components of a LOC record.
	Data DNSRecordEditResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseLOCType `json:"type,required"`
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
	Meta DNSRecordEditResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseLOCJSON `json:"-"`
}

// dnsRecordEditResponseLOCJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseLOC]
type dnsRecordEditResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseLOC) implementsDNSRecordEditResponse() {}

// Components of a LOC record.
type DNSRecordEditResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordEditResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordEditResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                          `json:"size"`
	JSON dnsRecordEditResponseLOCDataJSON `json:"-"`
}

// dnsRecordEditResponseLOCDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseLOCData]
type dnsRecordEditResponseLOCDataJSON struct {
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

func (r *DNSRecordEditResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordEditResponseLOCDataLatDirection string

const (
	DNSRecordEditResponseLOCDataLatDirectionN DNSRecordEditResponseLOCDataLatDirection = "N"
	DNSRecordEditResponseLOCDataLatDirectionS DNSRecordEditResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordEditResponseLOCDataLongDirection string

const (
	DNSRecordEditResponseLOCDataLongDirectionE DNSRecordEditResponseLOCDataLongDirection = "E"
	DNSRecordEditResponseLOCDataLongDirectionW DNSRecordEditResponseLOCDataLongDirection = "W"
)

// Record type.
type DNSRecordEditResponseLOCType string

const (
	DNSRecordEditResponseLOCTypeLOC DNSRecordEditResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordEditResponseLOCMetaJSON `json:"-"`
}

// dnsRecordEditResponseLOCMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseLOCMeta]
type dnsRecordEditResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseLOCTTLNumber].
type DNSRecordEditResponseLOCTTL interface {
	ImplementsDNSRecordEditResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseLOCTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseLOCTTLNumber float64

const (
	DNSRecordEditResponseLOCTTLNumber1 DNSRecordEditResponseLOCTTLNumber = 1
)

type DNSRecordEditResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordEditResponseMXType `json:"type,required"`
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
	Meta DNSRecordEditResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseMXJSON `json:"-"`
}

// dnsRecordEditResponseMXJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseMX]
type dnsRecordEditResponseMXJSON struct {
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

func (r *DNSRecordEditResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseMX) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseMXType string

const (
	DNSRecordEditResponseMXTypeMX DNSRecordEditResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordEditResponseMXMetaJSON `json:"-"`
}

// dnsRecordEditResponseMXMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseMXMeta]
type dnsRecordEditResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseMXTTLNumber].
type DNSRecordEditResponseMXTTL interface {
	ImplementsDNSRecordEditResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseMXTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseMXTTLNumber float64

const (
	DNSRecordEditResponseMXTTLNumber1 DNSRecordEditResponseMXTTLNumber = 1
)

type DNSRecordEditResponseNAPTR struct {
	// Components of a NAPTR record.
	Data DNSRecordEditResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseNAPTRType `json:"type,required"`
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
	Meta DNSRecordEditResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseNAPTRJSON `json:"-"`
}

// dnsRecordEditResponseNAPTRJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseNAPTR]
type dnsRecordEditResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseNAPTR) implementsDNSRecordEditResponse() {}

// Components of a NAPTR record.
type DNSRecordEditResponseNAPTRData struct {
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
	Service string                             `json:"service"`
	JSON    dnsRecordEditResponseNAPTRDataJSON `json:"-"`
}

// dnsRecordEditResponseNAPTRDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseNAPTRData]
type dnsRecordEditResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseNAPTRType string

const (
	DNSRecordEditResponseNAPTRTypeNAPTR DNSRecordEditResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordEditResponseNAPTRMetaJSON `json:"-"`
}

// dnsRecordEditResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseNAPTRMeta]
type dnsRecordEditResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseNAPTRTTLNumber].
type DNSRecordEditResponseNAPTRTTL interface {
	ImplementsDNSRecordEditResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseNAPTRTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseNAPTRTTLNumber float64

const (
	DNSRecordEditResponseNAPTRTTLNumber1 DNSRecordEditResponseNAPTRTTLNumber = 1
)

type DNSRecordEditResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseNSType `json:"type,required"`
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
	Meta DNSRecordEditResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseNSJSON `json:"-"`
}

// dnsRecordEditResponseNSJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseNS]
type dnsRecordEditResponseNSJSON struct {
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

func (r *DNSRecordEditResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseNS) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseNSType string

const (
	DNSRecordEditResponseNSTypeNS DNSRecordEditResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordEditResponseNSMetaJSON `json:"-"`
}

// dnsRecordEditResponseNSMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseNSMeta]
type dnsRecordEditResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseNSTTLNumber].
type DNSRecordEditResponseNSTTL interface {
	ImplementsDNSRecordEditResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseNSTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseNSTTLNumber float64

const (
	DNSRecordEditResponseNSTTLNumber1 DNSRecordEditResponseNSTTLNumber = 1
)

type DNSRecordEditResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponsePTRType `json:"type,required"`
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
	Meta DNSRecordEditResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponsePTRJSON `json:"-"`
}

// dnsRecordEditResponsePTRJSON contains the JSON metadata for the struct
// [DNSRecordEditResponsePTR]
type dnsRecordEditResponsePTRJSON struct {
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

func (r *DNSRecordEditResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponsePTR) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponsePTRType string

const (
	DNSRecordEditResponsePTRTypePTR DNSRecordEditResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordEditResponsePTRMetaJSON `json:"-"`
}

// dnsRecordEditResponsePTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponsePTRMeta]
type dnsRecordEditResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponsePTRTTLNumber].
type DNSRecordEditResponsePTRTTL interface {
	ImplementsDNSRecordEditResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponsePTRTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponsePTRTTLNumber float64

const (
	DNSRecordEditResponsePTRTTLNumber1 DNSRecordEditResponsePTRTTLNumber = 1
)

type DNSRecordEditResponseSmimea struct {
	// Components of a SMIMEA record.
	Data DNSRecordEditResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseSmimeaType `json:"type,required"`
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
	Meta DNSRecordEditResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                          `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseSmimeaJSON `json:"-"`
}

// dnsRecordEditResponseSmimeaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSmimea]
type dnsRecordEditResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseSmimea) implementsDNSRecordEditResponse() {}

// Components of a SMIMEA record.
type DNSRecordEditResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                             `json:"usage"`
	JSON  dnsRecordEditResponseSmimeaDataJSON `json:"-"`
}

// dnsRecordEditResponseSmimeaDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSmimeaData]
type dnsRecordEditResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordEditResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseSmimeaType string

const (
	DNSRecordEditResponseSmimeaTypeSmimea DNSRecordEditResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                              `json:"source"`
	JSON   dnsRecordEditResponseSmimeaMetaJSON `json:"-"`
}

// dnsRecordEditResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSmimeaMeta]
type dnsRecordEditResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [DNSRecordEditResponseSmimeaTTLNumber].
type DNSRecordEditResponseSmimeaTTL interface {
	ImplementsDNSRecordEditResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseSmimeaTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseSmimeaTTLNumber float64

const (
	DNSRecordEditResponseSmimeaTTLNumber1 DNSRecordEditResponseSmimeaTTLNumber = 1
)

type DNSRecordEditResponseSRV struct {
	// Components of a SRV record.
	Data DNSRecordEditResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseSRVType `json:"type,required"`
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
	Meta DNSRecordEditResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseSRVJSON `json:"-"`
}

// dnsRecordEditResponseSRVJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSRV]
type dnsRecordEditResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseSRV) implementsDNSRecordEditResponse() {}

// Components of a SRV record.
type DNSRecordEditResponseSRVData struct {
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
	Weight float64                          `json:"weight"`
	JSON   dnsRecordEditResponseSRVDataJSON `json:"-"`
}

// dnsRecordEditResponseSRVDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSRVData]
type dnsRecordEditResponseSRVDataJSON struct {
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

func (r *DNSRecordEditResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseSRVType string

const (
	DNSRecordEditResponseSRVTypeSRV DNSRecordEditResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordEditResponseSRVMetaJSON `json:"-"`
}

// dnsRecordEditResponseSRVMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSRVMeta]
type dnsRecordEditResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseSRVTTLNumber].
type DNSRecordEditResponseSRVTTL interface {
	ImplementsDNSRecordEditResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseSRVTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseSRVTTLNumber float64

const (
	DNSRecordEditResponseSRVTTLNumber1 DNSRecordEditResponseSRVTTLNumber = 1
)

type DNSRecordEditResponseSSHFP struct {
	// Components of a SSHFP record.
	Data DNSRecordEditResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseSSHFPType `json:"type,required"`
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
	Meta DNSRecordEditResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseSSHFPJSON `json:"-"`
}

// dnsRecordEditResponseSSHFPJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSSHFP]
type dnsRecordEditResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseSSHFP) implementsDNSRecordEditResponse() {}

// Components of a SSHFP record.
type DNSRecordEditResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                            `json:"type"`
	JSON dnsRecordEditResponseSSHFPDataJSON `json:"-"`
}

// dnsRecordEditResponseSSHFPDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSSHFPData]
type dnsRecordEditResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseSSHFPType string

const (
	DNSRecordEditResponseSSHFPTypeSSHFP DNSRecordEditResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordEditResponseSSHFPMetaJSON `json:"-"`
}

// dnsRecordEditResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSSHFPMeta]
type dnsRecordEditResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseSSHFPTTLNumber].
type DNSRecordEditResponseSSHFPTTL interface {
	ImplementsDNSRecordEditResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseSSHFPTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseSSHFPTTLNumber float64

const (
	DNSRecordEditResponseSSHFPTTLNumber1 DNSRecordEditResponseSSHFPTTLNumber = 1
)

type DNSRecordEditResponseSVCB struct {
	// Components of a SVCB record.
	Data DNSRecordEditResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseSVCBType `json:"type,required"`
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
	Meta DNSRecordEditResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseSVCBJSON `json:"-"`
}

// dnsRecordEditResponseSVCBJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSVCB]
type dnsRecordEditResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseSVCB) implementsDNSRecordEditResponse() {}

// Components of a SVCB record.
type DNSRecordEditResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                            `json:"value"`
	JSON  dnsRecordEditResponseSVCBDataJSON `json:"-"`
}

// dnsRecordEditResponseSVCBDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSVCBData]
type dnsRecordEditResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseSVCBType string

const (
	DNSRecordEditResponseSVCBTypeSVCB DNSRecordEditResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordEditResponseSVCBMetaJSON `json:"-"`
}

// dnsRecordEditResponseSVCBMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseSVCBMeta]
type dnsRecordEditResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseSVCBTTLNumber].
type DNSRecordEditResponseSVCBTTL interface {
	ImplementsDNSRecordEditResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseSVCBTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseSVCBTTLNumber float64

const (
	DNSRecordEditResponseSVCBTTLNumber1 DNSRecordEditResponseSVCBTTLNumber = 1
)

type DNSRecordEditResponseTLSA struct {
	// Components of a TLSA record.
	Data DNSRecordEditResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseTLSAType `json:"type,required"`
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
	Meta DNSRecordEditResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseTLSAJSON `json:"-"`
}

// dnsRecordEditResponseTLSAJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseTLSA]
type dnsRecordEditResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseTLSA) implementsDNSRecordEditResponse() {}

// Components of a TLSA record.
type DNSRecordEditResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                           `json:"usage"`
	JSON  dnsRecordEditResponseTLSADataJSON `json:"-"`
}

// dnsRecordEditResponseTLSADataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseTLSAData]
type dnsRecordEditResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordEditResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseTLSAType string

const (
	DNSRecordEditResponseTLSATypeTLSA DNSRecordEditResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordEditResponseTLSAMetaJSON `json:"-"`
}

// dnsRecordEditResponseTLSAMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseTLSAMeta]
type dnsRecordEditResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseTLSATTLNumber].
type DNSRecordEditResponseTLSATTL interface {
	ImplementsDNSRecordEditResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseTLSATTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseTLSATTLNumber float64

const (
	DNSRecordEditResponseTLSATTLNumber1 DNSRecordEditResponseTLSATTLNumber = 1
)

type DNSRecordEditResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordEditResponseTXTType `json:"type,required"`
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
	Meta DNSRecordEditResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseTXTJSON `json:"-"`
}

// dnsRecordEditResponseTXTJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseTXT]
type dnsRecordEditResponseTXTJSON struct {
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

func (r *DNSRecordEditResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseTXT) implementsDNSRecordEditResponse() {}

// Record type.
type DNSRecordEditResponseTXTType string

const (
	DNSRecordEditResponseTXTTypeTXT DNSRecordEditResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordEditResponseTXTMetaJSON `json:"-"`
}

// dnsRecordEditResponseTXTMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseTXTMeta]
type dnsRecordEditResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseTXTTTLNumber].
type DNSRecordEditResponseTXTTTL interface {
	ImplementsDNSRecordEditResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseTXTTTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseTXTTTLNumber float64

const (
	DNSRecordEditResponseTXTTTLNumber1 DNSRecordEditResponseTXTTTLNumber = 1
)

type DNSRecordEditResponseURI struct {
	// Components of a URI record.
	Data DNSRecordEditResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordEditResponseURIType `json:"type,required"`
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
	Meta DNSRecordEditResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordEditResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordEditResponseURIJSON `json:"-"`
}

// dnsRecordEditResponseURIJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseURI]
type dnsRecordEditResponseURIJSON struct {
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

func (r *DNSRecordEditResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordEditResponseURI) implementsDNSRecordEditResponse() {}

// Components of a URI record.
type DNSRecordEditResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                          `json:"weight"`
	JSON   dnsRecordEditResponseURIDataJSON `json:"-"`
}

// dnsRecordEditResponseURIDataJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseURIData]
type dnsRecordEditResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordEditResponseURIType string

const (
	DNSRecordEditResponseURITypeURI DNSRecordEditResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordEditResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordEditResponseURIMetaJSON `json:"-"`
}

// dnsRecordEditResponseURIMetaJSON contains the JSON metadata for the struct
// [DNSRecordEditResponseURIMeta]
type dnsRecordEditResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordEditResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordEditResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordEditResponseURITTLNumber].
type DNSRecordEditResponseURITTL interface {
	ImplementsDNSRecordEditResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordEditResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordEditResponseURITTLNumber(0)),
		},
	)
}

type DNSRecordEditResponseURITTLNumber float64

const (
	DNSRecordEditResponseURITTLNumber1 DNSRecordEditResponseURITTLNumber = 1
)

// Union satisfied by [DNSRecordGetResponseA], [DNSRecordGetResponseAAAA],
// [DNSRecordGetResponseCAA], [DNSRecordGetResponseCert],
// [DNSRecordGetResponseCNAME], [DNSRecordGetResponseDNSKEY],
// [DNSRecordGetResponseDS], [DNSRecordGetResponseHTTPS],
// [DNSRecordGetResponseLOC], [DNSRecordGetResponseMX],
// [DNSRecordGetResponseNAPTR], [DNSRecordGetResponseNS],
// [DNSRecordGetResponsePTR], [DNSRecordGetResponseSmimea],
// [DNSRecordGetResponseSRV], [DNSRecordGetResponseSSHFP],
// [DNSRecordGetResponseSVCB], [DNSRecordGetResponseTLSA],
// [DNSRecordGetResponseTXT] or [DNSRecordGetResponseURI].
type DNSRecordGetResponse interface {
	implementsDNSRecordGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSRecordGetResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type DNSRecordGetResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseAType `json:"type,required"`
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
	Meta DNSRecordGetResponseAMeta `json:"meta"`
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
	TTL DNSRecordGetResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseAJSON `json:"-"`
}

// dnsRecordGetResponseAJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseA]
type dnsRecordGetResponseAJSON struct {
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

func (r *DNSRecordGetResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseA) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseAType string

const (
	DNSRecordGetResponseATypeA DNSRecordGetResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   dnsRecordGetResponseAMetaJSON `json:"-"`
}

// dnsRecordGetResponseAMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseAMeta]
type dnsRecordGetResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseATTLNumber].
type DNSRecordGetResponseATTL interface {
	ImplementsDNSRecordGetResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseATTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseATTLNumber float64

const (
	DNSRecordGetResponseATTLNumber1 DNSRecordGetResponseATTLNumber = 1
)

type DNSRecordGetResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseAAAAType `json:"type,required"`
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
	Meta DNSRecordGetResponseAAAAMeta `json:"meta"`
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
	TTL DNSRecordGetResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseAAAAJSON `json:"-"`
}

// dnsRecordGetResponseAAAAJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseAAAA]
type dnsRecordGetResponseAAAAJSON struct {
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

func (r *DNSRecordGetResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseAAAA) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseAAAAType string

const (
	DNSRecordGetResponseAAAATypeAAAA DNSRecordGetResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordGetResponseAAAAMetaJSON `json:"-"`
}

// dnsRecordGetResponseAAAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseAAAAMeta]
type dnsRecordGetResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseAAAATTLNumber].
type DNSRecordGetResponseAAAATTL interface {
	ImplementsDNSRecordGetResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseAAAATTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseAAAATTLNumber float64

const (
	DNSRecordGetResponseAAAATTLNumber1 DNSRecordGetResponseAAAATTLNumber = 1
)

type DNSRecordGetResponseCAA struct {
	// Components of a CAA record.
	Data DNSRecordGetResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseCAAType `json:"type,required"`
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
	Meta DNSRecordGetResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseCAAJSON `json:"-"`
}

// dnsRecordGetResponseCAAJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCAA]
type dnsRecordGetResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseCAA) implementsDNSRecordGetResponse() {}

// Components of a CAA record.
type DNSRecordGetResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                          `json:"value"`
	JSON  dnsRecordGetResponseCAADataJSON `json:"-"`
}

// dnsRecordGetResponseCAADataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCAAData]
type dnsRecordGetResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseCAAType string

const (
	DNSRecordGetResponseCAATypeCAA DNSRecordGetResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordGetResponseCAAMetaJSON `json:"-"`
}

// dnsRecordGetResponseCAAMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCAAMeta]
type dnsRecordGetResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseCAATTLNumber].
type DNSRecordGetResponseCAATTL interface {
	ImplementsDNSRecordGetResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseCAATTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseCAATTLNumber float64

const (
	DNSRecordGetResponseCAATTLNumber1 DNSRecordGetResponseCAATTLNumber = 1
)

type DNSRecordGetResponseCert struct {
	// Components of a CERT record.
	Data DNSRecordGetResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseCertType `json:"type,required"`
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
	Meta DNSRecordGetResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseCertJSON `json:"-"`
}

// dnsRecordGetResponseCertJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCert]
type dnsRecordGetResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseCert) implementsDNSRecordGetResponse() {}

// Components of a CERT record.
type DNSRecordGetResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                          `json:"type"`
	JSON dnsRecordGetResponseCertDataJSON `json:"-"`
}

// dnsRecordGetResponseCertDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCertData]
type dnsRecordGetResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseCertType string

const (
	DNSRecordGetResponseCertTypeCert DNSRecordGetResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordGetResponseCertMetaJSON `json:"-"`
}

// dnsRecordGetResponseCertMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCertMeta]
type dnsRecordGetResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseCertTTLNumber].
type DNSRecordGetResponseCertTTL interface {
	ImplementsDNSRecordGetResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseCertTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseCertTTLNumber float64

const (
	DNSRecordGetResponseCertTTLNumber1 DNSRecordGetResponseCertTTLNumber = 1
)

type DNSRecordGetResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseCNAMEType `json:"type,required"`
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
	Meta DNSRecordGetResponseCNAMEMeta `json:"meta"`
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
	TTL DNSRecordGetResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseCNAMEJSON `json:"-"`
}

// dnsRecordGetResponseCNAMEJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCNAME]
type dnsRecordGetResponseCNAMEJSON struct {
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

func (r *DNSRecordGetResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseCNAME) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseCNAMEType string

const (
	DNSRecordGetResponseCNAMETypeCNAME DNSRecordGetResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordGetResponseCNAMEMetaJSON `json:"-"`
}

// dnsRecordGetResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseCNAMEMeta]
type dnsRecordGetResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseCNAMETTLNumber].
type DNSRecordGetResponseCNAMETTL interface {
	ImplementsDNSRecordGetResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseCNAMETTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseCNAMETTLNumber float64

const (
	DNSRecordGetResponseCNAMETTLNumber1 DNSRecordGetResponseCNAMETTLNumber = 1
)

type DNSRecordGetResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data DNSRecordGetResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDNSKEYType `json:"type,required"`
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
	Meta DNSRecordGetResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDNSKEYJSON `json:"-"`
}

// dnsRecordGetResponseDNSKEYJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseDNSKEY]
type dnsRecordGetResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseDNSKEY) implementsDNSRecordGetResponse() {}

// Components of a DNSKEY record.
type DNSRecordGetResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                             `json:"public_key"`
	JSON      dnsRecordGetResponseDNSKEYDataJSON `json:"-"`
}

// dnsRecordGetResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseDNSKEYData]
type dnsRecordGetResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseDNSKEYType string

const (
	DNSRecordGetResponseDNSKEYTypeDNSKEY DNSRecordGetResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordGetResponseDNSKEYMetaJSON `json:"-"`
}

// dnsRecordGetResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseDNSKEYMeta]
type dnsRecordGetResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseDNSKEYTTLNumber].
type DNSRecordGetResponseDNSKEYTTL interface {
	ImplementsDNSRecordGetResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseDNSKEYTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseDNSKEYTTLNumber float64

const (
	DNSRecordGetResponseDNSKEYTTLNumber1 DNSRecordGetResponseDNSKEYTTLNumber = 1
)

type DNSRecordGetResponseDS struct {
	// Components of a DS record.
	Data DNSRecordGetResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseDSType `json:"type,required"`
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
	Meta DNSRecordGetResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseDSJSON `json:"-"`
}

// dnsRecordGetResponseDSJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseDS]
type dnsRecordGetResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseDS) implementsDNSRecordGetResponse() {}

// Components of a DS record.
type DNSRecordGetResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                        `json:"key_tag"`
	JSON   dnsRecordGetResponseDSDataJSON `json:"-"`
}

// dnsRecordGetResponseDSDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseDSData]
type dnsRecordGetResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseDSType string

const (
	DNSRecordGetResponseDSTypeDS DNSRecordGetResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordGetResponseDSMetaJSON `json:"-"`
}

// dnsRecordGetResponseDSMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseDSMeta]
type dnsRecordGetResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseDSTTLNumber].
type DNSRecordGetResponseDSTTL interface {
	ImplementsDNSRecordGetResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseDSTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseDSTTLNumber float64

const (
	DNSRecordGetResponseDSTTLNumber1 DNSRecordGetResponseDSTTLNumber = 1
)

type DNSRecordGetResponseHTTPS struct {
	// Components of a HTTPS record.
	Data DNSRecordGetResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseHTTPSType `json:"type,required"`
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
	Meta DNSRecordGetResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseHTTPSJSON `json:"-"`
}

// dnsRecordGetResponseHTTPSJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseHTTPS]
type dnsRecordGetResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseHTTPS) implementsDNSRecordGetResponse() {}

// Components of a HTTPS record.
type DNSRecordGetResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                            `json:"value"`
	JSON  dnsRecordGetResponseHTTPSDataJSON `json:"-"`
}

// dnsRecordGetResponseHTTPSDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseHTTPSData]
type dnsRecordGetResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseHTTPSType string

const (
	DNSRecordGetResponseHTTPSTypeHTTPS DNSRecordGetResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordGetResponseHTTPSMetaJSON `json:"-"`
}

// dnsRecordGetResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseHTTPSMeta]
type dnsRecordGetResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseHTTPSTTLNumber].
type DNSRecordGetResponseHTTPSTTL interface {
	ImplementsDNSRecordGetResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseHTTPSTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseHTTPSTTLNumber float64

const (
	DNSRecordGetResponseHTTPSTTLNumber1 DNSRecordGetResponseHTTPSTTLNumber = 1
)

type DNSRecordGetResponseLOC struct {
	// Components of a LOC record.
	Data DNSRecordGetResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseLOCType `json:"type,required"`
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
	Meta DNSRecordGetResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseLOCJSON `json:"-"`
}

// dnsRecordGetResponseLOCJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseLOC]
type dnsRecordGetResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseLOC) implementsDNSRecordGetResponse() {}

// Components of a LOC record.
type DNSRecordGetResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection DNSRecordGetResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection DNSRecordGetResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                         `json:"size"`
	JSON dnsRecordGetResponseLOCDataJSON `json:"-"`
}

// dnsRecordGetResponseLOCDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseLOCData]
type dnsRecordGetResponseLOCDataJSON struct {
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

func (r *DNSRecordGetResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type DNSRecordGetResponseLOCDataLatDirection string

const (
	DNSRecordGetResponseLOCDataLatDirectionN DNSRecordGetResponseLOCDataLatDirection = "N"
	DNSRecordGetResponseLOCDataLatDirectionS DNSRecordGetResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type DNSRecordGetResponseLOCDataLongDirection string

const (
	DNSRecordGetResponseLOCDataLongDirectionE DNSRecordGetResponseLOCDataLongDirection = "E"
	DNSRecordGetResponseLOCDataLongDirectionW DNSRecordGetResponseLOCDataLongDirection = "W"
)

// Record type.
type DNSRecordGetResponseLOCType string

const (
	DNSRecordGetResponseLOCTypeLOC DNSRecordGetResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordGetResponseLOCMetaJSON `json:"-"`
}

// dnsRecordGetResponseLOCMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseLOCMeta]
type dnsRecordGetResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseLOCTTLNumber].
type DNSRecordGetResponseLOCTTL interface {
	ImplementsDNSRecordGetResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseLOCTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseLOCTTLNumber float64

const (
	DNSRecordGetResponseLOCTTLNumber1 DNSRecordGetResponseLOCTTLNumber = 1
)

type DNSRecordGetResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseMXType `json:"type,required"`
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
	Meta DNSRecordGetResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseMXJSON `json:"-"`
}

// dnsRecordGetResponseMXJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseMX]
type dnsRecordGetResponseMXJSON struct {
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

func (r *DNSRecordGetResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseMX) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseMXType string

const (
	DNSRecordGetResponseMXTypeMX DNSRecordGetResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordGetResponseMXMetaJSON `json:"-"`
}

// dnsRecordGetResponseMXMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseMXMeta]
type dnsRecordGetResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseMXTTLNumber].
type DNSRecordGetResponseMXTTL interface {
	ImplementsDNSRecordGetResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseMXTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseMXTTLNumber float64

const (
	DNSRecordGetResponseMXTTLNumber1 DNSRecordGetResponseMXTTLNumber = 1
)

type DNSRecordGetResponseNAPTR struct {
	// Components of a NAPTR record.
	Data DNSRecordGetResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseNAPTRType `json:"type,required"`
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
	Meta DNSRecordGetResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseNAPTRJSON `json:"-"`
}

// dnsRecordGetResponseNAPTRJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseNAPTR]
type dnsRecordGetResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseNAPTR) implementsDNSRecordGetResponse() {}

// Components of a NAPTR record.
type DNSRecordGetResponseNAPTRData struct {
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
	Service string                            `json:"service"`
	JSON    dnsRecordGetResponseNAPTRDataJSON `json:"-"`
}

// dnsRecordGetResponseNAPTRDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseNAPTRData]
type dnsRecordGetResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseNAPTRType string

const (
	DNSRecordGetResponseNAPTRTypeNAPTR DNSRecordGetResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordGetResponseNAPTRMetaJSON `json:"-"`
}

// dnsRecordGetResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseNAPTRMeta]
type dnsRecordGetResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseNAPTRTTLNumber].
type DNSRecordGetResponseNAPTRTTL interface {
	ImplementsDNSRecordGetResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseNAPTRTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseNAPTRTTLNumber float64

const (
	DNSRecordGetResponseNAPTRTTLNumber1 DNSRecordGetResponseNAPTRTTLNumber = 1
)

type DNSRecordGetResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseNSType `json:"type,required"`
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
	Meta DNSRecordGetResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseNSJSON `json:"-"`
}

// dnsRecordGetResponseNSJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseNS]
type dnsRecordGetResponseNSJSON struct {
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

func (r *DNSRecordGetResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseNS) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseNSType string

const (
	DNSRecordGetResponseNSTypeNS DNSRecordGetResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   dnsRecordGetResponseNSMetaJSON `json:"-"`
}

// dnsRecordGetResponseNSMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseNSMeta]
type dnsRecordGetResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseNSTTLNumber].
type DNSRecordGetResponseNSTTL interface {
	ImplementsDNSRecordGetResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseNSTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseNSTTLNumber float64

const (
	DNSRecordGetResponseNSTTLNumber1 DNSRecordGetResponseNSTTLNumber = 1
)

type DNSRecordGetResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponsePTRType `json:"type,required"`
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
	Meta DNSRecordGetResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponsePTRJSON `json:"-"`
}

// dnsRecordGetResponsePTRJSON contains the JSON metadata for the struct
// [DNSRecordGetResponsePTR]
type dnsRecordGetResponsePTRJSON struct {
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

func (r *DNSRecordGetResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponsePTR) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponsePTRType string

const (
	DNSRecordGetResponsePTRTypePTR DNSRecordGetResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordGetResponsePTRMetaJSON `json:"-"`
}

// dnsRecordGetResponsePTRMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponsePTRMeta]
type dnsRecordGetResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponsePTRTTLNumber].
type DNSRecordGetResponsePTRTTL interface {
	ImplementsDNSRecordGetResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponsePTRTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponsePTRTTLNumber float64

const (
	DNSRecordGetResponsePTRTTLNumber1 DNSRecordGetResponsePTRTTLNumber = 1
)

type DNSRecordGetResponseSmimea struct {
	// Components of a SMIMEA record.
	Data DNSRecordGetResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseSmimeaType `json:"type,required"`
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
	Meta DNSRecordGetResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseSmimeaJSON `json:"-"`
}

// dnsRecordGetResponseSmimeaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSmimea]
type dnsRecordGetResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseSmimea) implementsDNSRecordGetResponse() {}

// Components of a SMIMEA record.
type DNSRecordGetResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                            `json:"usage"`
	JSON  dnsRecordGetResponseSmimeaDataJSON `json:"-"`
}

// dnsRecordGetResponseSmimeaDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSmimeaData]
type dnsRecordGetResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseSmimeaType string

const (
	DNSRecordGetResponseSmimeaTypeSmimea DNSRecordGetResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   dnsRecordGetResponseSmimeaMetaJSON `json:"-"`
}

// dnsRecordGetResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSmimeaMeta]
type dnsRecordGetResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseSmimeaTTLNumber].
type DNSRecordGetResponseSmimeaTTL interface {
	ImplementsDNSRecordGetResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseSmimeaTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseSmimeaTTLNumber float64

const (
	DNSRecordGetResponseSmimeaTTLNumber1 DNSRecordGetResponseSmimeaTTLNumber = 1
)

type DNSRecordGetResponseSRV struct {
	// Components of a SRV record.
	Data DNSRecordGetResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseSRVType `json:"type,required"`
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
	Meta DNSRecordGetResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseSRVJSON `json:"-"`
}

// dnsRecordGetResponseSRVJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSRV]
type dnsRecordGetResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseSRV) implementsDNSRecordGetResponse() {}

// Components of a SRV record.
type DNSRecordGetResponseSRVData struct {
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
	Weight float64                         `json:"weight"`
	JSON   dnsRecordGetResponseSRVDataJSON `json:"-"`
}

// dnsRecordGetResponseSRVDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSRVData]
type dnsRecordGetResponseSRVDataJSON struct {
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

func (r *DNSRecordGetResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseSRVType string

const (
	DNSRecordGetResponseSRVTypeSRV DNSRecordGetResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordGetResponseSRVMetaJSON `json:"-"`
}

// dnsRecordGetResponseSRVMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSRVMeta]
type dnsRecordGetResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseSRVTTLNumber].
type DNSRecordGetResponseSRVTTL interface {
	ImplementsDNSRecordGetResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseSRVTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseSRVTTLNumber float64

const (
	DNSRecordGetResponseSRVTTLNumber1 DNSRecordGetResponseSRVTTLNumber = 1
)

type DNSRecordGetResponseSSHFP struct {
	// Components of a SSHFP record.
	Data DNSRecordGetResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseSSHFPType `json:"type,required"`
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
	Meta DNSRecordGetResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseSSHFPJSON `json:"-"`
}

// dnsRecordGetResponseSSHFPJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSSHFP]
type dnsRecordGetResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseSSHFP) implementsDNSRecordGetResponse() {}

// Components of a SSHFP record.
type DNSRecordGetResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                           `json:"type"`
	JSON dnsRecordGetResponseSSHFPDataJSON `json:"-"`
}

// dnsRecordGetResponseSSHFPDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSSHFPData]
type dnsRecordGetResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseSSHFPType string

const (
	DNSRecordGetResponseSSHFPTypeSSHFP DNSRecordGetResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   dnsRecordGetResponseSSHFPMetaJSON `json:"-"`
}

// dnsRecordGetResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSSHFPMeta]
type dnsRecordGetResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseSSHFPTTLNumber].
type DNSRecordGetResponseSSHFPTTL interface {
	ImplementsDNSRecordGetResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseSSHFPTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseSSHFPTTLNumber float64

const (
	DNSRecordGetResponseSSHFPTTLNumber1 DNSRecordGetResponseSSHFPTTLNumber = 1
)

type DNSRecordGetResponseSVCB struct {
	// Components of a SVCB record.
	Data DNSRecordGetResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseSVCBType `json:"type,required"`
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
	Meta DNSRecordGetResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseSVCBJSON `json:"-"`
}

// dnsRecordGetResponseSVCBJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSVCB]
type dnsRecordGetResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseSVCB) implementsDNSRecordGetResponse() {}

// Components of a SVCB record.
type DNSRecordGetResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                           `json:"value"`
	JSON  dnsRecordGetResponseSVCBDataJSON `json:"-"`
}

// dnsRecordGetResponseSVCBDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSVCBData]
type dnsRecordGetResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseSVCBType string

const (
	DNSRecordGetResponseSVCBTypeSVCB DNSRecordGetResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordGetResponseSVCBMetaJSON `json:"-"`
}

// dnsRecordGetResponseSVCBMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseSVCBMeta]
type dnsRecordGetResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseSVCBTTLNumber].
type DNSRecordGetResponseSVCBTTL interface {
	ImplementsDNSRecordGetResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseSVCBTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseSVCBTTLNumber float64

const (
	DNSRecordGetResponseSVCBTTLNumber1 DNSRecordGetResponseSVCBTTLNumber = 1
)

type DNSRecordGetResponseTLSA struct {
	// Components of a TLSA record.
	Data DNSRecordGetResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseTLSAType `json:"type,required"`
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
	Meta DNSRecordGetResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseTLSAJSON `json:"-"`
}

// dnsRecordGetResponseTLSAJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseTLSA]
type dnsRecordGetResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseTLSA) implementsDNSRecordGetResponse() {}

// Components of a TLSA record.
type DNSRecordGetResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                          `json:"usage"`
	JSON  dnsRecordGetResponseTLSADataJSON `json:"-"`
}

// dnsRecordGetResponseTLSADataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseTLSAData]
type dnsRecordGetResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSRecordGetResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseTLSAType string

const (
	DNSRecordGetResponseTLSATypeTLSA DNSRecordGetResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   dnsRecordGetResponseTLSAMetaJSON `json:"-"`
}

// dnsRecordGetResponseTLSAMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseTLSAMeta]
type dnsRecordGetResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseTLSATTLNumber].
type DNSRecordGetResponseTLSATTL interface {
	ImplementsDNSRecordGetResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseTLSATTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseTLSATTLNumber float64

const (
	DNSRecordGetResponseTLSATTLNumber1 DNSRecordGetResponseTLSATTLNumber = 1
)

type DNSRecordGetResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type DNSRecordGetResponseTXTType `json:"type,required"`
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
	Meta DNSRecordGetResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseTXTJSON `json:"-"`
}

// dnsRecordGetResponseTXTJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseTXT]
type dnsRecordGetResponseTXTJSON struct {
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

func (r *DNSRecordGetResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseTXT) implementsDNSRecordGetResponse() {}

// Record type.
type DNSRecordGetResponseTXTType string

const (
	DNSRecordGetResponseTXTTypeTXT DNSRecordGetResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordGetResponseTXTMetaJSON `json:"-"`
}

// dnsRecordGetResponseTXTMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseTXTMeta]
type dnsRecordGetResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseTXTTTLNumber].
type DNSRecordGetResponseTXTTTL interface {
	ImplementsDNSRecordGetResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseTXTTTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseTXTTTLNumber float64

const (
	DNSRecordGetResponseTXTTTLNumber1 DNSRecordGetResponseTXTTTLNumber = 1
)

type DNSRecordGetResponseURI struct {
	// Components of a URI record.
	Data DNSRecordGetResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type DNSRecordGetResponseURIType `json:"type,required"`
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
	Meta DNSRecordGetResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL DNSRecordGetResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     dnsRecordGetResponseURIJSON `json:"-"`
}

// dnsRecordGetResponseURIJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseURI]
type dnsRecordGetResponseURIJSON struct {
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

func (r *DNSRecordGetResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r DNSRecordGetResponseURI) implementsDNSRecordGetResponse() {}

// Components of a URI record.
type DNSRecordGetResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                         `json:"weight"`
	JSON   dnsRecordGetResponseURIDataJSON `json:"-"`
}

// dnsRecordGetResponseURIDataJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseURIData]
type dnsRecordGetResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSRecordGetResponseURIType string

const (
	DNSRecordGetResponseURITypeURI DNSRecordGetResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type DNSRecordGetResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   dnsRecordGetResponseURIMetaJSON `json:"-"`
}

// dnsRecordGetResponseURIMetaJSON contains the JSON metadata for the struct
// [DNSRecordGetResponseURIMeta]
type dnsRecordGetResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordGetResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsRecordGetResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [DNSRecordGetResponseURITTLNumber].
type DNSRecordGetResponseURITTL interface {
	ImplementsDNSRecordGetResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DNSRecordGetResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(DNSRecordGetResponseURITTLNumber(0)),
		},
	)
}

type DNSRecordGetResponseURITTLNumber float64

const (
	DNSRecordGetResponseURITTLNumber1 DNSRecordGetResponseURITTLNumber = 1
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

func (r dnsRecordImportResponseJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordScanResponseJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordImportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordImportResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordImportResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordImportResponseEnvelopeTimingJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordScanResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordScanResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordScanResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r dnsRecordScanResponseEnvelopeTimingJSON) RawJSON() string {
	return r.raw
}
