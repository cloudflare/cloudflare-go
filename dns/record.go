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

	"github.com/cloudflare/cloudflare-go/internal/apiform"
	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/pagination"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
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

// Overwrite an existing DNS record. Notes:
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

// Update an existing DNS record. Notes:
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
