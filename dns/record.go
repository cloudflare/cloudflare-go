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
func (r *RecordService) New(ctx context.Context, params RecordNewParams, opts ...option.RequestOption) (res *RecordResponse, err error) {
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
func (r *RecordService) Update(ctx context.Context, dnsRecordID string, params RecordUpdateParams, opts ...option.RequestOption) (res *RecordResponse, err error) {
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
func (r *RecordService) List(ctx context.Context, params RecordListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RecordResponse], err error) {
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
func (r *RecordService) ListAutoPaging(ctx context.Context, params RecordListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RecordResponse] {
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
func (r *RecordService) Edit(ctx context.Context, dnsRecordID string, params RecordEditParams, opts ...option.RequestOption) (res *RecordResponse, err error) {
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
func (r *RecordService) Get(ctx context.Context, dnsRecordID string, query RecordGetParams, opts ...option.RequestOption) (res *RecordResponse, err error) {
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
	// Settings for the DNS record.
	Settings ARecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type ARecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                `json:"ipv6_only"`
	JSON     aRecordSettingsJSON `json:"-"`
}

// aRecordSettingsJSON contains the JSON metadata for the struct [ARecordSettings]
type aRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ARecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[ARecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type ARecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r ARecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// Settings for the DNS record.
	Settings AAAARecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type AAAARecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                   `json:"ipv6_only"`
	JSON     aaaaRecordSettingsJSON `json:"-"`
}

// aaaaRecordSettingsJSON contains the JSON metadata for the struct
// [AAAARecordSettings]
type aaaaRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AAAARecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaaRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[AAAARecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type AAAARecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r AAAARecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [dns.BatchPatchARecordParam], [dns.BatchPatchAAAARecordParam],
// [dns.BatchPatchCAARecordParam], [dns.BatchPatchCERTRecordParam],
// [dns.BatchPatchCNAMERecordParam], [dns.BatchPatchDNSKEYRecordParam],
// [dns.BatchPatchDSRecordParam], [dns.BatchPatchHTTPSRecordParam],
// [dns.BatchPatchLOCRecordParam], [dns.BatchPatchMXRecordParam],
// [dns.BatchPatchNAPTRRecordParam], [dns.BatchPatchNSRecordParam],
// [dns.BatchPatchOpenpgpkeyParam], [dns.BatchPatchPTRRecordParam],
// [dns.BatchPatchSMIMEARecordParam], [dns.BatchPatchSRVRecordParam],
// [dns.BatchPatchSSHFPRecordParam], [dns.BatchPatchSVCBRecordParam],
// [dns.BatchPatchTLSARecordParam], [dns.BatchPatchTXTRecordParam],
// [dns.BatchPatchURIRecordParam].
type BatchPatchUnionParam interface {
	implementsDNSBatchPatchUnionParam()
}

type BatchPatchARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	ARecordParam
}

func (r BatchPatchARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchARecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchAAAARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	AAAARecordParam
}

func (r BatchPatchAAAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchAAAARecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchCAARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CAARecordParam
}

func (r BatchPatchCAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchCAARecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchCERTRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CERTRecordParam
}

func (r BatchPatchCERTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchCERTRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchCNAMERecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	CNAMERecordParam
}

func (r BatchPatchCNAMERecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchCNAMERecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchDNSKEYRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	DNSKEYRecordParam
}

func (r BatchPatchDNSKEYRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchDNSKEYRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchDSRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	DSRecordParam
}

func (r BatchPatchDSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchDSRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchHTTPSRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	HTTPSRecordParam
}

func (r BatchPatchHTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchHTTPSRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchLOCRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	LOCRecordParam
}

func (r BatchPatchLOCRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchLOCRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchMXRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	MXRecordParam
}

func (r BatchPatchMXRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchMXRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchNAPTRRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	NAPTRRecordParam
}

func (r BatchPatchNAPTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchNAPTRRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchNSRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	NSRecordParam
}

func (r BatchPatchNSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchNSRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchOpenpgpkeyParam struct {
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
	// Settings for the DNS record.
	Settings param.Field[BatchPatchOpenpgpkeySettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[BatchPatchOpenpgpkeyType] `json:"type"`
}

func (r BatchPatchOpenpgpkeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchOpenpgpkeyParam) implementsDNSBatchPatchUnionParam() {}

// Settings for the DNS record.
type BatchPatchOpenpgpkeySettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r BatchPatchOpenpgpkeySettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Record type.
type BatchPatchOpenpgpkeyType string

const (
	BatchPatchOpenpgpkeyTypeOpenpgpkey BatchPatchOpenpgpkeyType = "OPENPGPKEY"
)

func (r BatchPatchOpenpgpkeyType) IsKnown() bool {
	switch r {
	case BatchPatchOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

type BatchPatchPTRRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	PTRRecordParam
}

func (r BatchPatchPTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchPTRRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchSMIMEARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SMIMEARecordParam
}

func (r BatchPatchSMIMEARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchSMIMEARecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchSRVRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SRVRecordParam
}

func (r BatchPatchSRVRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchSRVRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchSSHFPRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SSHFPRecordParam
}

func (r BatchPatchSSHFPRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchSSHFPRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchSVCBRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	SVCBRecordParam
}

func (r BatchPatchSVCBRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchSVCBRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchTLSARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	TLSARecordParam
}

func (r BatchPatchTLSARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchTLSARecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchTXTRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	TXTRecordParam
}

func (r BatchPatchTXTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchTXTRecordParam) implementsDNSBatchPatchUnionParam() {}

type BatchPatchURIRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id,required"`
	URIRecordParam
}

func (r BatchPatchURIRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPatchURIRecordParam) implementsDNSBatchPatchUnionParam() {}

// Record type.
type BatchPatchType string

const (
	BatchPatchTypeA          BatchPatchType = "A"
	BatchPatchTypeAAAA       BatchPatchType = "AAAA"
	BatchPatchTypeCAA        BatchPatchType = "CAA"
	BatchPatchTypeCERT       BatchPatchType = "CERT"
	BatchPatchTypeCNAME      BatchPatchType = "CNAME"
	BatchPatchTypeDNSKEY     BatchPatchType = "DNSKEY"
	BatchPatchTypeDS         BatchPatchType = "DS"
	BatchPatchTypeHTTPS      BatchPatchType = "HTTPS"
	BatchPatchTypeLOC        BatchPatchType = "LOC"
	BatchPatchTypeMX         BatchPatchType = "MX"
	BatchPatchTypeNAPTR      BatchPatchType = "NAPTR"
	BatchPatchTypeNS         BatchPatchType = "NS"
	BatchPatchTypeOpenpgpkey BatchPatchType = "OPENPGPKEY"
	BatchPatchTypePTR        BatchPatchType = "PTR"
	BatchPatchTypeSMIMEA     BatchPatchType = "SMIMEA"
	BatchPatchTypeSRV        BatchPatchType = "SRV"
	BatchPatchTypeSSHFP      BatchPatchType = "SSHFP"
	BatchPatchTypeSVCB       BatchPatchType = "SVCB"
	BatchPatchTypeTLSA       BatchPatchType = "TLSA"
	BatchPatchTypeTXT        BatchPatchType = "TXT"
	BatchPatchTypeURI        BatchPatchType = "URI"
)

func (r BatchPatchType) IsKnown() bool {
	switch r {
	case BatchPatchTypeA, BatchPatchTypeAAAA, BatchPatchTypeCAA, BatchPatchTypeCERT, BatchPatchTypeCNAME, BatchPatchTypeDNSKEY, BatchPatchTypeDS, BatchPatchTypeHTTPS, BatchPatchTypeLOC, BatchPatchTypeMX, BatchPatchTypeNAPTR, BatchPatchTypeNS, BatchPatchTypeOpenpgpkey, BatchPatchTypePTR, BatchPatchTypeSMIMEA, BatchPatchTypeSRV, BatchPatchTypeSSHFP, BatchPatchTypeSVCB, BatchPatchTypeTLSA, BatchPatchTypeTXT, BatchPatchTypeURI:
		return true
	}
	return false
}

// Satisfied by [dns.BatchPutARecordParam], [dns.BatchPutAAAARecordParam],
// [dns.BatchPutCAARecordParam], [dns.BatchPutCERTRecordParam],
// [dns.BatchPutCNAMERecordParam], [dns.BatchPutDNSKEYRecordParam],
// [dns.BatchPutDSRecordParam], [dns.BatchPutHTTPSRecordParam],
// [dns.BatchPutLOCRecordParam], [dns.BatchPutMXRecordParam],
// [dns.BatchPutNAPTRRecordParam], [dns.BatchPutNSRecordParam],
// [dns.BatchPutOpenpgpkeyParam], [dns.BatchPutPTRRecordParam],
// [dns.BatchPutSMIMEARecordParam], [dns.BatchPutSRVRecordParam],
// [dns.BatchPutSSHFPRecordParam], [dns.BatchPutSVCBRecordParam],
// [dns.BatchPutTLSARecordParam], [dns.BatchPutTXTRecordParam],
// [dns.BatchPutURIRecordParam].
type BatchPutUnionParam interface {
	implementsDNSBatchPutUnionParam()
}

type BatchPutARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	ARecordParam
}

func (r BatchPutARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutARecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutAAAARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	AAAARecordParam
}

func (r BatchPutAAAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutAAAARecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutCAARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	CAARecordParam
}

func (r BatchPutCAARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutCAARecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutCERTRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	CERTRecordParam
}

func (r BatchPutCERTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutCERTRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutCNAMERecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	CNAMERecordParam
}

func (r BatchPutCNAMERecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutCNAMERecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutDNSKEYRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	DNSKEYRecordParam
}

func (r BatchPutDNSKEYRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutDNSKEYRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutDSRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	DSRecordParam
}

func (r BatchPutDSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutDSRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutHTTPSRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	HTTPSRecordParam
}

func (r BatchPutHTTPSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutHTTPSRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutLOCRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	LOCRecordParam
}

func (r BatchPutLOCRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutLOCRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutMXRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	MXRecordParam
}

func (r BatchPutMXRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutMXRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutNAPTRRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	NAPTRRecordParam
}

func (r BatchPutNAPTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutNAPTRRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutNSRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	NSRecordParam
}

func (r BatchPutNSRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutNSRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutOpenpgpkeyParam struct {
	// A single Base64-encoded OpenPGP Transferable Public Key (RFC 4880 Section 11.1)
	Content param.Field[string] `json:"content,required"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name,required"`
	// Record type.
	Type param.Field[BatchPutOpenpgpkeyType] `json:"type,required"`
	// Identifier
	ID param.Field[string] `json:"id"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[BatchPutOpenpgpkeySettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
}

func (r BatchPutOpenpgpkeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutOpenpgpkeyParam) implementsDNSBatchPutUnionParam() {}

// Record type.
type BatchPutOpenpgpkeyType string

const (
	BatchPutOpenpgpkeyTypeOpenpgpkey BatchPutOpenpgpkeyType = "OPENPGPKEY"
)

func (r BatchPutOpenpgpkeyType) IsKnown() bool {
	switch r {
	case BatchPutOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

// Settings for the DNS record.
type BatchPutOpenpgpkeySettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r BatchPutOpenpgpkeySettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BatchPutPTRRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	PTRRecordParam
}

func (r BatchPutPTRRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutPTRRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutSMIMEARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	SMIMEARecordParam
}

func (r BatchPutSMIMEARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutSMIMEARecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutSRVRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	SRVRecordParam
}

func (r BatchPutSRVRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutSRVRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutSSHFPRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	SSHFPRecordParam
}

func (r BatchPutSSHFPRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutSSHFPRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutSVCBRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	SVCBRecordParam
}

func (r BatchPutSVCBRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutSVCBRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutTLSARecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	TLSARecordParam
}

func (r BatchPutTLSARecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutTLSARecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutTXTRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	TXTRecordParam
}

func (r BatchPutTXTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutTXTRecordParam) implementsDNSBatchPutUnionParam() {}

type BatchPutURIRecordParam struct {
	// Identifier
	ID param.Field[string] `json:"id"`
	URIRecordParam
}

func (r BatchPutURIRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BatchPutURIRecordParam) implementsDNSBatchPutUnionParam() {}

// Record type.
type BatchPutType string

const (
	BatchPutTypeA          BatchPutType = "A"
	BatchPutTypeAAAA       BatchPutType = "AAAA"
	BatchPutTypeCAA        BatchPutType = "CAA"
	BatchPutTypeCERT       BatchPutType = "CERT"
	BatchPutTypeCNAME      BatchPutType = "CNAME"
	BatchPutTypeDNSKEY     BatchPutType = "DNSKEY"
	BatchPutTypeDS         BatchPutType = "DS"
	BatchPutTypeHTTPS      BatchPutType = "HTTPS"
	BatchPutTypeLOC        BatchPutType = "LOC"
	BatchPutTypeMX         BatchPutType = "MX"
	BatchPutTypeNAPTR      BatchPutType = "NAPTR"
	BatchPutTypeNS         BatchPutType = "NS"
	BatchPutTypeOpenpgpkey BatchPutType = "OPENPGPKEY"
	BatchPutTypePTR        BatchPutType = "PTR"
	BatchPutTypeSMIMEA     BatchPutType = "SMIMEA"
	BatchPutTypeSRV        BatchPutType = "SRV"
	BatchPutTypeSSHFP      BatchPutType = "SSHFP"
	BatchPutTypeSVCB       BatchPutType = "SVCB"
	BatchPutTypeTLSA       BatchPutType = "TLSA"
	BatchPutTypeTXT        BatchPutType = "TXT"
	BatchPutTypeURI        BatchPutType = "URI"
)

func (r BatchPutType) IsKnown() bool {
	switch r {
	case BatchPutTypeA, BatchPutTypeAAAA, BatchPutTypeCAA, BatchPutTypeCERT, BatchPutTypeCNAME, BatchPutTypeDNSKEY, BatchPutTypeDS, BatchPutTypeHTTPS, BatchPutTypeLOC, BatchPutTypeMX, BatchPutTypeNAPTR, BatchPutTypeNS, BatchPutTypeOpenpgpkey, BatchPutTypePTR, BatchPutTypeSMIMEA, BatchPutTypeSRV, BatchPutTypeSSHFP, BatchPutTypeSVCB, BatchPutTypeTLSA, BatchPutTypeTXT, BatchPutTypeURI:
		return true
	}
	return false
}

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
	// Settings for the DNS record.
	Settings CAARecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type CAARecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                  `json:"ipv6_only"`
	JSON     caaRecordSettingsJSON `json:"-"`
}

// caaRecordSettingsJSON contains the JSON metadata for the struct
// [CAARecordSettings]
type caaRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CAARecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r caaRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[CAARecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type CAARecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r CAARecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings CERTRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type CERTRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                   `json:"ipv6_only"`
	JSON     certRecordSettingsJSON `json:"-"`
}

// certRecordSettingsJSON contains the JSON metadata for the struct
// [CERTRecordSettings]
type certRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CERTRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[CERTRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type CERTRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r CERTRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	Proxied bool `json:"proxied"`
	// Settings for the DNS record.
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

// Settings for the DNS record.
type CNAMERecordSettings struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting has no effect on proxied records, which are always
	// flattened.
	FlattenCNAME bool `json:"flatten_cname"`
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                    `json:"ipv6_only"`
	JSON     cnameRecordSettingsJSON `json:"-"`
}

// cnameRecordSettingsJSON contains the JSON metadata for the struct
// [CNAMERecordSettings]
type cnameRecordSettingsJSON struct {
	FlattenCNAME apijson.Field
	IPV4Only     apijson.Field
	IPV6Only     apijson.Field
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
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
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

// Settings for the DNS record.
type CNAMERecordSettingsParam struct {
	// If enabled, causes the CNAME record to be resolved externally and the resulting
	// address records (e.g., A and AAAA) to be returned instead of the CNAME record
	// itself. This setting has no effect on proxied records, which are always
	// flattened.
	FlattenCNAME param.Field[bool] `json:"flatten_cname"`
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
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
	// Settings for the DNS record.
	Settings DNSKEYRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type DNSKEYRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                     `json:"ipv6_only"`
	JSON     dnskeyRecordSettingsJSON `json:"-"`
}

// dnskeyRecordSettingsJSON contains the JSON metadata for the struct
// [DNSKEYRecordSettings]
type dnskeyRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSKEYRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnskeyRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[DNSKEYRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type DNSKEYRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r DNSKEYRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings DSRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type DSRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                 `json:"ipv6_only"`
	JSON     dsRecordSettingsJSON `json:"-"`
}

// dsRecordSettingsJSON contains the JSON metadata for the struct
// [DSRecordSettings]
type dsRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DSRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dsRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[DSRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type DSRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r DSRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings HTTPSRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type HTTPSRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                    `json:"ipv6_only"`
	JSON     httpsRecordSettingsJSON `json:"-"`
}

// httpsRecordSettingsJSON contains the JSON metadata for the struct
// [HTTPSRecordSettings]
type httpsRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPSRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpsRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[HTTPSRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type HTTPSRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r HTTPSRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings LOCRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type LOCRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                  `json:"ipv6_only"`
	JSON     locRecordSettingsJSON `json:"-"`
}

// locRecordSettingsJSON contains the JSON metadata for the struct
// [LOCRecordSettings]
type locRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOCRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locRecordSettingsJSON) RawJSON() string {
	return r.raw
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
	// Settings for the DNS record.
	Settings param.Field[LOCRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type LOCRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r LOCRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings MXRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type MXRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                 `json:"ipv6_only"`
	JSON     mxRecordSettingsJSON `json:"-"`
}

// mxRecordSettingsJSON contains the JSON metadata for the struct
// [MXRecordSettings]
type mxRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MXRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mxRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[MXRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type MXRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r MXRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// Settings for the DNS record.
	Settings NAPTRRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type NAPTRRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                    `json:"ipv6_only"`
	JSON     naptrRecordSettingsJSON `json:"-"`
}

// naptrRecordSettingsJSON contains the JSON metadata for the struct
// [NAPTRRecordSettings]
type naptrRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NAPTRRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r naptrRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[NAPTRRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type NAPTRRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r NAPTRRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings NSRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type NSRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                 `json:"ipv6_only"`
	JSON     nsRecordSettingsJSON `json:"-"`
}

// nsRecordSettingsJSON contains the JSON metadata for the struct
// [NSRecordSettings]
type nsRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NSRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nsRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[NSRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type NSRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r NSRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// Settings for the DNS record.
	Settings PTRRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type PTRRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                  `json:"ipv6_only"`
	JSON     ptrRecordSettingsJSON `json:"-"`
}

// ptrRecordSettingsJSON contains the JSON metadata for the struct
// [PTRRecordSettings]
type ptrRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PTRRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ptrRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[PTRRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type PTRRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r PTRRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
// [dns.RecordOpenpgpkeyParam], [dns.PTRRecordParam], [dns.SMIMEARecordParam],
// [dns.SRVRecordParam], [dns.SSHFPRecordParam], [dns.SVCBRecordParam],
// [dns.TLSARecordParam], [dns.TXTRecordParam], [dns.URIRecordParam],
// [RecordParam].
type RecordUnionParam interface {
	implementsDNSRecordUnionParam()
}

type RecordOpenpgpkeyParam struct {
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
	// Settings for the DNS record.
	Settings param.Field[RecordOpenpgpkeySettingsParam] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
	// Record type.
	Type param.Field[RecordOpenpgpkeyType] `json:"type"`
}

func (r RecordOpenpgpkeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RecordOpenpgpkeyParam) implementsDNSRecordUnionParam() {}

// Settings for the DNS record.
type RecordOpenpgpkeySettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r RecordOpenpgpkeySettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

type RecordResponse struct {
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
	// This field can have the runtime type of [ARecordSettings], [AAAARecordSettings],
	// [CAARecordSettings], [CERTRecordSettings], [CNAMERecordSettings],
	// [DNSKEYRecordSettings], [DSRecordSettings], [HTTPSRecordSettings],
	// [LOCRecordSettings], [MXRecordSettings], [NAPTRRecordSettings],
	// [NSRecordSettings], [RecordResponseOpenpgpkeySettings], [PTRRecordSettings],
	// [SMIMEARecordSettings], [SRVRecordSettings], [SSHFPRecordSettings],
	// [SVCBRecordSettings], [TLSARecordSettings], [TXTRecordSettings],
	// [URIRecordSettings].
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
	Type  RecordResponseType `json:"type"`
	JSON  recordResponseJSON `json:"-"`
	union RecordResponseUnion
}

// recordResponseJSON contains the JSON metadata for the struct [RecordResponse]
type recordResponseJSON struct {
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

func (r recordResponseJSON) RawJSON() string {
	return r.raw
}

func (r *RecordResponse) UnmarshalJSON(data []byte) (err error) {
	*r = RecordResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RecordResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [dns.RecordResponseARecord],
// [dns.RecordResponseAAAARecord], [dns.RecordResponseCAARecord],
// [dns.RecordResponseCERTRecord], [dns.RecordResponseCNAMERecord],
// [dns.RecordResponseDNSKEYRecord], [dns.RecordResponseDSRecord],
// [dns.RecordResponseHTTPSRecord], [dns.RecordResponseLOCRecord],
// [dns.RecordResponseMXRecord], [dns.RecordResponseNAPTRRecord],
// [dns.RecordResponseNSRecord], [dns.RecordResponseOpenpgpkey],
// [dns.RecordResponsePTRRecord], [dns.RecordResponseSMIMEARecord],
// [dns.RecordResponseSRVRecord], [dns.RecordResponseSSHFPRecord],
// [dns.RecordResponseSVCBRecord], [dns.RecordResponseTLSARecord],
// [dns.RecordResponseTXTRecord], [dns.RecordResponseURIRecord].
func (r RecordResponse) AsUnion() RecordResponseUnion {
	return r.union
}

// Union satisfied by [dns.RecordResponseARecord], [dns.RecordResponseAAAARecord],
// [dns.RecordResponseCAARecord], [dns.RecordResponseCERTRecord],
// [dns.RecordResponseCNAMERecord], [dns.RecordResponseDNSKEYRecord],
// [dns.RecordResponseDSRecord], [dns.RecordResponseHTTPSRecord],
// [dns.RecordResponseLOCRecord], [dns.RecordResponseMXRecord],
// [dns.RecordResponseNAPTRRecord], [dns.RecordResponseNSRecord],
// [dns.RecordResponseOpenpgpkey], [dns.RecordResponsePTRRecord],
// [dns.RecordResponseSMIMEARecord], [dns.RecordResponseSRVRecord],
// [dns.RecordResponseSSHFPRecord], [dns.RecordResponseSVCBRecord],
// [dns.RecordResponseTLSARecord], [dns.RecordResponseTXTRecord] or
// [dns.RecordResponseURIRecord].
type RecordResponseUnion interface {
	implementsDNSRecordResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RecordResponseUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseARecord{}),
			DiscriminatorValue: "A",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseAAAARecord{}),
			DiscriminatorValue: "AAAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseCAARecord{}),
			DiscriminatorValue: "CAA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseCERTRecord{}),
			DiscriminatorValue: "CERT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseCNAMERecord{}),
			DiscriminatorValue: "CNAME",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseDNSKEYRecord{}),
			DiscriminatorValue: "DNSKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseDSRecord{}),
			DiscriminatorValue: "DS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseHTTPSRecord{}),
			DiscriminatorValue: "HTTPS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseLOCRecord{}),
			DiscriminatorValue: "LOC",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseMXRecord{}),
			DiscriminatorValue: "MX",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseNAPTRRecord{}),
			DiscriminatorValue: "NAPTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseNSRecord{}),
			DiscriminatorValue: "NS",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseOpenpgpkey{}),
			DiscriminatorValue: "OPENPGPKEY",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponsePTRRecord{}),
			DiscriminatorValue: "PTR",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseSMIMEARecord{}),
			DiscriminatorValue: "SMIMEA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseSRVRecord{}),
			DiscriminatorValue: "SRV",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseSSHFPRecord{}),
			DiscriminatorValue: "SSHFP",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseSVCBRecord{}),
			DiscriminatorValue: "SVCB",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseTLSARecord{}),
			DiscriminatorValue: "TLSA",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseTXTRecord{}),
			DiscriminatorValue: "TXT",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RecordResponseURIRecord{}),
			DiscriminatorValue: "URI",
		},
	)
}

type RecordResponseARecord struct {
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
	TagsModifiedOn time.Time                 `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseARecordJSON `json:"-"`
	ARecord
}

// recordResponseARecordJSON contains the JSON metadata for the struct
// [RecordResponseARecord]
type recordResponseARecordJSON struct {
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

func (r *RecordResponseARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseARecord) implementsDNSRecordResponse() {}

type RecordResponseAAAARecord struct {
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
	JSON           recordResponseAAAARecordJSON `json:"-"`
	AAAARecord
}

// recordResponseAAAARecordJSON contains the JSON metadata for the struct
// [RecordResponseAAAARecord]
type recordResponseAAAARecordJSON struct {
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

func (r *RecordResponseAAAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseAAAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseAAAARecord) implementsDNSRecordResponse() {}

type RecordResponseCAARecord struct {
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
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseCAARecordJSON `json:"-"`
	CAARecord
}

// recordResponseCAARecordJSON contains the JSON metadata for the struct
// [RecordResponseCAARecord]
type recordResponseCAARecordJSON struct {
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

func (r *RecordResponseCAARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseCAARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseCAARecord) implementsDNSRecordResponse() {}

type RecordResponseCERTRecord struct {
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
	JSON           recordResponseCERTRecordJSON `json:"-"`
	CERTRecord
}

// recordResponseCERTRecordJSON contains the JSON metadata for the struct
// [RecordResponseCERTRecord]
type recordResponseCERTRecordJSON struct {
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

func (r *RecordResponseCERTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseCERTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseCERTRecord) implementsDNSRecordResponse() {}

type RecordResponseCNAMERecord struct {
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
	JSON           recordResponseCNAMERecordJSON `json:"-"`
	CNAMERecord
}

// recordResponseCNAMERecordJSON contains the JSON metadata for the struct
// [RecordResponseCNAMERecord]
type recordResponseCNAMERecordJSON struct {
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

func (r *RecordResponseCNAMERecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseCNAMERecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseCNAMERecord) implementsDNSRecordResponse() {}

type RecordResponseDNSKEYRecord struct {
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
	JSON           recordResponseDNSKEYRecordJSON `json:"-"`
	DNSKEYRecord
}

// recordResponseDNSKEYRecordJSON contains the JSON metadata for the struct
// [RecordResponseDNSKEYRecord]
type recordResponseDNSKEYRecordJSON struct {
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

func (r *RecordResponseDNSKEYRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseDNSKEYRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseDNSKEYRecord) implementsDNSRecordResponse() {}

type RecordResponseDSRecord struct {
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
	TagsModifiedOn time.Time                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseDSRecordJSON `json:"-"`
	DSRecord
}

// recordResponseDSRecordJSON contains the JSON metadata for the struct
// [RecordResponseDSRecord]
type recordResponseDSRecordJSON struct {
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

func (r *RecordResponseDSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseDSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseDSRecord) implementsDNSRecordResponse() {}

type RecordResponseHTTPSRecord struct {
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
	JSON           recordResponseHTTPSRecordJSON `json:"-"`
	HTTPSRecord
}

// recordResponseHTTPSRecordJSON contains the JSON metadata for the struct
// [RecordResponseHTTPSRecord]
type recordResponseHTTPSRecordJSON struct {
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

func (r *RecordResponseHTTPSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseHTTPSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseHTTPSRecord) implementsDNSRecordResponse() {}

type RecordResponseLOCRecord struct {
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
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseLOCRecordJSON `json:"-"`
	LOCRecord
}

// recordResponseLOCRecordJSON contains the JSON metadata for the struct
// [RecordResponseLOCRecord]
type recordResponseLOCRecordJSON struct {
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

func (r *RecordResponseLOCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseLOCRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseLOCRecord) implementsDNSRecordResponse() {}

type RecordResponseMXRecord struct {
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
	TagsModifiedOn time.Time                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseMXRecordJSON `json:"-"`
	MXRecord
}

// recordResponseMXRecordJSON contains the JSON metadata for the struct
// [RecordResponseMXRecord]
type recordResponseMXRecordJSON struct {
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

func (r *RecordResponseMXRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseMXRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseMXRecord) implementsDNSRecordResponse() {}

type RecordResponseNAPTRRecord struct {
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
	JSON           recordResponseNAPTRRecordJSON `json:"-"`
	NAPTRRecord
}

// recordResponseNAPTRRecordJSON contains the JSON metadata for the struct
// [RecordResponseNAPTRRecord]
type recordResponseNAPTRRecordJSON struct {
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

func (r *RecordResponseNAPTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseNAPTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseNAPTRRecord) implementsDNSRecordResponse() {}

type RecordResponseNSRecord struct {
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
	TagsModifiedOn time.Time                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseNSRecordJSON `json:"-"`
	NSRecord
}

// recordResponseNSRecordJSON contains the JSON metadata for the struct
// [RecordResponseNSRecord]
type recordResponseNSRecordJSON struct {
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

func (r *RecordResponseNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseNSRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseNSRecord) implementsDNSRecordResponse() {}

type RecordResponseOpenpgpkey struct {
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
	// Settings for the DNS record.
	Settings RecordResponseOpenpgpkeySettings `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// Record type.
	Type RecordResponseOpenpgpkeyType `json:"type,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                    `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseOpenpgpkeyJSON `json:"-"`
}

// recordResponseOpenpgpkeyJSON contains the JSON metadata for the struct
// [RecordResponseOpenpgpkey]
type recordResponseOpenpgpkeyJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	Content           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	Type              apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordResponseOpenpgpkey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseOpenpgpkeyJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseOpenpgpkey) implementsDNSRecordResponse() {}

// Settings for the DNS record.
type RecordResponseOpenpgpkeySettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                                 `json:"ipv6_only"`
	JSON     recordResponseOpenpgpkeySettingsJSON `json:"-"`
}

// recordResponseOpenpgpkeySettingsJSON contains the JSON metadata for the struct
// [RecordResponseOpenpgpkeySettings]
type recordResponseOpenpgpkeySettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RecordResponseOpenpgpkeySettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseOpenpgpkeySettingsJSON) RawJSON() string {
	return r.raw
}

// Record type.
type RecordResponseOpenpgpkeyType string

const (
	RecordResponseOpenpgpkeyTypeOpenpgpkey RecordResponseOpenpgpkeyType = "OPENPGPKEY"
)

func (r RecordResponseOpenpgpkeyType) IsKnown() bool {
	switch r {
	case RecordResponseOpenpgpkeyTypeOpenpgpkey:
		return true
	}
	return false
}

type RecordResponsePTRRecord struct {
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
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponsePTRRecordJSON `json:"-"`
	PTRRecord
}

// recordResponsePTRRecordJSON contains the JSON metadata for the struct
// [RecordResponsePTRRecord]
type recordResponsePTRRecordJSON struct {
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

func (r *RecordResponsePTRRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponsePTRRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponsePTRRecord) implementsDNSRecordResponse() {}

type RecordResponseSMIMEARecord struct {
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
	JSON           recordResponseSMIMEARecordJSON `json:"-"`
	SMIMEARecord
}

// recordResponseSMIMEARecordJSON contains the JSON metadata for the struct
// [RecordResponseSMIMEARecord]
type recordResponseSMIMEARecordJSON struct {
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

func (r *RecordResponseSMIMEARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseSMIMEARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseSMIMEARecord) implementsDNSRecordResponse() {}

type RecordResponseSRVRecord struct {
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
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseSRVRecordJSON `json:"-"`
	SRVRecord
}

// recordResponseSRVRecordJSON contains the JSON metadata for the struct
// [RecordResponseSRVRecord]
type recordResponseSRVRecordJSON struct {
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

func (r *RecordResponseSRVRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseSRVRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseSRVRecord) implementsDNSRecordResponse() {}

type RecordResponseSSHFPRecord struct {
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
	JSON           recordResponseSSHFPRecordJSON `json:"-"`
	SSHFPRecord
}

// recordResponseSSHFPRecordJSON contains the JSON metadata for the struct
// [RecordResponseSSHFPRecord]
type recordResponseSSHFPRecordJSON struct {
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

func (r *RecordResponseSSHFPRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseSSHFPRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseSSHFPRecord) implementsDNSRecordResponse() {}

type RecordResponseSVCBRecord struct {
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
	JSON           recordResponseSVCBRecordJSON `json:"-"`
	SVCBRecord
}

// recordResponseSVCBRecordJSON contains the JSON metadata for the struct
// [RecordResponseSVCBRecord]
type recordResponseSVCBRecordJSON struct {
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

func (r *RecordResponseSVCBRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseSVCBRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseSVCBRecord) implementsDNSRecordResponse() {}

type RecordResponseTLSARecord struct {
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
	JSON           recordResponseTLSARecordJSON `json:"-"`
	TLSARecord
}

// recordResponseTLSARecordJSON contains the JSON metadata for the struct
// [RecordResponseTLSARecord]
type recordResponseTLSARecordJSON struct {
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

func (r *RecordResponseTLSARecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseTLSARecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseTLSARecord) implementsDNSRecordResponse() {}

type RecordResponseTXTRecord struct {
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
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseTXTRecordJSON `json:"-"`
	TXTRecord
}

// recordResponseTXTRecordJSON contains the JSON metadata for the struct
// [RecordResponseTXTRecord]
type recordResponseTXTRecordJSON struct {
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

func (r *RecordResponseTXTRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseTXTRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseTXTRecord) implementsDNSRecordResponse() {}

type RecordResponseURIRecord struct {
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
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordResponseURIRecordJSON `json:"-"`
	URIRecord
}

// recordResponseURIRecordJSON contains the JSON metadata for the struct
// [RecordResponseURIRecord]
type recordResponseURIRecordJSON struct {
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

func (r *RecordResponseURIRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordResponseURIRecordJSON) RawJSON() string {
	return r.raw
}

func (r RecordResponseURIRecord) implementsDNSRecordResponse() {}

// Record type.
type RecordResponseType string

const (
	RecordResponseTypeA          RecordResponseType = "A"
	RecordResponseTypeAAAA       RecordResponseType = "AAAA"
	RecordResponseTypeCAA        RecordResponseType = "CAA"
	RecordResponseTypeCERT       RecordResponseType = "CERT"
	RecordResponseTypeCNAME      RecordResponseType = "CNAME"
	RecordResponseTypeDNSKEY     RecordResponseType = "DNSKEY"
	RecordResponseTypeDS         RecordResponseType = "DS"
	RecordResponseTypeHTTPS      RecordResponseType = "HTTPS"
	RecordResponseTypeLOC        RecordResponseType = "LOC"
	RecordResponseTypeMX         RecordResponseType = "MX"
	RecordResponseTypeNAPTR      RecordResponseType = "NAPTR"
	RecordResponseTypeNS         RecordResponseType = "NS"
	RecordResponseTypeOpenpgpkey RecordResponseType = "OPENPGPKEY"
	RecordResponseTypePTR        RecordResponseType = "PTR"
	RecordResponseTypeSMIMEA     RecordResponseType = "SMIMEA"
	RecordResponseTypeSRV        RecordResponseType = "SRV"
	RecordResponseTypeSSHFP      RecordResponseType = "SSHFP"
	RecordResponseTypeSVCB       RecordResponseType = "SVCB"
	RecordResponseTypeTLSA       RecordResponseType = "TLSA"
	RecordResponseTypeTXT        RecordResponseType = "TXT"
	RecordResponseTypeURI        RecordResponseType = "URI"
)

func (r RecordResponseType) IsKnown() bool {
	switch r {
	case RecordResponseTypeA, RecordResponseTypeAAAA, RecordResponseTypeCAA, RecordResponseTypeCERT, RecordResponseTypeCNAME, RecordResponseTypeDNSKEY, RecordResponseTypeDS, RecordResponseTypeHTTPS, RecordResponseTypeLOC, RecordResponseTypeMX, RecordResponseTypeNAPTR, RecordResponseTypeNS, RecordResponseTypeOpenpgpkey, RecordResponseTypePTR, RecordResponseTypeSMIMEA, RecordResponseTypeSRV, RecordResponseTypeSSHFP, RecordResponseTypeSVCB, RecordResponseTypeTLSA, RecordResponseTypeTXT, RecordResponseTypeURI:
		return true
	}
	return false
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
	// Settings for the DNS record.
	Settings SMIMEARecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type SMIMEARecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                     `json:"ipv6_only"`
	JSON     smimeaRecordSettingsJSON `json:"-"`
}

// smimeaRecordSettingsJSON contains the JSON metadata for the struct
// [SMIMEARecordSettings]
type smimeaRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SMIMEARecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smimeaRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[SMIMEARecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type SMIMEARecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r SMIMEARecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings SRVRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type SRVRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                  `json:"ipv6_only"`
	JSON     srvRecordSettingsJSON `json:"-"`
}

// srvRecordSettingsJSON contains the JSON metadata for the struct
// [SRVRecordSettings]
type srvRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SRVRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r srvRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[SRVRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type SRVRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r SRVRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings SSHFPRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type SSHFPRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                    `json:"ipv6_only"`
	JSON     sshfpRecordSettingsJSON `json:"-"`
}

// sshfpRecordSettingsJSON contains the JSON metadata for the struct
// [SSHFPRecordSettings]
type sshfpRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSHFPRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sshfpRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[SSHFPRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type SSHFPRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r SSHFPRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings SVCBRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type SVCBRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                   `json:"ipv6_only"`
	JSON     svcbRecordSettingsJSON `json:"-"`
}

// svcbRecordSettingsJSON contains the JSON metadata for the struct
// [SVCBRecordSettings]
type svcbRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SVCBRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r svcbRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[SVCBRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type SVCBRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r SVCBRecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings TLSARecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type TLSARecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                   `json:"ipv6_only"`
	JSON     tlsaRecordSettingsJSON `json:"-"`
}

// tlsaRecordSettingsJSON contains the JSON metadata for the struct
// [TLSARecordSettings]
type tlsaRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSARecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsaRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[TLSARecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type TLSARecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r TLSARecordSettingsParam) MarshalJSON() (data []byte, err error) {
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
	// Settings for the DNS record.
	Settings TXTRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type TXTRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                  `json:"ipv6_only"`
	JSON     txtRecordSettingsJSON `json:"-"`
}

// txtRecordSettingsJSON contains the JSON metadata for the struct
// [TXTRecordSettings]
type txtRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TXTRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r txtRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[TXTRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type TXTRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r TXTRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	// Settings for the DNS record.
	Settings URIRecordSettings `json:"settings"`
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
	Settings    apijson.Field
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

// Settings for the DNS record.
type URIRecordSettings struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only bool `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only bool                  `json:"ipv6_only"`
	JSON     uriRecordSettingsJSON `json:"-"`
}

// uriRecordSettingsJSON contains the JSON metadata for the struct
// [URIRecordSettings]
type uriRecordSettingsJSON struct {
	IPV4Only    apijson.Field
	IPV6Only    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URIRecordSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uriRecordSettingsJSON) RawJSON() string {
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
	// Settings for the DNS record.
	Settings param.Field[URIRecordSettingsParam] `json:"settings"`
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

// Settings for the DNS record.
type URIRecordSettingsParam struct {
	// When enabled, only A records will be generated, and AAAA records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV4Only param.Field[bool] `json:"ipv4_only"`
	// When enabled, only AAAA records will be generated, and A records will not be
	// created. This setting is intended for exceptional cases. Note that this option
	// only applies to proxied records and it has no effect on whether Cloudflare
	// communicates with the origin using IPv4 or IPv6.
	IPV6Only param.Field[bool] `json:"ipv6_only"`
}

func (r URIRecordSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Deletes []RecordResponse        `json:"deletes"`
	Patches []RecordResponse        `json:"patches"`
	Posts   []RecordResponse        `json:"posts"`
	Puts    []RecordResponse        `json:"puts"`
	JSON    recordBatchResponseJSON `json:"-"`
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
	Result  RecordResponse                   `json:"result"`
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
	Result  RecordResponse                      `json:"result"`
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
	Patches param.Field[[]BatchPatchUnionParam]    `json:"patches"`
	Posts   param.Field[[]RecordUnionParam]        `json:"posts"`
	Puts    param.Field[[]BatchPutUnionParam]      `json:"puts"`
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
	Result  RecordResponse                    `json:"result"`
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
	Result  RecordResponse                   `json:"result"`
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
	JSON    recordImportResponseEnvelopeJSON    `json:"-"`
}

// recordImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordImportResponseEnvelope]
type recordImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	JSON    recordScanResponseEnvelopeJSON    `json:"-"`
}

// recordScanResponseEnvelopeJSON contains the JSON metadata for the struct
// [RecordScanResponseEnvelope]
type recordScanResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
