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

	"github.com/cloudflare/cloudflare-go/v3/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Record type.
	Type ARecordType `json:"type"`
	JSON aRecordJSON `json:"-"`
}

// aRecordJSON contains the JSON metadata for the struct [ARecord]
type aRecordJSON struct {
	Content     apijson.Field
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

func (r ARecord) implementsDNSRecordNewResponse() {}

func (r ARecord) implementsDNSRecordUpdateResponse() {}

func (r ARecord) implementsDNSRecordListResponse() {}

func (r ARecord) implementsDNSRecordBatchResponseDelete() {}

func (r ARecord) implementsDNSRecordBatchResponsePatch() {}

func (r ARecord) implementsDNSRecordBatchResponsePost() {}

func (r ARecord) implementsDNSRecordBatchResponsePut() {}

func (r ARecord) implementsDNSRecordEditResponse() {}

func (r ARecord) implementsDNSRecordGetResponse() {}

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
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Record type.
	Type param.Field[ARecordType] `json:"type"`
}

func (r ARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ARecordParam) implementsDNSRecordUnionParam() {}

type AAAARecord struct {
	// A valid IPv6 address.
	Content string `json:"content" format:"ipv6"`
	// Record type.
	Type AAAARecordType `json:"type"`
	JSON aaaaRecordJSON `json:"-"`
}

// aaaaRecordJSON contains the JSON metadata for the struct [AAAARecord]
type aaaaRecordJSON struct {
	Content     apijson.Field
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

func (r AAAARecord) implementsDNSRecordNewResponse() {}

func (r AAAARecord) implementsDNSRecordUpdateResponse() {}

func (r AAAARecord) implementsDNSRecordListResponse() {}

func (r AAAARecord) implementsDNSRecordBatchResponseDelete() {}

func (r AAAARecord) implementsDNSRecordBatchResponsePatch() {}

func (r AAAARecord) implementsDNSRecordBatchResponsePost() {}

func (r AAAARecord) implementsDNSRecordBatchResponsePut() {}

func (r AAAARecord) implementsDNSRecordEditResponse() {}

func (r AAAARecord) implementsDNSRecordGetResponse() {}

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
	// A valid IPv6 address.
	Content param.Field[string] `json:"content" format:"ipv6"`
	// Record type.
	Type param.Field[AAAARecordType] `json:"type"`
}

func (r AAAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AAAARecordParam) implementsDNSRecordUnionParam() {}

type CAARecord struct {
	// Formatted CAA content. See 'data' to set CAA properties.
	Content string `json:"content"`
	// Components of a CAA record.
	Data CAARecordData `json:"data"`
	// Record type.
	Type CAARecordType `json:"type"`
	JSON caaRecordJSON `json:"-"`
}

// caaRecordJSON contains the JSON metadata for the struct [CAARecord]
type caaRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r CAARecord) implementsDNSRecordNewResponse() {}

func (r CAARecord) implementsDNSRecordUpdateResponse() {}

func (r CAARecord) implementsDNSRecordListResponse() {}

func (r CAARecord) implementsDNSRecordBatchResponseDelete() {}

func (r CAARecord) implementsDNSRecordBatchResponsePatch() {}

func (r CAARecord) implementsDNSRecordBatchResponsePost() {}

func (r CAARecord) implementsDNSRecordBatchResponsePut() {}

func (r CAARecord) implementsDNSRecordEditResponse() {}

func (r CAARecord) implementsDNSRecordGetResponse() {}

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
	// Components of a CAA record.
	Data param.Field[CAARecordDataParam] `json:"data"`
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
	// Formatted CERT content. See 'data' to set CERT properties.
	Content string `json:"content"`
	// Components of a CERT record.
	Data CERTRecordData `json:"data"`
	// Record type.
	Type CERTRecordType `json:"type"`
	JSON certRecordJSON `json:"-"`
}

// certRecordJSON contains the JSON metadata for the struct [CERTRecord]
type certRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r CERTRecord) implementsDNSRecordNewResponse() {}

func (r CERTRecord) implementsDNSRecordUpdateResponse() {}

func (r CERTRecord) implementsDNSRecordListResponse() {}

func (r CERTRecord) implementsDNSRecordBatchResponseDelete() {}

func (r CERTRecord) implementsDNSRecordBatchResponsePatch() {}

func (r CERTRecord) implementsDNSRecordBatchResponsePost() {}

func (r CERTRecord) implementsDNSRecordBatchResponsePut() {}

func (r CERTRecord) implementsDNSRecordEditResponse() {}

func (r CERTRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a CERT record.
	Data param.Field[CERTRecordDataParam] `json:"data"`
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
	// A valid hostname. Must not match the record's name.
	Content  string              `json:"content"`
	Settings CNAMERecordSettings `json:"settings"`
	// Record type.
	Type CNAMERecordType `json:"type"`
	JSON cnameRecordJSON `json:"-"`
}

// cnameRecordJSON contains the JSON metadata for the struct [CNAMERecord]
type cnameRecordJSON struct {
	Content     apijson.Field
	Settings    apijson.Field
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

func (r CNAMERecord) implementsDNSRecordNewResponse() {}

func (r CNAMERecord) implementsDNSRecordUpdateResponse() {}

func (r CNAMERecord) implementsDNSRecordListResponse() {}

func (r CNAMERecord) implementsDNSRecordBatchResponseDelete() {}

func (r CNAMERecord) implementsDNSRecordBatchResponsePatch() {}

func (r CNAMERecord) implementsDNSRecordBatchResponsePost() {}

func (r CNAMERecord) implementsDNSRecordBatchResponsePut() {}

func (r CNAMERecord) implementsDNSRecordEditResponse() {}

func (r CNAMERecord) implementsDNSRecordGetResponse() {}

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
	// A valid hostname. Must not match the record's name.
	Content  param.Field[string]                   `json:"content"`
	Settings param.Field[CNAMERecordSettingsParam] `json:"settings"`
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
	// Formatted DNSKEY content. See 'data' to set DNSKEY properties.
	Content string `json:"content"`
	// Components of a DNSKEY record.
	Data DNSKEYRecordData `json:"data"`
	// Record type.
	Type DNSKEYRecordType `json:"type"`
	JSON dnskeyRecordJSON `json:"-"`
}

// dnskeyRecordJSON contains the JSON metadata for the struct [DNSKEYRecord]
type dnskeyRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r DNSKEYRecord) implementsDNSRecordNewResponse() {}

func (r DNSKEYRecord) implementsDNSRecordUpdateResponse() {}

func (r DNSKEYRecord) implementsDNSRecordListResponse() {}

func (r DNSKEYRecord) implementsDNSRecordBatchResponseDelete() {}

func (r DNSKEYRecord) implementsDNSRecordBatchResponsePatch() {}

func (r DNSKEYRecord) implementsDNSRecordBatchResponsePost() {}

func (r DNSKEYRecord) implementsDNSRecordBatchResponsePut() {}

func (r DNSKEYRecord) implementsDNSRecordEditResponse() {}

func (r DNSKEYRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a DNSKEY record.
	Data param.Field[DNSKEYRecordDataParam] `json:"data"`
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
	// Formatted DS content. See 'data' to set DS properties.
	Content string `json:"content"`
	// Components of a DS record.
	Data DSRecordData `json:"data"`
	// Record type.
	Type DSRecordType `json:"type"`
	JSON dsRecordJSON `json:"-"`
}

// dsRecordJSON contains the JSON metadata for the struct [DSRecord]
type dsRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r DSRecord) implementsDNSRecordNewResponse() {}

func (r DSRecord) implementsDNSRecordUpdateResponse() {}

func (r DSRecord) implementsDNSRecordListResponse() {}

func (r DSRecord) implementsDNSRecordBatchResponseDelete() {}

func (r DSRecord) implementsDNSRecordBatchResponsePatch() {}

func (r DSRecord) implementsDNSRecordBatchResponsePost() {}

func (r DSRecord) implementsDNSRecordBatchResponsePut() {}

func (r DSRecord) implementsDNSRecordEditResponse() {}

func (r DSRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a DS record.
	Data param.Field[DSRecordDataParam] `json:"data"`
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
	// Formatted HTTPS content. See 'data' to set HTTPS properties.
	Content string `json:"content"`
	// Components of a HTTPS record.
	Data HTTPSRecordData `json:"data"`
	// Record type.
	Type HTTPSRecordType `json:"type"`
	JSON httpsRecordJSON `json:"-"`
}

// httpsRecordJSON contains the JSON metadata for the struct [HTTPSRecord]
type httpsRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r HTTPSRecord) implementsDNSRecordNewResponse() {}

func (r HTTPSRecord) implementsDNSRecordUpdateResponse() {}

func (r HTTPSRecord) implementsDNSRecordListResponse() {}

func (r HTTPSRecord) implementsDNSRecordBatchResponseDelete() {}

func (r HTTPSRecord) implementsDNSRecordBatchResponsePatch() {}

func (r HTTPSRecord) implementsDNSRecordBatchResponsePost() {}

func (r HTTPSRecord) implementsDNSRecordBatchResponsePut() {}

func (r HTTPSRecord) implementsDNSRecordEditResponse() {}

func (r HTTPSRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a HTTPS record.
	Data param.Field[HTTPSRecordDataParam] `json:"data"`
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
	// Formatted LOC content. See 'data' to set LOC properties.
	Content string `json:"content"`
	// Components of a LOC record.
	Data LOCRecordData `json:"data"`
	// Record type.
	Type LOCRecordType `json:"type"`
	JSON locRecordJSON `json:"-"`
}

// locRecordJSON contains the JSON metadata for the struct [LOCRecord]
type locRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r LOCRecord) implementsDNSRecordNewResponse() {}

func (r LOCRecord) implementsDNSRecordUpdateResponse() {}

func (r LOCRecord) implementsDNSRecordListResponse() {}

func (r LOCRecord) implementsDNSRecordBatchResponseDelete() {}

func (r LOCRecord) implementsDNSRecordBatchResponsePatch() {}

func (r LOCRecord) implementsDNSRecordBatchResponsePost() {}

func (r LOCRecord) implementsDNSRecordBatchResponsePut() {}

func (r LOCRecord) implementsDNSRecordEditResponse() {}

func (r LOCRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a LOC record.
	Data param.Field[LOCRecordDataParam] `json:"data"`
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
	// A valid mail server hostname.
	Content string `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type MXRecordType `json:"type"`
	JSON mxRecordJSON `json:"-"`
}

// mxRecordJSON contains the JSON metadata for the struct [MXRecord]
type mxRecordJSON struct {
	Content     apijson.Field
	Priority    apijson.Field
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

func (r MXRecord) implementsDNSRecordNewResponse() {}

func (r MXRecord) implementsDNSRecordUpdateResponse() {}

func (r MXRecord) implementsDNSRecordListResponse() {}

func (r MXRecord) implementsDNSRecordBatchResponseDelete() {}

func (r MXRecord) implementsDNSRecordBatchResponsePatch() {}

func (r MXRecord) implementsDNSRecordBatchResponsePost() {}

func (r MXRecord) implementsDNSRecordBatchResponsePut() {}

func (r MXRecord) implementsDNSRecordEditResponse() {}

func (r MXRecord) implementsDNSRecordGetResponse() {}

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
	// A valid mail server hostname.
	Content param.Field[string] `json:"content" format:"hostname"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
	// Record type.
	Type param.Field[MXRecordType] `json:"type"`
}

func (r MXRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MXRecordParam) implementsDNSRecordUnionParam() {}

type NAPTRRecord struct {
	// Formatted NAPTR content. See 'data' to set NAPTR properties.
	Content string `json:"content"`
	// Components of a NAPTR record.
	Data NAPTRRecordData `json:"data"`
	// Record type.
	Type NAPTRRecordType `json:"type"`
	JSON naptrRecordJSON `json:"-"`
}

// naptrRecordJSON contains the JSON metadata for the struct [NAPTRRecord]
type naptrRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r NAPTRRecord) implementsDNSRecordNewResponse() {}

func (r NAPTRRecord) implementsDNSRecordUpdateResponse() {}

func (r NAPTRRecord) implementsDNSRecordListResponse() {}

func (r NAPTRRecord) implementsDNSRecordBatchResponseDelete() {}

func (r NAPTRRecord) implementsDNSRecordBatchResponsePatch() {}

func (r NAPTRRecord) implementsDNSRecordBatchResponsePost() {}

func (r NAPTRRecord) implementsDNSRecordBatchResponsePut() {}

func (r NAPTRRecord) implementsDNSRecordEditResponse() {}

func (r NAPTRRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a NAPTR record.
	Data param.Field[NAPTRRecordDataParam] `json:"data"`
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
	// A valid name server host name.
	Content string `json:"content"`
	// Record type.
	Type NSRecordType `json:"type"`
	JSON nsRecordJSON `json:"-"`
}

// nsRecordJSON contains the JSON metadata for the struct [NSRecord]
type nsRecordJSON struct {
	Content     apijson.Field
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

func (r NSRecord) implementsDNSRecordNewResponse() {}

func (r NSRecord) implementsDNSRecordUpdateResponse() {}

func (r NSRecord) implementsDNSRecordListResponse() {}

func (r NSRecord) implementsDNSRecordBatchResponseDelete() {}

func (r NSRecord) implementsDNSRecordBatchResponsePatch() {}

func (r NSRecord) implementsDNSRecordBatchResponsePost() {}

func (r NSRecord) implementsDNSRecordBatchResponsePut() {}

func (r NSRecord) implementsDNSRecordEditResponse() {}

func (r NSRecord) implementsDNSRecordGetResponse() {}

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
	// A valid name server host name.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[NSRecordType] `json:"type"`
}

func (r NSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NSRecordParam) implementsDNSRecordUnionParam() {}

type PTRRecord struct {
	// Domain name pointing to the address.
	Content string `json:"content"`
	// Record type.
	Type PTRRecordType `json:"type"`
	JSON ptrRecordJSON `json:"-"`
}

// ptrRecordJSON contains the JSON metadata for the struct [PTRRecord]
type ptrRecordJSON struct {
	Content     apijson.Field
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

func (r PTRRecord) implementsDNSRecordNewResponse() {}

func (r PTRRecord) implementsDNSRecordUpdateResponse() {}

func (r PTRRecord) implementsDNSRecordListResponse() {}

func (r PTRRecord) implementsDNSRecordBatchResponseDelete() {}

func (r PTRRecord) implementsDNSRecordBatchResponsePatch() {}

func (r PTRRecord) implementsDNSRecordBatchResponsePost() {}

func (r PTRRecord) implementsDNSRecordBatchResponsePut() {}

func (r PTRRecord) implementsDNSRecordEditResponse() {}

func (r PTRRecord) implementsDNSRecordGetResponse() {}

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
	// Domain name pointing to the address.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[PTRRecordType] `json:"type"`
}

func (r PTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PTRRecordParam) implementsDNSRecordUnionParam() {}

type RecordParam struct {
	Data     param.Field[interface{}] `json:"data,required"`
	Settings param.Field[interface{}] `json:"settings,required"`
	// A valid IPv4 address.
	Content param.Field[string] `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
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
// [dns.RecordOpenpgpkeyParam], [dns.PTRRecordParam], [dns.SMIMEARecordParam],
// [dns.SRVRecordParam], [dns.SSHFPRecordParam], [dns.SVCBRecordParam],
// [dns.TLSARecordParam], [dns.TXTRecordParam], [dns.URIRecordParam],
// [RecordParam].
type RecordUnionParam interface {
	implementsDNSRecordUnionParam()
}

type RecordOpenpgpkeyParam struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[RecordOpenpgpkeyType] `json:"type"`
}

func (r RecordOpenpgpkeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordOpenpgpkeyParam) implementsDNSRecordUnionParam() {}

// Record type.
type RecordOpenpgpkeyType string

const (
	RecordOpenpgpkeyTypeOpenpgpkey RecordOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordOpenpgpkeyTypeOpenpgpkey:
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

type SMIMEARecord struct {
	// Formatted SMIMEA content. See 'data' to set SMIMEA properties.
	Content string `json:"content"`
	// Components of a SMIMEA record.
	Data SMIMEARecordData `json:"data"`
	// Record type.
	Type SMIMEARecordType `json:"type"`
	JSON smimeaRecordJSON `json:"-"`
}

// smimeaRecordJSON contains the JSON metadata for the struct [SMIMEARecord]
type smimeaRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r SMIMEARecord) implementsDNSRecordNewResponse() {}

func (r SMIMEARecord) implementsDNSRecordUpdateResponse() {}

func (r SMIMEARecord) implementsDNSRecordListResponse() {}

func (r SMIMEARecord) implementsDNSRecordBatchResponseDelete() {}

func (r SMIMEARecord) implementsDNSRecordBatchResponsePatch() {}

func (r SMIMEARecord) implementsDNSRecordBatchResponsePost() {}

func (r SMIMEARecord) implementsDNSRecordBatchResponsePut() {}

func (r SMIMEARecord) implementsDNSRecordEditResponse() {}

func (r SMIMEARecord) implementsDNSRecordGetResponse() {}

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
	// Components of a SMIMEA record.
	Data param.Field[SMIMEARecordDataParam] `json:"data"`
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
	// Priority, weight, port, and SRV target. See 'data' for setting the individual
	// component values.
	Content string `json:"content"`
	// Components of a SRV record.
	Data SRVRecordData `json:"data"`
	// Record type.
	Type SRVRecordType `json:"type"`
	JSON srvRecordJSON `json:"-"`
}

// srvRecordJSON contains the JSON metadata for the struct [SRVRecord]
type srvRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r SRVRecord) implementsDNSRecordNewResponse() {}

func (r SRVRecord) implementsDNSRecordUpdateResponse() {}

func (r SRVRecord) implementsDNSRecordListResponse() {}

func (r SRVRecord) implementsDNSRecordBatchResponseDelete() {}

func (r SRVRecord) implementsDNSRecordBatchResponsePatch() {}

func (r SRVRecord) implementsDNSRecordBatchResponsePost() {}

func (r SRVRecord) implementsDNSRecordBatchResponsePut() {}

func (r SRVRecord) implementsDNSRecordEditResponse() {}

func (r SRVRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a SRV record.
	Data param.Field[SRVRecordDataParam] `json:"data"`
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
	// Formatted SSHFP content. See 'data' to set SSHFP properties.
	Content string `json:"content"`
	// Components of a SSHFP record.
	Data SSHFPRecordData `json:"data"`
	// Record type.
	Type SSHFPRecordType `json:"type"`
	JSON sshfpRecordJSON `json:"-"`
}

// sshfpRecordJSON contains the JSON metadata for the struct [SSHFPRecord]
type sshfpRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r SSHFPRecord) implementsDNSRecordNewResponse() {}

func (r SSHFPRecord) implementsDNSRecordUpdateResponse() {}

func (r SSHFPRecord) implementsDNSRecordListResponse() {}

func (r SSHFPRecord) implementsDNSRecordBatchResponseDelete() {}

func (r SSHFPRecord) implementsDNSRecordBatchResponsePatch() {}

func (r SSHFPRecord) implementsDNSRecordBatchResponsePost() {}

func (r SSHFPRecord) implementsDNSRecordBatchResponsePut() {}

func (r SSHFPRecord) implementsDNSRecordEditResponse() {}

func (r SSHFPRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a SSHFP record.
	Data param.Field[SSHFPRecordDataParam] `json:"data"`
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
	// Formatted SVCB content. See 'data' to set SVCB properties.
	Content string `json:"content"`
	// Components of a SVCB record.
	Data SVCBRecordData `json:"data"`
	// Record type.
	Type SVCBRecordType `json:"type"`
	JSON svcbRecordJSON `json:"-"`
}

// svcbRecordJSON contains the JSON metadata for the struct [SVCBRecord]
type svcbRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r SVCBRecord) implementsDNSRecordNewResponse() {}

func (r SVCBRecord) implementsDNSRecordUpdateResponse() {}

func (r SVCBRecord) implementsDNSRecordListResponse() {}

func (r SVCBRecord) implementsDNSRecordBatchResponseDelete() {}

func (r SVCBRecord) implementsDNSRecordBatchResponsePatch() {}

func (r SVCBRecord) implementsDNSRecordBatchResponsePost() {}

func (r SVCBRecord) implementsDNSRecordBatchResponsePut() {}

func (r SVCBRecord) implementsDNSRecordEditResponse() {}

func (r SVCBRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a SVCB record.
	Data param.Field[SVCBRecordDataParam] `json:"data"`
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
	// Formatted TLSA content. See 'data' to set TLSA properties.
	Content string `json:"content"`
	// Components of a TLSA record.
	Data TLSARecordData `json:"data"`
	// Record type.
	Type TLSARecordType `json:"type"`
	JSON tlsaRecordJSON `json:"-"`
}

// tlsaRecordJSON contains the JSON metadata for the struct [TLSARecord]
type tlsaRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
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

func (r TLSARecord) implementsDNSRecordNewResponse() {}

func (r TLSARecord) implementsDNSRecordUpdateResponse() {}

func (r TLSARecord) implementsDNSRecordListResponse() {}

func (r TLSARecord) implementsDNSRecordBatchResponseDelete() {}

func (r TLSARecord) implementsDNSRecordBatchResponsePatch() {}

func (r TLSARecord) implementsDNSRecordBatchResponsePost() {}

func (r TLSARecord) implementsDNSRecordBatchResponsePut() {}

func (r TLSARecord) implementsDNSRecordEditResponse() {}

func (r TLSARecord) implementsDNSRecordGetResponse() {}

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
	// Components of a TLSA record.
	Data param.Field[TLSARecordDataParam] `json:"data"`
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

type TXTRecord struct {
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content string `json:"content"`
	// Record type.
	Type TXTRecordType `json:"type"`
	JSON txtRecordJSON `json:"-"`
}

// txtRecordJSON contains the JSON metadata for the struct [TXTRecord]
type txtRecordJSON struct {
	Content     apijson.Field
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

func (r TXTRecord) implementsDNSRecordNewResponse() {}

func (r TXTRecord) implementsDNSRecordUpdateResponse() {}

func (r TXTRecord) implementsDNSRecordListResponse() {}

func (r TXTRecord) implementsDNSRecordBatchResponseDelete() {}

func (r TXTRecord) implementsDNSRecordBatchResponsePatch() {}

func (r TXTRecord) implementsDNSRecordBatchResponsePost() {}

func (r TXTRecord) implementsDNSRecordBatchResponsePut() {}

func (r TXTRecord) implementsDNSRecordEditResponse() {}

func (r TXTRecord) implementsDNSRecordGetResponse() {}

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
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	//
	// Learn more at
	// <https://www.cloudflare.com/learning/dns/dns-records/dns-txt-record/>.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[TXTRecordType] `json:"type"`
}

func (r TXTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TXTRecordParam) implementsDNSRecordUnionParam() {}

type URIRecord struct {
	// Formatted URI content. See 'data' to set URI properties.
	Content string `json:"content"`
	// Components of a URI record.
	Data URIRecordData `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type URIRecordType `json:"type"`
	JSON uriRecordJSON `json:"-"`
}

// uriRecordJSON contains the JSON metadata for the struct [URIRecord]
type uriRecordJSON struct {
	Content     apijson.Field
	Data        apijson.Field
	Priority    apijson.Field
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

func (r URIRecord) implementsDNSRecordNewResponse() {}

func (r URIRecord) implementsDNSRecordUpdateResponse() {}

func (r URIRecord) implementsDNSRecordListResponse() {}

func (r URIRecord) implementsDNSRecordBatchResponseDelete() {}

func (r URIRecord) implementsDNSRecordBatchResponsePatch() {}

func (r URIRecord) implementsDNSRecordBatchResponsePost() {}

func (r URIRecord) implementsDNSRecordBatchResponsePut() {}

func (r URIRecord) implementsDNSRecordEditResponse() {}

func (r URIRecord) implementsDNSRecordGetResponse() {}

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
	// Components of a URI record.
	Data param.Field[URIRecordDataParam] `json:"data"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority param.Field[float64] `json:"priority"`
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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordNewResponseType `json:"type"`
	JSON  recordNewResponseJSON `json:"-"`
	union RecordNewResponseUnion
}

// recordNewResponseJSON contains the JSON metadata for the struct
// [RecordNewResponse]
type recordNewResponseJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordNewResponseOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordNewResponse) AsUnion() RecordNewResponseUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordNewResponseOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordNewResponseUnion interface {
	implementsDNSRecordNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordNewResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordNewResponseOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordNewResponseOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordNewResponseOpenpgpkeyType `json:"type"`
	JSON recordNewResponseOpenpgpkeyJSON `json:"-"`
}

// recordNewResponseOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordNewResponseOpenpgpkey]
type recordNewResponseOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordNewResponseOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordNewResponseOpenpgpkey) implementsDNSRecordNewResponse() {}

// Record type.
type RecordNewResponseOpenpgpkeyType string

const (
	RecordNewResponseOpenpgpkeyTypeOpenpgpkey RecordNewResponseOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordNewResponseOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordNewResponseOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordUpdateResponseType `json:"type"`
	JSON  recordUpdateResponseJSON `json:"-"`
	union RecordUpdateResponseUnion
}

// recordUpdateResponseJSON contains the JSON metadata for the struct
// [RecordUpdateResponse]
type recordUpdateResponseJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordUpdateResponseOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordUpdateResponse) AsUnion() RecordUpdateResponseUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordUpdateResponseOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordUpdateResponseUnion interface {
	implementsDNSRecordUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordUpdateResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordUpdateResponseOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordUpdateResponseOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordUpdateResponseOpenpgpkeyType `json:"type"`
	JSON recordUpdateResponseOpenpgpkeyJSON `json:"-"`
}

// recordUpdateResponseOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordUpdateResponseOpenpgpkey]
type recordUpdateResponseOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordUpdateResponseOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordUpdateResponseOpenpgpkey) implementsDNSRecordUpdateResponse() {}

// Record type.
type RecordUpdateResponseOpenpgpkeyType string

const (
	RecordUpdateResponseOpenpgpkeyTypeOpenpgpkey RecordUpdateResponseOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordUpdateResponseOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordUpdateResponseOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordListResponseType `json:"type"`
	JSON  recordListResponseJSON `json:"-"`
	union RecordListResponseUnion
}

// recordListResponseJSON contains the JSON metadata for the struct
// [RecordListResponse]
type recordListResponseJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordListResponseOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordListResponse) AsUnion() RecordListResponseUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordListResponseOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordListResponseUnion interface {
	implementsDNSRecordListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordListResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordListResponseOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordListResponseOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordListResponseOpenpgpkeyType `json:"type"`
	JSON recordListResponseOpenpgpkeyJSON `json:"-"`
}

// recordListResponseOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordListResponseOpenpgpkey]
type recordListResponseOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordListResponseOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordListResponseOpenpgpkey) implementsDNSRecordListResponse() {}

// Record type.
type RecordListResponseOpenpgpkeyType string

const (
	RecordListResponseOpenpgpkeyTypeOpenpgpkey RecordListResponseOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordListResponseOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordListResponseOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordBatchResponseDeletesType `json:"type"`
	JSON  recordBatchResponseDeleteJSON  `json:"-"`
	union RecordBatchResponseDeletesUnion
}

// recordBatchResponseDeleteJSON contains the JSON metadata for the struct
// [RecordBatchResponseDelete]
type recordBatchResponseDeleteJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordBatchResponseDeletesOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordBatchResponseDelete) AsUnion() RecordBatchResponseDeletesUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordBatchResponseDeletesOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordBatchResponseDeletesUnion interface {
	implementsDNSRecordBatchResponseDelete()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponseDeletesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordBatchResponseDeletesOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordBatchResponseDeletesOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordBatchResponseDeletesOpenpgpkeyType `json:"type"`
	JSON recordBatchResponseDeletesOpenpgpkeyJSON `json:"-"`
}

// recordBatchResponseDeletesOpenpgpkeyJSON contains the JSON metadata for the
// struct [RecordBatchResponseDeletesOpenpgpkey]
type recordBatchResponseDeletesOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordBatchResponseDeletesOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeletesOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponseDeletesOpenpgpkey) implementsDNSRecordBatchResponseDelete() {}

// Record type.
type RecordBatchResponseDeletesOpenpgpkeyType string

const (
	RecordBatchResponseDeletesOpenpgpkeyTypeOpenpgpkey RecordBatchResponseDeletesOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordBatchResponseDeletesOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordBatchResponseDeletesOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordBatchResponsePatchesType `json:"type"`
	JSON  recordBatchResponsePatchJSON   `json:"-"`
	union RecordBatchResponsePatchesUnion
}

// recordBatchResponsePatchJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatch]
type recordBatchResponsePatchJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordBatchResponsePatchesOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordBatchResponsePatch) AsUnion() RecordBatchResponsePatchesUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordBatchResponsePatchesOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordBatchResponsePatchesUnion interface {
	implementsDNSRecordBatchResponsePatch()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponsePatchesUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordBatchResponsePatchesOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordBatchResponsePatchesOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordBatchResponsePatchesOpenpgpkeyType `json:"type"`
	JSON recordBatchResponsePatchesOpenpgpkeyJSON `json:"-"`
}

// recordBatchResponsePatchesOpenpgpkeyJSON contains the JSON metadata for the
// struct [RecordBatchResponsePatchesOpenpgpkey]
type recordBatchResponsePatchesOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordBatchResponsePatchesOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchesOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePatchesOpenpgpkey) implementsDNSRecordBatchResponsePatch() {}

// Record type.
type RecordBatchResponsePatchesOpenpgpkeyType string

const (
	RecordBatchResponsePatchesOpenpgpkeyTypeOpenpgpkey RecordBatchResponsePatchesOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordBatchResponsePatchesOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePatchesOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordBatchResponsePostsType `json:"type"`
	JSON  recordBatchResponsePostJSON  `json:"-"`
	union RecordBatchResponsePostsUnion
}

// recordBatchResponsePostJSON contains the JSON metadata for the struct
// [RecordBatchResponsePost]
type recordBatchResponsePostJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordBatchResponsePostsOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordBatchResponsePost) AsUnion() RecordBatchResponsePostsUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordBatchResponsePostsOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordBatchResponsePostsUnion interface {
	implementsDNSRecordBatchResponsePost()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponsePostsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordBatchResponsePostsOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordBatchResponsePostsOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordBatchResponsePostsOpenpgpkeyType `json:"type"`
	JSON recordBatchResponsePostsOpenpgpkeyJSON `json:"-"`
}

// recordBatchResponsePostsOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordBatchResponsePostsOpenpgpkey]
type recordBatchResponsePostsOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordBatchResponsePostsOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostsOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePostsOpenpgpkey) implementsDNSRecordBatchResponsePost() {}

// Record type.
type RecordBatchResponsePostsOpenpgpkeyType string

const (
	RecordBatchResponsePostsOpenpgpkeyTypeOpenpgpkey RecordBatchResponsePostsOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordBatchResponsePostsOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePostsOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordBatchResponsePutsType `json:"type"`
	JSON  recordBatchResponsePutJSON  `json:"-"`
	union RecordBatchResponsePutsUnion
}

// recordBatchResponsePutJSON contains the JSON metadata for the struct
// [RecordBatchResponsePut]
type recordBatchResponsePutJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordBatchResponsePutsOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordBatchResponsePut) AsUnion() RecordBatchResponsePutsUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordBatchResponsePutsOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordBatchResponsePutsUnion interface {
	implementsDNSRecordBatchResponsePut()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordBatchResponsePutsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordBatchResponsePutsOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordBatchResponsePutsOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordBatchResponsePutsOpenpgpkeyType `json:"type"`
	JSON recordBatchResponsePutsOpenpgpkeyJSON `json:"-"`
}

// recordBatchResponsePutsOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordBatchResponsePutsOpenpgpkey]
type recordBatchResponsePutsOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordBatchResponsePutsOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutsOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordBatchResponsePutsOpenpgpkey) implementsDNSRecordBatchResponsePut() {}

// Record type.
type RecordBatchResponsePutsOpenpgpkeyType string

const (
	RecordBatchResponsePutsOpenpgpkeyTypeOpenpgpkey RecordBatchResponsePutsOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordBatchResponsePutsOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordBatchResponsePutsOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordEditResponseType `json:"type"`
	JSON  recordEditResponseJSON `json:"-"`
	union RecordEditResponseUnion
}

// recordEditResponseJSON contains the JSON metadata for the struct
// [RecordEditResponse]
type recordEditResponseJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordEditResponseOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordEditResponse) AsUnion() RecordEditResponseUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordEditResponseOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordEditResponseUnion interface {
	implementsDNSRecordEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordEditResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordEditResponseOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordEditResponseOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordEditResponseOpenpgpkeyType `json:"type"`
	JSON recordEditResponseOpenpgpkeyJSON `json:"-"`
}

// recordEditResponseOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordEditResponseOpenpgpkey]
type recordEditResponseOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordEditResponseOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordEditResponseOpenpgpkey) implementsDNSRecordEditResponse() {}

// Record type.
type RecordEditResponseOpenpgpkeyType string

const (
	RecordEditResponseOpenpgpkeyTypeOpenpgpkey RecordEditResponseOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordEditResponseOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordEditResponseOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	// This field can have the runtime type of [CAARecordData], [CERTRecordData],
	// [DNSKEYRecordData], [DSRecordData], [HTTPSRecordData], [LOCRecordData],
	// [NAPTRRecordData], [SMIMEARecordData], [SRVRecordData], [SSHFPRecordData],
	// [SVCBRecordData], [TLSARecordData], [URIRecordData].
	Data interface{} `json:"data,required"`
	// This field can have the runtime type of [CNAMERecordSettings].
	Settings interface{} `json:"settings,required"`
	// A valid IPv4 address.
	Content string `json:"content" format:"ipv4"`
	// Required for MX, SRV and URI records; unused by other record types. Records with
	// lower priorities are preferred.
	Priority float64 `json:"priority"`
	// Record type.
	Type  RecordGetResponseType `json:"type"`
	JSON  recordGetResponseJSON `json:"-"`
	union RecordGetResponseUnion
}

// recordGetResponseJSON contains the JSON metadata for the struct
// [RecordGetResponse]
type recordGetResponseJSON struct {
	Data        apijson.Field
	Settings    apijson.Field
	Content     apijson.Field
	Priority    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
// Possible runtime types of the union are [dns.ARecord], [dns.AAAARecord],
// [dns.CAARecord], [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord],
// [dns.DSRecord], [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord],
// [dns.NAPTRRecord], [dns.NSRecord], [dns.RecordGetResponseOpenpgpkey],
// [dns.PTRRecord], [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord],
// [dns.SVCBRecord], [dns.TLSARecord], [dns.TXTRecord], [dns.URIRecord].
func (r RecordGetResponse) AsUnion() RecordGetResponseUnion {
	return r.union
}

// Union satisfied by [dns.ARecord], [dns.AAAARecord], [dns.CAARecord],
// [dns.CERTRecord], [dns.CNAMERecord], [dns.DNSKEYRecord], [dns.DSRecord],
// [dns.HTTPSRecord], [dns.LOCRecord], [dns.MXRecord], [dns.NAPTRRecord],
// [dns.NSRecord], [dns.RecordGetResponseOpenpgpkey], [dns.PTRRecord],
// [dns.SMIMEARecord], [dns.SRVRecord], [dns.SSHFPRecord], [dns.SVCBRecord],
// [dns.TLSARecord], [dns.TXTRecord] or [dns.URIRecord].
type RecordGetResponseUnion interface {
	implementsDNSRecordGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordGetResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(AAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(HTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(NSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordGetResponseOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(TXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(URIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordGetResponseOpenpgpkey struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content string `json:"content"`
	// Record type.
	Type RecordGetResponseOpenpgpkeyType `json:"type"`
	JSON recordGetResponseOpenpgpkeyJSON `json:"-"`
}

// recordGetResponseOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordGetResponseOpenpgpkey]
type recordGetResponseOpenpgpkeyJSON struct {
	Content     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordGetResponseOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordGetResponseOpenpgpkey) implementsDNSRecordGetResponse() {}

// Record type.
type RecordGetResponseOpenpgpkeyType string

const (
	RecordGetResponseOpenpgpkeyTypeOpenpgpkey RecordGetResponseOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordGetResponseOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordGetResponseOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

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
	ZoneID  param.Field[string]                    `path:"zone_id,required"`
	Deletes param.Field[[]RecordBatchParamsDelete] `json:"deletes"`
	Patches param.Field[[]RecordUnionParam]        `json:"patches"`
	Posts   param.Field[[]RecordUnionParam]        `json:"posts"`
	Puts    param.Field[[]RecordUnionParam]        `json:"puts"`
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
