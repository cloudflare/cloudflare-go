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
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
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

type CNAMERecordParam struct {
	// A valid hostname. Must not match the record's name.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[CNAMERecordType] `json:"type"`
}

func (r CNAMERecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CNAMERecordParam) implementsDNSRecordUnionParam() {}

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

type RecordParam struct {
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment param.Field[string] `json:"comment"`
	// DNS record name (or @ for the zone apex) in Punycode.
	Name param.Field[string] `json:"name"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
	// Settings for the DNS record.
	Settings param.Field[interface{}] `json:"settings"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags param.Field[[]RecordTagsParam] `json:"tags"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL param.Field[TTL] `json:"ttl"`
}

func (r RecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type TXTRecordParam struct {
	// Text content for the record. The content must consist of quoted "character
	// strings" (RFC 1035), each with a length of up to 255 bytes. Strings exceeding
	// this allowed maximum length are automatically split.
	Content param.Field[string] `json:"content"`
	// Record type.
	Type param.Field[TXTRecordType] `json:"type"`
}

func (r TXTRecordParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r TXTRecordParam) implementsDNSRecordUnionParam() {}

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

type RecordNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time             `json:"tags_modified_on" format:"date-time"`
	JSON           recordNewResponseJSON `json:"-"`
}

// recordNewResponseJSON contains the JSON metadata for the struct
// [RecordNewResponse]
type recordNewResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordNewResponseJSON) RawJSON() string {
	return r.raw
}

type RecordUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                `json:"tags_modified_on" format:"date-time"`
	JSON           recordUpdateResponseJSON `json:"-"`
}

// recordUpdateResponseJSON contains the JSON metadata for the struct
// [RecordUpdateResponse]
type recordUpdateResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RecordListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time              `json:"tags_modified_on" format:"date-time"`
	JSON           recordListResponseJSON `json:"-"`
}

// recordListResponseJSON contains the JSON metadata for the struct
// [RecordListResponse]
type recordListResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordListResponseJSON) RawJSON() string {
	return r.raw
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
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                     `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponseDeleteJSON `json:"-"`
}

// recordBatchResponseDeleteJSON contains the JSON metadata for the struct
// [RecordBatchResponseDelete]
type recordBatchResponseDeleteJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponseDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponseDeleteJSON) RawJSON() string {
	return r.raw
}

type RecordBatchResponsePatch struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                    `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePatchJSON `json:"-"`
}

// recordBatchResponsePatchJSON contains the JSON metadata for the struct
// [RecordBatchResponsePatch]
type recordBatchResponsePatchJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePatchJSON) RawJSON() string {
	return r.raw
}

type RecordBatchResponsePost struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                   `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePostJSON `json:"-"`
}

// recordBatchResponsePostJSON contains the JSON metadata for the struct
// [RecordBatchResponsePost]
type recordBatchResponsePostJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePostJSON) RawJSON() string {
	return r.raw
}

type RecordBatchResponsePut struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time                  `json:"tags_modified_on" format:"date-time"`
	JSON           recordBatchResponsePutJSON `json:"-"`
}

// recordBatchResponsePutJSON contains the JSON metadata for the struct
// [RecordBatchResponsePut]
type recordBatchResponsePutJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordBatchResponsePut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordBatchResponsePutJSON) RawJSON() string {
	return r.raw
}

type RecordEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time              `json:"tags_modified_on" format:"date-time"`
	JSON           recordEditResponseJSON `json:"-"`
}

// recordEditResponseJSON contains the JSON metadata for the struct
// [RecordEditResponse]
type recordEditResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordEditResponseJSON) RawJSON() string {
	return r.raw
}

type RecordGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// Comments or notes about the DNS record. This field has no effect on DNS
	// responses.
	Comment string `json:"comment,required"`
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
	Settings interface{} `json:"settings,required"`
	// Custom tags for the DNS record. This field has no effect on DNS responses.
	Tags []RecordTags `json:"tags,required"`
	// Time To Live (TTL) of the DNS record in seconds. Setting to 1 means 'automatic'.
	// Value must be between 60 and 86400, with the minimum reduced to 30 for
	// Enterprise zones.
	TTL TTL `json:"ttl,required"`
	// When the record comment was last modified. Omitted if there is no comment.
	CommentModifiedOn time.Time `json:"comment_modified_on" format:"date-time"`
	// When the record tags were last modified. Omitted if there are no tags.
	TagsModifiedOn time.Time             `json:"tags_modified_on" format:"date-time"`
	JSON           recordGetResponseJSON `json:"-"`
}

// recordGetResponseJSON contains the JSON metadata for the struct
// [RecordGetResponse]
type recordGetResponseJSON struct {
	ID                apijson.Field
	Comment           apijson.Field
	CreatedOn         apijson.Field
	Meta              apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	Proxiable         apijson.Field
	Proxied           apijson.Field
	Settings          apijson.Field
	Tags              apijson.Field
	TTL               apijson.Field
	CommentModifiedOn apijson.Field
	TagsModifiedOn    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RecordGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r recordGetResponseJSON) RawJSON() string {
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
	Record RecordParam         `json:"record,required"`
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
	Record RecordParam         `json:"record,required"`
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
	// DNS record content.
	Content param.Field[string] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[shared.SortDirection] `query:"direction"`
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
	Posts   param.Field[[]RecordParam]             `json:"posts"`
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
	Record RecordParam         `json:"record,required"`
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
