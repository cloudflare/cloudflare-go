// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// RecordService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordService] method instead.
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
	var env RecordNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Overwrite an existing DNS record.
//
// Notes:
//
//   - A/AAAA records cannot exist on the same name as CNAME records.
//   - NS records cannot exist on the same name as any other record type.
//   - Domain names are always represented in Punycode, even if Unicode characters
//     were used when creating the record.
func (r *RecordService) Update(ctx context.Context, dnsRecordID string, params RecordUpdateParams, opts ...option.RequestOption) (res *RecordUpdateResponse, err error) {
	var env RecordUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/%s", params.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *RecordService) List(ctx context.Context, params RecordListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RecordListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
func (r *RecordService) ListAutoPaging(ctx context.Context, params RecordListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RecordListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete DNS Record
func (r *RecordService) Delete(ctx context.Context, dnsRecordID string, body RecordDeleteParams, opts ...option.RequestOption) (res *RecordDeleteResponse, err error) {
	var env RecordDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/%s", body.ZoneID, dnsRecordID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Send a Batch of DNS Record API calls to be executed together.
//
// Notes:
//
//   - Although Cloudflare will execute the batched operations in a single database
//     transaction, Cloudflare's distributed KV store must treat each record change
//     as a single key-value pair. This means that the propagation of changes is not
//     atomic. See
//     [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/batch-record-changes/ "Batch DNS records")
//     for more information.
//
//   - The operations you specify within the /batch request body are always executed
//     in the following order:
//
//   - Deletes
//
//   - Patches
//
//   - Puts
//
//   - Posts
func (r *RecordService) Batch(ctx context.Context, params RecordBatchParams, opts ...option.RequestOption) (res *RecordBatchResponse, err error) {
	var env RecordBatchResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/batch", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
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
func (r *RecordService) Edit(ctx context.Context, dnsRecordID string, params RecordEditParams, opts ...option.RequestOption) (res *RecordEditResponse, err error) {
	var env RecordEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
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
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/export", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// DNS Record Details
func (r *RecordService) Get(ctx context.Context, dnsRecordID string, query RecordGetParams, opts ...option.RequestOption) (res *RecordGetResponse, err error) {
	var env RecordGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if dnsRecordID == "" {
		err = errors.New("missing required dns_record_id parameter")
		return
	}
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
	var env RecordImportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
func (r *RecordService) Scan(ctx context.Context, params RecordScanParams, opts ...option.RequestOption) (res *RecordScanResponse, err error) {
	var env RecordScanResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records/scan", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ARecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type ARecordType `json:"type"`
	JSON aRecordJSON `json:"-"`
}

// aRecordJSON contains the JSON metadata for the struct [ARecord]
type aRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aRecordJSON) RawJSON() string {
	return r.raw
}

// Record type.
type ARecordType string

const (
	ARecordTypeA ARecordType = "A"
)

func (r ARecordType) IsKnown() bool {
	switch r {
	case ARecordTypeA:
		return true
	}
	return false
}

type ARecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[ARecordType] `json:"type"`
}

func (r ARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ARecordParam) implementsDNSRecordUnionParam() {}

type AAAARecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// A valid IPv6 address.
	Content string `json:"content" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type AAAARecordType `json:"type"`
	JSON aaaaRecordJSON `json:"-"`
}

// aaaaRecordJSON contains the JSON metadata for the struct [AAAARecord]
type aaaaRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaaRecordJSON) RawJSON() string {
	return r.raw
}

// Record type.
type AAAARecordType string

const (
	AAAARecordTypeAAAA AAAARecordType = "AAAA"
)

func (r AAAARecordType) IsKnown() bool {
	switch r {
	case AAAARecordTypeAAAA:
		return true
	}
	return false
}

type AAAARecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[AAAARecordType] `json:"type"`
}

func (r AAAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AAAARecordParam) implementsDNSRecordUnionParam() {}

type CAARecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted CAA content. See 'data' to set CAA properties.
	Content string `json:"content"`
	// Components of a CAA record.
	Data CAARecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type CAARecordType `json:"type"`
	JSON caaRecordJSON `json:"-"`
}

// caaRecordJSON contains the JSON metadata for the struct [CAARecord]
type caaRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caaRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a CAA record.
type CAARecordData struct {
	// Flags for the CAA record.
	Flags float64 `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag string `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value string            `json:"value"`
	JSON  caaRecordDataJSON `json:"-"`
}

// caaRecordDataJSON contains the JSON metadata for the struct [CAARecordData]
type caaRecordDataJSON struct {
	Flags       apijson.Field
	Tag         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CAARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type CAARecordType string

const (
	CAARecordTypeCAA CAARecordType = "CAA"
)

func (r CAARecordType) IsKnown() bool {
	switch r {
	case CAARecordTypeCAA:
		return true
	}
	return false
}

type CAARecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CAA record.
	Data param.Field[CAARecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[CAARecordType] `json:"type"`
}

func (r CAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CAARecordParam) implementsDNSRecordUnionParam() {}

// Components of a CAA record.
type CAARecordDataParam struct {
	// Flags for the CAA record.
	Flags param.Field[float64] `json:"flags"`
	// Name of the property controlled by this record (e.g.: issue, issuewild, iodef).
	Tag param.Field[string] `json:"tag"`
	// Value of the record. This field's semantics depend on the chosen tag.
	Value param.Field[string] `json:"value"`
}

func (r CAARecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CERTRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted CERT content. See 'data' to set CERT properties.
	Content string `json:"content"`
	// Components of a CERT record.
	Data CERTRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type CERTRecordType `json:"type"`
	JSON certRecordJSON `json:"-"`
}

// certRecordJSON contains the JSON metadata for the struct [CERTRecord]
type certRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a CERT record.
type CERTRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Certificate.
	Certificate string `json:"certificate"`
	// Key Tag.
	KeyTag float64 `json:"key_tag"`
	// Type.
	Type float64            `json:"type"`
	JSON certRecordDataJSON `json:"-"`
}

// certRecordDataJSON contains the JSON metadata for the struct [CERTRecordData]
type certRecordDataJSON struct {
	Algorithm   apijson.Field
	Certificate apijson.Field
	KeyTag      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CERTRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type CERTRecordType string

const (
	CERTRecordTypeCERT CERTRecordType = "CERT"
)

func (r CERTRecordType) IsKnown() bool {
	switch r {
	case CERTRecordTypeCERT:
		return true
	}
	return false
}

type CERTRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a CERT record.
	Data param.Field[CERTRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[CERTRecordType] `json:"type"`
}

func (r CERTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CERTRecordParam) implementsDNSRecordUnionParam() {}

// Components of a CERT record.
type CERTRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
	// Type.
	Type param.Field[float64] `json:"type"`
}

func (r CERTRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNAMERecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// A valid hostname. Must not match the record's name.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  bool                `json:"proxied"`
	Settings CNAMERecordSettings `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type CNAMERecordType `json:"type"`
	JSON cnameRecordJSON `json:"-"`
}

// cnameRecordJSON contains the JSON metadata for the struct [CNAMERecord]
type cnameRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Settings    apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnameRecordJSON) RawJSON() string {
	return r.raw
}

type CNAMERecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting has no effect on proxied records, which are always
	// flattened.
	FlattenCNAME bool                    `json:"flatten_cname"`
	JSON         cnameRecordSettingsJSON `json:"-"`
}

// cnameRecordSettingsJSON contains the JSON metadata for the struct
// [CNAMERecordSettings]
type cnameRecordSettingsJSON struct {
	FlattenCNAME apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CNAMERecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cnameRecordSettingsJSON) RawJSON() string {
	return r.raw
}

// Record type.
type CNAMERecordType string

const (
	CNAMERecordTypeCNAME CNAMERecordType = "CNAME"
)

func (r CNAMERecordType) IsKnown() bool {
	switch r {
	case CNAMERecordTypeCNAME:
		return true
	}
	return false
}

type CNAMERecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid hostname. Must not match the record's name.
	Content param.Field[string] `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  param.Field[bool]                     `json:"proxied"`
	Settings param.Field[CNAMERecordSettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[CNAMERecordType] `json:"type"`
}

func (r CNAMERecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CNAMERecordParam) implementsDNSRecordUnionParam() {}

type CNAMERecordSettingsParam struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting has no effect on proxied records, which are always
	// flattened.
	FlattenCNAME param.Field[bool] `json:"flatten_cname"`
}

func (r CNAMERecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSKEYRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted DNSKEY content. See 'data' to set DNSKEY properties.
	Content string `json:"content"`
	// Components of a DNSKEY record.
	Data DNSKEYRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type DNSKEYRecordType `json:"type"`
	JSON dnskeyRecordJSON `json:"-"`
}

// dnskeyRecordJSON contains the JSON metadata for the struct [DNSKEYRecord]
type dnskeyRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnskeyRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a DNSKEY record.
type DNSKEYRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Flags.
	Flags float64 `json:"flags"`
	// Protocol.
	Protocol float64 `json:"protocol"`
	// Public Key.
	PublicKey string               `json:"public_key"`
	JSON      dnskeyRecordDataJSON `json:"-"`
}

// dnskeyRecordDataJSON contains the JSON metadata for the struct
// [DNSKEYRecordData]
type dnskeyRecordDataJSON struct {
	Algorithm   apijson.Field
	Flags       apijson.Field
	Protocol    apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSKEYRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnskeyRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DNSKEYRecordType string

const (
	DNSKEYRecordTypeDNSKEY DNSKEYRecordType = "DNSKEY"
)

func (r DNSKEYRecordType) IsKnown() bool {
	switch r {
	case DNSKEYRecordTypeDNSKEY:
		return true
	}
	return false
}

type DNSKEYRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DNSKEY record.
	Data param.Field[DNSKEYRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[DNSKEYRecordType] `json:"type"`
}

func (r DNSKEYRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DNSKEYRecordParam) implementsDNSRecordUnionParam() {}

// Components of a DNSKEY record.
type DNSKEYRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Flags.
	Flags param.Field[float64] `json:"flags"`
	// Protocol.
	Protocol param.Field[float64] `json:"protocol"`
	// Public Key.
	PublicKey param.Field[string] `json:"public_key"`
}

func (r DNSKEYRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DSRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted DS content. See 'data' to set DS properties.
	Content string `json:"content"`
	// Components of a DS record.
	Data DSRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type DSRecordType `json:"type"`
	JSON dsRecordJSON `json:"-"`
}

// dsRecordJSON contains the JSON metadata for the struct [DSRecord]
type dsRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dsRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a DS record.
type DSRecordData struct {
	// Algorithm.
	Algorithm float64 `json:"algorithm"`
	// Digest.
	Digest string `json:"digest"`
	// Digest Type.
	DigestType float64 `json:"digest_type"`
	// Key Tag.
	KeyTag float64          `json:"key_tag"`
	JSON   dsRecordDataJSON `json:"-"`
}

// dsRecordDataJSON contains the JSON metadata for the struct [DSRecordData]
type dsRecordDataJSON struct {
	Algorithm   apijson.Field
	Digest      apijson.Field
	DigestType  apijson.Field
	KeyTag      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dsRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type DSRecordType string

const (
	DSRecordTypeDS DSRecordType = "DS"
)

func (r DSRecordType) IsKnown() bool {
	switch r {
	case DSRecordTypeDS:
		return true
	}
	return false
}

type DSRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a DS record.
	Data param.Field[DSRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[DSRecordType] `json:"type"`
}

func (r DSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DSRecordParam) implementsDNSRecordUnionParam() {}

// Components of a DS record.
type DSRecordDataParam struct {
	// Algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// Digest.
	Digest param.Field[string] `json:"digest"`
	// Digest Type.
	DigestType param.Field[float64] `json:"digest_type"`
	// Key Tag.
	KeyTag param.Field[float64] `json:"key_tag"`
}

func (r DSRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HTTPSRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted HTTPS content. See 'data' to set HTTPS properties.
	Content string `json:"content"`
	// Components of a HTTPS record.
	Data HTTPSRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type HTTPSRecordType `json:"type"`
	JSON httpsRecordJSON `json:"-"`
}

// httpsRecordJSON contains the JSON metadata for the struct [HTTPSRecord]
type httpsRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpsRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a HTTPS record.
type HTTPSRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string              `json:"value"`
	JSON  httpsRecordDataJSON `json:"-"`
}

// httpsRecordDataJSON contains the JSON metadata for the struct [HTTPSRecordData]
type httpsRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpsRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type HTTPSRecordType string

const (
	HTTPSRecordTypeHTTPS HTTPSRecordType = "HTTPS"
)

func (r HTTPSRecordType) IsKnown() bool {
	switch r {
	case HTTPSRecordTypeHTTPS:
		return true
	}
	return false
}

type HTTPSRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a HTTPS record.
	Data param.Field[HTTPSRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[HTTPSRecordType] `json:"type"`
}

func (r HTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HTTPSRecordParam) implementsDNSRecordUnionParam() {}

// Components of a HTTPS record.
type HTTPSRecordDataParam struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r HTTPSRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LOCRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted LOC content. See 'data' to set LOC properties.
	Content string `json:"content"`
	// Components of a LOC record.
	Data LOCRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type LOCRecordType `json:"type"`
	JSON locRecordJSON `json:"-"`
}

// locRecordJSON contains the JSON metadata for the struct [LOCRecord]
type locRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a LOC record.
type LOCRecordData struct {
	// Altitude of location in meters.
	Altitude float64 `json:"altitude"`
	// Degrees of latitude.
	LatDegrees float64 `json:"lat_degrees"`
	// Latitude direction.
	LatDirection LOCRecordDataLatDirection `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes float64 `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds float64 `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees float64 `json:"long_degrees"`
	// Longitude direction.
	LongDirection LOCRecordDataLongDirection `json:"long_direction"`
	// Minutes of longitude.
	LongMinutes float64 `json:"long_minutes"`
	// Seconds of longitude.
	LongSeconds float64 `json:"long_seconds"`
	// Horizontal precision of location.
	PrecisionHorz float64 `json:"precision_horz"`
	// Vertical precision of location.
	PrecisionVert float64 `json:"precision_vert"`
	// Size of location in meters.
	Size float64           `json:"size"`
	JSON locRecordDataJSON `json:"-"`
}

// locRecordDataJSON contains the JSON metadata for the struct [LOCRecordData]
type locRecordDataJSON struct {
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

func (r *LOCRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locRecordDataJSON) RawJSON() string {
	return r.raw
}

// Latitude direction.
type LOCRecordDataLatDirection string

const (
	LOCRecordDataLatDirectionN LOCRecordDataLatDirection = "N"
	LOCRecordDataLatDirectionS LOCRecordDataLatDirection = "S"
)

func (r LOCRecordDataLatDirection) IsKnown() bool {
	switch r {
	case LOCRecordDataLatDirectionN, LOCRecordDataLatDirectionS:
		return true
	}
	return false
}

// Longitude direction.
type LOCRecordDataLongDirection string

const (
	LOCRecordDataLongDirectionE LOCRecordDataLongDirection = "E"
	LOCRecordDataLongDirectionW LOCRecordDataLongDirection = "W"
)

func (r LOCRecordDataLongDirection) IsKnown() bool {
	switch r {
	case LOCRecordDataLongDirectionE, LOCRecordDataLongDirectionW:
		return true
	}
	return false
}

// Record type.
type LOCRecordType string

const (
	LOCRecordTypeLOC LOCRecordType = "LOC"
)

func (r LOCRecordType) IsKnown() bool {
	switch r {
	case LOCRecordTypeLOC:
		return true
	}
	return false
}

type LOCRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a LOC record.
	Data param.Field[LOCRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[LOCRecordType] `json:"type"`
}

func (r LOCRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LOCRecordParam) implementsDNSRecordUnionParam() {}

// Components of a LOC record.
type LOCRecordDataParam struct {
	// Altitude of location in meters.
	Altitude param.Field[float64] `json:"altitude"`
	// Degrees of latitude.
	LatDegrees param.Field[float64] `json:"lat_degrees"`
	// Latitude direction.
	LatDirection param.Field[LOCRecordDataLatDirection] `json:"lat_direction"`
	// Minutes of latitude.
	LatMinutes param.Field[float64] `json:"lat_minutes"`
	// Seconds of latitude.
	LatSeconds param.Field[float64] `json:"lat_seconds"`
	// Degrees of longitude.
	LongDegrees param.Field[float64] `json:"long_degrees"`
	// Longitude direction.
	LongDirection param.Field[LOCRecordDataLongDirection] `json:"long_direction"`
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

func (r LOCRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MXRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// A valid mail server hostname.
	Content string `json:"content" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type MXRecordType `json:"type"`
	JSON mxRecordJSON `json:"-"`
}

// mxRecordJSON contains the JSON metadata for the struct [MXRecord]
type mxRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mxRecordJSON) RawJSON() string {
	return r.raw
}

// Record type.
type MXRecordType string

const (
	MXRecordTypeMX MXRecordType = "MX"
)

func (r MXRecordType) IsKnown() bool {
	switch r {
	case MXRecordTypeMX:
		return true
	}
	return false
}

type MXRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[MXRecordType] `json:"type"`
}

func (r MXRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MXRecordParam) implementsDNSRecordUnionParam() {}

type NAPTRRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted NAPTR content. See 'data' to set NAPTR properties.
	Content string `json:"content"`
	// Components of a NAPTR record.
	Data NAPTRRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type NAPTRRecordType `json:"type"`
	JSON naptrRecordJSON `json:"-"`
}

// naptrRecordJSON contains the JSON metadata for the struct [NAPTRRecord]
type naptrRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r naptrRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a NAPTR record.
type NAPTRRecordData struct {
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
	Service string              `json:"service"`
	JSON    naptrRecordDataJSON `json:"-"`
}

// naptrRecordDataJSON contains the JSON metadata for the struct [NAPTRRecordData]
type naptrRecordDataJSON struct {
	Flags       apijson.Field
	Order       apijson.Field
	Preference  apijson.Field
	Regex       apijson.Field
	Replacement apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NAPTRRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r naptrRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type NAPTRRecordType string

const (
	NAPTRRecordTypeNAPTR NAPTRRecordType = "NAPTR"
)

func (r NAPTRRecordType) IsKnown() bool {
	switch r {
	case NAPTRRecordTypeNAPTR:
		return true
	}
	return false
}

type NAPTRRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a NAPTR record.
	Data param.Field[NAPTRRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[NAPTRRecordType] `json:"type"`
}

func (r NAPTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NAPTRRecordParam) implementsDNSRecordUnionParam() {}

// Components of a NAPTR record.
type NAPTRRecordDataParam struct {
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

func (r NAPTRRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NSRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// A valid name server host name.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type NSRecordType `json:"type"`
	JSON nsRecordJSON `json:"-"`
}

// nsRecordJSON contains the JSON metadata for the struct [NSRecord]
type nsRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nsRecordJSON) RawJSON() string {
	return r.raw
}

// Record type.
type NSRecordType string

const (
	NSRecordTypeNS NSRecordType = "NS"
)

func (r NSRecordType) IsKnown() bool {
	switch r {
	case NSRecordTypeNS:
		return true
	}
	return false
}

type NSRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[NSRecordType] `json:"type"`
}

func (r NSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NSRecordParam) implementsDNSRecordUnionParam() {}

type PTRRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Domain name pointing to the address.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type PTRRecordType `json:"type"`
	JSON ptrRecordJSON `json:"-"`
}

// ptrRecordJSON contains the JSON metadata for the struct [PTRRecord]
type ptrRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptrRecordJSON) RawJSON() string {
	return r.raw
}

// Record type.
type PTRRecordType string

const (
	PTRRecordTypePTR PTRRecordType = "PTR"
)

func (r PTRRecordType) IsKnown() bool {
	switch r {
	case PTRRecordTypePTR:
		return true
	}
	return false
}

type PTRRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[PTRRecordType] `json:"type"`
}

func (r PTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PTRRecordParam) implementsDNSRecordUnionParam() {}

type RecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A valid IPv4 address.
	Content param.Field[string]      `json:"content" format:"ipv4"`
	Data    param.Field[interface{}] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied  param.Field[bool]        `json:"proxied"`
	Settings param.Field[interface{}] `json:"settings"`
	Tags     param.Field[interface{}] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[RecordType] `json:"type"`
}

func (r RecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordParam) implementsDNSRecordUnionParam() {}

// Satisfied by [dns.ARecordParam], [dns.AAAARecordParam], [dns.CAARecordParam],
// [dns.CERTRecordParam], [dns.CNAMERecordParam], [dns.DNSKEYRecordParam],
// [dns.DSRecordParam], [dns.HTTPSRecordParam], [dns.LOCRecordParam],
// [dns.MXRecordParam], [dns.NAPTRRecordParam], [dns.NSRecordParam],
// [dns.RecordDNSRecordsOpenpgpkeyRecordParam], [dns.PTRRecordParam],
// [dns.SMIMEARecordParam], [dns.SRVRecordParam], [dns.SSHFPRecordParam],
// [dns.SVCBRecordParam], [dns.TLSARecordParam], [dns.TXTRecordParam],
// [dns.URIRecordParam], [RecordParam].
type RecordUnionParam interface {
	implementsDNSRecordUnionParam()
}

type RecordDNSRecordsOpenpgpkeyRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[RecordDNSRecordsOpenpgpkeyRecordType] `json:"type"`
}

func (r RecordDNSRecordsOpenpgpkeyRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordDNSRecordsOpenpgpkeyRecordParam) implementsDNSRecordUnionParam() {}

// Record type.
type RecordDNSRecordsOpenpgpkeyRecordType string

const (
	RecordDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey RecordDNSRecordsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordDNSRecordsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordDNSRecordsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

// Record type.
type RecordType string

const (
	RecordTypeA          RecordType = "A"
	RecordTypeAAAA       RecordType = "AAAA"
	RecordTypeCAA        RecordType = "CAA"
	RecordTypeCERT       RecordType = "CERT"
	RecordTypeCNAME      RecordType = "CNAME"
	RecordTypeDNSKEY     RecordType = "DNSKEY"
	RecordTypeDS         RecordType = "DS"
	RecordTypeHTTPS      RecordType = "HTTPS"
	RecordTypeLOC        RecordType = "LOC"
	RecordTypeMX         RecordType = "MX"
	RecordTypeNAPTR      RecordType = "NAPTR"
	RecordTypeNS         RecordType = "NS"
	RecordTypeOpenpgpkey RecordType = "OPENPGPKEY"
	RecordTypePTR        RecordType = "PTR"
	RecordTypeSMIMEA     RecordType = "SMIMEA"
	RecordTypeSRV        RecordType = "SRV"
	RecordTypeSSHFP      RecordType = "SSHFP"
	RecordTypeSVCB       RecordType = "SVCB"
	RecordTypeTLSA       RecordType = "TLSA"
	RecordTypeTXT        RecordType = "TXT"
	RecordTypeURI        RecordType = "URI"
)

func (r RecordType) IsKnown() bool {
	switch r {
	case RecordTypeA, RecordTypeAAAA, RecordTypeCAA, RecordTypeCERT, RecordTypeCNAME, RecordTypeDNSKEY, RecordTypeDS, RecordTypeHTTPS, RecordTypeLOC, RecordTypeMX, RecordTypeNAPTR, RecordTypeNS, RecordTypeOpenpgpkey, RecordTypePTR, RecordTypeSMIMEA, RecordTypeSRV, RecordTypeSSHFP, RecordTypeSVCB, RecordTypeTLSA, RecordTypeTXT, RecordTypeURI:
		return true
	}
	return false
}

type RecordProcessTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time               `json:"start_time" format:"date-time"`
	JSON      recordProcessTimingJSON `json:"-"`
}

// recordProcessTimingJSON contains the JSON metadata for the struct
// [RecordProcessTiming]
type recordProcessTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordProcessTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordProcessTimingJSON) RawJSON() string {
	return r.raw
}

type RecordTags = string

type RecordTagsParam = string

type SMIMEARecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted SMIMEA content. See 'data' to set SMIMEA properties.
	Content string `json:"content"`
	// Components of a SMIMEA record.
	Data SMIMEARecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type SMIMEARecordType `json:"type"`
	JSON smimeaRecordJSON `json:"-"`
}

// smimeaRecordJSON contains the JSON metadata for the struct [SMIMEARecord]
type smimeaRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smimeaRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a SMIMEA record.
type SMIMEARecordData struct {
	// Certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64              `json:"usage"`
	JSON  smimeaRecordDataJSON `json:"-"`
}

// smimeaRecordDataJSON contains the JSON metadata for the struct
// [SMIMEARecordData]
type smimeaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SMIMEARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smimeaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SMIMEARecordType string

const (
	SMIMEARecordTypeSMIMEA SMIMEARecordType = "SMIMEA"
)

func (r SMIMEARecordType) IsKnown() bool {
	switch r {
	case SMIMEARecordTypeSMIMEA:
		return true
	}
	return false
}

type SMIMEARecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SMIMEA record.
	Data param.Field[SMIMEARecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[SMIMEARecordType] `json:"type"`
}

func (r SMIMEARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SMIMEARecordParam) implementsDNSRecordUnionParam() {}

// Components of a SMIMEA record.
type SMIMEARecordDataParam struct {
	// Certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r SMIMEARecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SRVRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Priority, weight, port, and SRV target. See 'data' for setting the individual
	// component values.
	Content string `json:"content"`
	// Components of a SRV record.
	Data SRVRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type SRVRecordType `json:"type"`
	JSON srvRecordJSON `json:"-"`
}

// srvRecordJSON contains the JSON metadata for the struct [SRVRecord]
type srvRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r srvRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a SRV record.
type SRVRecordData struct {
	// The port of the service.
	Port float64 `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// A valid hostname.
	Target string `json:"target" format:"hostname"`
	// The record weight.
	Weight float64           `json:"weight"`
	JSON   srvRecordDataJSON `json:"-"`
}

// srvRecordDataJSON contains the JSON metadata for the struct [SRVRecordData]
type srvRecordDataJSON struct {
	Port        apijson.Field
	Priority    apijson.Field
	Target      apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SRVRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r srvRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SRVRecordType string

const (
	SRVRecordTypeSRV SRVRecordType = "SRV"
)

func (r SRVRecordType) IsKnown() bool {
	switch r {
	case SRVRecordTypeSRV:
		return true
	}
	return false
}

type SRVRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SRV record.
	Data param.Field[SRVRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[SRVRecordType] `json:"type"`
}

func (r SRVRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SRVRecordParam) implementsDNSRecordUnionParam() {}

// Components of a SRV record.
type SRVRecordDataParam struct {
	// The port of the service.
	Port param.Field[float64] `json:"port"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// A valid hostname.
	Target param.Field[string] `json:"target" format:"hostname"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r SRVRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SSHFPRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted SSHFP content. See 'data' to set SSHFP properties.
	Content string `json:"content"`
	// Components of a SSHFP record.
	Data SSHFPRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type SSHFPRecordType `json:"type"`
	JSON sshfpRecordJSON `json:"-"`
}

// sshfpRecordJSON contains the JSON metadata for the struct [SSHFPRecord]
type sshfpRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sshfpRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a SSHFP record.
type SSHFPRecordData struct {
	// algorithm.
	Algorithm float64 `json:"algorithm"`
	// fingerprint.
	Fingerprint string `json:"fingerprint"`
	// type.
	Type float64             `json:"type"`
	JSON sshfpRecordDataJSON `json:"-"`
}

// sshfpRecordDataJSON contains the JSON metadata for the struct [SSHFPRecordData]
type sshfpRecordDataJSON struct {
	Algorithm   apijson.Field
	Fingerprint apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSHFPRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sshfpRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SSHFPRecordType string

const (
	SSHFPRecordTypeSSHFP SSHFPRecordType = "SSHFP"
)

func (r SSHFPRecordType) IsKnown() bool {
	switch r {
	case SSHFPRecordTypeSSHFP:
		return true
	}
	return false
}

type SSHFPRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SSHFP record.
	Data param.Field[SSHFPRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[SSHFPRecordType] `json:"type"`
}

func (r SSHFPRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SSHFPRecordParam) implementsDNSRecordUnionParam() {}

// Components of a SSHFP record.
type SSHFPRecordDataParam struct {
	// algorithm.
	Algorithm param.Field[float64] `json:"algorithm"`
	// fingerprint.
	Fingerprint param.Field[string] `json:"fingerprint"`
	// type.
	Type param.Field[float64] `json:"type"`
}

func (r SSHFPRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SVCBRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted SVCB content. See 'data' to set SVCB properties.
	Content string `json:"content"`
	// Components of a SVCB record.
	Data SVCBRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type SVCBRecordType `json:"type"`
	JSON svcbRecordJSON `json:"-"`
}

// svcbRecordJSON contains the JSON metadata for the struct [SVCBRecord]
type svcbRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r svcbRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a SVCB record.
type SVCBRecordData struct {
	// priority.
	Priority float64 `json:"priority"`
	// target.
	Target string `json:"target"`
	// value.
	Value string             `json:"value"`
	JSON  svcbRecordDataJSON `json:"-"`
}

// svcbRecordDataJSON contains the JSON metadata for the struct [SVCBRecordData]
type svcbRecordDataJSON struct {
	Priority    apijson.Field
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SVCBRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r svcbRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type SVCBRecordType string

const (
	SVCBRecordTypeSVCB SVCBRecordType = "SVCB"
)

func (r SVCBRecordType) IsKnown() bool {
	switch r {
	case SVCBRecordTypeSVCB:
		return true
	}
	return false
}

type SVCBRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a SVCB record.
	Data param.Field[SVCBRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[SVCBRecordType] `json:"type"`
}

func (r SVCBRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SVCBRecordParam) implementsDNSRecordUnionParam() {}

// Components of a SVCB record.
type SVCBRecordDataParam struct {
	// priority.
	Priority param.Field[float64] `json:"priority"`
	// target.
	Target param.Field[string] `json:"target"`
	// value.
	Value param.Field[string] `json:"value"`
}

func (r SVCBRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TLSARecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted TLSA content. See 'data' to set TLSA properties.
	Content string `json:"content"`
	// Components of a TLSA record.
	Data TLSARecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type TLSARecordType `json:"type"`
	JSON tlsaRecordJSON `json:"-"`
}

// tlsaRecordJSON contains the JSON metadata for the struct [TLSARecord]
type tlsaRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsaRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a TLSA record.
type TLSARecordData struct {
	// certificate.
	Certificate string `json:"certificate"`
	// Matching Type.
	MatchingType float64 `json:"matching_type"`
	// Selector.
	Selector float64 `json:"selector"`
	// Usage.
	Usage float64            `json:"usage"`
	JSON  tlsaRecordDataJSON `json:"-"`
}

// tlsaRecordDataJSON contains the JSON metadata for the struct [TLSARecordData]
type tlsaRecordDataJSON struct {
	Certificate  apijson.Field
	MatchingType apijson.Field
	Selector     apijson.Field
	Usage        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TLSARecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsaRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type TLSARecordType string

const (
	TLSARecordTypeTLSA TLSARecordType = "TLSA"
)

func (r TLSARecordType) IsKnown() bool {
	switch r {
	case TLSARecordTypeTLSA:
		return true
	}
	return false
}

type TLSARecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a TLSA record.
	Data param.Field[TLSARecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[TLSARecordType] `json:"type"`
}

func (r TLSARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TLSARecordParam) implementsDNSRecordUnionParam() {}

// Components of a TLSA record.
type TLSARecordDataParam struct {
	// certificate.
	Certificate param.Field[string] `json:"certificate"`
	// Matching Type.
	MatchingType param.Field[float64] `json:"matching_type"`
	// Selector.
	Selector param.Field[float64] `json:"selector"`
	// Usage.
	Usage param.Field[float64] `json:"usage"`
}

func (r TLSARecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
// Value must be between 60 and 86400, with the minimum reduced to 30 for
// Enterprise zones.
type TTL float64

const (
	TTL1 TTL = 1
)

func (r TTL) IsKnown() bool {
	switch r {
	case TTL1:
		return true
	}
	return false
}

type TXTRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content string `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type TXTRecordType `json:"type"`
	JSON txtRecordJSON `json:"-"`
}

// txtRecordJSON contains the JSON metadata for the struct [TXTRecord]
type txtRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r txtRecordJSON) RawJSON() string {
	return r.raw
}

// Record type.
type TXTRecordType string

const (
	TXTRecordTypeTXT TXTRecordType = "TXT"
)

func (r TXTRecordType) IsKnown() bool {
	switch r {
	case TXTRecordTypeTXT:
		return true
	}
	return false
}

type TXTRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[TXTRecordType] `json:"type"`
}

func (r TXTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TXTRecordParam) implementsDNSRecordUnionParam() {}

type URIRecord struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// Formatted URI content. See 'data' to set URI properties.
	Content string `json:"content"`
	// Components of a URI record.
	Data URIRecordData `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type URIRecordType `json:"type"`
	JSON uriRecordJSON `json:"-"`
}

// uriRecordJSON contains the JSON metadata for the struct [URIRecord]
type uriRecordJSON struct {
	Comment     apijson.Field
	Content     apijson.Field
	Data        apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Proxied     apijson.Field
	Tags        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uriRecordJSON) RawJSON() string {
	return r.raw
}

// Components of a URI record.
type URIRecordData struct {
	// The record content.
	Target string `json:"target"`
	// The record weight.
	Weight float64           `json:"weight"`
	JSON   uriRecordDataJSON `json:"-"`
}

// uriRecordDataJSON contains the JSON metadata for the struct [URIRecordData]
type uriRecordDataJSON struct {
	Target      apijson.Field
	Weight      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URIRecordData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uriRecordDataJSON) RawJSON() string {
	return r.raw
}

// Record type.
type URIRecordType string

const (
	URIRecordTypeURI URIRecordType = "URI"
)

func (r URIRecordType) IsKnown() bool {
	switch r {
	case URIRecordTypeURI:
		return true
	}
	return false
}

type URIRecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Components of a URI record.
	Data param.Field[URIRecordDataParam] `json:"data"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[URIRecordType] `json:"type"`
}

func (r URIRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r URIRecordParam) implementsDNSRecordUnionParam() {}

// Components of a URI record.
type URIRecordDataParam struct {
	// The record content.
	Target param.Field[string] `json:"target"`
	// The record weight.
	Weight param.Field[float64] `json:"weight"`
}

func (r URIRecordDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordNewResponseType `json:"type"`
	JSON  recordNewResponseJSON `json:"-"`
	union RecordNewResponseUnion
}

// recordNewResponseJSON contains the JSON metadata for the struct
// [RecordNewResponse]
type recordNewResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RecordNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RecordNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordNewResponseARecord],
// [dns.RecordNewResponseAAAARecord], [dns.RecordNewResponseCAARecord],
// [dns.RecordNewResponseCERTRecord], [dns.RecordNewResponseCNAMERecord],
// [dns.RecordNewResponseDNSKEYRecord], [dns.RecordNewResponseDSRecord],
// [dns.RecordNewResponseHTTPSRecord], [dns.RecordNewResponseLOCRecord],
// [dns.RecordNewResponseMXRecord], [dns.RecordNewResponseNAPTRRecord],
// [dns.RecordNewResponseNSRecord], [dns.RecordNewResponseOpenpgpkeyRecord],
// [dns.RecordNewResponsePTRRecord], [dns.RecordNewResponseSMIMEARecord],
// [dns.RecordNewResponseSRVRecord], [dns.RecordNewResponseSSHFPRecord],
// [dns.RecordNewResponseSVCBRecord], [dns.RecordNewResponseTLSARecord],
// [dns.RecordNewResponseTXTRecord], [dns.RecordNewResponseURIRecord].
func (r RecordNewResponse) AsUnion() RecordNewResponseUnion {
	return r.union
}

// Union satisfied by [dns.RecordNewResponseARecord],
// [dns.RecordNewResponseAAAARecord], [dns.RecordNewResponseCAARecord],
// [dns.RecordNewResponseCERTRecord], [dns.RecordNewResponseCNAMERecord],
// [dns.RecordNewResponseDNSKEYRecord], [dns.RecordNewResponseDSRecord],
// [dns.RecordNewResponseHTTPSRecord], [dns.RecordNewResponseLOCRecord],
// [dns.RecordNewResponseMXRecord], [dns.RecordNewResponseNAPTRRecord],
// [dns.RecordNewResponseNSRecord], [dns.RecordNewResponseOpenpgpkeyRecord],
// [dns.RecordNewResponsePTRRecord], [dns.RecordNewResponseSMIMEARecord],
// [dns.RecordNewResponseSRVRecord], [dns.RecordNewResponseSSHFPRecord],
// [dns.RecordNewResponseSVCBRecord], [dns.RecordNewResponseTLSARecord],
// [dns.RecordNewResponseTXTRecord] or [dns.RecordNewResponseURIRecord].
type RecordNewResponseUnion interface {
	implementsDNSRecordNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponsePTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordNewResponseURIRecord{}),
		},
	)
}

type RecordNewResponseARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                    `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseARecordJSON `json:"-"`
	ARecord
}

// recordNewResponseARecordJSON contains the JSON metadata for the struct
// [RecordNewResponseARecord]
type recordNewResponseARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseARecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordNewResponseAAAARecordJSON contains the JSON metadata for the struct
// [RecordNewResponseAAAARecord]
type recordNewResponseAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseAAAARecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseCAARecordJSON `json:"-"`
	CAARecord
}

// recordNewResponseCAARecordJSON contains the JSON metadata for the struct
// [RecordNewResponseCAARecord]
type recordNewResponseCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseCAARecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordNewResponseCERTRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseCERTRecord]
type recordNewResponseCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseCERTRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordNewResponseCNAMERecordJSON contains the JSON metadata for the struct
// [RecordNewResponseCNAMERecord]
type recordNewResponseCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseCNAMERecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordNewResponseDNSKEYRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseDNSKEYRecord]
type recordNewResponseDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseDNSKEYRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseDSRecordJSON `json:"-"`
	DSRecord
}

// recordNewResponseDSRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseDSRecord]
type recordNewResponseDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseDSRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordNewResponseHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseHTTPSRecord]
type recordNewResponseHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseHTTPSRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordNewResponseLOCRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseLOCRecord]
type recordNewResponseLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseLOCRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseMXRecordJSON `json:"-"`
	MXRecord
}

// recordNewResponseMXRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseMXRecord]
type recordNewResponseMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseMXRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordNewResponseNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseNAPTRRecord]
type recordNewResponseNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseNAPTRRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseNSRecordJSON `json:"-"`
	NSRecord
}

// recordNewResponseNSRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseNSRecord]
type recordNewResponseNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseNSRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordNewResponseOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseOpenpgpkeyRecordJSON `json:"-"`
}

// recordNewResponseOpenpgpkeyRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseOpenpgpkeyRecord]
type recordNewResponseOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseOpenpgpkeyRecord) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseOpenpgpkeyRecordType string

const (
	RecordNewResponseOpenpgpkeyRecordTypeOpenpgpkey RecordNewResponseOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordNewResponseOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordNewResponseOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordNewResponsePTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponsePTRRecordJSON `json:"-"`
	PTRRecord
}

// recordNewResponsePTRRecordJSON contains the JSON metadata for the struct
// [RecordNewResponsePTRRecord]
type recordNewResponsePTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponsePTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponsePTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponsePTRRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordNewResponseSMIMEARecordJSON contains the JSON metadata for the struct
// [RecordNewResponseSMIMEARecord]
type recordNewResponseSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSMIMEARecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordNewResponseSRVRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseSRVRecord]
type recordNewResponseSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSRVRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordNewResponseSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseSSHFPRecord]
type recordNewResponseSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSSHFPRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordNewResponseSVCBRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseSVCBRecord]
type recordNewResponseSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseSVCBRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordNewResponseTLSARecordJSON contains the JSON metadata for the struct
// [RecordNewResponseTLSARecord]
type recordNewResponseTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseTLSARecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordNewResponseTXTRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseTXTRecord]
type recordNewResponseTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseTXTRecord) implementsDNSRecordNewResponse() {}

type RecordNewResponseURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseURIRecordJSON `json:"-"`
	URIRecord
}

// recordNewResponseURIRecordJSON contains the JSON metadata for the struct
// [RecordNewResponseURIRecord]
type recordNewResponseURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponseURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseURIRecord) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseType string

const (
	RecordNewResponseTypeA          RecordNewResponseType = "A"
	RecordNewResponseTypeAAAA       RecordNewResponseType = "AAAA"
	RecordNewResponseTypeCAA        RecordNewResponseType = "CAA"
	RecordNewResponseTypeCERT       RecordNewResponseType = "CERT"
	RecordNewResponseTypeCNAME      RecordNewResponseType = "CNAME"
	RecordNewResponseTypeDNSKEY     RecordNewResponseType = "DNSKEY"
	RecordNewResponseTypeDS         RecordNewResponseType = "DS"
	RecordNewResponseTypeHTTPS      RecordNewResponseType = "HTTPS"
	RecordNewResponseTypeLOC        RecordNewResponseType = "LOC"
	RecordNewResponseTypeMX         RecordNewResponseType = "MX"
	RecordNewResponseTypeNAPTR      RecordNewResponseType = "NAPTR"
	RecordNewResponseTypeNS         RecordNewResponseType = "NS"
	RecordNewResponseTypeOpenpgpkey RecordNewResponseType = "OPENPGPKEY"
	RecordNewResponseTypePTR        RecordNewResponseType = "PTR"
	RecordNewResponseTypeSMIMEA     RecordNewResponseType = "SMIMEA"
	RecordNewResponseTypeSRV        RecordNewResponseType = "SRV"
	RecordNewResponseTypeSSHFP      RecordNewResponseType = "SSHFP"
	RecordNewResponseTypeSVCB       RecordNewResponseType = "SVCB"
	RecordNewResponseTypeTLSA       RecordNewResponseType = "TLSA"
	RecordNewResponseTypeTXT        RecordNewResponseType = "TXT"
	RecordNewResponseTypeURI        RecordNewResponseType = "URI"
)

func (r RecordNewResponseType) IsKnown() bool {
	switch r {
	case RecordNewResponseTypeA, RecordNewResponseTypeAAAA, RecordNewResponseTypeCAA, RecordNewResponseTypeCERT, RecordNewResponseTypeCNAME, RecordNewResponseTypeDNSKEY, RecordNewResponseTypeDS, RecordNewResponseTypeHTTPS, RecordNewResponseTypeLOC, RecordNewResponseTypeMX, RecordNewResponseTypeNAPTR, RecordNewResponseTypeNS, RecordNewResponseTypeOpenpgpkey, RecordNewResponseTypePTR, RecordNewResponseTypeSMIMEA, RecordNewResponseTypeSRV, RecordNewResponseTypeSSHFP, RecordNewResponseTypeSVCB, RecordNewResponseTypeTLSA, RecordNewResponseTypeTXT, RecordNewResponseTypeURI:
		return true
	}
	return false
}

type RecordUpdateResponse struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordUpdateResponseType `json:"type"`
	JSON  recordUpdateResponseJSON `json:"-"`
	union RecordUpdateResponseUnion
}

// recordUpdateResponseJSON contains the JSON metadata for the struct
// [RecordUpdateResponse]
type recordUpdateResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RecordUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RecordUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordUpdateResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordUpdateResponseARecord],
// [dns.RecordUpdateResponseAAAARecord], [dns.RecordUpdateResponseCAARecord],
// [dns.RecordUpdateResponseCERTRecord], [dns.RecordUpdateResponseCNAMERecord],
// [dns.RecordUpdateResponseDNSKEYRecord], [dns.RecordUpdateResponseDSRecord],
// [dns.RecordUpdateResponseHTTPSRecord], [dns.RecordUpdateResponseLOCRecord],
// [dns.RecordUpdateResponseMXRecord], [dns.RecordUpdateResponseNAPTRRecord],
// [dns.RecordUpdateResponseNSRecord], [dns.RecordUpdateResponseOpenpgpkeyRecord],
// [dns.RecordUpdateResponsePTRRecord], [dns.RecordUpdateResponseSMIMEARecord],
// [dns.RecordUpdateResponseSRVRecord], [dns.RecordUpdateResponseSSHFPRecord],
// [dns.RecordUpdateResponseSVCBRecord], [dns.RecordUpdateResponseTLSARecord],
// [dns.RecordUpdateResponseTXTRecord], [dns.RecordUpdateResponseURIRecord].
func (r RecordUpdateResponse) AsUnion() RecordUpdateResponseUnion {
	return r.union
}

// Union satisfied by [dns.RecordUpdateResponseARecord],
// [dns.RecordUpdateResponseAAAARecord], [dns.RecordUpdateResponseCAARecord],
// [dns.RecordUpdateResponseCERTRecord], [dns.RecordUpdateResponseCNAMERecord],
// [dns.RecordUpdateResponseDNSKEYRecord], [dns.RecordUpdateResponseDSRecord],
// [dns.RecordUpdateResponseHTTPSRecord], [dns.RecordUpdateResponseLOCRecord],
// [dns.RecordUpdateResponseMXRecord], [dns.RecordUpdateResponseNAPTRRecord],
// [dns.RecordUpdateResponseNSRecord], [dns.RecordUpdateResponseOpenpgpkeyRecord],
// [dns.RecordUpdateResponsePTRRecord], [dns.RecordUpdateResponseSMIMEARecord],
// [dns.RecordUpdateResponseSRVRecord], [dns.RecordUpdateResponseSSHFPRecord],
// [dns.RecordUpdateResponseSVCBRecord], [dns.RecordUpdateResponseTLSARecord],
// [dns.RecordUpdateResponseTXTRecord] or [dns.RecordUpdateResponseURIRecord].
type RecordUpdateResponseUnion interface {
	implementsDNSRecordUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponsePTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordUpdateResponseURIRecord{}),
		},
	)
}

type RecordUpdateResponseARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseARecordJSON `json:"-"`
	ARecord
}

// recordUpdateResponseARecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseARecord]
type recordUpdateResponseARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseARecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordUpdateResponseAAAARecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseAAAARecord]
type recordUpdateResponseAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseAAAARecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseCAARecordJSON `json:"-"`
	CAARecord
}

// recordUpdateResponseCAARecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCAARecord]
type recordUpdateResponseCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseCAARecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordUpdateResponseCERTRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCERTRecord]
type recordUpdateResponseCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseCERTRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordUpdateResponseCNAMERecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseCNAMERecord]
type recordUpdateResponseCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseCNAMERecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordUpdateResponseDNSKEYRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDNSKEYRecord]
type recordUpdateResponseDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseDNSKEYRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseDSRecordJSON `json:"-"`
	DSRecord
}

// recordUpdateResponseDSRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseDSRecord]
type recordUpdateResponseDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseDSRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordUpdateResponseHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseHTTPSRecord]
type recordUpdateResponseHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseHTTPSRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordUpdateResponseLOCRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseLOCRecord]
type recordUpdateResponseLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseLOCRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseMXRecordJSON `json:"-"`
	MXRecord
}

// recordUpdateResponseMXRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseMXRecord]
type recordUpdateResponseMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseMXRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordUpdateResponseNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNAPTRRecord]
type recordUpdateResponseNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseNAPTRRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseNSRecordJSON `json:"-"`
	NSRecord
}

// recordUpdateResponseNSRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseNSRecord]
type recordUpdateResponseNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseNSRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordUpdateResponseOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseOpenpgpkeyRecordJSON `json:"-"`
}

// recordUpdateResponseOpenpgpkeyRecordJSON contains the JSON metadata for the
// struct [RecordUpdateResponseOpenpgpkeyRecord]
type recordUpdateResponseOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseOpenpgpkeyRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseOpenpgpkeyRecordType string

const (
	RecordUpdateResponseOpenpgpkeyRecordTypeOpenpgpkey RecordUpdateResponseOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordUpdateResponseOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordUpdateResponseOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordUpdateResponsePTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponsePTRRecordJSON `json:"-"`
	PTRRecord
}

// recordUpdateResponsePTRRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponsePTRRecord]
type recordUpdateResponsePTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponsePTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponsePTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponsePTRRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordUpdateResponseSMIMEARecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSMIMEARecord]
type recordUpdateResponseSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSMIMEARecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordUpdateResponseSRVRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSRVRecord]
type recordUpdateResponseSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSRVRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordUpdateResponseSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSSHFPRecord]
type recordUpdateResponseSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSSHFPRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordUpdateResponseSVCBRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseSVCBRecord]
type recordUpdateResponseSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseSVCBRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordUpdateResponseTLSARecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTLSARecord]
type recordUpdateResponseTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseTLSARecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordUpdateResponseTXTRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseTXTRecord]
type recordUpdateResponseTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseTXTRecord) implementsDNSRecordUpdateResponse() {}

type RecordUpdateResponseURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseURIRecordJSON `json:"-"`
	URIRecord
}

// recordUpdateResponseURIRecordJSON contains the JSON metadata for the struct
// [RecordUpdateResponseURIRecord]
type recordUpdateResponseURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponseURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseURIRecord) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseType string

const (
	RecordUpdateResponseTypeA          RecordUpdateResponseType = "A"
	RecordUpdateResponseTypeAAAA       RecordUpdateResponseType = "AAAA"
	RecordUpdateResponseTypeCAA        RecordUpdateResponseType = "CAA"
	RecordUpdateResponseTypeCERT       RecordUpdateResponseType = "CERT"
	RecordUpdateResponseTypeCNAME      RecordUpdateResponseType = "CNAME"
	RecordUpdateResponseTypeDNSKEY     RecordUpdateResponseType = "DNSKEY"
	RecordUpdateResponseTypeDS         RecordUpdateResponseType = "DS"
	RecordUpdateResponseTypeHTTPS      RecordUpdateResponseType = "HTTPS"
	RecordUpdateResponseTypeLOC        RecordUpdateResponseType = "LOC"
	RecordUpdateResponseTypeMX         RecordUpdateResponseType = "MX"
	RecordUpdateResponseTypeNAPTR      RecordUpdateResponseType = "NAPTR"
	RecordUpdateResponseTypeNS         RecordUpdateResponseType = "NS"
	RecordUpdateResponseTypeOpenpgpkey RecordUpdateResponseType = "OPENPGPKEY"
	RecordUpdateResponseTypePTR        RecordUpdateResponseType = "PTR"
	RecordUpdateResponseTypeSMIMEA     RecordUpdateResponseType = "SMIMEA"
	RecordUpdateResponseTypeSRV        RecordUpdateResponseType = "SRV"
	RecordUpdateResponseTypeSSHFP      RecordUpdateResponseType = "SSHFP"
	RecordUpdateResponseTypeSVCB       RecordUpdateResponseType = "SVCB"
	RecordUpdateResponseTypeTLSA       RecordUpdateResponseType = "TLSA"
	RecordUpdateResponseTypeTXT        RecordUpdateResponseType = "TXT"
	RecordUpdateResponseTypeURI        RecordUpdateResponseType = "URI"
)

func (r RecordUpdateResponseType) IsKnown() bool {
	switch r {
	case RecordUpdateResponseTypeA, RecordUpdateResponseTypeAAAA, RecordUpdateResponseTypeCAA, RecordUpdateResponseTypeCERT, RecordUpdateResponseTypeCNAME, RecordUpdateResponseTypeDNSKEY, RecordUpdateResponseTypeDS, RecordUpdateResponseTypeHTTPS, RecordUpdateResponseTypeLOC, RecordUpdateResponseTypeMX, RecordUpdateResponseTypeNAPTR, RecordUpdateResponseTypeNS, RecordUpdateResponseTypeOpenpgpkey, RecordUpdateResponseTypePTR, RecordUpdateResponseTypeSMIMEA, RecordUpdateResponseTypeSRV, RecordUpdateResponseTypeSSHFP, RecordUpdateResponseTypeSVCB, RecordUpdateResponseTypeTLSA, RecordUpdateResponseTypeTXT, RecordUpdateResponseTypeURI:
		return true
	}
	return false
}

type RecordListResponse struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordListResponseType `json:"type"`
	JSON  recordListResponseJSON `json:"-"`
	union RecordListResponseUnion
}

// recordListResponseJSON contains the JSON metadata for the struct
// [RecordListResponse]
type recordListResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RecordListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RecordListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordListResponseARecord],
// [dns.RecordListResponseAAAARecord], [dns.RecordListResponseCAARecord],
// [dns.RecordListResponseCERTRecord], [dns.RecordListResponseCNAMERecord],
// [dns.RecordListResponseDNSKEYRecord], [dns.RecordListResponseDSRecord],
// [dns.RecordListResponseHTTPSRecord], [dns.RecordListResponseLOCRecord],
// [dns.RecordListResponseMXRecord], [dns.RecordListResponseNAPTRRecord],
// [dns.RecordListResponseNSRecord], [dns.RecordListResponseOpenpgpkeyRecord],
// [dns.RecordListResponsePTRRecord], [dns.RecordListResponseSMIMEARecord],
// [dns.RecordListResponseSRVRecord], [dns.RecordListResponseSSHFPRecord],
// [dns.RecordListResponseSVCBRecord], [dns.RecordListResponseTLSARecord],
// [dns.RecordListResponseTXTRecord], [dns.RecordListResponseURIRecord].
func (r RecordListResponse) AsUnion() RecordListResponseUnion {
	return r.union
}

// Union satisfied by [dns.RecordListResponseARecord],
// [dns.RecordListResponseAAAARecord], [dns.RecordListResponseCAARecord],
// [dns.RecordListResponseCERTRecord], [dns.RecordListResponseCNAMERecord],
// [dns.RecordListResponseDNSKEYRecord], [dns.RecordListResponseDSRecord],
// [dns.RecordListResponseHTTPSRecord], [dns.RecordListResponseLOCRecord],
// [dns.RecordListResponseMXRecord], [dns.RecordListResponseNAPTRRecord],
// [dns.RecordListResponseNSRecord], [dns.RecordListResponseOpenpgpkeyRecord],
// [dns.RecordListResponsePTRRecord], [dns.RecordListResponseSMIMEARecord],
// [dns.RecordListResponseSRVRecord], [dns.RecordListResponseSSHFPRecord],
// [dns.RecordListResponseSVCBRecord], [dns.RecordListResponseTLSARecord],
// [dns.RecordListResponseTXTRecord] or [dns.RecordListResponseURIRecord].
type RecordListResponseUnion interface {
	implementsDNSRecordListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponsePTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordListResponseURIRecord{}),
		},
	)
}

type RecordListResponseARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseARecordJSON `json:"-"`
	ARecord
}

// recordListResponseARecordJSON contains the JSON metadata for the struct
// [RecordListResponseARecord]
type recordListResponseARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseARecord) implementsDNSRecordListResponse() {}

type RecordListResponseAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordListResponseAAAARecordJSON contains the JSON metadata for the struct
// [RecordListResponseAAAARecord]
type recordListResponseAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseAAAARecord) implementsDNSRecordListResponse() {}

type RecordListResponseCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseCAARecordJSON `json:"-"`
	CAARecord
}

// recordListResponseCAARecordJSON contains the JSON metadata for the struct
// [RecordListResponseCAARecord]
type recordListResponseCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseCAARecord) implementsDNSRecordListResponse() {}

type RecordListResponseCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordListResponseCERTRecordJSON contains the JSON metadata for the struct
// [RecordListResponseCERTRecord]
type recordListResponseCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseCERTRecord) implementsDNSRecordListResponse() {}

type RecordListResponseCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordListResponseCNAMERecordJSON contains the JSON metadata for the struct
// [RecordListResponseCNAMERecord]
type recordListResponseCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseCNAMERecord) implementsDNSRecordListResponse() {}

type RecordListResponseDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordListResponseDNSKEYRecordJSON contains the JSON metadata for the struct
// [RecordListResponseDNSKEYRecord]
type recordListResponseDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseDNSKEYRecord) implementsDNSRecordListResponse() {}

type RecordListResponseDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseDSRecordJSON `json:"-"`
	DSRecord
}

// recordListResponseDSRecordJSON contains the JSON metadata for the struct
// [RecordListResponseDSRecord]
type recordListResponseDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseDSRecord) implementsDNSRecordListResponse() {}

type RecordListResponseHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordListResponseHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordListResponseHTTPSRecord]
type recordListResponseHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseHTTPSRecord) implementsDNSRecordListResponse() {}

type RecordListResponseLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordListResponseLOCRecordJSON contains the JSON metadata for the struct
// [RecordListResponseLOCRecord]
type recordListResponseLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseLOCRecord) implementsDNSRecordListResponse() {}

type RecordListResponseMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseMXRecordJSON `json:"-"`
	MXRecord
}

// recordListResponseMXRecordJSON contains the JSON metadata for the struct
// [RecordListResponseMXRecord]
type recordListResponseMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseMXRecord) implementsDNSRecordListResponse() {}

type RecordListResponseNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordListResponseNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordListResponseNAPTRRecord]
type recordListResponseNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseNAPTRRecord) implementsDNSRecordListResponse() {}

type RecordListResponseNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseNSRecordJSON `json:"-"`
	NSRecord
}

// recordListResponseNSRecordJSON contains the JSON metadata for the struct
// [RecordListResponseNSRecord]
type recordListResponseNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseNSRecord) implementsDNSRecordListResponse() {}

type RecordListResponseOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordListResponseOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseOpenpgpkeyRecordJSON `json:"-"`
}

// recordListResponseOpenpgpkeyRecordJSON contains the JSON metadata for the struct
// [RecordListResponseOpenpgpkeyRecord]
type recordListResponseOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseOpenpgpkeyRecord) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseOpenpgpkeyRecordType string

const (
	RecordListResponseOpenpgpkeyRecordTypeOpenpgpkey RecordListResponseOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordListResponseOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordListResponseOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordListResponsePTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponsePTRRecordJSON `json:"-"`
	PTRRecord
}

// recordListResponsePTRRecordJSON contains the JSON metadata for the struct
// [RecordListResponsePTRRecord]
type recordListResponsePTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponsePTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponsePTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponsePTRRecord) implementsDNSRecordListResponse() {}

type RecordListResponseSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordListResponseSMIMEARecordJSON contains the JSON metadata for the struct
// [RecordListResponseSMIMEARecord]
type recordListResponseSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSMIMEARecord) implementsDNSRecordListResponse() {}

type RecordListResponseSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordListResponseSRVRecordJSON contains the JSON metadata for the struct
// [RecordListResponseSRVRecord]
type recordListResponseSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSRVRecord) implementsDNSRecordListResponse() {}

type RecordListResponseSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordListResponseSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordListResponseSSHFPRecord]
type recordListResponseSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSSHFPRecord) implementsDNSRecordListResponse() {}

type RecordListResponseSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordListResponseSVCBRecordJSON contains the JSON metadata for the struct
// [RecordListResponseSVCBRecord]
type recordListResponseSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseSVCBRecord) implementsDNSRecordListResponse() {}

type RecordListResponseTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordListResponseTLSARecordJSON contains the JSON metadata for the struct
// [RecordListResponseTLSARecord]
type recordListResponseTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseTLSARecord) implementsDNSRecordListResponse() {}

type RecordListResponseTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordListResponseTXTRecordJSON contains the JSON metadata for the struct
// [RecordListResponseTXTRecord]
type recordListResponseTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseTXTRecord) implementsDNSRecordListResponse() {}

type RecordListResponseURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseURIRecordJSON `json:"-"`
	URIRecord
}

// recordListResponseURIRecordJSON contains the JSON metadata for the struct
// [RecordListResponseURIRecord]
type recordListResponseURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponseURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseURIRecord) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseType string

const (
	RecordListResponseTypeA          RecordListResponseType = "A"
	RecordListResponseTypeAAAA       RecordListResponseType = "AAAA"
	RecordListResponseTypeCAA        RecordListResponseType = "CAA"
	RecordListResponseTypeCERT       RecordListResponseType = "CERT"
	RecordListResponseTypeCNAME      RecordListResponseType = "CNAME"
	RecordListResponseTypeDNSKEY     RecordListResponseType = "DNSKEY"
	RecordListResponseTypeDS         RecordListResponseType = "DS"
	RecordListResponseTypeHTTPS      RecordListResponseType = "HTTPS"
	RecordListResponseTypeLOC        RecordListResponseType = "LOC"
	RecordListResponseTypeMX         RecordListResponseType = "MX"
	RecordListResponseTypeNAPTR      RecordListResponseType = "NAPTR"
	RecordListResponseTypeNS         RecordListResponseType = "NS"
	RecordListResponseTypeOpenpgpkey RecordListResponseType = "OPENPGPKEY"
	RecordListResponseTypePTR        RecordListResponseType = "PTR"
	RecordListResponseTypeSMIMEA     RecordListResponseType = "SMIMEA"
	RecordListResponseTypeSRV        RecordListResponseType = "SRV"
	RecordListResponseTypeSSHFP      RecordListResponseType = "SSHFP"
	RecordListResponseTypeSVCB       RecordListResponseType = "SVCB"
	RecordListResponseTypeTLSA       RecordListResponseType = "TLSA"
	RecordListResponseTypeTXT        RecordListResponseType = "TXT"
	RecordListResponseTypeURI        RecordListResponseType = "URI"
)

func (r RecordListResponseType) IsKnown() bool {
	switch r {
	case RecordListResponseTypeA, RecordListResponseTypeAAAA, RecordListResponseTypeCAA, RecordListResponseTypeCERT, RecordListResponseTypeCNAME, RecordListResponseTypeDNSKEY, RecordListResponseTypeDS, RecordListResponseTypeHTTPS, RecordListResponseTypeLOC, RecordListResponseTypeMX, RecordListResponseTypeNAPTR, RecordListResponseTypeNS, RecordListResponseTypeOpenpgpkey, RecordListResponseTypePTR, RecordListResponseTypeSMIMEA, RecordListResponseTypeSRV, RecordListResponseTypeSSHFP, RecordListResponseTypeSVCB, RecordListResponseTypeTLSA, RecordListResponseTypeTXT, RecordListResponseTypeURI:
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

type RecordBatchResponse struct {
	Deletes []RecordBatchResponseDelete `json:"deletes"`
	Patches []RecordBatchResponsePatch  `json:"patches"`
	Posts   []RecordBatchResponsePost   `json:"posts"`
	Puts    []RecordBatchResponsePut    `json:"puts"`
	JSON    recordBatchResponseJSON     `json:"-"`
}

// recordBatchResponseJSON contains the JSON metadata for the struct
// [RecordBatchResponse]
type recordBatchResponseJSON struct {
	Deletes     apijson.Field
	Patches     apijson.Field
	Posts       apijson.Field
	Puts        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseJSON) RawJSON() string {
	return r.raw
}

type RecordBatchResponseDelete struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordBatchResponseDeletesType `json:"type"`
	JSON  recordBatchResponseDeleteJSON  `json:"-"`
	union RecordBatchResponseDeletesUnion
}

// recordBatchResponseDeleteJSON contains the JSON metadata for the struct
// [RecordBatchResponseDelete]
type recordBatchResponseDeleteJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordBatchResponseDeleteJSON) RawJSON() string {
	return r.raw
}

func (r *RecordBatchResponseDelete) UnmarshalJSON(data []byte) (err error) {
	*r = RecordBatchResponseDelete{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordBatchResponseDeletesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordBatchResponseDeletesARecord],
// [dns.RecordBatchResponseDeletesAAAARecord],
// [dns.RecordBatchResponseDeletesCAARecord],
// [dns.RecordBatchResponseDeletesCERTRecord],
// [dns.RecordBatchResponseDeletesCNAMERecord],
// [dns.RecordBatchResponseDeletesDNSKEYRecord],
// [dns.RecordBatchResponseDeletesDSRecord],
// [dns.RecordBatchResponseDeletesHTTPSRecord],
// [dns.RecordBatchResponseDeletesLOCRecord],
// [dns.RecordBatchResponseDeletesMXRecord],
// [dns.RecordBatchResponseDeletesNAPTRRecord],
// [dns.RecordBatchResponseDeletesNSRecord],
// [dns.RecordBatchResponseDeletesOpenpgpkeyRecord],
// [dns.RecordBatchResponseDeletesPTRRecord],
// [dns.RecordBatchResponseDeletesSMIMEARecord],
// [dns.RecordBatchResponseDeletesSRVRecord],
// [dns.RecordBatchResponseDeletesSSHFPRecord],
// [dns.RecordBatchResponseDeletesSVCBRecord],
// [dns.RecordBatchResponseDeletesTLSARecord],
// [dns.RecordBatchResponseDeletesTXTRecord],
// [dns.RecordBatchResponseDeletesURIRecord].
func (r RecordBatchResponseDelete) AsUnion() RecordBatchResponseDeletesUnion {
	return r.union
}

// Union satisfied by [dns.RecordBatchResponseDeletesARecord],
// [dns.RecordBatchResponseDeletesAAAARecord],
// [dns.RecordBatchResponseDeletesCAARecord],
// [dns.RecordBatchResponseDeletesCERTRecord],
// [dns.RecordBatchResponseDeletesCNAMERecord],
// [dns.RecordBatchResponseDeletesDNSKEYRecord],
// [dns.RecordBatchResponseDeletesDSRecord],
// [dns.RecordBatchResponseDeletesHTTPSRecord],
// [dns.RecordBatchResponseDeletesLOCRecord],
// [dns.RecordBatchResponseDeletesMXRecord],
// [dns.RecordBatchResponseDeletesNAPTRRecord],
// [dns.RecordBatchResponseDeletesNSRecord],
// [dns.RecordBatchResponseDeletesOpenpgpkeyRecord],
// [dns.RecordBatchResponseDeletesPTRRecord],
// [dns.RecordBatchResponseDeletesSMIMEARecord],
// [dns.RecordBatchResponseDeletesSRVRecord],
// [dns.RecordBatchResponseDeletesSSHFPRecord],
// [dns.RecordBatchResponseDeletesSVCBRecord],
// [dns.RecordBatchResponseDeletesTLSARecord],
// [dns.RecordBatchResponseDeletesTXTRecord] or
// [dns.RecordBatchResponseDeletesURIRecord].
type RecordBatchResponseDeletesUnion interface {
	implementsDNSRecordBatchResponseDelete()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponseDeletesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponseDeletesURIRecord{}),
		},
	)
}

type RecordBatchResponseDeletesARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesARecordJSON `json:"-"`
	ARecord
}

// recordBatchResponseDeletesARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponseDeletesARecord]
type recordBatchResponseDeletesARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesARecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordBatchResponseDeletesAAAARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesAAAARecord]
type recordBatchResponseDeletesAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesAAAARecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesCAARecordJSON `json:"-"`
	CAARecord
}

// recordBatchResponseDeletesCAARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesCAARecord]
type recordBatchResponseDeletesCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesCAARecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordBatchResponseDeletesCERTRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesCERTRecord]
type recordBatchResponseDeletesCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesCERTRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordBatchResponseDeletesCNAMERecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesCNAMERecord]
type recordBatchResponseDeletesCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesCNAMERecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordBatchResponseDeletesDNSKEYRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesDNSKEYRecord]
type recordBatchResponseDeletesDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesDNSKEYRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesDSRecordJSON `json:"-"`
	DSRecord
}

// recordBatchResponseDeletesDSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponseDeletesDSRecord]
type recordBatchResponseDeletesDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesDSRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordBatchResponseDeletesHTTPSRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesHTTPSRecord]
type recordBatchResponseDeletesHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesHTTPSRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordBatchResponseDeletesLOCRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesLOCRecord]
type recordBatchResponseDeletesLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesLOCRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesMXRecordJSON `json:"-"`
	MXRecord
}

// recordBatchResponseDeletesMXRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponseDeletesMXRecord]
type recordBatchResponseDeletesMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesMXRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordBatchResponseDeletesNAPTRRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesNAPTRRecord]
type recordBatchResponseDeletesNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesNAPTRRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesNSRecordJSON `json:"-"`
	NSRecord
}

// recordBatchResponseDeletesNSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponseDeletesNSRecord]
type recordBatchResponseDeletesNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesNSRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordBatchResponseDeletesOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesOpenpgpkeyRecordJSON `json:"-"`
}

// recordBatchResponseDeletesOpenpgpkeyRecordJSON contains the JSON metadata for
// the struct [RecordBatchResponseDeletesOpenpgpkeyRecord]
type recordBatchResponseDeletesOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesOpenpgpkeyRecord) implementsDNSRecordBatchResponseDelete() {}

// Record type.
type RecordBatchResponseDeletesOpenpgpkeyRecordType string

const (
	RecordBatchResponseDeletesOpenpgpkeyRecordTypeOpenpgpkey RecordBatchResponseDeletesOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordBatchResponseDeletesOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordBatchResponseDeletesOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordBatchResponseDeletesPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesPTRRecordJSON `json:"-"`
	PTRRecord
}

// recordBatchResponseDeletesPTRRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesPTRRecord]
type recordBatchResponseDeletesPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesPTRRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordBatchResponseDeletesSMIMEARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesSMIMEARecord]
type recordBatchResponseDeletesSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesSMIMEARecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordBatchResponseDeletesSRVRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesSRVRecord]
type recordBatchResponseDeletesSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesSRVRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordBatchResponseDeletesSSHFPRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesSSHFPRecord]
type recordBatchResponseDeletesSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesSSHFPRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordBatchResponseDeletesSVCBRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesSVCBRecord]
type recordBatchResponseDeletesSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesSVCBRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordBatchResponseDeletesTLSARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesTLSARecord]
type recordBatchResponseDeletesTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesTLSARecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordBatchResponseDeletesTXTRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesTXTRecord]
type recordBatchResponseDeletesTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesTXTRecord) implementsDNSRecordBatchResponseDelete() {}

type RecordBatchResponseDeletesURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeletesURIRecordJSON `json:"-"`
	URIRecord
}

// recordBatchResponseDeletesURIRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesURIRecord]
type recordBatchResponseDeletesURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesURIRecord) implementsDNSRecordBatchResponseDelete() {}

// Record type.
type RecordBatchResponseDeletesType string

const (
	RecordBatchResponseDeletesTypeA          RecordBatchResponseDeletesType = "A"
	RecordBatchResponseDeletesTypeAAAA       RecordBatchResponseDeletesType = "AAAA"
	RecordBatchResponseDeletesTypeCAA        RecordBatchResponseDeletesType = "CAA"
	RecordBatchResponseDeletesTypeCERT       RecordBatchResponseDeletesType = "CERT"
	RecordBatchResponseDeletesTypeCNAME      RecordBatchResponseDeletesType = "CNAME"
	RecordBatchResponseDeletesTypeDNSKEY     RecordBatchResponseDeletesType = "DNSKEY"
	RecordBatchResponseDeletesTypeDS         RecordBatchResponseDeletesType = "DS"
	RecordBatchResponseDeletesTypeHTTPS      RecordBatchResponseDeletesType = "HTTPS"
	RecordBatchResponseDeletesTypeLOC        RecordBatchResponseDeletesType = "LOC"
	RecordBatchResponseDeletesTypeMX         RecordBatchResponseDeletesType = "MX"
	RecordBatchResponseDeletesTypeNAPTR      RecordBatchResponseDeletesType = "NAPTR"
	RecordBatchResponseDeletesTypeNS         RecordBatchResponseDeletesType = "NS"
	RecordBatchResponseDeletesTypeOpenpgpkey RecordBatchResponseDeletesType = "OPENPGPKEY"
	RecordBatchResponseDeletesTypePTR        RecordBatchResponseDeletesType = "PTR"
	RecordBatchResponseDeletesTypeSMIMEA     RecordBatchResponseDeletesType = "SMIMEA"
	RecordBatchResponseDeletesTypeSRV        RecordBatchResponseDeletesType = "SRV"
	RecordBatchResponseDeletesTypeSSHFP      RecordBatchResponseDeletesType = "SSHFP"
	RecordBatchResponseDeletesTypeSVCB       RecordBatchResponseDeletesType = "SVCB"
	RecordBatchResponseDeletesTypeTLSA       RecordBatchResponseDeletesType = "TLSA"
	RecordBatchResponseDeletesTypeTXT        RecordBatchResponseDeletesType = "TXT"
	RecordBatchResponseDeletesTypeURI        RecordBatchResponseDeletesType = "URI"
)

func (r RecordBatchResponseDeletesType) IsKnown() bool {
	switch r {
	case RecordBatchResponseDeletesTypeA, RecordBatchResponseDeletesTypeAAAA, RecordBatchResponseDeletesTypeCAA, RecordBatchResponseDeletesTypeCERT, RecordBatchResponseDeletesTypeCNAME, RecordBatchResponseDeletesTypeDNSKEY, RecordBatchResponseDeletesTypeDS, RecordBatchResponseDeletesTypeHTTPS, RecordBatchResponseDeletesTypeLOC, RecordBatchResponseDeletesTypeMX, RecordBatchResponseDeletesTypeNAPTR, RecordBatchResponseDeletesTypeNS, RecordBatchResponseDeletesTypeOpenpgpkey, RecordBatchResponseDeletesTypePTR, RecordBatchResponseDeletesTypeSMIMEA, RecordBatchResponseDeletesTypeSRV, RecordBatchResponseDeletesTypeSSHFP, RecordBatchResponseDeletesTypeSVCB, RecordBatchResponseDeletesTypeTLSA, RecordBatchResponseDeletesTypeTXT, RecordBatchResponseDeletesTypeURI:
		return true
	}
	return false
}

type RecordBatchResponsePatch struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordBatchResponsePatchesType `json:"type"`
	JSON  recordBatchResponsePatchJSON   `json:"-"`
	union RecordBatchResponsePatchesUnion
}

// recordBatchResponsePatchJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatch]
type recordBatchResponsePatchJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordBatchResponsePatchJSON) RawJSON() string {
	return r.raw
}

func (r *RecordBatchResponsePatch) UnmarshalJSON(data []byte) (err error) {
	*r = RecordBatchResponsePatch{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordBatchResponsePatchesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordBatchResponsePatchesARecord],
// [dns.RecordBatchResponsePatchesAAAARecord],
// [dns.RecordBatchResponsePatchesCAARecord],
// [dns.RecordBatchResponsePatchesCERTRecord],
// [dns.RecordBatchResponsePatchesCNAMERecord],
// [dns.RecordBatchResponsePatchesDNSKEYRecord],
// [dns.RecordBatchResponsePatchesDSRecord],
// [dns.RecordBatchResponsePatchesHTTPSRecord],
// [dns.RecordBatchResponsePatchesLOCRecord],
// [dns.RecordBatchResponsePatchesMXRecord],
// [dns.RecordBatchResponsePatchesNAPTRRecord],
// [dns.RecordBatchResponsePatchesNSRecord],
// [dns.RecordBatchResponsePatchesOpenpgpkeyRecord],
// [dns.RecordBatchResponsePatchesPTRRecord],
// [dns.RecordBatchResponsePatchesSMIMEARecord],
// [dns.RecordBatchResponsePatchesSRVRecord],
// [dns.RecordBatchResponsePatchesSSHFPRecord],
// [dns.RecordBatchResponsePatchesSVCBRecord],
// [dns.RecordBatchResponsePatchesTLSARecord],
// [dns.RecordBatchResponsePatchesTXTRecord],
// [dns.RecordBatchResponsePatchesURIRecord].
func (r RecordBatchResponsePatch) AsUnion() RecordBatchResponsePatchesUnion {
	return r.union
}

// Union satisfied by [dns.RecordBatchResponsePatchesARecord],
// [dns.RecordBatchResponsePatchesAAAARecord],
// [dns.RecordBatchResponsePatchesCAARecord],
// [dns.RecordBatchResponsePatchesCERTRecord],
// [dns.RecordBatchResponsePatchesCNAMERecord],
// [dns.RecordBatchResponsePatchesDNSKEYRecord],
// [dns.RecordBatchResponsePatchesDSRecord],
// [dns.RecordBatchResponsePatchesHTTPSRecord],
// [dns.RecordBatchResponsePatchesLOCRecord],
// [dns.RecordBatchResponsePatchesMXRecord],
// [dns.RecordBatchResponsePatchesNAPTRRecord],
// [dns.RecordBatchResponsePatchesNSRecord],
// [dns.RecordBatchResponsePatchesOpenpgpkeyRecord],
// [dns.RecordBatchResponsePatchesPTRRecord],
// [dns.RecordBatchResponsePatchesSMIMEARecord],
// [dns.RecordBatchResponsePatchesSRVRecord],
// [dns.RecordBatchResponsePatchesSSHFPRecord],
// [dns.RecordBatchResponsePatchesSVCBRecord],
// [dns.RecordBatchResponsePatchesTLSARecord],
// [dns.RecordBatchResponsePatchesTXTRecord] or
// [dns.RecordBatchResponsePatchesURIRecord].
type RecordBatchResponsePatchesUnion interface {
	implementsDNSRecordBatchResponsePatch()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponsePatchesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePatchesURIRecord{}),
		},
	)
}

type RecordBatchResponsePatchesARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesARecordJSON `json:"-"`
	ARecord
}

// recordBatchResponsePatchesARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatchesARecord]
type recordBatchResponsePatchesARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesARecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordBatchResponsePatchesAAAARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesAAAARecord]
type recordBatchResponsePatchesAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesAAAARecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesCAARecordJSON `json:"-"`
	CAARecord
}

// recordBatchResponsePatchesCAARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesCAARecord]
type recordBatchResponsePatchesCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesCAARecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordBatchResponsePatchesCERTRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesCERTRecord]
type recordBatchResponsePatchesCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesCERTRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordBatchResponsePatchesCNAMERecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesCNAMERecord]
type recordBatchResponsePatchesCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesCNAMERecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordBatchResponsePatchesDNSKEYRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesDNSKEYRecord]
type recordBatchResponsePatchesDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesDNSKEYRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesDSRecordJSON `json:"-"`
	DSRecord
}

// recordBatchResponsePatchesDSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatchesDSRecord]
type recordBatchResponsePatchesDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesDSRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordBatchResponsePatchesHTTPSRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesHTTPSRecord]
type recordBatchResponsePatchesHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesHTTPSRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordBatchResponsePatchesLOCRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesLOCRecord]
type recordBatchResponsePatchesLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesLOCRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesMXRecordJSON `json:"-"`
	MXRecord
}

// recordBatchResponsePatchesMXRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatchesMXRecord]
type recordBatchResponsePatchesMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesMXRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordBatchResponsePatchesNAPTRRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesNAPTRRecord]
type recordBatchResponsePatchesNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesNAPTRRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesNSRecordJSON `json:"-"`
	NSRecord
}

// recordBatchResponsePatchesNSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatchesNSRecord]
type recordBatchResponsePatchesNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesNSRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordBatchResponsePatchesOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesOpenpgpkeyRecordJSON `json:"-"`
}

// recordBatchResponsePatchesOpenpgpkeyRecordJSON contains the JSON metadata for
// the struct [RecordBatchResponsePatchesOpenpgpkeyRecord]
type recordBatchResponsePatchesOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesOpenpgpkeyRecord) implementsDNSRecordBatchResponsePatch() {}

// Record type.
type RecordBatchResponsePatchesOpenpgpkeyRecordType string

const (
	RecordBatchResponsePatchesOpenpgpkeyRecordTypeOpenpgpkey RecordBatchResponsePatchesOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordBatchResponsePatchesOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePatchesOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordBatchResponsePatchesPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesPTRRecordJSON `json:"-"`
	PTRRecord
}

// recordBatchResponsePatchesPTRRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesPTRRecord]
type recordBatchResponsePatchesPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesPTRRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordBatchResponsePatchesSMIMEARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesSMIMEARecord]
type recordBatchResponsePatchesSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesSMIMEARecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordBatchResponsePatchesSRVRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesSRVRecord]
type recordBatchResponsePatchesSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesSRVRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordBatchResponsePatchesSSHFPRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesSSHFPRecord]
type recordBatchResponsePatchesSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesSSHFPRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordBatchResponsePatchesSVCBRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesSVCBRecord]
type recordBatchResponsePatchesSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesSVCBRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordBatchResponsePatchesTLSARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesTLSARecord]
type recordBatchResponsePatchesTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesTLSARecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordBatchResponsePatchesTXTRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesTXTRecord]
type recordBatchResponsePatchesTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesTXTRecord) implementsDNSRecordBatchResponsePatch() {}

type RecordBatchResponsePatchesURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchesURIRecordJSON `json:"-"`
	URIRecord
}

// recordBatchResponsePatchesURIRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesURIRecord]
type recordBatchResponsePatchesURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesURIRecord) implementsDNSRecordBatchResponsePatch() {}

// Record type.
type RecordBatchResponsePatchesType string

const (
	RecordBatchResponsePatchesTypeA          RecordBatchResponsePatchesType = "A"
	RecordBatchResponsePatchesTypeAAAA       RecordBatchResponsePatchesType = "AAAA"
	RecordBatchResponsePatchesTypeCAA        RecordBatchResponsePatchesType = "CAA"
	RecordBatchResponsePatchesTypeCERT       RecordBatchResponsePatchesType = "CERT"
	RecordBatchResponsePatchesTypeCNAME      RecordBatchResponsePatchesType = "CNAME"
	RecordBatchResponsePatchesTypeDNSKEY     RecordBatchResponsePatchesType = "DNSKEY"
	RecordBatchResponsePatchesTypeDS         RecordBatchResponsePatchesType = "DS"
	RecordBatchResponsePatchesTypeHTTPS      RecordBatchResponsePatchesType = "HTTPS"
	RecordBatchResponsePatchesTypeLOC        RecordBatchResponsePatchesType = "LOC"
	RecordBatchResponsePatchesTypeMX         RecordBatchResponsePatchesType = "MX"
	RecordBatchResponsePatchesTypeNAPTR      RecordBatchResponsePatchesType = "NAPTR"
	RecordBatchResponsePatchesTypeNS         RecordBatchResponsePatchesType = "NS"
	RecordBatchResponsePatchesTypeOpenpgpkey RecordBatchResponsePatchesType = "OPENPGPKEY"
	RecordBatchResponsePatchesTypePTR        RecordBatchResponsePatchesType = "PTR"
	RecordBatchResponsePatchesTypeSMIMEA     RecordBatchResponsePatchesType = "SMIMEA"
	RecordBatchResponsePatchesTypeSRV        RecordBatchResponsePatchesType = "SRV"
	RecordBatchResponsePatchesTypeSSHFP      RecordBatchResponsePatchesType = "SSHFP"
	RecordBatchResponsePatchesTypeSVCB       RecordBatchResponsePatchesType = "SVCB"
	RecordBatchResponsePatchesTypeTLSA       RecordBatchResponsePatchesType = "TLSA"
	RecordBatchResponsePatchesTypeTXT        RecordBatchResponsePatchesType = "TXT"
	RecordBatchResponsePatchesTypeURI        RecordBatchResponsePatchesType = "URI"
)

func (r RecordBatchResponsePatchesType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePatchesTypeA, RecordBatchResponsePatchesTypeAAAA, RecordBatchResponsePatchesTypeCAA, RecordBatchResponsePatchesTypeCERT, RecordBatchResponsePatchesTypeCNAME, RecordBatchResponsePatchesTypeDNSKEY, RecordBatchResponsePatchesTypeDS, RecordBatchResponsePatchesTypeHTTPS, RecordBatchResponsePatchesTypeLOC, RecordBatchResponsePatchesTypeMX, RecordBatchResponsePatchesTypeNAPTR, RecordBatchResponsePatchesTypeNS, RecordBatchResponsePatchesTypeOpenpgpkey, RecordBatchResponsePatchesTypePTR, RecordBatchResponsePatchesTypeSMIMEA, RecordBatchResponsePatchesTypeSRV, RecordBatchResponsePatchesTypeSSHFP, RecordBatchResponsePatchesTypeSVCB, RecordBatchResponsePatchesTypeTLSA, RecordBatchResponsePatchesTypeTXT, RecordBatchResponsePatchesTypeURI:
		return true
	}
	return false
}

type RecordBatchResponsePost struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordBatchResponsePostsType `json:"type"`
	JSON  recordBatchResponsePostJSON  `json:"-"`
	union RecordBatchResponsePostsUnion
}

// recordBatchResponsePostJSON contains the JSON metadata for the struct
// [RecordBatchResponsePost]
type recordBatchResponsePostJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordBatchResponsePostJSON) RawJSON() string {
	return r.raw
}

func (r *RecordBatchResponsePost) UnmarshalJSON(data []byte) (err error) {
	*r = RecordBatchResponsePost{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordBatchResponsePostsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordBatchResponsePostsARecord],
// [dns.RecordBatchResponsePostsAAAARecord],
// [dns.RecordBatchResponsePostsCAARecord],
// [dns.RecordBatchResponsePostsCERTRecord],
// [dns.RecordBatchResponsePostsCNAMERecord],
// [dns.RecordBatchResponsePostsDNSKEYRecord],
// [dns.RecordBatchResponsePostsDSRecord],
// [dns.RecordBatchResponsePostsHTTPSRecord],
// [dns.RecordBatchResponsePostsLOCRecord], [dns.RecordBatchResponsePostsMXRecord],
// [dns.RecordBatchResponsePostsNAPTRRecord],
// [dns.RecordBatchResponsePostsNSRecord],
// [dns.RecordBatchResponsePostsOpenpgpkeyRecord],
// [dns.RecordBatchResponsePostsPTRRecord],
// [dns.RecordBatchResponsePostsSMIMEARecord],
// [dns.RecordBatchResponsePostsSRVRecord],
// [dns.RecordBatchResponsePostsSSHFPRecord],
// [dns.RecordBatchResponsePostsSVCBRecord],
// [dns.RecordBatchResponsePostsTLSARecord],
// [dns.RecordBatchResponsePostsTXTRecord],
// [dns.RecordBatchResponsePostsURIRecord].
func (r RecordBatchResponsePost) AsUnion() RecordBatchResponsePostsUnion {
	return r.union
}

// Union satisfied by [dns.RecordBatchResponsePostsARecord],
// [dns.RecordBatchResponsePostsAAAARecord],
// [dns.RecordBatchResponsePostsCAARecord],
// [dns.RecordBatchResponsePostsCERTRecord],
// [dns.RecordBatchResponsePostsCNAMERecord],
// [dns.RecordBatchResponsePostsDNSKEYRecord],
// [dns.RecordBatchResponsePostsDSRecord],
// [dns.RecordBatchResponsePostsHTTPSRecord],
// [dns.RecordBatchResponsePostsLOCRecord], [dns.RecordBatchResponsePostsMXRecord],
// [dns.RecordBatchResponsePostsNAPTRRecord],
// [dns.RecordBatchResponsePostsNSRecord],
// [dns.RecordBatchResponsePostsOpenpgpkeyRecord],
// [dns.RecordBatchResponsePostsPTRRecord],
// [dns.RecordBatchResponsePostsSMIMEARecord],
// [dns.RecordBatchResponsePostsSRVRecord],
// [dns.RecordBatchResponsePostsSSHFPRecord],
// [dns.RecordBatchResponsePostsSVCBRecord],
// [dns.RecordBatchResponsePostsTLSARecord],
// [dns.RecordBatchResponsePostsTXTRecord] or
// [dns.RecordBatchResponsePostsURIRecord].
type RecordBatchResponsePostsUnion interface {
	implementsDNSRecordBatchResponsePost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponsePostsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePostsURIRecord{}),
		},
	)
}

type RecordBatchResponsePostsARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsARecordJSON `json:"-"`
	ARecord
}

// recordBatchResponsePostsARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsARecord]
type recordBatchResponsePostsARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsARecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordBatchResponsePostsAAAARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsAAAARecord]
type recordBatchResponsePostsAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsAAAARecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsCAARecordJSON `json:"-"`
	CAARecord
}

// recordBatchResponsePostsCAARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsCAARecord]
type recordBatchResponsePostsCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsCAARecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordBatchResponsePostsCERTRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsCERTRecord]
type recordBatchResponsePostsCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsCERTRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordBatchResponsePostsCNAMERecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsCNAMERecord]
type recordBatchResponsePostsCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsCNAMERecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordBatchResponsePostsDNSKEYRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsDNSKEYRecord]
type recordBatchResponsePostsDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsDNSKEYRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsDSRecordJSON `json:"-"`
	DSRecord
}

// recordBatchResponsePostsDSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsDSRecord]
type recordBatchResponsePostsDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsDSRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordBatchResponsePostsHTTPSRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsHTTPSRecord]
type recordBatchResponsePostsHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsHTTPSRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordBatchResponsePostsLOCRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsLOCRecord]
type recordBatchResponsePostsLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsLOCRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsMXRecordJSON `json:"-"`
	MXRecord
}

// recordBatchResponsePostsMXRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsMXRecord]
type recordBatchResponsePostsMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsMXRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordBatchResponsePostsNAPTRRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsNAPTRRecord]
type recordBatchResponsePostsNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsNAPTRRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsNSRecordJSON `json:"-"`
	NSRecord
}

// recordBatchResponsePostsNSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsNSRecord]
type recordBatchResponsePostsNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsNSRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordBatchResponsePostsOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                    `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsOpenpgpkeyRecordJSON `json:"-"`
}

// recordBatchResponsePostsOpenpgpkeyRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsOpenpgpkeyRecord]
type recordBatchResponsePostsOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsOpenpgpkeyRecord) implementsDNSRecordBatchResponsePost() {}

// Record type.
type RecordBatchResponsePostsOpenpgpkeyRecordType string

const (
	RecordBatchResponsePostsOpenpgpkeyRecordTypeOpenpgpkey RecordBatchResponsePostsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordBatchResponsePostsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePostsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordBatchResponsePostsPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsPTRRecordJSON `json:"-"`
	PTRRecord
}

// recordBatchResponsePostsPTRRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsPTRRecord]
type recordBatchResponsePostsPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsPTRRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordBatchResponsePostsSMIMEARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsSMIMEARecord]
type recordBatchResponsePostsSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsSMIMEARecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordBatchResponsePostsSRVRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsSRVRecord]
type recordBatchResponsePostsSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsSRVRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordBatchResponsePostsSSHFPRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePostsSSHFPRecord]
type recordBatchResponsePostsSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsSSHFPRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordBatchResponsePostsSVCBRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsSVCBRecord]
type recordBatchResponsePostsSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsSVCBRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordBatchResponsePostsTLSARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsTLSARecord]
type recordBatchResponsePostsTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsTLSARecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordBatchResponsePostsTXTRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsTXTRecord]
type recordBatchResponsePostsTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsTXTRecord) implementsDNSRecordBatchResponsePost() {}

type RecordBatchResponsePostsURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostsURIRecordJSON `json:"-"`
	URIRecord
}

// recordBatchResponsePostsURIRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsURIRecord]
type recordBatchResponsePostsURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePostsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsURIRecord) implementsDNSRecordBatchResponsePost() {}

// Record type.
type RecordBatchResponsePostsType string

const (
	RecordBatchResponsePostsTypeA          RecordBatchResponsePostsType = "A"
	RecordBatchResponsePostsTypeAAAA       RecordBatchResponsePostsType = "AAAA"
	RecordBatchResponsePostsTypeCAA        RecordBatchResponsePostsType = "CAA"
	RecordBatchResponsePostsTypeCERT       RecordBatchResponsePostsType = "CERT"
	RecordBatchResponsePostsTypeCNAME      RecordBatchResponsePostsType = "CNAME"
	RecordBatchResponsePostsTypeDNSKEY     RecordBatchResponsePostsType = "DNSKEY"
	RecordBatchResponsePostsTypeDS         RecordBatchResponsePostsType = "DS"
	RecordBatchResponsePostsTypeHTTPS      RecordBatchResponsePostsType = "HTTPS"
	RecordBatchResponsePostsTypeLOC        RecordBatchResponsePostsType = "LOC"
	RecordBatchResponsePostsTypeMX         RecordBatchResponsePostsType = "MX"
	RecordBatchResponsePostsTypeNAPTR      RecordBatchResponsePostsType = "NAPTR"
	RecordBatchResponsePostsTypeNS         RecordBatchResponsePostsType = "NS"
	RecordBatchResponsePostsTypeOpenpgpkey RecordBatchResponsePostsType = "OPENPGPKEY"
	RecordBatchResponsePostsTypePTR        RecordBatchResponsePostsType = "PTR"
	RecordBatchResponsePostsTypeSMIMEA     RecordBatchResponsePostsType = "SMIMEA"
	RecordBatchResponsePostsTypeSRV        RecordBatchResponsePostsType = "SRV"
	RecordBatchResponsePostsTypeSSHFP      RecordBatchResponsePostsType = "SSHFP"
	RecordBatchResponsePostsTypeSVCB       RecordBatchResponsePostsType = "SVCB"
	RecordBatchResponsePostsTypeTLSA       RecordBatchResponsePostsType = "TLSA"
	RecordBatchResponsePostsTypeTXT        RecordBatchResponsePostsType = "TXT"
	RecordBatchResponsePostsTypeURI        RecordBatchResponsePostsType = "URI"
)

func (r RecordBatchResponsePostsType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePostsTypeA, RecordBatchResponsePostsTypeAAAA, RecordBatchResponsePostsTypeCAA, RecordBatchResponsePostsTypeCERT, RecordBatchResponsePostsTypeCNAME, RecordBatchResponsePostsTypeDNSKEY, RecordBatchResponsePostsTypeDS, RecordBatchResponsePostsTypeHTTPS, RecordBatchResponsePostsTypeLOC, RecordBatchResponsePostsTypeMX, RecordBatchResponsePostsTypeNAPTR, RecordBatchResponsePostsTypeNS, RecordBatchResponsePostsTypeOpenpgpkey, RecordBatchResponsePostsTypePTR, RecordBatchResponsePostsTypeSMIMEA, RecordBatchResponsePostsTypeSRV, RecordBatchResponsePostsTypeSSHFP, RecordBatchResponsePostsTypeSVCB, RecordBatchResponsePostsTypeTLSA, RecordBatchResponsePostsTypeTXT, RecordBatchResponsePostsTypeURI:
		return true
	}
	return false
}

type RecordBatchResponsePut struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordBatchResponsePutsType `json:"type"`
	JSON  recordBatchResponsePutJSON  `json:"-"`
	union RecordBatchResponsePutsUnion
}

// recordBatchResponsePutJSON contains the JSON metadata for the struct
// [RecordBatchResponsePut]
type recordBatchResponsePutJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordBatchResponsePutJSON) RawJSON() string {
	return r.raw
}

func (r *RecordBatchResponsePut) UnmarshalJSON(data []byte) (err error) {
	*r = RecordBatchResponsePut{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordBatchResponsePutsUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordBatchResponsePutsARecord],
// [dns.RecordBatchResponsePutsAAAARecord], [dns.RecordBatchResponsePutsCAARecord],
// [dns.RecordBatchResponsePutsCERTRecord],
// [dns.RecordBatchResponsePutsCNAMERecord],
// [dns.RecordBatchResponsePutsDNSKEYRecord],
// [dns.RecordBatchResponsePutsDSRecord], [dns.RecordBatchResponsePutsHTTPSRecord],
// [dns.RecordBatchResponsePutsLOCRecord], [dns.RecordBatchResponsePutsMXRecord],
// [dns.RecordBatchResponsePutsNAPTRRecord], [dns.RecordBatchResponsePutsNSRecord],
// [dns.RecordBatchResponsePutsOpenpgpkeyRecord],
// [dns.RecordBatchResponsePutsPTRRecord],
// [dns.RecordBatchResponsePutsSMIMEARecord],
// [dns.RecordBatchResponsePutsSRVRecord],
// [dns.RecordBatchResponsePutsSSHFPRecord],
// [dns.RecordBatchResponsePutsSVCBRecord],
// [dns.RecordBatchResponsePutsTLSARecord], [dns.RecordBatchResponsePutsTXTRecord],
// [dns.RecordBatchResponsePutsURIRecord].
func (r RecordBatchResponsePut) AsUnion() RecordBatchResponsePutsUnion {
	return r.union
}

// Union satisfied by [dns.RecordBatchResponsePutsARecord],
// [dns.RecordBatchResponsePutsAAAARecord], [dns.RecordBatchResponsePutsCAARecord],
// [dns.RecordBatchResponsePutsCERTRecord],
// [dns.RecordBatchResponsePutsCNAMERecord],
// [dns.RecordBatchResponsePutsDNSKEYRecord],
// [dns.RecordBatchResponsePutsDSRecord], [dns.RecordBatchResponsePutsHTTPSRecord],
// [dns.RecordBatchResponsePutsLOCRecord], [dns.RecordBatchResponsePutsMXRecord],
// [dns.RecordBatchResponsePutsNAPTRRecord], [dns.RecordBatchResponsePutsNSRecord],
// [dns.RecordBatchResponsePutsOpenpgpkeyRecord],
// [dns.RecordBatchResponsePutsPTRRecord],
// [dns.RecordBatchResponsePutsSMIMEARecord],
// [dns.RecordBatchResponsePutsSRVRecord],
// [dns.RecordBatchResponsePutsSSHFPRecord],
// [dns.RecordBatchResponsePutsSVCBRecord],
// [dns.RecordBatchResponsePutsTLSARecord], [dns.RecordBatchResponsePutsTXTRecord]
// or [dns.RecordBatchResponsePutsURIRecord].
type RecordBatchResponsePutsUnion interface {
	implementsDNSRecordBatchResponsePut()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponsePutsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordBatchResponsePutsURIRecord{}),
		},
	)
}

type RecordBatchResponsePutsARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsARecordJSON `json:"-"`
	ARecord
}

// recordBatchResponsePutsARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsARecord]
type recordBatchResponsePutsARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsARecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordBatchResponsePutsAAAARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsAAAARecord]
type recordBatchResponsePutsAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsAAAARecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsCAARecordJSON `json:"-"`
	CAARecord
}

// recordBatchResponsePutsCAARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsCAARecord]
type recordBatchResponsePutsCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsCAARecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordBatchResponsePutsCERTRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsCERTRecord]
type recordBatchResponsePutsCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsCERTRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordBatchResponsePutsCNAMERecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsCNAMERecord]
type recordBatchResponsePutsCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsCNAMERecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordBatchResponsePutsDNSKEYRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePutsDNSKEYRecord]
type recordBatchResponsePutsDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsDNSKEYRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsDSRecordJSON `json:"-"`
	DSRecord
}

// recordBatchResponsePutsDSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsDSRecord]
type recordBatchResponsePutsDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsDSRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordBatchResponsePutsHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsHTTPSRecord]
type recordBatchResponsePutsHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsHTTPSRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordBatchResponsePutsLOCRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsLOCRecord]
type recordBatchResponsePutsLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsLOCRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsMXRecordJSON `json:"-"`
	MXRecord
}

// recordBatchResponsePutsMXRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsMXRecord]
type recordBatchResponsePutsMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsMXRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordBatchResponsePutsNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsNAPTRRecord]
type recordBatchResponsePutsNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsNAPTRRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                           `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsNSRecordJSON `json:"-"`
	NSRecord
}

// recordBatchResponsePutsNSRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsNSRecord]
type recordBatchResponsePutsNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsNSRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordBatchResponsePutsOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsOpenpgpkeyRecordJSON `json:"-"`
}

// recordBatchResponsePutsOpenpgpkeyRecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePutsOpenpgpkeyRecord]
type recordBatchResponsePutsOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsOpenpgpkeyRecord) implementsDNSRecordBatchResponsePut() {}

// Record type.
type RecordBatchResponsePutsOpenpgpkeyRecordType string

const (
	RecordBatchResponsePutsOpenpgpkeyRecordTypeOpenpgpkey RecordBatchResponsePutsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordBatchResponsePutsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePutsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordBatchResponsePutsPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsPTRRecordJSON `json:"-"`
	PTRRecord
}

// recordBatchResponsePutsPTRRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsPTRRecord]
type recordBatchResponsePutsPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsPTRRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                               `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordBatchResponsePutsSMIMEARecordJSON contains the JSON metadata for the
// struct [RecordBatchResponsePutsSMIMEARecord]
type recordBatchResponsePutsSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsSMIMEARecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordBatchResponsePutsSRVRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsSRVRecord]
type recordBatchResponsePutsSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsSRVRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordBatchResponsePutsSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsSSHFPRecord]
type recordBatchResponsePutsSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsSSHFPRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordBatchResponsePutsSVCBRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsSVCBRecord]
type recordBatchResponsePutsSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsSVCBRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordBatchResponsePutsTLSARecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsTLSARecord]
type recordBatchResponsePutsTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsTLSARecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordBatchResponsePutsTXTRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsTXTRecord]
type recordBatchResponsePutsTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsTXTRecord) implementsDNSRecordBatchResponsePut() {}

type RecordBatchResponsePutsURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                            `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutsURIRecordJSON `json:"-"`
	URIRecord
}

// recordBatchResponsePutsURIRecordJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsURIRecord]
type recordBatchResponsePutsURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePutsURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsURIRecord) implementsDNSRecordBatchResponsePut() {}

// Record type.
type RecordBatchResponsePutsType string

const (
	RecordBatchResponsePutsTypeA          RecordBatchResponsePutsType = "A"
	RecordBatchResponsePutsTypeAAAA       RecordBatchResponsePutsType = "AAAA"
	RecordBatchResponsePutsTypeCAA        RecordBatchResponsePutsType = "CAA"
	RecordBatchResponsePutsTypeCERT       RecordBatchResponsePutsType = "CERT"
	RecordBatchResponsePutsTypeCNAME      RecordBatchResponsePutsType = "CNAME"
	RecordBatchResponsePutsTypeDNSKEY     RecordBatchResponsePutsType = "DNSKEY"
	RecordBatchResponsePutsTypeDS         RecordBatchResponsePutsType = "DS"
	RecordBatchResponsePutsTypeHTTPS      RecordBatchResponsePutsType = "HTTPS"
	RecordBatchResponsePutsTypeLOC        RecordBatchResponsePutsType = "LOC"
	RecordBatchResponsePutsTypeMX         RecordBatchResponsePutsType = "MX"
	RecordBatchResponsePutsTypeNAPTR      RecordBatchResponsePutsType = "NAPTR"
	RecordBatchResponsePutsTypeNS         RecordBatchResponsePutsType = "NS"
	RecordBatchResponsePutsTypeOpenpgpkey RecordBatchResponsePutsType = "OPENPGPKEY"
	RecordBatchResponsePutsTypePTR        RecordBatchResponsePutsType = "PTR"
	RecordBatchResponsePutsTypeSMIMEA     RecordBatchResponsePutsType = "SMIMEA"
	RecordBatchResponsePutsTypeSRV        RecordBatchResponsePutsType = "SRV"
	RecordBatchResponsePutsTypeSSHFP      RecordBatchResponsePutsType = "SSHFP"
	RecordBatchResponsePutsTypeSVCB       RecordBatchResponsePutsType = "SVCB"
	RecordBatchResponsePutsTypeTLSA       RecordBatchResponsePutsType = "TLSA"
	RecordBatchResponsePutsTypeTXT        RecordBatchResponsePutsType = "TXT"
	RecordBatchResponsePutsTypeURI        RecordBatchResponsePutsType = "URI"
)

func (r RecordBatchResponsePutsType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePutsTypeA, RecordBatchResponsePutsTypeAAAA, RecordBatchResponsePutsTypeCAA, RecordBatchResponsePutsTypeCERT, RecordBatchResponsePutsTypeCNAME, RecordBatchResponsePutsTypeDNSKEY, RecordBatchResponsePutsTypeDS, RecordBatchResponsePutsTypeHTTPS, RecordBatchResponsePutsTypeLOC, RecordBatchResponsePutsTypeMX, RecordBatchResponsePutsTypeNAPTR, RecordBatchResponsePutsTypeNS, RecordBatchResponsePutsTypeOpenpgpkey, RecordBatchResponsePutsTypePTR, RecordBatchResponsePutsTypeSMIMEA, RecordBatchResponsePutsTypeSRV, RecordBatchResponsePutsTypeSSHFP, RecordBatchResponsePutsTypeSVCB, RecordBatchResponsePutsTypeTLSA, RecordBatchResponsePutsTypeTXT, RecordBatchResponsePutsTypeURI:
		return true
	}
	return false
}

type RecordEditResponse struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordEditResponseType `json:"type"`
	JSON  recordEditResponseJSON `json:"-"`
	union RecordEditResponseUnion
}

// recordEditResponseJSON contains the JSON metadata for the struct
// [RecordEditResponse]
type recordEditResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RecordEditResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RecordEditResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordEditResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordEditResponseARecord],
// [dns.RecordEditResponseAAAARecord], [dns.RecordEditResponseCAARecord],
// [dns.RecordEditResponseCERTRecord], [dns.RecordEditResponseCNAMERecord],
// [dns.RecordEditResponseDNSKEYRecord], [dns.RecordEditResponseDSRecord],
// [dns.RecordEditResponseHTTPSRecord], [dns.RecordEditResponseLOCRecord],
// [dns.RecordEditResponseMXRecord], [dns.RecordEditResponseNAPTRRecord],
// [dns.RecordEditResponseNSRecord], [dns.RecordEditResponseOpenpgpkeyRecord],
// [dns.RecordEditResponsePTRRecord], [dns.RecordEditResponseSMIMEARecord],
// [dns.RecordEditResponseSRVRecord], [dns.RecordEditResponseSSHFPRecord],
// [dns.RecordEditResponseSVCBRecord], [dns.RecordEditResponseTLSARecord],
// [dns.RecordEditResponseTXTRecord], [dns.RecordEditResponseURIRecord].
func (r RecordEditResponse) AsUnion() RecordEditResponseUnion {
	return r.union
}

// Union satisfied by [dns.RecordEditResponseARecord],
// [dns.RecordEditResponseAAAARecord], [dns.RecordEditResponseCAARecord],
// [dns.RecordEditResponseCERTRecord], [dns.RecordEditResponseCNAMERecord],
// [dns.RecordEditResponseDNSKEYRecord], [dns.RecordEditResponseDSRecord],
// [dns.RecordEditResponseHTTPSRecord], [dns.RecordEditResponseLOCRecord],
// [dns.RecordEditResponseMXRecord], [dns.RecordEditResponseNAPTRRecord],
// [dns.RecordEditResponseNSRecord], [dns.RecordEditResponseOpenpgpkeyRecord],
// [dns.RecordEditResponsePTRRecord], [dns.RecordEditResponseSMIMEARecord],
// [dns.RecordEditResponseSRVRecord], [dns.RecordEditResponseSSHFPRecord],
// [dns.RecordEditResponseSVCBRecord], [dns.RecordEditResponseTLSARecord],
// [dns.RecordEditResponseTXTRecord] or [dns.RecordEditResponseURIRecord].
type RecordEditResponseUnion interface {
	implementsDNSRecordEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponsePTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordEditResponseURIRecord{}),
		},
	)
}

type RecordEditResponseARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseARecordJSON `json:"-"`
	ARecord
}

// recordEditResponseARecordJSON contains the JSON metadata for the struct
// [RecordEditResponseARecord]
type recordEditResponseARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseARecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordEditResponseAAAARecordJSON contains the JSON metadata for the struct
// [RecordEditResponseAAAARecord]
type recordEditResponseAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseAAAARecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseCAARecordJSON `json:"-"`
	CAARecord
}

// recordEditResponseCAARecordJSON contains the JSON metadata for the struct
// [RecordEditResponseCAARecord]
type recordEditResponseCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseCAARecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordEditResponseCERTRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseCERTRecord]
type recordEditResponseCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseCERTRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordEditResponseCNAMERecordJSON contains the JSON metadata for the struct
// [RecordEditResponseCNAMERecord]
type recordEditResponseCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseCNAMERecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordEditResponseDNSKEYRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseDNSKEYRecord]
type recordEditResponseDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseDNSKEYRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseDSRecordJSON `json:"-"`
	DSRecord
}

// recordEditResponseDSRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseDSRecord]
type recordEditResponseDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseDSRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordEditResponseHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseHTTPSRecord]
type recordEditResponseHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseHTTPSRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordEditResponseLOCRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseLOCRecord]
type recordEditResponseLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseLOCRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseMXRecordJSON `json:"-"`
	MXRecord
}

// recordEditResponseMXRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseMXRecord]
type recordEditResponseMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseMXRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordEditResponseNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseNAPTRRecord]
type recordEditResponseNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseNAPTRRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseNSRecordJSON `json:"-"`
	NSRecord
}

// recordEditResponseNSRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseNSRecord]
type recordEditResponseNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseNSRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordEditResponseOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                              `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseOpenpgpkeyRecordJSON `json:"-"`
}

// recordEditResponseOpenpgpkeyRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseOpenpgpkeyRecord]
type recordEditResponseOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseOpenpgpkeyRecord) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseOpenpgpkeyRecordType string

const (
	RecordEditResponseOpenpgpkeyRecordTypeOpenpgpkey RecordEditResponseOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordEditResponseOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordEditResponseOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordEditResponsePTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponsePTRRecordJSON `json:"-"`
	PTRRecord
}

// recordEditResponsePTRRecordJSON contains the JSON metadata for the struct
// [RecordEditResponsePTRRecord]
type recordEditResponsePTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponsePTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponsePTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponsePTRRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                          `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordEditResponseSMIMEARecordJSON contains the JSON metadata for the struct
// [RecordEditResponseSMIMEARecord]
type recordEditResponseSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSMIMEARecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordEditResponseSRVRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseSRVRecord]
type recordEditResponseSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSRVRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordEditResponseSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseSSHFPRecord]
type recordEditResponseSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSSHFPRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordEditResponseSVCBRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseSVCBRecord]
type recordEditResponseSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseSVCBRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordEditResponseTLSARecordJSON contains the JSON metadata for the struct
// [RecordEditResponseTLSARecord]
type recordEditResponseTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseTLSARecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordEditResponseTXTRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseTXTRecord]
type recordEditResponseTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseTXTRecord) implementsDNSRecordEditResponse() {}

type RecordEditResponseURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseURIRecordJSON `json:"-"`
	URIRecord
}

// recordEditResponseURIRecordJSON contains the JSON metadata for the struct
// [RecordEditResponseURIRecord]
type recordEditResponseURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponseURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseURIRecord) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseType string

const (
	RecordEditResponseTypeA          RecordEditResponseType = "A"
	RecordEditResponseTypeAAAA       RecordEditResponseType = "AAAA"
	RecordEditResponseTypeCAA        RecordEditResponseType = "CAA"
	RecordEditResponseTypeCERT       RecordEditResponseType = "CERT"
	RecordEditResponseTypeCNAME      RecordEditResponseType = "CNAME"
	RecordEditResponseTypeDNSKEY     RecordEditResponseType = "DNSKEY"
	RecordEditResponseTypeDS         RecordEditResponseType = "DS"
	RecordEditResponseTypeHTTPS      RecordEditResponseType = "HTTPS"
	RecordEditResponseTypeLOC        RecordEditResponseType = "LOC"
	RecordEditResponseTypeMX         RecordEditResponseType = "MX"
	RecordEditResponseTypeNAPTR      RecordEditResponseType = "NAPTR"
	RecordEditResponseTypeNS         RecordEditResponseType = "NS"
	RecordEditResponseTypeOpenpgpkey RecordEditResponseType = "OPENPGPKEY"
	RecordEditResponseTypePTR        RecordEditResponseType = "PTR"
	RecordEditResponseTypeSMIMEA     RecordEditResponseType = "SMIMEA"
	RecordEditResponseTypeSRV        RecordEditResponseType = "SRV"
	RecordEditResponseTypeSSHFP      RecordEditResponseType = "SSHFP"
	RecordEditResponseTypeSVCB       RecordEditResponseType = "SVCB"
	RecordEditResponseTypeTLSA       RecordEditResponseType = "TLSA"
	RecordEditResponseTypeTXT        RecordEditResponseType = "TXT"
	RecordEditResponseTypeURI        RecordEditResponseType = "URI"
)

func (r RecordEditResponseType) IsKnown() bool {
	switch r {
	case RecordEditResponseTypeA, RecordEditResponseTypeAAAA, RecordEditResponseTypeCAA, RecordEditResponseTypeCERT, RecordEditResponseTypeCNAME, RecordEditResponseTypeDNSKEY, RecordEditResponseTypeDS, RecordEditResponseTypeHTTPS, RecordEditResponseTypeLOC, RecordEditResponseTypeMX, RecordEditResponseTypeNAPTR, RecordEditResponseTypeNS, RecordEditResponseTypeOpenpgpkey, RecordEditResponseTypePTR, RecordEditResponseTypeSMIMEA, RecordEditResponseTypeSRV, RecordEditResponseTypeSSHFP, RecordEditResponseTypeSVCB, RecordEditResponseTypeTLSA, RecordEditResponseTypeTXT, RecordEditResponseTypeURI:
		return true
	}
	return false
}

type RecordGetResponse struct {
	// Identifier
	ID string `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data"`
	// This field can have the runtime type of [interface{}].
	Meta interface{} `json:"meta"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings"`
	// This field can have the runtime type of [[]RecordTags].
	Tags interface{} `json:"tags"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time `json:"tags_modified_on" format:"date-time"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl"`
	// Record type.
	Type  RecordGetResponseType `json:"type"`
	JSON  recordGetResponseJSON `json:"-"`
	union RecordGetResponseUnion
}

// recordGetResponseJSON contains the JSON metadata for the struct
// [RecordGetResponse]
type recordGetResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CommentModifiedOn apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Data              apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Priority          apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TagsModifiedOn    apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r recordGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RecordGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RecordGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordGetResponseARecord],
// [dns.RecordGetResponseAAAARecord], [dns.RecordGetResponseCAARecord],
// [dns.RecordGetResponseCERTRecord], [dns.RecordGetResponseCNAMERecord],
// [dns.RecordGetResponseDNSKEYRecord], [dns.RecordGetResponseDSRecord],
// [dns.RecordGetResponseHTTPSRecord], [dns.RecordGetResponseLOCRecord],
// [dns.RecordGetResponseMXRecord], [dns.RecordGetResponseNAPTRRecord],
// [dns.RecordGetResponseNSRecord], [dns.RecordGetResponseOpenpgpkeyRecord],
// [dns.RecordGetResponsePTRRecord], [dns.RecordGetResponseSMIMEARecord],
// [dns.RecordGetResponseSRVRecord], [dns.RecordGetResponseSSHFPRecord],
// [dns.RecordGetResponseSVCBRecord], [dns.RecordGetResponseTLSARecord],
// [dns.RecordGetResponseTXTRecord], [dns.RecordGetResponseURIRecord].
func (r RecordGetResponse) AsUnion() RecordGetResponseUnion {
	return r.union
}

// Union satisfied by [dns.RecordGetResponseARecord],
// [dns.RecordGetResponseAAAARecord], [dns.RecordGetResponseCAARecord],
// [dns.RecordGetResponseCERTRecord], [dns.RecordGetResponseCNAMERecord],
// [dns.RecordGetResponseDNSKEYRecord], [dns.RecordGetResponseDSRecord],
// [dns.RecordGetResponseHTTPSRecord], [dns.RecordGetResponseLOCRecord],
// [dns.RecordGetResponseMXRecord], [dns.RecordGetResponseNAPTRRecord],
// [dns.RecordGetResponseNSRecord], [dns.RecordGetResponseOpenpgpkeyRecord],
// [dns.RecordGetResponsePTRRecord], [dns.RecordGetResponseSMIMEARecord],
// [dns.RecordGetResponseSRVRecord], [dns.RecordGetResponseSSHFPRecord],
// [dns.RecordGetResponseSVCBRecord], [dns.RecordGetResponseTLSARecord],
// [dns.RecordGetResponseTXTRecord] or [dns.RecordGetResponseURIRecord].
type RecordGetResponseUnion interface {
	implementsDNSRecordGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseAAAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseCAARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseCERTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseCNAMERecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseDNSKEYRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseDSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseHTTPSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseLOCRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseMXRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseNAPTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseNSRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseOpenpgpkeyRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponsePTRRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseSMIMEARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseSRVRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseSSHFPRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseSVCBRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseTLSARecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseTXTRecord{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RecordGetResponseURIRecord{}),
		},
	)
}

type RecordGetResponseARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                    `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseARecordJSON `json:"-"`
	ARecord
}

// recordGetResponseARecordJSON contains the JSON metadata for the struct
// [RecordGetResponseARecord]
type recordGetResponseARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseARecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseAAAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordGetResponseAAAARecordJSON contains the JSON metadata for the struct
// [RecordGetResponseAAAARecord]
type recordGetResponseAAAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseAAAARecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseCAARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseCAARecordJSON `json:"-"`
	CAARecord
}

// recordGetResponseCAARecordJSON contains the JSON metadata for the struct
// [RecordGetResponseCAARecord]
type recordGetResponseCAARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseCAARecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseCERTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordGetResponseCERTRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseCERTRecord]
type recordGetResponseCERTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseCERTRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseCNAMERecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordGetResponseCNAMERecordJSON contains the JSON metadata for the struct
// [RecordGetResponseCNAMERecord]
type recordGetResponseCNAMERecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseCNAMERecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseDNSKEYRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordGetResponseDNSKEYRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseDNSKEYRecord]
type recordGetResponseDNSKEYRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseDNSKEYRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseDSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseDSRecordJSON `json:"-"`
	DSRecord
}

// recordGetResponseDSRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseDSRecord]
type recordGetResponseDSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseDSRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseHTTPSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordGetResponseHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseHTTPSRecord]
type recordGetResponseHTTPSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseHTTPSRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseLOCRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordGetResponseLOCRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseLOCRecord]
type recordGetResponseLOCRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseLOCRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseMXRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseMXRecordJSON `json:"-"`
	MXRecord
}

// recordGetResponseMXRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseMXRecord]
type recordGetResponseMXRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseMXRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseNAPTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordGetResponseNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseNAPTRRecord]
type recordGetResponseNAPTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseNAPTRRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseNSRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseNSRecordJSON `json:"-"`
	NSRecord
}

// recordGetResponseNSRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseNSRecord]
type recordGetResponseNSRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseNSRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseOpenpgpkeyRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name string `json:"name,required"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied bool `json:"proxied,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordGetResponseOpenpgpkeyRecordType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                             `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseOpenpgpkeyRecordJSON `json:"-"`
}

// recordGetResponseOpenpgpkeyRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseOpenpgpkeyRecord]
type recordGetResponseOpenpgpkeyRecordJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseOpenpgpkeyRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseOpenpgpkeyRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseOpenpgpkeyRecord) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseOpenpgpkeyRecordType string

const (
	RecordGetResponseOpenpgpkeyRecordTypeOpenpgpkey RecordGetResponseOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordGetResponseOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordGetResponseOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordGetResponsePTRRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponsePTRRecordJSON `json:"-"`
	PTRRecord
}

// recordGetResponsePTRRecordJSON contains the JSON metadata for the struct
// [RecordGetResponsePTRRecord]
type recordGetResponsePTRRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponsePTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponsePTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponsePTRRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseSMIMEARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                         `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordGetResponseSMIMEARecordJSON contains the JSON metadata for the struct
// [RecordGetResponseSMIMEARecord]
type recordGetResponseSMIMEARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSMIMEARecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseSRVRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordGetResponseSRVRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseSRVRecord]
type recordGetResponseSRVRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSRVRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseSSHFPRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                        `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordGetResponseSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseSSHFPRecord]
type recordGetResponseSSHFPRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSSHFPRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseSVCBRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordGetResponseSVCBRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseSVCBRecord]
type recordGetResponseSVCBRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseSVCBRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseTLSARecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                       `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordGetResponseTLSARecordJSON contains the JSON metadata for the struct
// [RecordGetResponseTLSARecord]
type recordGetResponseTLSARecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseTLSARecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseTXTRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordGetResponseTXTRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseTXTRecord]
type recordGetResponseTXTRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseTXTRecord) implementsDNSRecordGetResponse() {}

type RecordGetResponseURIRecord struct {
	// Identifier
	ID string `json:"id,required"`
	// When the record was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Extra Cloudflare-specific information about the record.
	Meta interface{} `json:"meta,required"`
	// When the record was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Whether the record can be proxied by Cloudflare or not.
	Proxiable bool `json:"proxiable,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                      `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseURIRecordJSON `json:"-"`
	URIRecord
}

// recordGetResponseURIRecordJSON contains the JSON metadata for the struct
// [RecordGetResponseURIRecord]
type recordGetResponseURIRecordJSON struct {
	ID                apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Proxiable         apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponseURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseURIRecord) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseType string

const (
	RecordGetResponseTypeA          RecordGetResponseType = "A"
	RecordGetResponseTypeAAAA       RecordGetResponseType = "AAAA"
	RecordGetResponseTypeCAA        RecordGetResponseType = "CAA"
	RecordGetResponseTypeCERT       RecordGetResponseType = "CERT"
	RecordGetResponseTypeCNAME      RecordGetResponseType = "CNAME"
	RecordGetResponseTypeDNSKEY     RecordGetResponseType = "DNSKEY"
	RecordGetResponseTypeDS         RecordGetResponseType = "DS"
	RecordGetResponseTypeHTTPS      RecordGetResponseType = "HTTPS"
	RecordGetResponseTypeLOC        RecordGetResponseType = "LOC"
	RecordGetResponseTypeMX         RecordGetResponseType = "MX"
	RecordGetResponseTypeNAPTR      RecordGetResponseType = "NAPTR"
	RecordGetResponseTypeNS         RecordGetResponseType = "NS"
	RecordGetResponseTypeOpenpgpkey RecordGetResponseType = "OPENPGPKEY"
	RecordGetResponseTypePTR        RecordGetResponseType = "PTR"
	RecordGetResponseTypeSMIMEA     RecordGetResponseType = "SMIMEA"
	RecordGetResponseTypeSRV        RecordGetResponseType = "SRV"
	RecordGetResponseTypeSSHFP      RecordGetResponseType = "SSHFP"
	RecordGetResponseTypeSVCB       RecordGetResponseType = "SVCB"
	RecordGetResponseTypeTLSA       RecordGetResponseType = "TLSA"
	RecordGetResponseTypeTXT        RecordGetResponseType = "TXT"
	RecordGetResponseTypeURI        RecordGetResponseType = "URI"
)

func (r RecordGetResponseType) IsKnown() bool {
	switch r {
	case RecordGetResponseTypeA, RecordGetResponseTypeAAAA, RecordGetResponseTypeCAA, RecordGetResponseTypeCERT, RecordGetResponseTypeCNAME, RecordGetResponseTypeDNSKEY, RecordGetResponseTypeDS, RecordGetResponseTypeHTTPS, RecordGetResponseTypeLOC, RecordGetResponseTypeMX, RecordGetResponseTypeNAPTR, RecordGetResponseTypeNS, RecordGetResponseTypeOpenpgpkey, RecordGetResponseTypePTR, RecordGetResponseTypeSMIMEA, RecordGetResponseTypeSRV, RecordGetResponseTypeSSHFP, RecordGetResponseTypeSVCB, RecordGetResponseTypeTLSA, RecordGetResponseTypeTXT, RecordGetResponseTypeURI:
		return true
	}
	return false
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
	Record RecordUnionParam    `json:"record,required"`
}

func (r RecordNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Record)
}

type RecordNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordNewResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordNewResponse                `json:"result"`
	JSON    recordNewResponseEnvelopeJSON    `json:"-"`
}

// recordNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordNewResponseEnvelope]
type recordNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseEnvelopeJSON) RawJSON() string {
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
	Record RecordUnionParam    `json:"record,required"`
}

func (r RecordUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Record)
}

type RecordUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordUpdateResponse                `json:"result"`
	JSON    recordUpdateResponseEnvelopeJSON    `json:"-"`
}

// recordUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordUpdateResponseEnvelope]
type recordUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseEnvelopeJSON) RawJSON() string {
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
	Content param.Field[RecordListParamsContent] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[shared.SortDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[RecordListParamsMatch] `query:"match"`
	Name  param.Field[RecordListParamsName]  `query:"name"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RecordListParamsContent struct {
	// Substring of the DNS record content. Content filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record content. Content filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record content. Content filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS record content. Content filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [RecordListParamsContent]'s query parameters as
// `url.Values`.
func (r RecordListParamsContent) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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

type RecordListParamsName struct {
	// Substring of the DNS record name. Name filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record name. Name filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record name. Name filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS record name. Name filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [RecordListParamsName]'s query parameters as `url.Values`.
func (r RecordListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
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
	RecordListParamsTypeA          RecordListParamsType = "A"
	RecordListParamsTypeAAAA       RecordListParamsType = "AAAA"
	RecordListParamsTypeCAA        RecordListParamsType = "CAA"
	RecordListParamsTypeCERT       RecordListParamsType = "CERT"
	RecordListParamsTypeCNAME      RecordListParamsType = "CNAME"
	RecordListParamsTypeDNSKEY     RecordListParamsType = "DNSKEY"
	RecordListParamsTypeDS         RecordListParamsType = "DS"
	RecordListParamsTypeHTTPS      RecordListParamsType = "HTTPS"
	RecordListParamsTypeLOC        RecordListParamsType = "LOC"
	RecordListParamsTypeMX         RecordListParamsType = "MX"
	RecordListParamsTypeNAPTR      RecordListParamsType = "NAPTR"
	RecordListParamsTypeNS         RecordListParamsType = "NS"
	RecordListParamsTypeOpenpgpkey RecordListParamsType = "OPENPGPKEY"
	RecordListParamsTypePTR        RecordListParamsType = "PTR"
	RecordListParamsTypeSMIMEA     RecordListParamsType = "SMIMEA"
	RecordListParamsTypeSRV        RecordListParamsType = "SRV"
	RecordListParamsTypeSSHFP      RecordListParamsType = "SSHFP"
	RecordListParamsTypeSVCB       RecordListParamsType = "SVCB"
	RecordListParamsTypeTLSA       RecordListParamsType = "TLSA"
	RecordListParamsTypeTXT        RecordListParamsType = "TXT"
	RecordListParamsTypeURI        RecordListParamsType = "URI"
)

func (r RecordListParamsType) IsKnown() bool {
	switch r {
	case RecordListParamsTypeA, RecordListParamsTypeAAAA, RecordListParamsTypeCAA, RecordListParamsTypeCERT, RecordListParamsTypeCNAME, RecordListParamsTypeDNSKEY, RecordListParamsTypeDS, RecordListParamsTypeHTTPS, RecordListParamsTypeLOC, RecordListParamsTypeMX, RecordListParamsTypeNAPTR, RecordListParamsTypeNS, RecordListParamsTypeOpenpgpkey, RecordListParamsTypePTR, RecordListParamsTypeSMIMEA, RecordListParamsTypeSRV, RecordListParamsTypeSSHFP, RecordListParamsTypeSVCB, RecordListParamsTypeTLSA, RecordListParamsTypeTXT, RecordListParamsTypeURI:
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

type RecordBatchParams struct {
	// Identifier
	ZoneID  param.Field[string]                        `path:"zone_id,required"`
	Deletes param.Field[[]RecordBatchParamsDelete]     `json:"deletes"`
	Patches param.Field[[]RecordBatchParamsPatchUnion] `json:"patches"`
	Posts   param.Field[[]RecordUnionParam]            `json:"posts"`
	Puts    param.Field[[]RecordBatchParamsPutUnion]   `json:"puts"`
}

func (r RecordBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RecordBatchParamsDelete struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
}

func (r RecordBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [dns.RecordBatchParamsPatchesARecord],
// [dns.RecordBatchParamsPatchesAAAARecord],
// [dns.RecordBatchParamsPatchesCAARecord],
// [dns.RecordBatchParamsPatchesCERTRecord],
// [dns.RecordBatchParamsPatchesCNAMERecord],
// [dns.RecordBatchParamsPatchesDNSKEYRecord],
// [dns.RecordBatchParamsPatchesDSRecord],
// [dns.RecordBatchParamsPatchesHTTPSRecord],
// [dns.RecordBatchParamsPatchesLOCRecord], [dns.RecordBatchParamsPatchesMXRecord],
// [dns.RecordBatchParamsPatchesNAPTRRecord],
// [dns.RecordBatchParamsPatchesNSRecord],
// [dns.RecordBatchParamsPatchesOpenpgpkeyRecord],
// [dns.RecordBatchParamsPatchesPTRRecord],
// [dns.RecordBatchParamsPatchesSMIMEARecord],
// [dns.RecordBatchParamsPatchesSRVRecord],
// [dns.RecordBatchParamsPatchesSSHFPRecord],
// [dns.RecordBatchParamsPatchesSVCBRecord],
// [dns.RecordBatchParamsPatchesTLSARecord],
// [dns.RecordBatchParamsPatchesTXTRecord],
// [dns.RecordBatchParamsPatchesURIRecord].
type RecordBatchParamsPatchUnion interface {
	implementsDNSRecordBatchParamsPatchUnion()
}

type RecordBatchParamsPatchesARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	ARecordParam
}

func (r RecordBatchParamsPatchesARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesARecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesAAAARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	AAAARecordParam
}

func (r RecordBatchParamsPatchesAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesAAAARecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesCAARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CAARecordParam
}

func (r RecordBatchParamsPatchesCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesCAARecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesCERTRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CERTRecordParam
}

func (r RecordBatchParamsPatchesCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesCERTRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesCNAMERecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CNAMERecordParam
}

func (r RecordBatchParamsPatchesCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesCNAMERecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesDNSKEYRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	DNSKEYRecordParam
}

func (r RecordBatchParamsPatchesDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesDNSKEYRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesDSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	DSRecordParam
}

func (r RecordBatchParamsPatchesDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesDSRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesHTTPSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	HTTPSRecordParam
}

func (r RecordBatchParamsPatchesHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesHTTPSRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesLOCRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	LOCRecordParam
}

func (r RecordBatchParamsPatchesLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesLOCRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesMXRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	MXRecordParam
}

func (r RecordBatchParamsPatchesMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesMXRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesNAPTRRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	NAPTRRecordParam
}

func (r RecordBatchParamsPatchesNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesNAPTRRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesNSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	NSRecordParam
}

func (r RecordBatchParamsPatchesNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesNSRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesOpenpgpkeyRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[RecordBatchParamsPatchesOpenpgpkeyRecordType] `json:"type"`
}

func (r RecordBatchParamsPatchesOpenpgpkeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesOpenpgpkeyRecord) implementsDNSRecordBatchParamsPatchUnion() {}

// Record type.
type RecordBatchParamsPatchesOpenpgpkeyRecordType string

const (
	RecordBatchParamsPatchesOpenpgpkeyRecordTypeOpenpgpkey RecordBatchParamsPatchesOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordBatchParamsPatchesOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordBatchParamsPatchesOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordBatchParamsPatchesPTRRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	PTRRecordParam
}

func (r RecordBatchParamsPatchesPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesPTRRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesSMIMEARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SMIMEARecordParam
}

func (r RecordBatchParamsPatchesSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesSMIMEARecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesSRVRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SRVRecordParam
}

func (r RecordBatchParamsPatchesSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesSRVRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesSSHFPRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SSHFPRecordParam
}

func (r RecordBatchParamsPatchesSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesSSHFPRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesSVCBRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SVCBRecordParam
}

func (r RecordBatchParamsPatchesSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesSVCBRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesTLSARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	TLSARecordParam
}

func (r RecordBatchParamsPatchesTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesTLSARecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesTXTRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	TXTRecordParam
}

func (r RecordBatchParamsPatchesTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesTXTRecord) implementsDNSRecordBatchParamsPatchUnion() {}

type RecordBatchParamsPatchesURIRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	URIRecordParam
}

func (r RecordBatchParamsPatchesURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPatchesURIRecord) implementsDNSRecordBatchParamsPatchUnion() {}

// Satisfied by [dns.RecordBatchParamsPutsARecord],
// [dns.RecordBatchParamsPutsAAAARecord], [dns.RecordBatchParamsPutsCAARecord],
// [dns.RecordBatchParamsPutsCERTRecord], [dns.RecordBatchParamsPutsCNAMERecord],
// [dns.RecordBatchParamsPutsDNSKEYRecord], [dns.RecordBatchParamsPutsDSRecord],
// [dns.RecordBatchParamsPutsHTTPSRecord], [dns.RecordBatchParamsPutsLOCRecord],
// [dns.RecordBatchParamsPutsMXRecord], [dns.RecordBatchParamsPutsNAPTRRecord],
// [dns.RecordBatchParamsPutsNSRecord],
// [dns.RecordBatchParamsPutsOpenpgpkeyRecord],
// [dns.RecordBatchParamsPutsPTRRecord], [dns.RecordBatchParamsPutsSMIMEARecord],
// [dns.RecordBatchParamsPutsSRVRecord], [dns.RecordBatchParamsPutsSSHFPRecord],
// [dns.RecordBatchParamsPutsSVCBRecord], [dns.RecordBatchParamsPutsTLSARecord],
// [dns.RecordBatchParamsPutsTXTRecord], [dns.RecordBatchParamsPutsURIRecord].
type RecordBatchParamsPutUnion interface {
	implementsDNSRecordBatchParamsPutUnion()
}

type RecordBatchParamsPutsARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	ARecordParam
}

func (r RecordBatchParamsPutsARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsARecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsAAAARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	AAAARecordParam
}

func (r RecordBatchParamsPutsAAAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsAAAARecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsCAARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CAARecordParam
}

func (r RecordBatchParamsPutsCAARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsCAARecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsCERTRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CERTRecordParam
}

func (r RecordBatchParamsPutsCERTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsCERTRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsCNAMERecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CNAMERecordParam
}

func (r RecordBatchParamsPutsCNAMERecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsCNAMERecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsDNSKEYRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	DNSKEYRecordParam
}

func (r RecordBatchParamsPutsDNSKEYRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsDNSKEYRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsDSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	DSRecordParam
}

func (r RecordBatchParamsPutsDSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsDSRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsHTTPSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	HTTPSRecordParam
}

func (r RecordBatchParamsPutsHTTPSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsHTTPSRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsLOCRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	LOCRecordParam
}

func (r RecordBatchParamsPutsLOCRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsLOCRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsMXRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	MXRecordParam
}

func (r RecordBatchParamsPutsMXRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsMXRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsNAPTRRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	NAPTRRecordParam
}

func (r RecordBatchParamsPutsNAPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsNAPTRRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsNSRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	NSRecordParam
}

func (r RecordBatchParamsPutsNSRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsNSRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsOpenpgpkeyRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[RecordBatchParamsPutsOpenpgpkeyRecordType] `json:"type,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
}

func (r RecordBatchParamsPutsOpenpgpkeyRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsOpenpgpkeyRecord) implementsDNSRecordBatchParamsPutUnion() {}

// Record type.
type RecordBatchParamsPutsOpenpgpkeyRecordType string

const (
	RecordBatchParamsPutsOpenpgpkeyRecordTypeOpenpgpkey RecordBatchParamsPutsOpenpgpkeyRecordType = "OPENPGPKEY"
)

func (r RecordBatchParamsPutsOpenpgpkeyRecordType) IsKnown() bool {
	switch r {
	case RecordBatchParamsPutsOpenpgpkeyRecordTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordBatchParamsPutsPTRRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	PTRRecordParam
}

func (r RecordBatchParamsPutsPTRRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsPTRRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsSMIMEARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SMIMEARecordParam
}

func (r RecordBatchParamsPutsSMIMEARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsSMIMEARecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsSRVRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SRVRecordParam
}

func (r RecordBatchParamsPutsSRVRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsSRVRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsSSHFPRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SSHFPRecordParam
}

func (r RecordBatchParamsPutsSSHFPRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsSSHFPRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsSVCBRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SVCBRecordParam
}

func (r RecordBatchParamsPutsSVCBRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsSVCBRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsTLSARecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	TLSARecordParam
}

func (r RecordBatchParamsPutsTLSARecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsTLSARecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsTXTRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	TXTRecordParam
}

func (r RecordBatchParamsPutsTXTRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsTXTRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchParamsPutsURIRecord struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	URIRecordParam
}

func (r RecordBatchParamsPutsURIRecord) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordBatchParamsPutsURIRecord) implementsDNSRecordBatchParamsPutUnion() {}

type RecordBatchResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordBatchResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordBatchResponse                `json:"result"`
	JSON    recordBatchResponseEnvelopeJSON    `json:"-"`
}

// recordBatchResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordBatchResponseEnvelope]
type recordBatchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordBatchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RecordBatchResponseEnvelopeSuccess bool

const (
	RecordBatchResponseEnvelopeSuccessTrue RecordBatchResponseEnvelopeSuccess = true
)

func (r RecordBatchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RecordBatchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RecordEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Record RecordUnionParam    `json:"record,required"`
}

func (r RecordEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Record)
}

type RecordEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordEditResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordEditResponse                `json:"result"`
	JSON    recordEditResponseEnvelopeJSON    `json:"-"`
}

// recordEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordEditResponseEnvelope]
type recordEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseEnvelopeJSON) RawJSON() string {
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordGetResponse                `json:"result"`
	JSON    recordGetResponseEnvelopeJSON    `json:"-"`
}

// recordGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordGetResponseEnvelope]
type recordGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseEnvelopeJSON) RawJSON() string {
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

func (r RecordImportParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type RecordImportResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordImportResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordImportResponse                `json:"result"`
	Timing  RecordProcessTiming                 `json:"timing"`
	JSON    recordImportResponseEnvelopeJSON    `json:"-"`
}

// recordImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordImportResponseEnvelope]
type recordImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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

type RecordScanParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r RecordScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RecordScanResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RecordScanResponseEnvelopeSuccess `json:"success,required"`
	Result  RecordScanResponse                `json:"result"`
	Timing  RecordProcessTiming               `json:"timing"`
	JSON    recordScanResponseEnvelopeJSON    `json:"-"`
}

// recordScanResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordScanResponseEnvelope]
type recordScanResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
