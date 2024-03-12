// File generated from our OpenAPI spec by Stainless.

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
func (r *RecordService) New(ctx context.Context, params RecordNewParams, opts ...option.RequestOption) (res *RecordNewResponse, err error) {
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
func (r *RecordService) Update(ctx context.Context, dnsRecordID string, params RecordUpdateParams, opts ...option.RequestOption) (res *RecordUpdateResponse, err error) {
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
func (r *RecordService) List(ctx context.Context, params RecordListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RecordListResponse], err error) {
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
func (r *RecordService) ListAutoPaging(ctx context.Context, params RecordListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RecordListResponse] {
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
func (r *RecordService) Edit(ctx context.Context, dnsRecordID string, params RecordEditParams, opts ...option.RequestOption) (res *RecordEditResponse, err error) {
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
func (r *RecordService) Get(ctx context.Context, dnsRecordID string, query RecordGetParams, opts ...option.RequestOption) (res *RecordGetResponse, err error) {
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

// Union satisfied by [dns.RecordNewResponseA], [dns.RecordNewResponseAAAA],
// [dns.RecordNewResponseCAA], [dns.RecordNewResponseCert],
// [dns.RecordNewResponseCNAME], [dns.RecordNewResponseDNSKEY],
// [dns.RecordNewResponseDS], [dns.RecordNewResponseHTTPS],
// [dns.RecordNewResponseLOC], [dns.RecordNewResponseMX],
// [dns.RecordNewResponseNAPTR], [dns.RecordNewResponseNS],
// [dns.RecordNewResponsePTR], [dns.RecordNewResponseSmimea],
// [dns.RecordNewResponseSRV], [dns.RecordNewResponseSSHFP],
// [dns.RecordNewResponseSVCB], [dns.RecordNewResponseTLSA],
// [dns.RecordNewResponseTXT] or [dns.RecordNewResponseURI].
type RecordNewResponse interface {
	implementsDNSRecordNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordNewResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseAType `json:"type,required"`
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
	Meta RecordNewResponseAMeta `json:"meta"`
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
	TTL RecordNewResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                 `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseAJSON `json:"-"`
}

// recordNewResponseAJSON contains the JSON metadata for the struct
// [RecordNewResponseA]
type recordNewResponseAJSON struct {
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

func (r *RecordNewResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseAJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseA) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseAType string

const (
	RecordNewResponseATypeA RecordNewResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                     `json:"source"`
	JSON   recordNewResponseAMetaJSON `json:"-"`
}

// recordNewResponseAMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseAMeta]
type recordNewResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseATTLNumber].
type RecordNewResponseATTL interface {
	ImplementsDNSRecordNewResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseATTLNumber(0)),
		},
	)
}

type RecordNewResponseATTLNumber float64

const (
	RecordNewResponseATTLNumber1 RecordNewResponseATTLNumber = 1
)

type RecordNewResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseAAAAType `json:"type,required"`
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
	Meta RecordNewResponseAAAAMeta `json:"meta"`
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
	TTL RecordNewResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseAAAAJSON `json:"-"`
}

// recordNewResponseAAAAJSON contains the JSON metadata for the struct
// [RecordNewResponseAAAA]
type recordNewResponseAAAAJSON struct {
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

func (r *RecordNewResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseAAAA) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseAAAAType string

const (
	RecordNewResponseAAAATypeAAAA RecordNewResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordNewResponseAAAAMetaJSON `json:"-"`
}

// recordNewResponseAAAAMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseAAAAMeta]
type recordNewResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseAAAATTLNumber].
type RecordNewResponseAAAATTL interface {
	ImplementsDNSRecordNewResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseAAAATTLNumber(0)),
		},
	)
}

type RecordNewResponseAAAATTLNumber float64

const (
	RecordNewResponseAAAATTLNumber1 RecordNewResponseAAAATTLNumber = 1
)

type RecordNewResponseCAA struct {
	// Components of a CAA record.
	Data RecordNewResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseCAAType `json:"type,required"`
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
	Meta RecordNewResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseCAAJSON `json:"-"`
}

// recordNewResponseCAAJSON contains the JSON metadata for the struct
// [RecordNewResponseCAA]
type recordNewResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseCAA) implementsDNSRecordNewResponse() {}

// Components of a CAA record.
type RecordNewResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                       `json:"value"`
	JSON  recordNewResponseCAADataJSON `json:"-"`
}

// recordNewResponseCAADataJSON contains the JSON metadata for the struct
// [RecordNewResponseCAAData]
type recordNewResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseCAAType string

const (
	RecordNewResponseCAATypeCAA RecordNewResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordNewResponseCAAMetaJSON `json:"-"`
}

// recordNewResponseCAAMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseCAAMeta]
type recordNewResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseCAATTLNumber].
type RecordNewResponseCAATTL interface {
	ImplementsDNSRecordNewResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseCAATTLNumber(0)),
		},
	)
}

type RecordNewResponseCAATTLNumber float64

const (
	RecordNewResponseCAATTLNumber1 RecordNewResponseCAATTLNumber = 1
)

type RecordNewResponseCert struct {
	// Components of a CERT record.
	Data RecordNewResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseCertType `json:"type,required"`
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
	Meta RecordNewResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseCertJSON `json:"-"`
}

// recordNewResponseCertJSON contains the JSON metadata for the struct
// [RecordNewResponseCert]
type recordNewResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseCert) implementsDNSRecordNewResponse() {}

// Components of a CERT record.
type RecordNewResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                       `json:"type"`
	JSON recordNewResponseCertDataJSON `json:"-"`
}

// recordNewResponseCertDataJSON contains the JSON metadata for the struct
// [RecordNewResponseCertData]
type recordNewResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseCertType string

const (
	RecordNewResponseCertTypeCert RecordNewResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordNewResponseCertMetaJSON `json:"-"`
}

// recordNewResponseCertMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseCertMeta]
type recordNewResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseCertTTLNumber].
type RecordNewResponseCertTTL interface {
	ImplementsDNSRecordNewResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseCertTTLNumber(0)),
		},
	)
}

type RecordNewResponseCertTTLNumber float64

const (
	RecordNewResponseCertTTLNumber1 RecordNewResponseCertTTLNumber = 1
)

type RecordNewResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseCNAMEType `json:"type,required"`
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
	Meta RecordNewResponseCNAMEMeta `json:"meta"`
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
	TTL RecordNewResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseCNAMEJSON `json:"-"`
}

// recordNewResponseCNAMEJSON contains the JSON metadata for the struct
// [RecordNewResponseCNAME]
type recordNewResponseCNAMEJSON struct {
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

func (r *RecordNewResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseCNAME) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseCNAMEType string

const (
	RecordNewResponseCNAMETypeCNAME RecordNewResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordNewResponseCNAMEMetaJSON `json:"-"`
}

// recordNewResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseCNAMEMeta]
type recordNewResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseCNAMETTLNumber].
type RecordNewResponseCNAMETTL interface {
	ImplementsDNSRecordNewResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseCNAMETTLNumber(0)),
		},
	)
}

type RecordNewResponseCNAMETTLNumber float64

const (
	RecordNewResponseCNAMETTLNumber1 RecordNewResponseCNAMETTLNumber = 1
)

type RecordNewResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data RecordNewResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseDNSKEYType `json:"type,required"`
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
	Meta RecordNewResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseDNSKEYJSON `json:"-"`
}

// recordNewResponseDNSKEYJSON contains the JSON metadata for the struct
// [RecordNewResponseDNSKEY]
type recordNewResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseDNSKEY) implementsDNSRecordNewResponse() {}

// Components of a DNSKEY record.
type RecordNewResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                          `json:"public_key"`
	JSON      recordNewResponseDNSKEYDataJSON `json:"-"`
}

// recordNewResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [RecordNewResponseDNSKEYData]
type recordNewResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseDNSKEYType string

const (
	RecordNewResponseDNSKEYTypeDNSKEY RecordNewResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordNewResponseDNSKEYMetaJSON `json:"-"`
}

// recordNewResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseDNSKEYMeta]
type recordNewResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordNewResponseDNSKEYTTLNumber].
type RecordNewResponseDNSKEYTTL interface {
	ImplementsDNSRecordNewResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseDNSKEYTTLNumber(0)),
		},
	)
}

type RecordNewResponseDNSKEYTTLNumber float64

const (
	RecordNewResponseDNSKEYTTLNumber1 RecordNewResponseDNSKEYTTLNumber = 1
)

type RecordNewResponseDS struct {
	// Components of a DS record.
	Data RecordNewResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseDSType `json:"type,required"`
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
	Meta RecordNewResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseDSJSON `json:"-"`
}

// recordNewResponseDSJSON contains the JSON metadata for the struct
// [RecordNewResponseDS]
type recordNewResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseDS) implementsDNSRecordNewResponse() {}

// Components of a DS record.
type RecordNewResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                     `json:"key_tag"`
	JSON   recordNewResponseDSDataJSON `json:"-"`
}

// recordNewResponseDSDataJSON contains the JSON metadata for the struct
// [RecordNewResponseDSData]
type recordNewResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseDSType string

const (
	RecordNewResponseDSTypeDS RecordNewResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordNewResponseDSMetaJSON `json:"-"`
}

// recordNewResponseDSMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseDSMeta]
type recordNewResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseDSTTLNumber].
type RecordNewResponseDSTTL interface {
	ImplementsDNSRecordNewResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseDSTTLNumber(0)),
		},
	)
}

type RecordNewResponseDSTTLNumber float64

const (
	RecordNewResponseDSTTLNumber1 RecordNewResponseDSTTLNumber = 1
)

type RecordNewResponseHTTPS struct {
	// Components of a HTTPS record.
	Data RecordNewResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseHTTPSType `json:"type,required"`
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
	Meta RecordNewResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseHTTPSJSON `json:"-"`
}

// recordNewResponseHTTPSJSON contains the JSON metadata for the struct
// [RecordNewResponseHTTPS]
type recordNewResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseHTTPS) implementsDNSRecordNewResponse() {}

// Components of a HTTPS record.
type RecordNewResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                         `json:"value"`
	JSON  recordNewResponseHTTPSDataJSON `json:"-"`
}

// recordNewResponseHTTPSDataJSON contains the JSON metadata for the struct
// [RecordNewResponseHTTPSData]
type recordNewResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseHTTPSType string

const (
	RecordNewResponseHTTPSTypeHTTPS RecordNewResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordNewResponseHTTPSMetaJSON `json:"-"`
}

// recordNewResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseHTTPSMeta]
type recordNewResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseHTTPSTTLNumber].
type RecordNewResponseHTTPSTTL interface {
	ImplementsDNSRecordNewResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseHTTPSTTLNumber(0)),
		},
	)
}

type RecordNewResponseHTTPSTTLNumber float64

const (
	RecordNewResponseHTTPSTTLNumber1 RecordNewResponseHTTPSTTLNumber = 1
)

type RecordNewResponseLOC struct {
	// Components of a LOC record.
	Data RecordNewResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseLOCType `json:"type,required"`
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
	Meta RecordNewResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseLOCJSON `json:"-"`
}

// recordNewResponseLOCJSON contains the JSON metadata for the struct
// [RecordNewResponseLOC]
type recordNewResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseLOC) implementsDNSRecordNewResponse() {}

// Components of a LOC record.
type RecordNewResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection RecordNewResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection RecordNewResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                      `json:"size"`
	JSON recordNewResponseLOCDataJSON `json:"-"`
}

// recordNewResponseLOCDataJSON contains the JSON metadata for the struct
// [RecordNewResponseLOCData]
type recordNewResponseLOCDataJSON struct {
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

func (r *RecordNewResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type RecordNewResponseLOCDataLatDirection string

const (
	RecordNewResponseLOCDataLatDirectionN RecordNewResponseLOCDataLatDirection = "N"
	RecordNewResponseLOCDataLatDirectionS RecordNewResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type RecordNewResponseLOCDataLongDirection string

const (
	RecordNewResponseLOCDataLongDirectionE RecordNewResponseLOCDataLongDirection = "E"
	RecordNewResponseLOCDataLongDirectionW RecordNewResponseLOCDataLongDirection = "W"
)

// Record type.
type RecordNewResponseLOCType string

const (
	RecordNewResponseLOCTypeLOC RecordNewResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordNewResponseLOCMetaJSON `json:"-"`
}

// recordNewResponseLOCMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseLOCMeta]
type recordNewResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseLOCTTLNumber].
type RecordNewResponseLOCTTL interface {
	ImplementsDNSRecordNewResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseLOCTTLNumber(0)),
		},
	)
}

type RecordNewResponseLOCTTLNumber float64

const (
	RecordNewResponseLOCTTLNumber1 RecordNewResponseLOCTTLNumber = 1
)

type RecordNewResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordNewResponseMXType `json:"type,required"`
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
	Meta RecordNewResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseMXJSON `json:"-"`
}

// recordNewResponseMXJSON contains the JSON metadata for the struct
// [RecordNewResponseMX]
type recordNewResponseMXJSON struct {
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

func (r *RecordNewResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseMX) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseMXType string

const (
	RecordNewResponseMXTypeMX RecordNewResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordNewResponseMXMetaJSON `json:"-"`
}

// recordNewResponseMXMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseMXMeta]
type recordNewResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseMXTTLNumber].
type RecordNewResponseMXTTL interface {
	ImplementsDNSRecordNewResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseMXTTLNumber(0)),
		},
	)
}

type RecordNewResponseMXTTLNumber float64

const (
	RecordNewResponseMXTTLNumber1 RecordNewResponseMXTTLNumber = 1
)

type RecordNewResponseNAPTR struct {
	// Components of a NAPTR record.
	Data RecordNewResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseNAPTRType `json:"type,required"`
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
	Meta RecordNewResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseNAPTRJSON `json:"-"`
}

// recordNewResponseNAPTRJSON contains the JSON metadata for the struct
// [RecordNewResponseNAPTR]
type recordNewResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseNAPTR) implementsDNSRecordNewResponse() {}

// Components of a NAPTR record.
type RecordNewResponseNAPTRData struct {
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
	Service string                         `json:"service"`
	JSON    recordNewResponseNAPTRDataJSON `json:"-"`
}

// recordNewResponseNAPTRDataJSON contains the JSON metadata for the struct
// [RecordNewResponseNAPTRData]
type recordNewResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseNAPTRType string

const (
	RecordNewResponseNAPTRTypeNAPTR RecordNewResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordNewResponseNAPTRMetaJSON `json:"-"`
}

// recordNewResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseNAPTRMeta]
type recordNewResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseNAPTRTTLNumber].
type RecordNewResponseNAPTRTTL interface {
	ImplementsDNSRecordNewResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseNAPTRTTLNumber(0)),
		},
	)
}

type RecordNewResponseNAPTRTTLNumber float64

const (
	RecordNewResponseNAPTRTTLNumber1 RecordNewResponseNAPTRTTLNumber = 1
)

type RecordNewResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseNSType `json:"type,required"`
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
	Meta RecordNewResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseNSJSON `json:"-"`
}

// recordNewResponseNSJSON contains the JSON metadata for the struct
// [RecordNewResponseNS]
type recordNewResponseNSJSON struct {
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

func (r *RecordNewResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseNS) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseNSType string

const (
	RecordNewResponseNSTypeNS RecordNewResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordNewResponseNSMetaJSON `json:"-"`
}

// recordNewResponseNSMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseNSMeta]
type recordNewResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseNSTTLNumber].
type RecordNewResponseNSTTL interface {
	ImplementsDNSRecordNewResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseNSTTLNumber(0)),
		},
	)
}

type RecordNewResponseNSTTLNumber float64

const (
	RecordNewResponseNSTTLNumber1 RecordNewResponseNSTTLNumber = 1
)

type RecordNewResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponsePTRType `json:"type,required"`
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
	Meta RecordNewResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordNewResponsePTRJSON `json:"-"`
}

// recordNewResponsePTRJSON contains the JSON metadata for the struct
// [RecordNewResponsePTR]
type recordNewResponsePTRJSON struct {
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

func (r *RecordNewResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponsePTR) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponsePTRType string

const (
	RecordNewResponsePTRTypePTR RecordNewResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordNewResponsePTRMetaJSON `json:"-"`
}

// recordNewResponsePTRMetaJSON contains the JSON metadata for the struct
// [RecordNewResponsePTRMeta]
type recordNewResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponsePTRTTLNumber].
type RecordNewResponsePTRTTL interface {
	ImplementsDNSRecordNewResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponsePTRTTLNumber(0)),
		},
	)
}

type RecordNewResponsePTRTTLNumber float64

const (
	RecordNewResponsePTRTTLNumber1 RecordNewResponsePTRTTLNumber = 1
)

type RecordNewResponseSmimea struct {
	// Components of a SMIMEA record.
	Data RecordNewResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseSmimeaType `json:"type,required"`
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
	Meta RecordNewResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseSmimeaJSON `json:"-"`
}

// recordNewResponseSmimeaJSON contains the JSON metadata for the struct
// [RecordNewResponseSmimea]
type recordNewResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSmimea) implementsDNSRecordNewResponse() {}

// Components of a SMIMEA record.
type RecordNewResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                         `json:"usage"`
	JSON  recordNewResponseSmimeaDataJSON `json:"-"`
}

// recordNewResponseSmimeaDataJSON contains the JSON metadata for the struct
// [RecordNewResponseSmimeaData]
type recordNewResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordNewResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseSmimeaType string

const (
	RecordNewResponseSmimeaTypeSmimea RecordNewResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordNewResponseSmimeaMetaJSON `json:"-"`
}

// recordNewResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseSmimeaMeta]
type recordNewResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordNewResponseSmimeaTTLNumber].
type RecordNewResponseSmimeaTTL interface {
	ImplementsDNSRecordNewResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseSmimeaTTLNumber(0)),
		},
	)
}

type RecordNewResponseSmimeaTTLNumber float64

const (
	RecordNewResponseSmimeaTTLNumber1 RecordNewResponseSmimeaTTLNumber = 1
)

type RecordNewResponseSRV struct {
	// Components of a SRV record.
	Data RecordNewResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseSRVType `json:"type,required"`
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
	Meta RecordNewResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseSRVJSON `json:"-"`
}

// recordNewResponseSRVJSON contains the JSON metadata for the struct
// [RecordNewResponseSRV]
type recordNewResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSRV) implementsDNSRecordNewResponse() {}

// Components of a SRV record.
type RecordNewResponseSRVData struct {
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
	Weight float64                      `json:"weight"`
	JSON   recordNewResponseSRVDataJSON `json:"-"`
}

// recordNewResponseSRVDataJSON contains the JSON metadata for the struct
// [RecordNewResponseSRVData]
type recordNewResponseSRVDataJSON struct {
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

func (r *RecordNewResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseSRVType string

const (
	RecordNewResponseSRVTypeSRV RecordNewResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordNewResponseSRVMetaJSON `json:"-"`
}

// recordNewResponseSRVMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseSRVMeta]
type recordNewResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseSRVTTLNumber].
type RecordNewResponseSRVTTL interface {
	ImplementsDNSRecordNewResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseSRVTTLNumber(0)),
		},
	)
}

type RecordNewResponseSRVTTLNumber float64

const (
	RecordNewResponseSRVTTLNumber1 RecordNewResponseSRVTTLNumber = 1
)

type RecordNewResponseSSHFP struct {
	// Components of a SSHFP record.
	Data RecordNewResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseSSHFPType `json:"type,required"`
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
	Meta RecordNewResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseSSHFPJSON `json:"-"`
}

// recordNewResponseSSHFPJSON contains the JSON metadata for the struct
// [RecordNewResponseSSHFP]
type recordNewResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSSHFP) implementsDNSRecordNewResponse() {}

// Components of a SSHFP record.
type RecordNewResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                        `json:"type"`
	JSON recordNewResponseSSHFPDataJSON `json:"-"`
}

// recordNewResponseSSHFPDataJSON contains the JSON metadata for the struct
// [RecordNewResponseSSHFPData]
type recordNewResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseSSHFPType string

const (
	RecordNewResponseSSHFPTypeSSHFP RecordNewResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordNewResponseSSHFPMetaJSON `json:"-"`
}

// recordNewResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseSSHFPMeta]
type recordNewResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseSSHFPTTLNumber].
type RecordNewResponseSSHFPTTL interface {
	ImplementsDNSRecordNewResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseSSHFPTTLNumber(0)),
		},
	)
}

type RecordNewResponseSSHFPTTLNumber float64

const (
	RecordNewResponseSSHFPTTLNumber1 RecordNewResponseSSHFPTTLNumber = 1
)

type RecordNewResponseSVCB struct {
	// Components of a SVCB record.
	Data RecordNewResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseSVCBType `json:"type,required"`
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
	Meta RecordNewResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseSVCBJSON `json:"-"`
}

// recordNewResponseSVCBJSON contains the JSON metadata for the struct
// [RecordNewResponseSVCB]
type recordNewResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSVCB) implementsDNSRecordNewResponse() {}

// Components of a SVCB record.
type RecordNewResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                        `json:"value"`
	JSON  recordNewResponseSVCBDataJSON `json:"-"`
}

// recordNewResponseSVCBDataJSON contains the JSON metadata for the struct
// [RecordNewResponseSVCBData]
type recordNewResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseSVCBType string

const (
	RecordNewResponseSVCBTypeSVCB RecordNewResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordNewResponseSVCBMetaJSON `json:"-"`
}

// recordNewResponseSVCBMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseSVCBMeta]
type recordNewResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseSVCBTTLNumber].
type RecordNewResponseSVCBTTL interface {
	ImplementsDNSRecordNewResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseSVCBTTLNumber(0)),
		},
	)
}

type RecordNewResponseSVCBTTLNumber float64

const (
	RecordNewResponseSVCBTTLNumber1 RecordNewResponseSVCBTTLNumber = 1
)

type RecordNewResponseTLSA struct {
	// Components of a TLSA record.
	Data RecordNewResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseTLSAType `json:"type,required"`
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
	Meta RecordNewResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseTLSAJSON `json:"-"`
}

// recordNewResponseTLSAJSON contains the JSON metadata for the struct
// [RecordNewResponseTLSA]
type recordNewResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseTLSA) implementsDNSRecordNewResponse() {}

// Components of a TLSA record.
type RecordNewResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                       `json:"usage"`
	JSON  recordNewResponseTLSADataJSON `json:"-"`
}

// recordNewResponseTLSADataJSON contains the JSON metadata for the struct
// [RecordNewResponseTLSAData]
type recordNewResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordNewResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseTLSAType string

const (
	RecordNewResponseTLSATypeTLSA RecordNewResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordNewResponseTLSAMetaJSON `json:"-"`
}

// recordNewResponseTLSAMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseTLSAMeta]
type recordNewResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseTLSATTLNumber].
type RecordNewResponseTLSATTL interface {
	ImplementsDNSRecordNewResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseTLSATTLNumber(0)),
		},
	)
}

type RecordNewResponseTLSATTLNumber float64

const (
	RecordNewResponseTLSATTLNumber1 RecordNewResponseTLSATTLNumber = 1
)

type RecordNewResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordNewResponseTXTType `json:"type,required"`
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
	Meta RecordNewResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseTXTJSON `json:"-"`
}

// recordNewResponseTXTJSON contains the JSON metadata for the struct
// [RecordNewResponseTXT]
type recordNewResponseTXTJSON struct {
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

func (r *RecordNewResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseTXT) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseTXTType string

const (
	RecordNewResponseTXTTypeTXT RecordNewResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordNewResponseTXTMetaJSON `json:"-"`
}

// recordNewResponseTXTMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseTXTMeta]
type recordNewResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseTXTTTLNumber].
type RecordNewResponseTXTTTL interface {
	ImplementsDNSRecordNewResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseTXTTTLNumber(0)),
		},
	)
}

type RecordNewResponseTXTTTLNumber float64

const (
	RecordNewResponseTXTTTLNumber1 RecordNewResponseTXTTTLNumber = 1
)

type RecordNewResponseURI struct {
	// Components of a URI record.
	Data RecordNewResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordNewResponseURIType `json:"type,required"`
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
	Meta RecordNewResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordNewResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordNewResponseURIJSON `json:"-"`
}

// recordNewResponseURIJSON contains the JSON metadata for the struct
// [RecordNewResponseURI]
type recordNewResponseURIJSON struct {
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

func (r *RecordNewResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseURI) implementsDNSRecordNewResponse() {}

// Components of a URI record.
type RecordNewResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                      `json:"weight"`
	JSON   recordNewResponseURIDataJSON `json:"-"`
}

// recordNewResponseURIDataJSON contains the JSON metadata for the struct
// [RecordNewResponseURIData]
type recordNewResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordNewResponseURIType string

const (
	RecordNewResponseURITypeURI RecordNewResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type RecordNewResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordNewResponseURIMetaJSON `json:"-"`
}

// recordNewResponseURIMetaJSON contains the JSON metadata for the struct
// [RecordNewResponseURIMeta]
type recordNewResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordNewResponseURITTLNumber].
type RecordNewResponseURITTL interface {
	ImplementsDNSRecordNewResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordNewResponseURITTLNumber(0)),
		},
	)
}

type RecordNewResponseURITTLNumber float64

const (
	RecordNewResponseURITTLNumber1 RecordNewResponseURITTLNumber = 1
)

// Union satisfied by [dns.RecordUpdateResponseA], [dns.RecordUpdateResponseAAAA],
// [dns.RecordUpdateResponseCAA], [dns.RecordUpdateResponseCert],
// [dns.RecordUpdateResponseCNAME], [dns.RecordUpdateResponseDNSKEY],
// [dns.RecordUpdateResponseDS], [dns.RecordUpdateResponseHTTPS],
// [dns.RecordUpdateResponseLOC], [dns.RecordUpdateResponseMX],
// [dns.RecordUpdateResponseNAPTR], [dns.RecordUpdateResponseNS],
// [dns.RecordUpdateResponsePTR], [dns.RecordUpdateResponseSmimea],
// [dns.RecordUpdateResponseSRV], [dns.RecordUpdateResponseSSHFP],
// [dns.RecordUpdateResponseSVCB], [dns.RecordUpdateResponseTLSA],
// [dns.RecordUpdateResponseTXT] or [dns.RecordUpdateResponseURI].
type RecordUpdateResponse interface {
	implementsDNSRecordUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordUpdateResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseAType `json:"type,required"`
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
	Meta RecordUpdateResponseAMeta `json:"meta"`
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
	TTL RecordUpdateResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseAJSON `json:"-"`
}

// recordUpdateResponseAJSON contains the JSON metadata for the struct
// [RecordUpdateResponseA]
type recordUpdateResponseAJSON struct {
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

func (r *RecordUpdateResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseAJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseA) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseAType string

const (
	RecordUpdateResponseATypeA RecordUpdateResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordUpdateResponseAMetaJSON `json:"-"`
}

// recordUpdateResponseAMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseAMeta]
type recordUpdateResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordUpdateResponseATTLNumber].
type RecordUpdateResponseATTL interface {
	ImplementsDNSRecordUpdateResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseATTLNumber(0)),
		},
	)
}

type RecordUpdateResponseATTLNumber float64

const (
	RecordUpdateResponseATTLNumber1 RecordUpdateResponseATTLNumber = 1
)

type RecordUpdateResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseAAAAType `json:"type,required"`
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
	Meta RecordUpdateResponseAAAAMeta `json:"meta"`
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
	TTL RecordUpdateResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseAAAAJSON `json:"-"`
}

// recordUpdateResponseAAAAJSON contains the JSON metadata for the struct
// [RecordUpdateResponseAAAA]
type recordUpdateResponseAAAAJSON struct {
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

func (r *RecordUpdateResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseAAAA) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseAAAAType string

const (
	RecordUpdateResponseAAAATypeAAAA RecordUpdateResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordUpdateResponseAAAAMetaJSON `json:"-"`
}

// recordUpdateResponseAAAAMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseAAAAMeta]
type recordUpdateResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseAAAATTLNumber].
type RecordUpdateResponseAAAATTL interface {
	ImplementsDNSRecordUpdateResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseAAAATTLNumber(0)),
		},
	)
}

type RecordUpdateResponseAAAATTLNumber float64

const (
	RecordUpdateResponseAAAATTLNumber1 RecordUpdateResponseAAAATTLNumber = 1
)

type RecordUpdateResponseCAA struct {
	// Components of a CAA record.
	Data RecordUpdateResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseCAAType `json:"type,required"`
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
	Meta RecordUpdateResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseCAAJSON `json:"-"`
}

// recordUpdateResponseCAAJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCAA]
type recordUpdateResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseCAA) implementsDNSRecordUpdateResponse() {}

// Components of a CAA record.
type RecordUpdateResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                          `json:"value"`
	JSON  recordUpdateResponseCAADataJSON `json:"-"`
}

// recordUpdateResponseCAADataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCAAData]
type recordUpdateResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseCAAType string

const (
	RecordUpdateResponseCAATypeCAA RecordUpdateResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordUpdateResponseCAAMetaJSON `json:"-"`
}

// recordUpdateResponseCAAMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCAAMeta]
type recordUpdateResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseCAATTLNumber].
type RecordUpdateResponseCAATTL interface {
	ImplementsDNSRecordUpdateResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseCAATTLNumber(0)),
		},
	)
}

type RecordUpdateResponseCAATTLNumber float64

const (
	RecordUpdateResponseCAATTLNumber1 RecordUpdateResponseCAATTLNumber = 1
)

type RecordUpdateResponseCert struct {
	// Components of a CERT record.
	Data RecordUpdateResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseCertType `json:"type,required"`
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
	Meta RecordUpdateResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseCertJSON `json:"-"`
}

// recordUpdateResponseCertJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCert]
type recordUpdateResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseCert) implementsDNSRecordUpdateResponse() {}

// Components of a CERT record.
type RecordUpdateResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                          `json:"type"`
	JSON recordUpdateResponseCertDataJSON `json:"-"`
}

// recordUpdateResponseCertDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCertData]
type recordUpdateResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseCertType string

const (
	RecordUpdateResponseCertTypeCert RecordUpdateResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordUpdateResponseCertMetaJSON `json:"-"`
}

// recordUpdateResponseCertMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCertMeta]
type recordUpdateResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseCertTTLNumber].
type RecordUpdateResponseCertTTL interface {
	ImplementsDNSRecordUpdateResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseCertTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseCertTTLNumber float64

const (
	RecordUpdateResponseCertTTLNumber1 RecordUpdateResponseCertTTLNumber = 1
)

type RecordUpdateResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseCNAMEType `json:"type,required"`
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
	Meta RecordUpdateResponseCNAMEMeta `json:"meta"`
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
	TTL RecordUpdateResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseCNAMEJSON `json:"-"`
}

// recordUpdateResponseCNAMEJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCNAME]
type recordUpdateResponseCNAMEJSON struct {
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

func (r *RecordUpdateResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseCNAME) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseCNAMEType string

const (
	RecordUpdateResponseCNAMETypeCNAME RecordUpdateResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   recordUpdateResponseCNAMEMetaJSON `json:"-"`
}

// recordUpdateResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCNAMEMeta]
type recordUpdateResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseCNAMETTLNumber].
type RecordUpdateResponseCNAMETTL interface {
	ImplementsDNSRecordUpdateResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseCNAMETTLNumber(0)),
		},
	)
}

type RecordUpdateResponseCNAMETTLNumber float64

const (
	RecordUpdateResponseCNAMETTLNumber1 RecordUpdateResponseCNAMETTLNumber = 1
)

type RecordUpdateResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data RecordUpdateResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseDNSKEYType `json:"type,required"`
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
	Meta RecordUpdateResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseDNSKEYJSON `json:"-"`
}

// recordUpdateResponseDNSKEYJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDNSKEY]
type recordUpdateResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseDNSKEY) implementsDNSRecordUpdateResponse() {}

// Components of a DNSKEY record.
type RecordUpdateResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                             `json:"public_key"`
	JSON      recordUpdateResponseDNSKEYDataJSON `json:"-"`
}

// recordUpdateResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDNSKEYData]
type recordUpdateResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseDNSKEYType string

const (
	RecordUpdateResponseDNSKEYTypeDNSKEY RecordUpdateResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   recordUpdateResponseDNSKEYMetaJSON `json:"-"`
}

// recordUpdateResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDNSKEYMeta]
type recordUpdateResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseDNSKEYTTLNumber].
type RecordUpdateResponseDNSKEYTTL interface {
	ImplementsDNSRecordUpdateResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseDNSKEYTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseDNSKEYTTLNumber float64

const (
	RecordUpdateResponseDNSKEYTTLNumber1 RecordUpdateResponseDNSKEYTTLNumber = 1
)

type RecordUpdateResponseDS struct {
	// Components of a DS record.
	Data RecordUpdateResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseDSType `json:"type,required"`
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
	Meta RecordUpdateResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseDSJSON `json:"-"`
}

// recordUpdateResponseDSJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDS]
type recordUpdateResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseDS) implementsDNSRecordUpdateResponse() {}

// Components of a DS record.
type RecordUpdateResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                        `json:"key_tag"`
	JSON   recordUpdateResponseDSDataJSON `json:"-"`
}

// recordUpdateResponseDSDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDSData]
type recordUpdateResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseDSType string

const (
	RecordUpdateResponseDSTypeDS RecordUpdateResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordUpdateResponseDSMetaJSON `json:"-"`
}

// recordUpdateResponseDSMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDSMeta]
type recordUpdateResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordUpdateResponseDSTTLNumber].
type RecordUpdateResponseDSTTL interface {
	ImplementsDNSRecordUpdateResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseDSTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseDSTTLNumber float64

const (
	RecordUpdateResponseDSTTLNumber1 RecordUpdateResponseDSTTLNumber = 1
)

type RecordUpdateResponseHTTPS struct {
	// Components of a HTTPS record.
	Data RecordUpdateResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseHTTPSType `json:"type,required"`
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
	Meta RecordUpdateResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseHTTPSJSON `json:"-"`
}

// recordUpdateResponseHTTPSJSON contains the JSON metadata for the struct
// [RecordUpdateResponseHTTPS]
type recordUpdateResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseHTTPS) implementsDNSRecordUpdateResponse() {}

// Components of a HTTPS record.
type RecordUpdateResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                            `json:"value"`
	JSON  recordUpdateResponseHTTPSDataJSON `json:"-"`
}

// recordUpdateResponseHTTPSDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseHTTPSData]
type recordUpdateResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseHTTPSType string

const (
	RecordUpdateResponseHTTPSTypeHTTPS RecordUpdateResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   recordUpdateResponseHTTPSMetaJSON `json:"-"`
}

// recordUpdateResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseHTTPSMeta]
type recordUpdateResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseHTTPSTTLNumber].
type RecordUpdateResponseHTTPSTTL interface {
	ImplementsDNSRecordUpdateResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseHTTPSTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseHTTPSTTLNumber float64

const (
	RecordUpdateResponseHTTPSTTLNumber1 RecordUpdateResponseHTTPSTTLNumber = 1
)

type RecordUpdateResponseLOC struct {
	// Components of a LOC record.
	Data RecordUpdateResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseLOCType `json:"type,required"`
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
	Meta RecordUpdateResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseLOCJSON `json:"-"`
}

// recordUpdateResponseLOCJSON contains the JSON metadata for the struct
// [RecordUpdateResponseLOC]
type recordUpdateResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseLOC) implementsDNSRecordUpdateResponse() {}

// Components of a LOC record.
type RecordUpdateResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection RecordUpdateResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection RecordUpdateResponseLOCDataLongDirection `json:"long_direction"`
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
	JSON recordUpdateResponseLOCDataJSON `json:"-"`
}

// recordUpdateResponseLOCDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseLOCData]
type recordUpdateResponseLOCDataJSON struct {
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

func (r *RecordUpdateResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type RecordUpdateResponseLOCDataLatDirection string

const (
	RecordUpdateResponseLOCDataLatDirectionN RecordUpdateResponseLOCDataLatDirection = "N"
	RecordUpdateResponseLOCDataLatDirectionS RecordUpdateResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type RecordUpdateResponseLOCDataLongDirection string

const (
	RecordUpdateResponseLOCDataLongDirectionE RecordUpdateResponseLOCDataLongDirection = "E"
	RecordUpdateResponseLOCDataLongDirectionW RecordUpdateResponseLOCDataLongDirection = "W"
)

// Record type.
type RecordUpdateResponseLOCType string

const (
	RecordUpdateResponseLOCTypeLOC RecordUpdateResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordUpdateResponseLOCMetaJSON `json:"-"`
}

// recordUpdateResponseLOCMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseLOCMeta]
type recordUpdateResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseLOCTTLNumber].
type RecordUpdateResponseLOCTTL interface {
	ImplementsDNSRecordUpdateResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseLOCTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseLOCTTLNumber float64

const (
	RecordUpdateResponseLOCTTLNumber1 RecordUpdateResponseLOCTTLNumber = 1
)

type RecordUpdateResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordUpdateResponseMXType `json:"type,required"`
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
	Meta RecordUpdateResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseMXJSON `json:"-"`
}

// recordUpdateResponseMXJSON contains the JSON metadata for the struct
// [RecordUpdateResponseMX]
type recordUpdateResponseMXJSON struct {
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

func (r *RecordUpdateResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseMX) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseMXType string

const (
	RecordUpdateResponseMXTypeMX RecordUpdateResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordUpdateResponseMXMetaJSON `json:"-"`
}

// recordUpdateResponseMXMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseMXMeta]
type recordUpdateResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordUpdateResponseMXTTLNumber].
type RecordUpdateResponseMXTTL interface {
	ImplementsDNSRecordUpdateResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseMXTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseMXTTLNumber float64

const (
	RecordUpdateResponseMXTTLNumber1 RecordUpdateResponseMXTTLNumber = 1
)

type RecordUpdateResponseNAPTR struct {
	// Components of a NAPTR record.
	Data RecordUpdateResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseNAPTRType `json:"type,required"`
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
	Meta RecordUpdateResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseNAPTRJSON `json:"-"`
}

// recordUpdateResponseNAPTRJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNAPTR]
type recordUpdateResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseNAPTR) implementsDNSRecordUpdateResponse() {}

// Components of a NAPTR record.
type RecordUpdateResponseNAPTRData struct {
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
	JSON    recordUpdateResponseNAPTRDataJSON `json:"-"`
}

// recordUpdateResponseNAPTRDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNAPTRData]
type recordUpdateResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseNAPTRType string

const (
	RecordUpdateResponseNAPTRTypeNAPTR RecordUpdateResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   recordUpdateResponseNAPTRMetaJSON `json:"-"`
}

// recordUpdateResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNAPTRMeta]
type recordUpdateResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseNAPTRTTLNumber].
type RecordUpdateResponseNAPTRTTL interface {
	ImplementsDNSRecordUpdateResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseNAPTRTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseNAPTRTTLNumber float64

const (
	RecordUpdateResponseNAPTRTTLNumber1 RecordUpdateResponseNAPTRTTLNumber = 1
)

type RecordUpdateResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseNSType `json:"type,required"`
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
	Meta RecordUpdateResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseNSJSON `json:"-"`
}

// recordUpdateResponseNSJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNS]
type recordUpdateResponseNSJSON struct {
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

func (r *RecordUpdateResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseNS) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseNSType string

const (
	RecordUpdateResponseNSTypeNS RecordUpdateResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordUpdateResponseNSMetaJSON `json:"-"`
}

// recordUpdateResponseNSMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNSMeta]
type recordUpdateResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordUpdateResponseNSTTLNumber].
type RecordUpdateResponseNSTTL interface {
	ImplementsDNSRecordUpdateResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseNSTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseNSTTLNumber float64

const (
	RecordUpdateResponseNSTTLNumber1 RecordUpdateResponseNSTTLNumber = 1
)

type RecordUpdateResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponsePTRType `json:"type,required"`
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
	Meta RecordUpdateResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponsePTRJSON `json:"-"`
}

// recordUpdateResponsePTRJSON contains the JSON metadata for the struct
// [RecordUpdateResponsePTR]
type recordUpdateResponsePTRJSON struct {
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

func (r *RecordUpdateResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponsePTR) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponsePTRType string

const (
	RecordUpdateResponsePTRTypePTR RecordUpdateResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordUpdateResponsePTRMetaJSON `json:"-"`
}

// recordUpdateResponsePTRMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponsePTRMeta]
type recordUpdateResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponsePTRTTLNumber].
type RecordUpdateResponsePTRTTL interface {
	ImplementsDNSRecordUpdateResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponsePTRTTLNumber(0)),
		},
	)
}

type RecordUpdateResponsePTRTTLNumber float64

const (
	RecordUpdateResponsePTRTTLNumber1 RecordUpdateResponsePTRTTLNumber = 1
)

type RecordUpdateResponseSmimea struct {
	// Components of a SMIMEA record.
	Data RecordUpdateResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseSmimeaType `json:"type,required"`
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
	Meta RecordUpdateResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                         `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseSmimeaJSON `json:"-"`
}

// recordUpdateResponseSmimeaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSmimea]
type recordUpdateResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSmimea) implementsDNSRecordUpdateResponse() {}

// Components of a SMIMEA record.
type RecordUpdateResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                            `json:"usage"`
	JSON  recordUpdateResponseSmimeaDataJSON `json:"-"`
}

// recordUpdateResponseSmimeaDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSmimeaData]
type recordUpdateResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordUpdateResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseSmimeaType string

const (
	RecordUpdateResponseSmimeaTypeSmimea RecordUpdateResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                             `json:"source"`
	JSON   recordUpdateResponseSmimeaMetaJSON `json:"-"`
}

// recordUpdateResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSmimeaMeta]
type recordUpdateResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseSmimeaTTLNumber].
type RecordUpdateResponseSmimeaTTL interface {
	ImplementsDNSRecordUpdateResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseSmimeaTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseSmimeaTTLNumber float64

const (
	RecordUpdateResponseSmimeaTTLNumber1 RecordUpdateResponseSmimeaTTLNumber = 1
)

type RecordUpdateResponseSRV struct {
	// Components of a SRV record.
	Data RecordUpdateResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseSRVType `json:"type,required"`
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
	Meta RecordUpdateResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseSRVJSON `json:"-"`
}

// recordUpdateResponseSRVJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSRV]
type recordUpdateResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSRV) implementsDNSRecordUpdateResponse() {}

// Components of a SRV record.
type RecordUpdateResponseSRVData struct {
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
	JSON   recordUpdateResponseSRVDataJSON `json:"-"`
}

// recordUpdateResponseSRVDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSRVData]
type recordUpdateResponseSRVDataJSON struct {
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

func (r *RecordUpdateResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseSRVType string

const (
	RecordUpdateResponseSRVTypeSRV RecordUpdateResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordUpdateResponseSRVMetaJSON `json:"-"`
}

// recordUpdateResponseSRVMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSRVMeta]
type recordUpdateResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseSRVTTLNumber].
type RecordUpdateResponseSRVTTL interface {
	ImplementsDNSRecordUpdateResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseSRVTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseSRVTTLNumber float64

const (
	RecordUpdateResponseSRVTTLNumber1 RecordUpdateResponseSRVTTLNumber = 1
)

type RecordUpdateResponseSSHFP struct {
	// Components of a SSHFP record.
	Data RecordUpdateResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseSSHFPType `json:"type,required"`
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
	Meta RecordUpdateResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                        `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseSSHFPJSON `json:"-"`
}

// recordUpdateResponseSSHFPJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSSHFP]
type recordUpdateResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSSHFP) implementsDNSRecordUpdateResponse() {}

// Components of a SSHFP record.
type RecordUpdateResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                           `json:"type"`
	JSON recordUpdateResponseSSHFPDataJSON `json:"-"`
}

// recordUpdateResponseSSHFPDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSSHFPData]
type recordUpdateResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseSSHFPType string

const (
	RecordUpdateResponseSSHFPTypeSSHFP RecordUpdateResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                            `json:"source"`
	JSON   recordUpdateResponseSSHFPMetaJSON `json:"-"`
}

// recordUpdateResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSSHFPMeta]
type recordUpdateResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseSSHFPTTLNumber].
type RecordUpdateResponseSSHFPTTL interface {
	ImplementsDNSRecordUpdateResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseSSHFPTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseSSHFPTTLNumber float64

const (
	RecordUpdateResponseSSHFPTTLNumber1 RecordUpdateResponseSSHFPTTLNumber = 1
)

type RecordUpdateResponseSVCB struct {
	// Components of a SVCB record.
	Data RecordUpdateResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseSVCBType `json:"type,required"`
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
	Meta RecordUpdateResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseSVCBJSON `json:"-"`
}

// recordUpdateResponseSVCBJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSVCB]
type recordUpdateResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSVCB) implementsDNSRecordUpdateResponse() {}

// Components of a SVCB record.
type RecordUpdateResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                           `json:"value"`
	JSON  recordUpdateResponseSVCBDataJSON `json:"-"`
}

// recordUpdateResponseSVCBDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSVCBData]
type recordUpdateResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseSVCBType string

const (
	RecordUpdateResponseSVCBTypeSVCB RecordUpdateResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordUpdateResponseSVCBMetaJSON `json:"-"`
}

// recordUpdateResponseSVCBMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSVCBMeta]
type recordUpdateResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseSVCBTTLNumber].
type RecordUpdateResponseSVCBTTL interface {
	ImplementsDNSRecordUpdateResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseSVCBTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseSVCBTTLNumber float64

const (
	RecordUpdateResponseSVCBTTLNumber1 RecordUpdateResponseSVCBTTLNumber = 1
)

type RecordUpdateResponseTLSA struct {
	// Components of a TLSA record.
	Data RecordUpdateResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseTLSAType `json:"type,required"`
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
	Meta RecordUpdateResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseTLSAJSON `json:"-"`
}

// recordUpdateResponseTLSAJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTLSA]
type recordUpdateResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseTLSA) implementsDNSRecordUpdateResponse() {}

// Components of a TLSA record.
type RecordUpdateResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                          `json:"usage"`
	JSON  recordUpdateResponseTLSADataJSON `json:"-"`
}

// recordUpdateResponseTLSADataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTLSAData]
type recordUpdateResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordUpdateResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseTLSAType string

const (
	RecordUpdateResponseTLSATypeTLSA RecordUpdateResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordUpdateResponseTLSAMetaJSON `json:"-"`
}

// recordUpdateResponseTLSAMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTLSAMeta]
type recordUpdateResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseTLSATTLNumber].
type RecordUpdateResponseTLSATTL interface {
	ImplementsDNSRecordUpdateResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseTLSATTLNumber(0)),
		},
	)
}

type RecordUpdateResponseTLSATTLNumber float64

const (
	RecordUpdateResponseTLSATTLNumber1 RecordUpdateResponseTLSATTLNumber = 1
)

type RecordUpdateResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordUpdateResponseTXTType `json:"type,required"`
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
	Meta RecordUpdateResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseTXTJSON `json:"-"`
}

// recordUpdateResponseTXTJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTXT]
type recordUpdateResponseTXTJSON struct {
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

func (r *RecordUpdateResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseTXT) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseTXTType string

const (
	RecordUpdateResponseTXTTypeTXT RecordUpdateResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordUpdateResponseTXTMetaJSON `json:"-"`
}

// recordUpdateResponseTXTMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTXTMeta]
type recordUpdateResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseTXTTTLNumber].
type RecordUpdateResponseTXTTTL interface {
	ImplementsDNSRecordUpdateResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseTXTTTLNumber(0)),
		},
	)
}

type RecordUpdateResponseTXTTTLNumber float64

const (
	RecordUpdateResponseTXTTTLNumber1 RecordUpdateResponseTXTTTLNumber = 1
)

type RecordUpdateResponseURI struct {
	// Components of a URI record.
	Data RecordUpdateResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordUpdateResponseURIType `json:"type,required"`
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
	Meta RecordUpdateResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordUpdateResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordUpdateResponseURIJSON `json:"-"`
}

// recordUpdateResponseURIJSON contains the JSON metadata for the struct
// [RecordUpdateResponseURI]
type recordUpdateResponseURIJSON struct {
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

func (r *RecordUpdateResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseURI) implementsDNSRecordUpdateResponse() {}

// Components of a URI record.
type RecordUpdateResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                         `json:"weight"`
	JSON   recordUpdateResponseURIDataJSON `json:"-"`
}

// recordUpdateResponseURIDataJSON contains the JSON metadata for the struct
// [RecordUpdateResponseURIData]
type recordUpdateResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordUpdateResponseURIType string

const (
	RecordUpdateResponseURITypeURI RecordUpdateResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type RecordUpdateResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordUpdateResponseURIMetaJSON `json:"-"`
}

// recordUpdateResponseURIMetaJSON contains the JSON metadata for the struct
// [RecordUpdateResponseURIMeta]
type recordUpdateResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordUpdateResponseURITTLNumber].
type RecordUpdateResponseURITTL interface {
	ImplementsDNSRecordUpdateResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordUpdateResponseURITTLNumber(0)),
		},
	)
}

type RecordUpdateResponseURITTLNumber float64

const (
	RecordUpdateResponseURITTLNumber1 RecordUpdateResponseURITTLNumber = 1
)

// Union satisfied by [dns.RecordListResponseA], [dns.RecordListResponseAAAA],
// [dns.RecordListResponseCAA], [dns.RecordListResponseCert],
// [dns.RecordListResponseCNAME], [dns.RecordListResponseDNSKEY],
// [dns.RecordListResponseDS], [dns.RecordListResponseHTTPS],
// [dns.RecordListResponseLOC], [dns.RecordListResponseMX],
// [dns.RecordListResponseNAPTR], [dns.RecordListResponseNS],
// [dns.RecordListResponsePTR], [dns.RecordListResponseSmimea],
// [dns.RecordListResponseSRV], [dns.RecordListResponseSSHFP],
// [dns.RecordListResponseSVCB], [dns.RecordListResponseTLSA],
// [dns.RecordListResponseTXT] or [dns.RecordListResponseURI].
type RecordListResponse interface {
	implementsDNSRecordListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordListResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseAType `json:"type,required"`
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
	Meta RecordListResponseAMeta `json:"meta"`
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
	TTL RecordListResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordListResponseAJSON `json:"-"`
}

// recordListResponseAJSON contains the JSON metadata for the struct
// [RecordListResponseA]
type recordListResponseAJSON struct {
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

func (r *RecordListResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseAJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseA) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseAType string

const (
	RecordListResponseATypeA RecordListResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordListResponseAMetaJSON `json:"-"`
}

// recordListResponseAMetaJSON contains the JSON metadata for the struct
// [RecordListResponseAMeta]
type recordListResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseATTLNumber].
type RecordListResponseATTL interface {
	ImplementsDNSRecordListResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseATTLNumber(0)),
		},
	)
}

type RecordListResponseATTLNumber float64

const (
	RecordListResponseATTLNumber1 RecordListResponseATTLNumber = 1
)

type RecordListResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseAAAAType `json:"type,required"`
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
	Meta RecordListResponseAAAAMeta `json:"meta"`
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
	TTL RecordListResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordListResponseAAAAJSON `json:"-"`
}

// recordListResponseAAAAJSON contains the JSON metadata for the struct
// [RecordListResponseAAAA]
type recordListResponseAAAAJSON struct {
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

func (r *RecordListResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseAAAA) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseAAAAType string

const (
	RecordListResponseAAAATypeAAAA RecordListResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordListResponseAAAAMetaJSON `json:"-"`
}

// recordListResponseAAAAMetaJSON contains the JSON metadata for the struct
// [RecordListResponseAAAAMeta]
type recordListResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseAAAATTLNumber].
type RecordListResponseAAAATTL interface {
	ImplementsDNSRecordListResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseAAAATTLNumber(0)),
		},
	)
}

type RecordListResponseAAAATTLNumber float64

const (
	RecordListResponseAAAATTLNumber1 RecordListResponseAAAATTLNumber = 1
)

type RecordListResponseCAA struct {
	// Components of a CAA record.
	Data RecordListResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseCAAType `json:"type,required"`
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
	Meta RecordListResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordListResponseCAAJSON `json:"-"`
}

// recordListResponseCAAJSON contains the JSON metadata for the struct
// [RecordListResponseCAA]
type recordListResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseCAA) implementsDNSRecordListResponse() {}

// Components of a CAA record.
type RecordListResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                        `json:"value"`
	JSON  recordListResponseCAADataJSON `json:"-"`
}

// recordListResponseCAADataJSON contains the JSON metadata for the struct
// [RecordListResponseCAAData]
type recordListResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseCAAType string

const (
	RecordListResponseCAATypeCAA RecordListResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordListResponseCAAMetaJSON `json:"-"`
}

// recordListResponseCAAMetaJSON contains the JSON metadata for the struct
// [RecordListResponseCAAMeta]
type recordListResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseCAATTLNumber].
type RecordListResponseCAATTL interface {
	ImplementsDNSRecordListResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseCAATTLNumber(0)),
		},
	)
}

type RecordListResponseCAATTLNumber float64

const (
	RecordListResponseCAATTLNumber1 RecordListResponseCAATTLNumber = 1
)

type RecordListResponseCert struct {
	// Components of a CERT record.
	Data RecordListResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseCertType `json:"type,required"`
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
	Meta RecordListResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordListResponseCertJSON `json:"-"`
}

// recordListResponseCertJSON contains the JSON metadata for the struct
// [RecordListResponseCert]
type recordListResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseCert) implementsDNSRecordListResponse() {}

// Components of a CERT record.
type RecordListResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                        `json:"type"`
	JSON recordListResponseCertDataJSON `json:"-"`
}

// recordListResponseCertDataJSON contains the JSON metadata for the struct
// [RecordListResponseCertData]
type recordListResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseCertType string

const (
	RecordListResponseCertTypeCert RecordListResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordListResponseCertMetaJSON `json:"-"`
}

// recordListResponseCertMetaJSON contains the JSON metadata for the struct
// [RecordListResponseCertMeta]
type recordListResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseCertTTLNumber].
type RecordListResponseCertTTL interface {
	ImplementsDNSRecordListResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseCertTTLNumber(0)),
		},
	)
}

type RecordListResponseCertTTLNumber float64

const (
	RecordListResponseCertTTLNumber1 RecordListResponseCertTTLNumber = 1
)

type RecordListResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseCNAMEType `json:"type,required"`
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
	Meta RecordListResponseCNAMEMeta `json:"meta"`
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
	TTL RecordListResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordListResponseCNAMEJSON `json:"-"`
}

// recordListResponseCNAMEJSON contains the JSON metadata for the struct
// [RecordListResponseCNAME]
type recordListResponseCNAMEJSON struct {
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

func (r *RecordListResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseCNAME) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseCNAMEType string

const (
	RecordListResponseCNAMETypeCNAME RecordListResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordListResponseCNAMEMetaJSON `json:"-"`
}

// recordListResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [RecordListResponseCNAMEMeta]
type recordListResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordListResponseCNAMETTLNumber].
type RecordListResponseCNAMETTL interface {
	ImplementsDNSRecordListResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseCNAMETTLNumber(0)),
		},
	)
}

type RecordListResponseCNAMETTLNumber float64

const (
	RecordListResponseCNAMETTLNumber1 RecordListResponseCNAMETTLNumber = 1
)

type RecordListResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data RecordListResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseDNSKEYType `json:"type,required"`
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
	Meta RecordListResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordListResponseDNSKEYJSON `json:"-"`
}

// recordListResponseDNSKEYJSON contains the JSON metadata for the struct
// [RecordListResponseDNSKEY]
type recordListResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseDNSKEY) implementsDNSRecordListResponse() {}

// Components of a DNSKEY record.
type RecordListResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                           `json:"public_key"`
	JSON      recordListResponseDNSKEYDataJSON `json:"-"`
}

// recordListResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [RecordListResponseDNSKEYData]
type recordListResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseDNSKEYType string

const (
	RecordListResponseDNSKEYTypeDNSKEY RecordListResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordListResponseDNSKEYMetaJSON `json:"-"`
}

// recordListResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [RecordListResponseDNSKEYMeta]
type recordListResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordListResponseDNSKEYTTLNumber].
type RecordListResponseDNSKEYTTL interface {
	ImplementsDNSRecordListResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseDNSKEYTTLNumber(0)),
		},
	)
}

type RecordListResponseDNSKEYTTLNumber float64

const (
	RecordListResponseDNSKEYTTLNumber1 RecordListResponseDNSKEYTTLNumber = 1
)

type RecordListResponseDS struct {
	// Components of a DS record.
	Data RecordListResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseDSType `json:"type,required"`
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
	Meta RecordListResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordListResponseDSJSON `json:"-"`
}

// recordListResponseDSJSON contains the JSON metadata for the struct
// [RecordListResponseDS]
type recordListResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseDS) implementsDNSRecordListResponse() {}

// Components of a DS record.
type RecordListResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                      `json:"key_tag"`
	JSON   recordListResponseDSDataJSON `json:"-"`
}

// recordListResponseDSDataJSON contains the JSON metadata for the struct
// [RecordListResponseDSData]
type recordListResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseDSType string

const (
	RecordListResponseDSTypeDS RecordListResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordListResponseDSMetaJSON `json:"-"`
}

// recordListResponseDSMetaJSON contains the JSON metadata for the struct
// [RecordListResponseDSMeta]
type recordListResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseDSTTLNumber].
type RecordListResponseDSTTL interface {
	ImplementsDNSRecordListResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseDSTTLNumber(0)),
		},
	)
}

type RecordListResponseDSTTLNumber float64

const (
	RecordListResponseDSTTLNumber1 RecordListResponseDSTTLNumber = 1
)

type RecordListResponseHTTPS struct {
	// Components of a HTTPS record.
	Data RecordListResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseHTTPSType `json:"type,required"`
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
	Meta RecordListResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordListResponseHTTPSJSON `json:"-"`
}

// recordListResponseHTTPSJSON contains the JSON metadata for the struct
// [RecordListResponseHTTPS]
type recordListResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseHTTPS) implementsDNSRecordListResponse() {}

// Components of a HTTPS record.
type RecordListResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                          `json:"value"`
	JSON  recordListResponseHTTPSDataJSON `json:"-"`
}

// recordListResponseHTTPSDataJSON contains the JSON metadata for the struct
// [RecordListResponseHTTPSData]
type recordListResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseHTTPSType string

const (
	RecordListResponseHTTPSTypeHTTPS RecordListResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordListResponseHTTPSMetaJSON `json:"-"`
}

// recordListResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [RecordListResponseHTTPSMeta]
type recordListResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordListResponseHTTPSTTLNumber].
type RecordListResponseHTTPSTTL interface {
	ImplementsDNSRecordListResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseHTTPSTTLNumber(0)),
		},
	)
}

type RecordListResponseHTTPSTTLNumber float64

const (
	RecordListResponseHTTPSTTLNumber1 RecordListResponseHTTPSTTLNumber = 1
)

type RecordListResponseLOC struct {
	// Components of a LOC record.
	Data RecordListResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseLOCType `json:"type,required"`
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
	Meta RecordListResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordListResponseLOCJSON `json:"-"`
}

// recordListResponseLOCJSON contains the JSON metadata for the struct
// [RecordListResponseLOC]
type recordListResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseLOC) implementsDNSRecordListResponse() {}

// Components of a LOC record.
type RecordListResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection RecordListResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection RecordListResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                       `json:"size"`
	JSON recordListResponseLOCDataJSON `json:"-"`
}

// recordListResponseLOCDataJSON contains the JSON metadata for the struct
// [RecordListResponseLOCData]
type recordListResponseLOCDataJSON struct {
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

func (r *RecordListResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type RecordListResponseLOCDataLatDirection string

const (
	RecordListResponseLOCDataLatDirectionN RecordListResponseLOCDataLatDirection = "N"
	RecordListResponseLOCDataLatDirectionS RecordListResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type RecordListResponseLOCDataLongDirection string

const (
	RecordListResponseLOCDataLongDirectionE RecordListResponseLOCDataLongDirection = "E"
	RecordListResponseLOCDataLongDirectionW RecordListResponseLOCDataLongDirection = "W"
)

// Record type.
type RecordListResponseLOCType string

const (
	RecordListResponseLOCTypeLOC RecordListResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordListResponseLOCMetaJSON `json:"-"`
}

// recordListResponseLOCMetaJSON contains the JSON metadata for the struct
// [RecordListResponseLOCMeta]
type recordListResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseLOCTTLNumber].
type RecordListResponseLOCTTL interface {
	ImplementsDNSRecordListResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseLOCTTLNumber(0)),
		},
	)
}

type RecordListResponseLOCTTLNumber float64

const (
	RecordListResponseLOCTTLNumber1 RecordListResponseLOCTTLNumber = 1
)

type RecordListResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordListResponseMXType `json:"type,required"`
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
	Meta RecordListResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordListResponseMXJSON `json:"-"`
}

// recordListResponseMXJSON contains the JSON metadata for the struct
// [RecordListResponseMX]
type recordListResponseMXJSON struct {
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

func (r *RecordListResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseMX) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseMXType string

const (
	RecordListResponseMXTypeMX RecordListResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordListResponseMXMetaJSON `json:"-"`
}

// recordListResponseMXMetaJSON contains the JSON metadata for the struct
// [RecordListResponseMXMeta]
type recordListResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseMXTTLNumber].
type RecordListResponseMXTTL interface {
	ImplementsDNSRecordListResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseMXTTLNumber(0)),
		},
	)
}

type RecordListResponseMXTTLNumber float64

const (
	RecordListResponseMXTTLNumber1 RecordListResponseMXTTLNumber = 1
)

type RecordListResponseNAPTR struct {
	// Components of a NAPTR record.
	Data RecordListResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseNAPTRType `json:"type,required"`
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
	Meta RecordListResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordListResponseNAPTRJSON `json:"-"`
}

// recordListResponseNAPTRJSON contains the JSON metadata for the struct
// [RecordListResponseNAPTR]
type recordListResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseNAPTR) implementsDNSRecordListResponse() {}

// Components of a NAPTR record.
type RecordListResponseNAPTRData struct {
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
	Service string                          `json:"service"`
	JSON    recordListResponseNAPTRDataJSON `json:"-"`
}

// recordListResponseNAPTRDataJSON contains the JSON metadata for the struct
// [RecordListResponseNAPTRData]
type recordListResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseNAPTRType string

const (
	RecordListResponseNAPTRTypeNAPTR RecordListResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordListResponseNAPTRMetaJSON `json:"-"`
}

// recordListResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [RecordListResponseNAPTRMeta]
type recordListResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordListResponseNAPTRTTLNumber].
type RecordListResponseNAPTRTTL interface {
	ImplementsDNSRecordListResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseNAPTRTTLNumber(0)),
		},
	)
}

type RecordListResponseNAPTRTTLNumber float64

const (
	RecordListResponseNAPTRTTLNumber1 RecordListResponseNAPTRTTLNumber = 1
)

type RecordListResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseNSType `json:"type,required"`
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
	Meta RecordListResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordListResponseNSJSON `json:"-"`
}

// recordListResponseNSJSON contains the JSON metadata for the struct
// [RecordListResponseNS]
type recordListResponseNSJSON struct {
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

func (r *RecordListResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseNS) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseNSType string

const (
	RecordListResponseNSTypeNS RecordListResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordListResponseNSMetaJSON `json:"-"`
}

// recordListResponseNSMetaJSON contains the JSON metadata for the struct
// [RecordListResponseNSMeta]
type recordListResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseNSTTLNumber].
type RecordListResponseNSTTL interface {
	ImplementsDNSRecordListResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseNSTTLNumber(0)),
		},
	)
}

type RecordListResponseNSTTLNumber float64

const (
	RecordListResponseNSTTLNumber1 RecordListResponseNSTTLNumber = 1
)

type RecordListResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponsePTRType `json:"type,required"`
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
	Meta RecordListResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordListResponsePTRJSON `json:"-"`
}

// recordListResponsePTRJSON contains the JSON metadata for the struct
// [RecordListResponsePTR]
type recordListResponsePTRJSON struct {
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

func (r *RecordListResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponsePTR) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponsePTRType string

const (
	RecordListResponsePTRTypePTR RecordListResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordListResponsePTRMetaJSON `json:"-"`
}

// recordListResponsePTRMetaJSON contains the JSON metadata for the struct
// [RecordListResponsePTRMeta]
type recordListResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponsePTRTTLNumber].
type RecordListResponsePTRTTL interface {
	ImplementsDNSRecordListResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponsePTRTTLNumber(0)),
		},
	)
}

type RecordListResponsePTRTTLNumber float64

const (
	RecordListResponsePTRTTLNumber1 RecordListResponsePTRTTLNumber = 1
)

type RecordListResponseSmimea struct {
	// Components of a SMIMEA record.
	Data RecordListResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseSmimeaType `json:"type,required"`
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
	Meta RecordListResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordListResponseSmimeaJSON `json:"-"`
}

// recordListResponseSmimeaJSON contains the JSON metadata for the struct
// [RecordListResponseSmimea]
type recordListResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSmimea) implementsDNSRecordListResponse() {}

// Components of a SMIMEA record.
type RecordListResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                          `json:"usage"`
	JSON  recordListResponseSmimeaDataJSON `json:"-"`
}

// recordListResponseSmimeaDataJSON contains the JSON metadata for the struct
// [RecordListResponseSmimeaData]
type recordListResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordListResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseSmimeaType string

const (
	RecordListResponseSmimeaTypeSmimea RecordListResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordListResponseSmimeaMetaJSON `json:"-"`
}

// recordListResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [RecordListResponseSmimeaMeta]
type recordListResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordListResponseSmimeaTTLNumber].
type RecordListResponseSmimeaTTL interface {
	ImplementsDNSRecordListResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseSmimeaTTLNumber(0)),
		},
	)
}

type RecordListResponseSmimeaTTLNumber float64

const (
	RecordListResponseSmimeaTTLNumber1 RecordListResponseSmimeaTTLNumber = 1
)

type RecordListResponseSRV struct {
	// Components of a SRV record.
	Data RecordListResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseSRVType `json:"type,required"`
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
	Meta RecordListResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordListResponseSRVJSON `json:"-"`
}

// recordListResponseSRVJSON contains the JSON metadata for the struct
// [RecordListResponseSRV]
type recordListResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSRV) implementsDNSRecordListResponse() {}

// Components of a SRV record.
type RecordListResponseSRVData struct {
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
	Weight float64                       `json:"weight"`
	JSON   recordListResponseSRVDataJSON `json:"-"`
}

// recordListResponseSRVDataJSON contains the JSON metadata for the struct
// [RecordListResponseSRVData]
type recordListResponseSRVDataJSON struct {
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

func (r *RecordListResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseSRVType string

const (
	RecordListResponseSRVTypeSRV RecordListResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordListResponseSRVMetaJSON `json:"-"`
}

// recordListResponseSRVMetaJSON contains the JSON metadata for the struct
// [RecordListResponseSRVMeta]
type recordListResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseSRVTTLNumber].
type RecordListResponseSRVTTL interface {
	ImplementsDNSRecordListResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseSRVTTLNumber(0)),
		},
	)
}

type RecordListResponseSRVTTLNumber float64

const (
	RecordListResponseSRVTTLNumber1 RecordListResponseSRVTTLNumber = 1
)

type RecordListResponseSSHFP struct {
	// Components of a SSHFP record.
	Data RecordListResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseSSHFPType `json:"type,required"`
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
	Meta RecordListResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordListResponseSSHFPJSON `json:"-"`
}

// recordListResponseSSHFPJSON contains the JSON metadata for the struct
// [RecordListResponseSSHFP]
type recordListResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSSHFP) implementsDNSRecordListResponse() {}

// Components of a SSHFP record.
type RecordListResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                         `json:"type"`
	JSON recordListResponseSSHFPDataJSON `json:"-"`
}

// recordListResponseSSHFPDataJSON contains the JSON metadata for the struct
// [RecordListResponseSSHFPData]
type recordListResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseSSHFPType string

const (
	RecordListResponseSSHFPTypeSSHFP RecordListResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordListResponseSSHFPMetaJSON `json:"-"`
}

// recordListResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [RecordListResponseSSHFPMeta]
type recordListResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordListResponseSSHFPTTLNumber].
type RecordListResponseSSHFPTTL interface {
	ImplementsDNSRecordListResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseSSHFPTTLNumber(0)),
		},
	)
}

type RecordListResponseSSHFPTTLNumber float64

const (
	RecordListResponseSSHFPTTLNumber1 RecordListResponseSSHFPTTLNumber = 1
)

type RecordListResponseSVCB struct {
	// Components of a SVCB record.
	Data RecordListResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseSVCBType `json:"type,required"`
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
	Meta RecordListResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordListResponseSVCBJSON `json:"-"`
}

// recordListResponseSVCBJSON contains the JSON metadata for the struct
// [RecordListResponseSVCB]
type recordListResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSVCB) implementsDNSRecordListResponse() {}

// Components of a SVCB record.
type RecordListResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                         `json:"value"`
	JSON  recordListResponseSVCBDataJSON `json:"-"`
}

// recordListResponseSVCBDataJSON contains the JSON metadata for the struct
// [RecordListResponseSVCBData]
type recordListResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseSVCBType string

const (
	RecordListResponseSVCBTypeSVCB RecordListResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordListResponseSVCBMetaJSON `json:"-"`
}

// recordListResponseSVCBMetaJSON contains the JSON metadata for the struct
// [RecordListResponseSVCBMeta]
type recordListResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseSVCBTTLNumber].
type RecordListResponseSVCBTTL interface {
	ImplementsDNSRecordListResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseSVCBTTLNumber(0)),
		},
	)
}

type RecordListResponseSVCBTTLNumber float64

const (
	RecordListResponseSVCBTTLNumber1 RecordListResponseSVCBTTLNumber = 1
)

type RecordListResponseTLSA struct {
	// Components of a TLSA record.
	Data RecordListResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseTLSAType `json:"type,required"`
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
	Meta RecordListResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordListResponseTLSAJSON `json:"-"`
}

// recordListResponseTLSAJSON contains the JSON metadata for the struct
// [RecordListResponseTLSA]
type recordListResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseTLSA) implementsDNSRecordListResponse() {}

// Components of a TLSA record.
type RecordListResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                        `json:"usage"`
	JSON  recordListResponseTLSADataJSON `json:"-"`
}

// recordListResponseTLSADataJSON contains the JSON metadata for the struct
// [RecordListResponseTLSAData]
type recordListResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordListResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseTLSAType string

const (
	RecordListResponseTLSATypeTLSA RecordListResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordListResponseTLSAMetaJSON `json:"-"`
}

// recordListResponseTLSAMetaJSON contains the JSON metadata for the struct
// [RecordListResponseTLSAMeta]
type recordListResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseTLSATTLNumber].
type RecordListResponseTLSATTL interface {
	ImplementsDNSRecordListResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseTLSATTLNumber(0)),
		},
	)
}

type RecordListResponseTLSATTLNumber float64

const (
	RecordListResponseTLSATTLNumber1 RecordListResponseTLSATTLNumber = 1
)

type RecordListResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordListResponseTXTType `json:"type,required"`
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
	Meta RecordListResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordListResponseTXTJSON `json:"-"`
}

// recordListResponseTXTJSON contains the JSON metadata for the struct
// [RecordListResponseTXT]
type recordListResponseTXTJSON struct {
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

func (r *RecordListResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseTXT) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseTXTType string

const (
	RecordListResponseTXTTypeTXT RecordListResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordListResponseTXTMetaJSON `json:"-"`
}

// recordListResponseTXTMetaJSON contains the JSON metadata for the struct
// [RecordListResponseTXTMeta]
type recordListResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseTXTTTLNumber].
type RecordListResponseTXTTTL interface {
	ImplementsDNSRecordListResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseTXTTTLNumber(0)),
		},
	)
}

type RecordListResponseTXTTTLNumber float64

const (
	RecordListResponseTXTTTLNumber1 RecordListResponseTXTTTLNumber = 1
)

type RecordListResponseURI struct {
	// Components of a URI record.
	Data RecordListResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordListResponseURIType `json:"type,required"`
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
	Meta RecordListResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordListResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordListResponseURIJSON `json:"-"`
}

// recordListResponseURIJSON contains the JSON metadata for the struct
// [RecordListResponseURI]
type recordListResponseURIJSON struct {
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

func (r *RecordListResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseURI) implementsDNSRecordListResponse() {}

// Components of a URI record.
type RecordListResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                       `json:"weight"`
	JSON   recordListResponseURIDataJSON `json:"-"`
}

// recordListResponseURIDataJSON contains the JSON metadata for the struct
// [RecordListResponseURIData]
type recordListResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordListResponseURIType string

const (
	RecordListResponseURITypeURI RecordListResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type RecordListResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordListResponseURIMetaJSON `json:"-"`
}

// recordListResponseURIMetaJSON contains the JSON metadata for the struct
// [RecordListResponseURIMeta]
type recordListResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordListResponseURITTLNumber].
type RecordListResponseURITTL interface {
	ImplementsDNSRecordListResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordListResponseURITTLNumber(0)),
		},
	)
}

type RecordListResponseURITTLNumber float64

const (
	RecordListResponseURITTLNumber1 RecordListResponseURITTLNumber = 1
)

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

// Union satisfied by [dns.RecordEditResponseA], [dns.RecordEditResponseAAAA],
// [dns.RecordEditResponseCAA], [dns.RecordEditResponseCert],
// [dns.RecordEditResponseCNAME], [dns.RecordEditResponseDNSKEY],
// [dns.RecordEditResponseDS], [dns.RecordEditResponseHTTPS],
// [dns.RecordEditResponseLOC], [dns.RecordEditResponseMX],
// [dns.RecordEditResponseNAPTR], [dns.RecordEditResponseNS],
// [dns.RecordEditResponsePTR], [dns.RecordEditResponseSmimea],
// [dns.RecordEditResponseSRV], [dns.RecordEditResponseSSHFP],
// [dns.RecordEditResponseSVCB], [dns.RecordEditResponseTLSA],
// [dns.RecordEditResponseTXT] or [dns.RecordEditResponseURI].
type RecordEditResponse interface {
	implementsDNSRecordEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordEditResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseAType `json:"type,required"`
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
	Meta RecordEditResponseAMeta `json:"meta"`
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
	TTL RecordEditResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseAJSON `json:"-"`
}

// recordEditResponseAJSON contains the JSON metadata for the struct
// [RecordEditResponseA]
type recordEditResponseAJSON struct {
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

func (r *RecordEditResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseAJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseA) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseAType string

const (
	RecordEditResponseATypeA RecordEditResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordEditResponseAMetaJSON `json:"-"`
}

// recordEditResponseAMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseAMeta]
type recordEditResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseATTLNumber].
type RecordEditResponseATTL interface {
	ImplementsDNSRecordEditResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseATTLNumber(0)),
		},
	)
}

type RecordEditResponseATTLNumber float64

const (
	RecordEditResponseATTLNumber1 RecordEditResponseATTLNumber = 1
)

type RecordEditResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseAAAAType `json:"type,required"`
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
	Meta RecordEditResponseAAAAMeta `json:"meta"`
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
	TTL RecordEditResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseAAAAJSON `json:"-"`
}

// recordEditResponseAAAAJSON contains the JSON metadata for the struct
// [RecordEditResponseAAAA]
type recordEditResponseAAAAJSON struct {
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

func (r *RecordEditResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseAAAA) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseAAAAType string

const (
	RecordEditResponseAAAATypeAAAA RecordEditResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordEditResponseAAAAMetaJSON `json:"-"`
}

// recordEditResponseAAAAMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseAAAAMeta]
type recordEditResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseAAAATTLNumber].
type RecordEditResponseAAAATTL interface {
	ImplementsDNSRecordEditResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseAAAATTLNumber(0)),
		},
	)
}

type RecordEditResponseAAAATTLNumber float64

const (
	RecordEditResponseAAAATTLNumber1 RecordEditResponseAAAATTLNumber = 1
)

type RecordEditResponseCAA struct {
	// Components of a CAA record.
	Data RecordEditResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseCAAType `json:"type,required"`
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
	Meta RecordEditResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseCAAJSON `json:"-"`
}

// recordEditResponseCAAJSON contains the JSON metadata for the struct
// [RecordEditResponseCAA]
type recordEditResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseCAA) implementsDNSRecordEditResponse() {}

// Components of a CAA record.
type RecordEditResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                        `json:"value"`
	JSON  recordEditResponseCAADataJSON `json:"-"`
}

// recordEditResponseCAADataJSON contains the JSON metadata for the struct
// [RecordEditResponseCAAData]
type recordEditResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseCAAType string

const (
	RecordEditResponseCAATypeCAA RecordEditResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordEditResponseCAAMetaJSON `json:"-"`
}

// recordEditResponseCAAMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseCAAMeta]
type recordEditResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseCAATTLNumber].
type RecordEditResponseCAATTL interface {
	ImplementsDNSRecordEditResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseCAATTLNumber(0)),
		},
	)
}

type RecordEditResponseCAATTLNumber float64

const (
	RecordEditResponseCAATTLNumber1 RecordEditResponseCAATTLNumber = 1
)

type RecordEditResponseCert struct {
	// Components of a CERT record.
	Data RecordEditResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseCertType `json:"type,required"`
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
	Meta RecordEditResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseCertJSON `json:"-"`
}

// recordEditResponseCertJSON contains the JSON metadata for the struct
// [RecordEditResponseCert]
type recordEditResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseCert) implementsDNSRecordEditResponse() {}

// Components of a CERT record.
type RecordEditResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                        `json:"type"`
	JSON recordEditResponseCertDataJSON `json:"-"`
}

// recordEditResponseCertDataJSON contains the JSON metadata for the struct
// [RecordEditResponseCertData]
type recordEditResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseCertType string

const (
	RecordEditResponseCertTypeCert RecordEditResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordEditResponseCertMetaJSON `json:"-"`
}

// recordEditResponseCertMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseCertMeta]
type recordEditResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseCertTTLNumber].
type RecordEditResponseCertTTL interface {
	ImplementsDNSRecordEditResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseCertTTLNumber(0)),
		},
	)
}

type RecordEditResponseCertTTLNumber float64

const (
	RecordEditResponseCertTTLNumber1 RecordEditResponseCertTTLNumber = 1
)

type RecordEditResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseCNAMEType `json:"type,required"`
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
	Meta RecordEditResponseCNAMEMeta `json:"meta"`
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
	TTL RecordEditResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseCNAMEJSON `json:"-"`
}

// recordEditResponseCNAMEJSON contains the JSON metadata for the struct
// [RecordEditResponseCNAME]
type recordEditResponseCNAMEJSON struct {
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

func (r *RecordEditResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseCNAME) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseCNAMEType string

const (
	RecordEditResponseCNAMETypeCNAME RecordEditResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordEditResponseCNAMEMetaJSON `json:"-"`
}

// recordEditResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseCNAMEMeta]
type recordEditResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordEditResponseCNAMETTLNumber].
type RecordEditResponseCNAMETTL interface {
	ImplementsDNSRecordEditResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseCNAMETTLNumber(0)),
		},
	)
}

type RecordEditResponseCNAMETTLNumber float64

const (
	RecordEditResponseCNAMETTLNumber1 RecordEditResponseCNAMETTLNumber = 1
)

type RecordEditResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data RecordEditResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseDNSKEYType `json:"type,required"`
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
	Meta RecordEditResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseDNSKEYJSON `json:"-"`
}

// recordEditResponseDNSKEYJSON contains the JSON metadata for the struct
// [RecordEditResponseDNSKEY]
type recordEditResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseDNSKEY) implementsDNSRecordEditResponse() {}

// Components of a DNSKEY record.
type RecordEditResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                           `json:"public_key"`
	JSON      recordEditResponseDNSKEYDataJSON `json:"-"`
}

// recordEditResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [RecordEditResponseDNSKEYData]
type recordEditResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseDNSKEYType string

const (
	RecordEditResponseDNSKEYTypeDNSKEY RecordEditResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordEditResponseDNSKEYMetaJSON `json:"-"`
}

// recordEditResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseDNSKEYMeta]
type recordEditResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordEditResponseDNSKEYTTLNumber].
type RecordEditResponseDNSKEYTTL interface {
	ImplementsDNSRecordEditResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseDNSKEYTTLNumber(0)),
		},
	)
}

type RecordEditResponseDNSKEYTTLNumber float64

const (
	RecordEditResponseDNSKEYTTLNumber1 RecordEditResponseDNSKEYTTLNumber = 1
)

type RecordEditResponseDS struct {
	// Components of a DS record.
	Data RecordEditResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseDSType `json:"type,required"`
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
	Meta RecordEditResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseDSJSON `json:"-"`
}

// recordEditResponseDSJSON contains the JSON metadata for the struct
// [RecordEditResponseDS]
type recordEditResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseDS) implementsDNSRecordEditResponse() {}

// Components of a DS record.
type RecordEditResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                      `json:"key_tag"`
	JSON   recordEditResponseDSDataJSON `json:"-"`
}

// recordEditResponseDSDataJSON contains the JSON metadata for the struct
// [RecordEditResponseDSData]
type recordEditResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseDSType string

const (
	RecordEditResponseDSTypeDS RecordEditResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordEditResponseDSMetaJSON `json:"-"`
}

// recordEditResponseDSMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseDSMeta]
type recordEditResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseDSTTLNumber].
type RecordEditResponseDSTTL interface {
	ImplementsDNSRecordEditResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseDSTTLNumber(0)),
		},
	)
}

type RecordEditResponseDSTTLNumber float64

const (
	RecordEditResponseDSTTLNumber1 RecordEditResponseDSTTLNumber = 1
)

type RecordEditResponseHTTPS struct {
	// Components of a HTTPS record.
	Data RecordEditResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseHTTPSType `json:"type,required"`
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
	Meta RecordEditResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseHTTPSJSON `json:"-"`
}

// recordEditResponseHTTPSJSON contains the JSON metadata for the struct
// [RecordEditResponseHTTPS]
type recordEditResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseHTTPS) implementsDNSRecordEditResponse() {}

// Components of a HTTPS record.
type RecordEditResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                          `json:"value"`
	JSON  recordEditResponseHTTPSDataJSON `json:"-"`
}

// recordEditResponseHTTPSDataJSON contains the JSON metadata for the struct
// [RecordEditResponseHTTPSData]
type recordEditResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseHTTPSType string

const (
	RecordEditResponseHTTPSTypeHTTPS RecordEditResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordEditResponseHTTPSMetaJSON `json:"-"`
}

// recordEditResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseHTTPSMeta]
type recordEditResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordEditResponseHTTPSTTLNumber].
type RecordEditResponseHTTPSTTL interface {
	ImplementsDNSRecordEditResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseHTTPSTTLNumber(0)),
		},
	)
}

type RecordEditResponseHTTPSTTLNumber float64

const (
	RecordEditResponseHTTPSTTLNumber1 RecordEditResponseHTTPSTTLNumber = 1
)

type RecordEditResponseLOC struct {
	// Components of a LOC record.
	Data RecordEditResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseLOCType `json:"type,required"`
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
	Meta RecordEditResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseLOCJSON `json:"-"`
}

// recordEditResponseLOCJSON contains the JSON metadata for the struct
// [RecordEditResponseLOC]
type recordEditResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseLOC) implementsDNSRecordEditResponse() {}

// Components of a LOC record.
type RecordEditResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection RecordEditResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection RecordEditResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                       `json:"size"`
	JSON recordEditResponseLOCDataJSON `json:"-"`
}

// recordEditResponseLOCDataJSON contains the JSON metadata for the struct
// [RecordEditResponseLOCData]
type recordEditResponseLOCDataJSON struct {
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

func (r *RecordEditResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type RecordEditResponseLOCDataLatDirection string

const (
	RecordEditResponseLOCDataLatDirectionN RecordEditResponseLOCDataLatDirection = "N"
	RecordEditResponseLOCDataLatDirectionS RecordEditResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type RecordEditResponseLOCDataLongDirection string

const (
	RecordEditResponseLOCDataLongDirectionE RecordEditResponseLOCDataLongDirection = "E"
	RecordEditResponseLOCDataLongDirectionW RecordEditResponseLOCDataLongDirection = "W"
)

// Record type.
type RecordEditResponseLOCType string

const (
	RecordEditResponseLOCTypeLOC RecordEditResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordEditResponseLOCMetaJSON `json:"-"`
}

// recordEditResponseLOCMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseLOCMeta]
type recordEditResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseLOCTTLNumber].
type RecordEditResponseLOCTTL interface {
	ImplementsDNSRecordEditResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseLOCTTLNumber(0)),
		},
	)
}

type RecordEditResponseLOCTTLNumber float64

const (
	RecordEditResponseLOCTTLNumber1 RecordEditResponseLOCTTLNumber = 1
)

type RecordEditResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordEditResponseMXType `json:"type,required"`
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
	Meta RecordEditResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseMXJSON `json:"-"`
}

// recordEditResponseMXJSON contains the JSON metadata for the struct
// [RecordEditResponseMX]
type recordEditResponseMXJSON struct {
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

func (r *RecordEditResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseMX) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseMXType string

const (
	RecordEditResponseMXTypeMX RecordEditResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordEditResponseMXMetaJSON `json:"-"`
}

// recordEditResponseMXMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseMXMeta]
type recordEditResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseMXTTLNumber].
type RecordEditResponseMXTTL interface {
	ImplementsDNSRecordEditResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseMXTTLNumber(0)),
		},
	)
}

type RecordEditResponseMXTTLNumber float64

const (
	RecordEditResponseMXTTLNumber1 RecordEditResponseMXTTLNumber = 1
)

type RecordEditResponseNAPTR struct {
	// Components of a NAPTR record.
	Data RecordEditResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseNAPTRType `json:"type,required"`
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
	Meta RecordEditResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseNAPTRJSON `json:"-"`
}

// recordEditResponseNAPTRJSON contains the JSON metadata for the struct
// [RecordEditResponseNAPTR]
type recordEditResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseNAPTR) implementsDNSRecordEditResponse() {}

// Components of a NAPTR record.
type RecordEditResponseNAPTRData struct {
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
	Service string                          `json:"service"`
	JSON    recordEditResponseNAPTRDataJSON `json:"-"`
}

// recordEditResponseNAPTRDataJSON contains the JSON metadata for the struct
// [RecordEditResponseNAPTRData]
type recordEditResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseNAPTRType string

const (
	RecordEditResponseNAPTRTypeNAPTR RecordEditResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordEditResponseNAPTRMetaJSON `json:"-"`
}

// recordEditResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseNAPTRMeta]
type recordEditResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordEditResponseNAPTRTTLNumber].
type RecordEditResponseNAPTRTTL interface {
	ImplementsDNSRecordEditResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseNAPTRTTLNumber(0)),
		},
	)
}

type RecordEditResponseNAPTRTTLNumber float64

const (
	RecordEditResponseNAPTRTTLNumber1 RecordEditResponseNAPTRTTLNumber = 1
)

type RecordEditResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseNSType `json:"type,required"`
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
	Meta RecordEditResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseNSJSON `json:"-"`
}

// recordEditResponseNSJSON contains the JSON metadata for the struct
// [RecordEditResponseNS]
type recordEditResponseNSJSON struct {
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

func (r *RecordEditResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseNS) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseNSType string

const (
	RecordEditResponseNSTypeNS RecordEditResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordEditResponseNSMetaJSON `json:"-"`
}

// recordEditResponseNSMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseNSMeta]
type recordEditResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseNSTTLNumber].
type RecordEditResponseNSTTL interface {
	ImplementsDNSRecordEditResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseNSTTLNumber(0)),
		},
	)
}

type RecordEditResponseNSTTLNumber float64

const (
	RecordEditResponseNSTTLNumber1 RecordEditResponseNSTTLNumber = 1
)

type RecordEditResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponsePTRType `json:"type,required"`
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
	Meta RecordEditResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordEditResponsePTRJSON `json:"-"`
}

// recordEditResponsePTRJSON contains the JSON metadata for the struct
// [RecordEditResponsePTR]
type recordEditResponsePTRJSON struct {
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

func (r *RecordEditResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponsePTR) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponsePTRType string

const (
	RecordEditResponsePTRTypePTR RecordEditResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordEditResponsePTRMetaJSON `json:"-"`
}

// recordEditResponsePTRMetaJSON contains the JSON metadata for the struct
// [RecordEditResponsePTRMeta]
type recordEditResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponsePTRTTLNumber].
type RecordEditResponsePTRTTL interface {
	ImplementsDNSRecordEditResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponsePTRTTLNumber(0)),
		},
	)
}

type RecordEditResponsePTRTTLNumber float64

const (
	RecordEditResponsePTRTTLNumber1 RecordEditResponsePTRTTLNumber = 1
)

type RecordEditResponseSmimea struct {
	// Components of a SMIMEA record.
	Data RecordEditResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseSmimeaType `json:"type,required"`
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
	Meta RecordEditResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                       `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseSmimeaJSON `json:"-"`
}

// recordEditResponseSmimeaJSON contains the JSON metadata for the struct
// [RecordEditResponseSmimea]
type recordEditResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSmimea) implementsDNSRecordEditResponse() {}

// Components of a SMIMEA record.
type RecordEditResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                          `json:"usage"`
	JSON  recordEditResponseSmimeaDataJSON `json:"-"`
}

// recordEditResponseSmimeaDataJSON contains the JSON metadata for the struct
// [RecordEditResponseSmimeaData]
type recordEditResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordEditResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseSmimeaType string

const (
	RecordEditResponseSmimeaTypeSmimea RecordEditResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                           `json:"source"`
	JSON   recordEditResponseSmimeaMetaJSON `json:"-"`
}

// recordEditResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseSmimeaMeta]
type recordEditResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordEditResponseSmimeaTTLNumber].
type RecordEditResponseSmimeaTTL interface {
	ImplementsDNSRecordEditResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseSmimeaTTLNumber(0)),
		},
	)
}

type RecordEditResponseSmimeaTTLNumber float64

const (
	RecordEditResponseSmimeaTTLNumber1 RecordEditResponseSmimeaTTLNumber = 1
)

type RecordEditResponseSRV struct {
	// Components of a SRV record.
	Data RecordEditResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseSRVType `json:"type,required"`
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
	Meta RecordEditResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseSRVJSON `json:"-"`
}

// recordEditResponseSRVJSON contains the JSON metadata for the struct
// [RecordEditResponseSRV]
type recordEditResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSRV) implementsDNSRecordEditResponse() {}

// Components of a SRV record.
type RecordEditResponseSRVData struct {
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
	Weight float64                       `json:"weight"`
	JSON   recordEditResponseSRVDataJSON `json:"-"`
}

// recordEditResponseSRVDataJSON contains the JSON metadata for the struct
// [RecordEditResponseSRVData]
type recordEditResponseSRVDataJSON struct {
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

func (r *RecordEditResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseSRVType string

const (
	RecordEditResponseSRVTypeSRV RecordEditResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordEditResponseSRVMetaJSON `json:"-"`
}

// recordEditResponseSRVMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseSRVMeta]
type recordEditResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseSRVTTLNumber].
type RecordEditResponseSRVTTL interface {
	ImplementsDNSRecordEditResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseSRVTTLNumber(0)),
		},
	)
}

type RecordEditResponseSRVTTLNumber float64

const (
	RecordEditResponseSRVTTLNumber1 RecordEditResponseSRVTTLNumber = 1
)

type RecordEditResponseSSHFP struct {
	// Components of a SSHFP record.
	Data RecordEditResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseSSHFPType `json:"type,required"`
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
	Meta RecordEditResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseSSHFPJSON `json:"-"`
}

// recordEditResponseSSHFPJSON contains the JSON metadata for the struct
// [RecordEditResponseSSHFP]
type recordEditResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSSHFP) implementsDNSRecordEditResponse() {}

// Components of a SSHFP record.
type RecordEditResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                         `json:"type"`
	JSON recordEditResponseSSHFPDataJSON `json:"-"`
}

// recordEditResponseSSHFPDataJSON contains the JSON metadata for the struct
// [RecordEditResponseSSHFPData]
type recordEditResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseSSHFPType string

const (
	RecordEditResponseSSHFPTypeSSHFP RecordEditResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordEditResponseSSHFPMetaJSON `json:"-"`
}

// recordEditResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseSSHFPMeta]
type recordEditResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordEditResponseSSHFPTTLNumber].
type RecordEditResponseSSHFPTTL interface {
	ImplementsDNSRecordEditResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseSSHFPTTLNumber(0)),
		},
	)
}

type RecordEditResponseSSHFPTTLNumber float64

const (
	RecordEditResponseSSHFPTTLNumber1 RecordEditResponseSSHFPTTLNumber = 1
)

type RecordEditResponseSVCB struct {
	// Components of a SVCB record.
	Data RecordEditResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseSVCBType `json:"type,required"`
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
	Meta RecordEditResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseSVCBJSON `json:"-"`
}

// recordEditResponseSVCBJSON contains the JSON metadata for the struct
// [RecordEditResponseSVCB]
type recordEditResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSVCB) implementsDNSRecordEditResponse() {}

// Components of a SVCB record.
type RecordEditResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                         `json:"value"`
	JSON  recordEditResponseSVCBDataJSON `json:"-"`
}

// recordEditResponseSVCBDataJSON contains the JSON metadata for the struct
// [RecordEditResponseSVCBData]
type recordEditResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseSVCBType string

const (
	RecordEditResponseSVCBTypeSVCB RecordEditResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordEditResponseSVCBMetaJSON `json:"-"`
}

// recordEditResponseSVCBMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseSVCBMeta]
type recordEditResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseSVCBTTLNumber].
type RecordEditResponseSVCBTTL interface {
	ImplementsDNSRecordEditResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseSVCBTTLNumber(0)),
		},
	)
}

type RecordEditResponseSVCBTTLNumber float64

const (
	RecordEditResponseSVCBTTLNumber1 RecordEditResponseSVCBTTLNumber = 1
)

type RecordEditResponseTLSA struct {
	// Components of a TLSA record.
	Data RecordEditResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseTLSAType `json:"type,required"`
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
	Meta RecordEditResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseTLSAJSON `json:"-"`
}

// recordEditResponseTLSAJSON contains the JSON metadata for the struct
// [RecordEditResponseTLSA]
type recordEditResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseTLSA) implementsDNSRecordEditResponse() {}

// Components of a TLSA record.
type RecordEditResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                        `json:"usage"`
	JSON  recordEditResponseTLSADataJSON `json:"-"`
}

// recordEditResponseTLSADataJSON contains the JSON metadata for the struct
// [RecordEditResponseTLSAData]
type recordEditResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordEditResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseTLSAType string

const (
	RecordEditResponseTLSATypeTLSA RecordEditResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordEditResponseTLSAMetaJSON `json:"-"`
}

// recordEditResponseTLSAMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseTLSAMeta]
type recordEditResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseTLSATTLNumber].
type RecordEditResponseTLSATTL interface {
	ImplementsDNSRecordEditResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseTLSATTLNumber(0)),
		},
	)
}

type RecordEditResponseTLSATTLNumber float64

const (
	RecordEditResponseTLSATTLNumber1 RecordEditResponseTLSATTLNumber = 1
)

type RecordEditResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordEditResponseTXTType `json:"type,required"`
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
	Meta RecordEditResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseTXTJSON `json:"-"`
}

// recordEditResponseTXTJSON contains the JSON metadata for the struct
// [RecordEditResponseTXT]
type recordEditResponseTXTJSON struct {
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

func (r *RecordEditResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseTXT) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseTXTType string

const (
	RecordEditResponseTXTTypeTXT RecordEditResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordEditResponseTXTMetaJSON `json:"-"`
}

// recordEditResponseTXTMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseTXTMeta]
type recordEditResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseTXTTTLNumber].
type RecordEditResponseTXTTTL interface {
	ImplementsDNSRecordEditResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseTXTTTLNumber(0)),
		},
	)
}

type RecordEditResponseTXTTTLNumber float64

const (
	RecordEditResponseTXTTTLNumber1 RecordEditResponseTXTTTLNumber = 1
)

type RecordEditResponseURI struct {
	// Components of a URI record.
	Data RecordEditResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordEditResponseURIType `json:"type,required"`
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
	Meta RecordEditResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordEditResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordEditResponseURIJSON `json:"-"`
}

// recordEditResponseURIJSON contains the JSON metadata for the struct
// [RecordEditResponseURI]
type recordEditResponseURIJSON struct {
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

func (r *RecordEditResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseURI) implementsDNSRecordEditResponse() {}

// Components of a URI record.
type RecordEditResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                       `json:"weight"`
	JSON   recordEditResponseURIDataJSON `json:"-"`
}

// recordEditResponseURIDataJSON contains the JSON metadata for the struct
// [RecordEditResponseURIData]
type recordEditResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordEditResponseURIType string

const (
	RecordEditResponseURITypeURI RecordEditResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type RecordEditResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordEditResponseURIMetaJSON `json:"-"`
}

// recordEditResponseURIMetaJSON contains the JSON metadata for the struct
// [RecordEditResponseURIMeta]
type recordEditResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordEditResponseURITTLNumber].
type RecordEditResponseURITTL interface {
	ImplementsDNSRecordEditResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordEditResponseURITTLNumber(0)),
		},
	)
}

type RecordEditResponseURITTLNumber float64

const (
	RecordEditResponseURITTLNumber1 RecordEditResponseURITTLNumber = 1
)

// Union satisfied by [dns.RecordGetResponseA], [dns.RecordGetResponseAAAA],
// [dns.RecordGetResponseCAA], [dns.RecordGetResponseCert],
// [dns.RecordGetResponseCNAME], [dns.RecordGetResponseDNSKEY],
// [dns.RecordGetResponseDS], [dns.RecordGetResponseHTTPS],
// [dns.RecordGetResponseLOC], [dns.RecordGetResponseMX],
// [dns.RecordGetResponseNAPTR], [dns.RecordGetResponseNS],
// [dns.RecordGetResponsePTR], [dns.RecordGetResponseSmimea],
// [dns.RecordGetResponseSRV], [dns.RecordGetResponseSSHFP],
// [dns.RecordGetResponseSVCB], [dns.RecordGetResponseTLSA],
// [dns.RecordGetResponseTXT] or [dns.RecordGetResponseURI].
type RecordGetResponse interface {
	implementsDNSRecordGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponse)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseA{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseAAAA{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseCAA{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseCert{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseCNAME{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseDNSKEY{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseDS{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseHTTPS{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseLOC{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseMX{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseNAPTR{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseNS{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponsePTR{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseSmimea{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseSRV{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseSSHFP{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseSVCB{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseTLSA{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseTXT{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseURI{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordGetResponseA struct {
	// A valid IPv4 address.
	Content string `json:"content,required" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseAType `json:"type,required"`
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
	Meta RecordGetResponseAMeta `json:"meta"`
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
	TTL RecordGetResponseATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                 `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseAJSON `json:"-"`
}

// recordGetResponseAJSON contains the JSON metadata for the struct
// [RecordGetResponseA]
type recordGetResponseAJSON struct {
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

func (r *RecordGetResponseA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseAJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseA) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseAType string

const (
	RecordGetResponseATypeA RecordGetResponseAType = "A"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                     `json:"source"`
	JSON   recordGetResponseAMetaJSON `json:"-"`
}

// recordGetResponseAMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseAMeta]
type recordGetResponseAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseATTLNumber].
type RecordGetResponseATTL interface {
	ImplementsDNSRecordGetResponseAttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseATTLNumber(0)),
		},
	)
}

type RecordGetResponseATTLNumber float64

const (
	RecordGetResponseATTLNumber1 RecordGetResponseATTLNumber = 1
)

type RecordGetResponseAAAA struct {
	// A valid IPv6 address.
	Content string `json:"content,required" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseAAAAType `json:"type,required"`
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
	Meta RecordGetResponseAAAAMeta `json:"meta"`
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
	TTL RecordGetResponseAAAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseAAAAJSON `json:"-"`
}

// recordGetResponseAAAAJSON contains the JSON metadata for the struct
// [RecordGetResponseAAAA]
type recordGetResponseAAAAJSON struct {
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

func (r *RecordGetResponseAAAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseAAAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseAAAA) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseAAAAType string

const (
	RecordGetResponseAAAATypeAAAA RecordGetResponseAAAAType = "AAAA"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseAAAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordGetResponseAAAAMetaJSON `json:"-"`
}

// recordGetResponseAAAAMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseAAAAMeta]
type recordGetResponseAAAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseAAAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseAAAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseAAAATTLNumber].
type RecordGetResponseAAAATTL interface {
	ImplementsDNSRecordGetResponseAaaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseAAAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseAAAATTLNumber(0)),
		},
	)
}

type RecordGetResponseAAAATTLNumber float64

const (
	RecordGetResponseAAAATTLNumber1 RecordGetResponseAAAATTLNumber = 1
)

type RecordGetResponseCAA struct {
	// Components of a CAA record.
	Data RecordGetResponseCAAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseCAAType `json:"type,required"`
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
	Meta RecordGetResponseCAAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseCAATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseCAAJSON `json:"-"`
}

// recordGetResponseCAAJSON contains the JSON metadata for the struct
// [RecordGetResponseCAA]
type recordGetResponseCAAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCAA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCAAJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseCAA) implementsDNSRecordGetResponse() {}

// Components of a CAA record.
type RecordGetResponseCAAData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string                       `json:"value"`
	JSON  recordGetResponseCAADataJSON `json:"-"`
}

// recordGetResponseCAADataJSON contains the JSON metadata for the struct
// [RecordGetResponseCAAData]
type recordGetResponseCAADataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCAAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCAADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseCAAType string

const (
	RecordGetResponseCAATypeCAA RecordGetResponseCAAType = "CAA"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseCAAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordGetResponseCAAMetaJSON `json:"-"`
}

// recordGetResponseCAAMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseCAAMeta]
type recordGetResponseCAAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCAAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCAAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseCAATTLNumber].
type RecordGetResponseCAATTL interface {
	ImplementsDNSRecordGetResponseCaattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseCAATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseCAATTLNumber(0)),
		},
	)
}

type RecordGetResponseCAATTLNumber float64

const (
	RecordGetResponseCAATTLNumber1 RecordGetResponseCAATTLNumber = 1
)

type RecordGetResponseCert struct {
	// Components of a CERT record.
	Data RecordGetResponseCertData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseCertType `json:"type,required"`
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
	Meta RecordGetResponseCertMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseCertTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseCertJSON `json:"-"`
}

// recordGetResponseCertJSON contains the JSON metadata for the struct
// [RecordGetResponseCert]
type recordGetResponseCertJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCertJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseCert) implementsDNSRecordGetResponse() {}

// Components of a CERT record.
type RecordGetResponseCertData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64                       `json:"type"`
	JSON recordGetResponseCertDataJSON `json:"-"`
}

// recordGetResponseCertDataJSON contains the JSON metadata for the struct
// [RecordGetResponseCertData]
type recordGetResponseCertDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCertData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCertDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseCertType string

const (
	RecordGetResponseCertTypeCert RecordGetResponseCertType = "CERT"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseCertMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordGetResponseCertMetaJSON `json:"-"`
}

// recordGetResponseCertMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseCertMeta]
type recordGetResponseCertMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCertMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCertMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseCertTTLNumber].
type RecordGetResponseCertTTL interface {
	ImplementsDNSRecordGetResponseCertTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseCertTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseCertTTLNumber(0)),
		},
	)
}

type RecordGetResponseCertTTLNumber float64

const (
	RecordGetResponseCertTTLNumber1 RecordGetResponseCertTTLNumber = 1
)

type RecordGetResponseCNAME struct {
	// A valid hostname. Must not match the record's name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseCNAMEType `json:"type,required"`
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
	Meta RecordGetResponseCNAMEMeta `json:"meta"`
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
	TTL RecordGetResponseCNAMETTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseCNAMEJSON `json:"-"`
}

// recordGetResponseCNAMEJSON contains the JSON metadata for the struct
// [RecordGetResponseCNAME]
type recordGetResponseCNAMEJSON struct {
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

func (r *RecordGetResponseCNAME) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCNAMEJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseCNAME) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseCNAMEType string

const (
	RecordGetResponseCNAMETypeCNAME RecordGetResponseCNAMEType = "CNAME"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseCNAMEMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordGetResponseCNAMEMetaJSON `json:"-"`
}

// recordGetResponseCNAMEMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseCNAMEMeta]
type recordGetResponseCNAMEMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseCNAMEMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCNAMEMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseCNAMETTLNumber].
type RecordGetResponseCNAMETTL interface {
	ImplementsDNSRecordGetResponseCnamettl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseCNAMETTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseCNAMETTLNumber(0)),
		},
	)
}

type RecordGetResponseCNAMETTLNumber float64

const (
	RecordGetResponseCNAMETTLNumber1 RecordGetResponseCNAMETTLNumber = 1
)

type RecordGetResponseDNSKEY struct {
	// Components of a DNSKEY record.
	Data RecordGetResponseDNSKEYData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseDNSKEYType `json:"type,required"`
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
	Meta RecordGetResponseDNSKEYMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseDNSKEYTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseDNSKEYJSON `json:"-"`
}

// recordGetResponseDNSKEYJSON contains the JSON metadata for the struct
// [RecordGetResponseDNSKEY]
type recordGetResponseDNSKEYJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseDNSKEY) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDNSKEYJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseDNSKEY) implementsDNSRecordGetResponse() {}

// Components of a DNSKEY record.
type RecordGetResponseDNSKEYData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string                          `json:"public_key"`
	JSON      recordGetResponseDNSKEYDataJSON `json:"-"`
}

// recordGetResponseDNSKEYDataJSON contains the JSON metadata for the struct
// [RecordGetResponseDNSKEYData]
type recordGetResponseDNSKEYDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseDNSKEYData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDNSKEYDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseDNSKEYType string

const (
	RecordGetResponseDNSKEYTypeDNSKEY RecordGetResponseDNSKEYType = "DNSKEY"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseDNSKEYMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordGetResponseDNSKEYMetaJSON `json:"-"`
}

// recordGetResponseDNSKEYMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseDNSKEYMeta]
type recordGetResponseDNSKEYMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseDNSKEYMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDNSKEYMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordGetResponseDNSKEYTTLNumber].
type RecordGetResponseDNSKEYTTL interface {
	ImplementsDNSRecordGetResponseDnskeyttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseDNSKEYTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseDNSKEYTTLNumber(0)),
		},
	)
}

type RecordGetResponseDNSKEYTTLNumber float64

const (
	RecordGetResponseDNSKEYTTLNumber1 RecordGetResponseDNSKEYTTLNumber = 1
)

type RecordGetResponseDS struct {
	// Components of a DS record.
	Data RecordGetResponseDSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseDSType `json:"type,required"`
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
	Meta RecordGetResponseDSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseDSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseDSJSON `json:"-"`
}

// recordGetResponseDSJSON contains the JSON metadata for the struct
// [RecordGetResponseDS]
type recordGetResponseDSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseDS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDSJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseDS) implementsDNSRecordGetResponse() {}

// Components of a DS record.
type RecordGetResponseDSData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64                     `json:"key_tag"`
	JSON   recordGetResponseDSDataJSON `json:"-"`
}

// recordGetResponseDSDataJSON contains the JSON metadata for the struct
// [RecordGetResponseDSData]
type recordGetResponseDSDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseDSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseDSType string

const (
	RecordGetResponseDSTypeDS RecordGetResponseDSType = "DS"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseDSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordGetResponseDSMetaJSON `json:"-"`
}

// recordGetResponseDSMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseDSMeta]
type recordGetResponseDSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseDSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseDSTTLNumber].
type RecordGetResponseDSTTL interface {
	ImplementsDNSRecordGetResponseDsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseDSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseDSTTLNumber(0)),
		},
	)
}

type RecordGetResponseDSTTLNumber float64

const (
	RecordGetResponseDSTTLNumber1 RecordGetResponseDSTTLNumber = 1
)

type RecordGetResponseHTTPS struct {
	// Components of a HTTPS record.
	Data RecordGetResponseHTTPSData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseHTTPSType `json:"type,required"`
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
	Meta RecordGetResponseHTTPSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseHTTPSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseHTTPSJSON `json:"-"`
}

// recordGetResponseHTTPSJSON contains the JSON metadata for the struct
// [RecordGetResponseHTTPS]
type recordGetResponseHTTPSJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseHTTPS) implementsDNSRecordGetResponse() {}

// Components of a HTTPS record.
type RecordGetResponseHTTPSData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                         `json:"value"`
	JSON  recordGetResponseHTTPSDataJSON `json:"-"`
}

// recordGetResponseHTTPSDataJSON contains the JSON metadata for the struct
// [RecordGetResponseHTTPSData]
type recordGetResponseHTTPSDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseHTTPSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseHTTPSDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseHTTPSType string

const (
	RecordGetResponseHTTPSTypeHTTPS RecordGetResponseHTTPSType = "HTTPS"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseHTTPSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordGetResponseHTTPSMetaJSON `json:"-"`
}

// recordGetResponseHTTPSMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseHTTPSMeta]
type recordGetResponseHTTPSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseHTTPSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseHTTPSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseHTTPSTTLNumber].
type RecordGetResponseHTTPSTTL interface {
	ImplementsDNSRecordGetResponseHttpsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseHTTPSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseHTTPSTTLNumber(0)),
		},
	)
}

type RecordGetResponseHTTPSTTLNumber float64

const (
	RecordGetResponseHTTPSTTLNumber1 RecordGetResponseHTTPSTTLNumber = 1
)

type RecordGetResponseLOC struct {
	// Components of a LOC record.
	Data RecordGetResponseLOCData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseLOCType `json:"type,required"`
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
	Meta RecordGetResponseLOCMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseLOCTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseLOCJSON `json:"-"`
}

// recordGetResponseLOCJSON contains the JSON metadata for the struct
// [RecordGetResponseLOC]
type recordGetResponseLOCJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseLOC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseLOCJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseLOC) implementsDNSRecordGetResponse() {}

// Components of a LOC record.
type RecordGetResponseLOCData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection RecordGetResponseLOCDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection RecordGetResponseLOCDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64                      `json:"size"`
	JSON recordGetResponseLOCDataJSON `json:"-"`
}

// recordGetResponseLOCDataJSON contains the JSON metadata for the struct
// [RecordGetResponseLOCData]
type recordGetResponseLOCDataJSON struct {
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

func (r *RecordGetResponseLOCData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseLOCDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type RecordGetResponseLOCDataLatDirection string

const (
	RecordGetResponseLOCDataLatDirectionN RecordGetResponseLOCDataLatDirection = "N"
	RecordGetResponseLOCDataLatDirectionS RecordGetResponseLOCDataLatDirection = "S"
)

// Longitude direction.
type RecordGetResponseLOCDataLongDirection string

const (
	RecordGetResponseLOCDataLongDirectionE RecordGetResponseLOCDataLongDirection = "E"
	RecordGetResponseLOCDataLongDirectionW RecordGetResponseLOCDataLongDirection = "W"
)

// Record type.
type RecordGetResponseLOCType string

const (
	RecordGetResponseLOCTypeLOC RecordGetResponseLOCType = "LOC"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseLOCMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordGetResponseLOCMetaJSON `json:"-"`
}

// recordGetResponseLOCMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseLOCMeta]
type recordGetResponseLOCMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseLOCMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseLOCMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseLOCTTLNumber].
type RecordGetResponseLOCTTL interface {
	ImplementsDNSRecordGetResponseLocttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseLOCTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseLOCTTLNumber(0)),
		},
	)
}

type RecordGetResponseLOCTTLNumber float64

const (
	RecordGetResponseLOCTTLNumber1 RecordGetResponseLOCTTLNumber = 1
)

type RecordGetResponseMX struct {
	// A valid mail server hostname.
	Content string `json:"content,required" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordGetResponseMXType `json:"type,required"`
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
	Meta RecordGetResponseMXMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseMXTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseMXJSON `json:"-"`
}

// recordGetResponseMXJSON contains the JSON metadata for the struct
// [RecordGetResponseMX]
type recordGetResponseMXJSON struct {
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

func (r *RecordGetResponseMX) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseMXJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseMX) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseMXType string

const (
	RecordGetResponseMXTypeMX RecordGetResponseMXType = "MX"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseMXMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordGetResponseMXMetaJSON `json:"-"`
}

// recordGetResponseMXMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseMXMeta]
type recordGetResponseMXMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseMXMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseMXMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseMXTTLNumber].
type RecordGetResponseMXTTL interface {
	ImplementsDNSRecordGetResponseMxttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseMXTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseMXTTLNumber(0)),
		},
	)
}

type RecordGetResponseMXTTLNumber float64

const (
	RecordGetResponseMXTTLNumber1 RecordGetResponseMXTTLNumber = 1
)

type RecordGetResponseNAPTR struct {
	// Components of a NAPTR record.
	Data RecordGetResponseNAPTRData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseNAPTRType `json:"type,required"`
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
	Meta RecordGetResponseNAPTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseNAPTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseNAPTRJSON `json:"-"`
}

// recordGetResponseNAPTRJSON contains the JSON metadata for the struct
// [RecordGetResponseNAPTR]
type recordGetResponseNAPTRJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseNAPTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNAPTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseNAPTR) implementsDNSRecordGetResponse() {}

// Components of a NAPTR record.
type RecordGetResponseNAPTRData struct {
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
	Service string                         `json:"service"`
	JSON    recordGetResponseNAPTRDataJSON `json:"-"`
}

// recordGetResponseNAPTRDataJSON contains the JSON metadata for the struct
// [RecordGetResponseNAPTRData]
type recordGetResponseNAPTRDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseNAPTRData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNAPTRDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseNAPTRType string

const (
	RecordGetResponseNAPTRTypeNAPTR RecordGetResponseNAPTRType = "NAPTR"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseNAPTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordGetResponseNAPTRMetaJSON `json:"-"`
}

// recordGetResponseNAPTRMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseNAPTRMeta]
type recordGetResponseNAPTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseNAPTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNAPTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseNAPTRTTLNumber].
type RecordGetResponseNAPTRTTL interface {
	ImplementsDNSRecordGetResponseNaptrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseNAPTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseNAPTRTTLNumber(0)),
		},
	)
}

type RecordGetResponseNAPTRTTLNumber float64

const (
	RecordGetResponseNAPTRTTLNumber1 RecordGetResponseNAPTRTTLNumber = 1
)

type RecordGetResponseNS struct {
	// A valid name server host name.
	Content interface{} `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseNSType `json:"type,required"`
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
	Meta RecordGetResponseNSMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseNSTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                  `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseNSJSON `json:"-"`
}

// recordGetResponseNSJSON contains the JSON metadata for the struct
// [RecordGetResponseNS]
type recordGetResponseNSJSON struct {
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

func (r *RecordGetResponseNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNSJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseNS) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseNSType string

const (
	RecordGetResponseNSTypeNS RecordGetResponseNSType = "NS"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseNSMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                      `json:"source"`
	JSON   recordGetResponseNSMetaJSON `json:"-"`
}

// recordGetResponseNSMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseNSMeta]
type recordGetResponseNSMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseNSMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNSMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseNSTTLNumber].
type RecordGetResponseNSTTL interface {
	ImplementsDNSRecordGetResponseNsttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseNSTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseNSTTLNumber(0)),
		},
	)
}

type RecordGetResponseNSTTLNumber float64

const (
	RecordGetResponseNSTTLNumber1 RecordGetResponseNSTTLNumber = 1
)

type RecordGetResponsePTR struct {
	// Domain name pointing to the address.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponsePTRType `json:"type,required"`
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
	Meta RecordGetResponsePTRMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponsePTRTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordGetResponsePTRJSON `json:"-"`
}

// recordGetResponsePTRJSON contains the JSON metadata for the struct
// [RecordGetResponsePTR]
type recordGetResponsePTRJSON struct {
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

func (r *RecordGetResponsePTR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponsePTRJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponsePTR) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponsePTRType string

const (
	RecordGetResponsePTRTypePTR RecordGetResponsePTRType = "PTR"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponsePTRMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordGetResponsePTRMetaJSON `json:"-"`
}

// recordGetResponsePTRMetaJSON contains the JSON metadata for the struct
// [RecordGetResponsePTRMeta]
type recordGetResponsePTRMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponsePTRMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponsePTRMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponsePTRTTLNumber].
type RecordGetResponsePTRTTL interface {
	ImplementsDNSRecordGetResponsePtrttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponsePTRTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponsePTRTTLNumber(0)),
		},
	)
}

type RecordGetResponsePTRTTLNumber float64

const (
	RecordGetResponsePTRTTLNumber1 RecordGetResponsePTRTTLNumber = 1
)

type RecordGetResponseSmimea struct {
	// Components of a SMIMEA record.
	Data RecordGetResponseSmimeaData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseSmimeaType `json:"type,required"`
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
	Meta RecordGetResponseSmimeaMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseSmimeaTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                      `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseSmimeaJSON `json:"-"`
}

// recordGetResponseSmimeaJSON contains the JSON metadata for the struct
// [RecordGetResponseSmimea]
type recordGetResponseSmimeaJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSmimea) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSmimeaJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSmimea) implementsDNSRecordGetResponse() {}

// Components of a SMIMEA record.
type RecordGetResponseSmimeaData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                         `json:"usage"`
	JSON  recordGetResponseSmimeaDataJSON `json:"-"`
}

// recordGetResponseSmimeaDataJSON contains the JSON metadata for the struct
// [RecordGetResponseSmimeaData]
type recordGetResponseSmimeaDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordGetResponseSmimeaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSmimeaDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseSmimeaType string

const (
	RecordGetResponseSmimeaTypeSmimea RecordGetResponseSmimeaType = "SMIMEA"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseSmimeaMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                          `json:"source"`
	JSON   recordGetResponseSmimeaMetaJSON `json:"-"`
}

// recordGetResponseSmimeaMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseSmimeaMeta]
type recordGetResponseSmimeaMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSmimeaMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSmimeaMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or
// [dns.RecordGetResponseSmimeaTTLNumber].
type RecordGetResponseSmimeaTTL interface {
	ImplementsDNSRecordGetResponseSmimeaTTL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseSmimeaTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseSmimeaTTLNumber(0)),
		},
	)
}

type RecordGetResponseSmimeaTTLNumber float64

const (
	RecordGetResponseSmimeaTTLNumber1 RecordGetResponseSmimeaTTLNumber = 1
)

type RecordGetResponseSRV struct {
	// Components of a SRV record.
	Data RecordGetResponseSRVData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode. For SRV records, the first
	// label is normally a service and the second a protocol name, each starting with
	// an underscore.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseSRVType `json:"type,required"`
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
	Meta RecordGetResponseSRVMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseSRVTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseSRVJSON `json:"-"`
}

// recordGetResponseSRVJSON contains the JSON metadata for the struct
// [RecordGetResponseSRV]
type recordGetResponseSRVJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSRV) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSRVJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSRV) implementsDNSRecordGetResponse() {}

// Components of a SRV record.
type RecordGetResponseSRVData struct {
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
	Weight float64                      `json:"weight"`
	JSON   recordGetResponseSRVDataJSON `json:"-"`
}

// recordGetResponseSRVDataJSON contains the JSON metadata for the struct
// [RecordGetResponseSRVData]
type recordGetResponseSRVDataJSON struct {
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

func (r *RecordGetResponseSRVData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSRVDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseSRVType string

const (
	RecordGetResponseSRVTypeSRV RecordGetResponseSRVType = "SRV"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseSRVMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordGetResponseSRVMetaJSON `json:"-"`
}

// recordGetResponseSRVMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseSRVMeta]
type recordGetResponseSRVMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSRVMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSRVMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseSRVTTLNumber].
type RecordGetResponseSRVTTL interface {
	ImplementsDNSRecordGetResponseSrvttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseSRVTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseSRVTTLNumber(0)),
		},
	)
}

type RecordGetResponseSRVTTLNumber float64

const (
	RecordGetResponseSRVTTLNumber1 RecordGetResponseSRVTTLNumber = 1
)

type RecordGetResponseSSHFP struct {
	// Components of a SSHFP record.
	Data RecordGetResponseSSHFPData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseSSHFPType `json:"type,required"`
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
	Meta RecordGetResponseSSHFPMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseSSHFPTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                     `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseSSHFPJSON `json:"-"`
}

// recordGetResponseSSHFPJSON contains the JSON metadata for the struct
// [RecordGetResponseSSHFP]
type recordGetResponseSSHFPJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSSHFP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSSHFPJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSSHFP) implementsDNSRecordGetResponse() {}

// Components of a SSHFP record.
type RecordGetResponseSSHFPData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64                        `json:"type"`
	JSON recordGetResponseSSHFPDataJSON `json:"-"`
}

// recordGetResponseSSHFPDataJSON contains the JSON metadata for the struct
// [RecordGetResponseSSHFPData]
type recordGetResponseSSHFPDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSSHFPData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSSHFPDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseSSHFPType string

const (
	RecordGetResponseSSHFPTypeSSHFP RecordGetResponseSSHFPType = "SSHFP"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseSSHFPMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                         `json:"source"`
	JSON   recordGetResponseSSHFPMetaJSON `json:"-"`
}

// recordGetResponseSSHFPMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseSSHFPMeta]
type recordGetResponseSSHFPMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSSHFPMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSSHFPMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseSSHFPTTLNumber].
type RecordGetResponseSSHFPTTL interface {
	ImplementsDNSRecordGetResponseSshfpttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseSSHFPTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseSSHFPTTLNumber(0)),
		},
	)
}

type RecordGetResponseSSHFPTTLNumber float64

const (
	RecordGetResponseSSHFPTTLNumber1 RecordGetResponseSSHFPTTLNumber = 1
)

type RecordGetResponseSVCB struct {
	// Components of a SVCB record.
	Data RecordGetResponseSVCBData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseSVCBType `json:"type,required"`
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
	Meta RecordGetResponseSVCBMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseSVCBTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseSVCBJSON `json:"-"`
}

// recordGetResponseSVCBJSON contains the JSON metadata for the struct
// [RecordGetResponseSVCB]
type recordGetResponseSVCBJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSVCB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSVCBJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSVCB) implementsDNSRecordGetResponse() {}

// Components of a SVCB record.
type RecordGetResponseSVCBData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string                        `json:"value"`
	JSON  recordGetResponseSVCBDataJSON `json:"-"`
}

// recordGetResponseSVCBDataJSON contains the JSON metadata for the struct
// [RecordGetResponseSVCBData]
type recordGetResponseSVCBDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSVCBData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSVCBDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseSVCBType string

const (
	RecordGetResponseSVCBTypeSVCB RecordGetResponseSVCBType = "SVCB"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseSVCBMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordGetResponseSVCBMetaJSON `json:"-"`
}

// recordGetResponseSVCBMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseSVCBMeta]
type recordGetResponseSVCBMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseSVCBMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSVCBMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseSVCBTTLNumber].
type RecordGetResponseSVCBTTL interface {
	ImplementsDNSRecordGetResponseSvcbttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseSVCBTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseSVCBTTLNumber(0)),
		},
	)
}

type RecordGetResponseSVCBTTLNumber float64

const (
	RecordGetResponseSVCBTTLNumber1 RecordGetResponseSVCBTTLNumber = 1
)

type RecordGetResponseTLSA struct {
	// Components of a TLSA record.
	Data RecordGetResponseTLSAData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseTLSAType `json:"type,required"`
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
	Meta RecordGetResponseTLSAMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseTLSATTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                    `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseTLSAJSON `json:"-"`
}

// recordGetResponseTLSAJSON contains the JSON metadata for the struct
// [RecordGetResponseTLSA]
type recordGetResponseTLSAJSON struct {
	Data        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	Comment     apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Locked      apijson.Field
	Meta        apijson.Field
	ModifiedOn  apijson.Field
	Proxiable   apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseTLSA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTLSAJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseTLSA) implementsDNSRecordGetResponse() {}

// Components of a TLSA record.
type RecordGetResponseTLSAData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64                       `json:"usage"`
	JSON  recordGetResponseTLSADataJSON `json:"-"`
}

// recordGetResponseTLSADataJSON contains the JSON metadata for the struct
// [RecordGetResponseTLSAData]
type recordGetResponseTLSADataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RecordGetResponseTLSAData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTLSADataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseTLSAType string

const (
	RecordGetResponseTLSATypeTLSA RecordGetResponseTLSAType = "TLSA"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseTLSAMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                        `json:"source"`
	JSON   recordGetResponseTLSAMetaJSON `json:"-"`
}

// recordGetResponseTLSAMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseTLSAMeta]
type recordGetResponseTLSAMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseTLSAMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTLSAMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseTLSATTLNumber].
type RecordGetResponseTLSATTL interface {
	ImplementsDNSRecordGetResponseTlsattl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseTLSATTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseTLSATTLNumber(0)),
		},
	)
}

type RecordGetResponseTLSATTLNumber float64

const (
	RecordGetResponseTLSATTLNumber1 RecordGetResponseTLSATTLNumber = 1
)

type RecordGetResponseTXT struct {
	// Text content for the record.
	Content string `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Record type.
	Type RecordGetResponseTXTType `json:"type,required"`
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
	Meta RecordGetResponseTXTMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseTXTTTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseTXTJSON `json:"-"`
}

// recordGetResponseTXTJSON contains the JSON metadata for the struct
// [RecordGetResponseTXT]
type recordGetResponseTXTJSON struct {
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

func (r *RecordGetResponseTXT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTXTJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseTXT) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseTXTType string

const (
	RecordGetResponseTXTTypeTXT RecordGetResponseTXTType = "TXT"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseTXTMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordGetResponseTXTMetaJSON `json:"-"`
}

// recordGetResponseTXTMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseTXTMeta]
type recordGetResponseTXTMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseTXTMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTXTMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseTXTTTLNumber].
type RecordGetResponseTXTTTL interface {
	ImplementsDNSRecordGetResponseTxtttl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseTXTTTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseTXTTTLNumber(0)),
		},
	)
}

type RecordGetResponseTXTTTLNumber float64

const (
	RecordGetResponseTXTTTLNumber1 RecordGetResponseTXTTTLNumber = 1
)

type RecordGetResponseURI struct {
	// Components of a URI record.
	Data RecordGetResponseURIData `json:"data,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority,required"`
	// Record type.
	Type RecordGetResponseURIType `json:"type,required"`
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
	Meta RecordGetResponseURIMeta `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []string `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL RecordGetResponseURITTL `json:"ttl"`
	// Identifier
	ZoneID string `json:"zone_id"`
	// The domain of the record.
	ZoneName string                   `json:"zone_name" format:"hostname"`
	JSON     recordGetResponseURIJSON `json:"-"`
}

// recordGetResponseURIJSON contains the JSON metadata for the struct
// [RecordGetResponseURI]
type recordGetResponseURIJSON struct {
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

func (r *RecordGetResponseURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseURIJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseURI) implementsDNSRecordGetResponse() {}

// Components of a URI record.
type RecordGetResponseURIData struct {
	// The record content.
	Content string `json:"content"`
	// The record weight.
	Weight float64                      `json:"weight"`
	JSON   recordGetResponseURIDataJSON `json:"-"`
}

// recordGetResponseURIDataJSON contains the JSON metadata for the struct
// [RecordGetResponseURIData]
type recordGetResponseURIDataJSON struct {
	Content     apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseURIData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseURIDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordGetResponseURIType string

const (
	RecordGetResponseURITypeURI RecordGetResponseURIType = "URI"
)

// Extra Cloudflare-specific information about the record.
type RecordGetResponseURIMeta struct {
	// Will exist if Cloudflare automatically added this DNS record during initial
	// setup.
	AutoAdded bool `json:"auto_added"`
	// Where the record originated from.
	Source string                       `json:"source"`
	JSON   recordGetResponseURIMetaJSON `json:"-"`
}

// recordGetResponseURIMetaJSON contains the JSON metadata for the struct
// [RecordGetResponseURIMeta]
type recordGetResponseURIMetaJSON struct {
	AutoAdded   apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseURIMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseURIMetaJSON) RawJSON() string {
	return r.raw
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
//
// Union satisfied by [shared.UnionFloat] or [dns.RecordGetResponseURITTLNumber].
type RecordGetResponseURITTL interface {
	ImplementsDNSRecordGetResponseUrittl()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseURITTL)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(RecordGetResponseURITTLNumber(0)),
		},
	)
}

type RecordGetResponseURITTLNumber float64

const (
	RecordGetResponseURITTLNumber1 RecordGetResponseURITTLNumber = 1
)

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
	Comment param.Field[string]              `json:"comment"`
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
	RecordNewParamsTypeURI RecordNewParamsType = "URI"
)

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
	Flags param.Field[string] `json:"flags"`
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

// Longitude direction.
type RecordNewParamsDataLongDirection string

const (
	RecordNewParamsDataLongDirectionE RecordNewParamsDataLongDirection = "E"
	RecordNewParamsDataLongDirectionW RecordNewParamsDataLongDirection = "W"
)

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

type RecordNewResponseEnvelope struct {
	Errors   []RecordNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RecordNewResponse                   `json:"result,required"`
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

type RecordUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordUpdateParamsType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string]                 `json:"comment"`
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
	RecordUpdateParamsTypeURI RecordUpdateParamsType = "URI"
)

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
	Flags param.Field[string] `json:"flags"`
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

// Longitude direction.
type RecordUpdateParamsDataLongDirection string

const (
	RecordUpdateParamsDataLongDirectionE RecordUpdateParamsDataLongDirection = "E"
	RecordUpdateParamsDataLongDirectionW RecordUpdateParamsDataLongDirection = "W"
)

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

type RecordUpdateResponseEnvelope struct {
	Errors   []RecordUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RecordUpdateResponse                   `json:"result,required"`
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

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type RecordListParamsMatch string

const (
	RecordListParamsMatchAny RecordListParamsMatch = "any"
	RecordListParamsMatchAll RecordListParamsMatch = "all"
)

// Field to order DNS records by.
type RecordListParamsOrder string

const (
	RecordListParamsOrderType    RecordListParamsOrder = "type"
	RecordListParamsOrderName    RecordListParamsOrder = "name"
	RecordListParamsOrderContent RecordListParamsOrder = "content"
	RecordListParamsOrderTTL     RecordListParamsOrder = "ttl"
	RecordListParamsOrderProxied RecordListParamsOrder = "proxied"
)

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
	Comment param.Field[string]               `json:"comment"`
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
	RecordEditParamsTypeURI RecordEditParamsType = "URI"
)

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
	Flags param.Field[string] `json:"flags"`
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

// Longitude direction.
type RecordEditParamsDataLongDirection string

const (
	RecordEditParamsDataLongDirectionE RecordEditParamsDataLongDirection = "E"
	RecordEditParamsDataLongDirectionW RecordEditParamsDataLongDirection = "W"
)

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

type RecordEditResponseEnvelope struct {
	Errors   []RecordEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RecordEditResponseEnvelopeMessages `json:"messages,required"`
	Result   RecordEditResponse                   `json:"result,required"`
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
	Result   RecordGetResponse                   `json:"result,required"`
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
