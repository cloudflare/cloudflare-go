// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// DMARCReportService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDMARCReportService] method instead.
type DMARCReportService struct {
	Options []option.RequestOption
}

// NewDMARCReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDMARCReportService(opts ...option.RequestOption) (r *DMARCReportService) {
	r = &DMARCReportService{}
	r.Options = opts
	return
}

// Updates the DMARC report configuration for a zone. At least one of `enabled` or
// `skip_wizard` must be provided. When enabling, the handler will ensure the DMARC
// RUA record exists in DNS.
func (r *DMARCReportService) Edit(ctx context.Context, params DMARCReportEditParams, opts ...option.RequestOption) (res *DMARCReportEditResponse, err error) {
	var env DMARCReportEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/auth/dmarc-reports", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the current DMARC report configuration and status for a zone. Returns
// the RUA prefix, enabled status, approved sources, and DNS records.
func (r *DMARCReportService) Get(ctx context.Context, query DMARCReportGetParams, opts ...option.RequestOption) (res *DMARCReportGetResponse, err error) {
	var env DMARCReportGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/email/auth/dmarc-reports", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Response for GET/PATCH /dmarc-reports
type DMARCReportEditResponse struct {
	// List of approved sending sources (omitted when empty)
	ApprovedSources []DMARCReportEditResponseApprovedSource `json:"approved_sources"`
	// Deprecated, use created_at
	//
	// Deprecated: deprecated
	Created time.Time `json:"created" format:"date-time"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether DMARC reports are enabled
	Enabled bool `json:"enabled"`
	// Deprecated, use modified_at
	//
	// Deprecated: deprecated
	Modified time.Time `json:"modified" format:"date-time"`
	// Last modification timestamp
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// Live DNS records for the zone, grouped by type
	Records DMARCReportEditResponseRecords `json:"records"`
	// Prefix for DMARC RUA addresses (32-char hex string)
	RuaPrefix string `json:"rua_prefix"`
	// Whether to skip the setup wizard
	SkipWizard bool `json:"skip_wizard"`
	// DMARC configuration status
	Status DMARCReportEditResponseStatus `json:"status"`
	// Use `zone_id` instead
	//
	// Deprecated: deprecated
	Tag string `json:"tag"`
	// Zone identifier
	ZoneID string                      `json:"zone_id"`
	JSON   dmarcReportEditResponseJSON `json:"-"`
}

// dmarcReportEditResponseJSON contains the JSON metadata for the struct
// [DMARCReportEditResponse]
type dmarcReportEditResponseJSON struct {
	ApprovedSources apijson.Field
	Created         apijson.Field
	CreatedAt       apijson.Field
	Enabled         apijson.Field
	Modified        apijson.Field
	ModifiedAt      apijson.Field
	Records         apijson.Field
	RuaPrefix       apijson.Field
	SkipWizard      apijson.Field
	Status          apijson.Field
	Tag             apijson.Field
	ZoneID          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DMARCReportEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseJSON) RawJSON() string {
	return r.raw
}

// A single approved sending source
type DMARCReportEditResponseApprovedSource struct {
	// Deprecated, use created_at
	//
	// Deprecated: deprecated
	Created time.Time `json:"created" format:"date-time"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The source domain
	Domain string `json:"domain"`
	// Resolved IP addresses from SPF
	IPs []string `json:"ips"`
	// Deprecated, use modified_at
	//
	// Deprecated: deprecated
	Modified time.Time `json:"modified" format:"date-time"`
	// Last modification timestamp
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// Source name (typically same as domain)
	Name string `json:"name"`
	// URL-friendly identifier
	Slug string `json:"slug"`
	// Source UUID
	Tag  string                                    `json:"tag"`
	JSON dmarcReportEditResponseApprovedSourceJSON `json:"-"`
}

// dmarcReportEditResponseApprovedSourceJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseApprovedSource]
type dmarcReportEditResponseApprovedSourceJSON struct {
	Created     apijson.Field
	CreatedAt   apijson.Field
	Domain      apijson.Field
	IPs         apijson.Field
	Modified    apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseApprovedSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseApprovedSourceJSON) RawJSON() string {
	return r.raw
}

// Live DNS records for the zone, grouped by type
type DMARCReportEditResponseRecords struct {
	// BIMI TXT records
	BimiRecords []DMARCReportEditResponseRecordsBimiRecord `json:"bimi_records"`
	// CNAME records for DKIM
	CNAMEDKIMRecords []DMARCReportEditResponseRecordsCnamedkimRecord `json:"cname_dkim_records"`
	// CNAME records at \_dmarc (problematic)
	CNAMEDMARCRecords []DMARCReportEditResponseRecordsCnamedmarcRecord `json:"cname_dmarc_records"`
	// CNAME records for SPF
	CNAMESPFRecords []DMARCReportEditResponseRecordsCnamespfRecord `json:"cname_spf_records"`
	// DKIM TXT records
	DKIMRecords []DMARCReportEditResponseRecordsDKIMRecord `json:"dkim_records"`
	// DMARC TXT records
	DMARCRecords []DMARCReportEditResponseRecordsDMARCRecord `json:"dmarc_records"`
	// SPF TXT records
	SPFRecords []DMARCReportEditResponseRecordsSPFRecord `json:"spf_records"`
	JSON       dmarcReportEditResponseRecordsJSON        `json:"-"`
}

// dmarcReportEditResponseRecordsJSON contains the JSON metadata for the struct
// [DMARCReportEditResponseRecords]
type dmarcReportEditResponseRecordsJSON struct {
	BimiRecords       apijson.Field
	CNAMEDKIMRecords  apijson.Field
	CNAMEDMARCRecords apijson.Field
	CNAMESPFRecords   apijson.Field
	DKIMRecords       apijson.Field
	DMARCRecords      apijson.Field
	SPFRecords        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsBimiRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                       `json:"type"`
	JSON dmarcReportEditResponseRecordsBimiRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsBimiRecordJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseRecordsBimiRecord]
type dmarcReportEditResponseRecordsBimiRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsBimiRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsBimiRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsCnamedkimRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                            `json:"type"`
	JSON dmarcReportEditResponseRecordsCnamedkimRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsCnamedkimRecordJSON contains the JSON metadata for
// the struct [DMARCReportEditResponseRecordsCnamedkimRecord]
type dmarcReportEditResponseRecordsCnamedkimRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsCnamedkimRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsCnamedkimRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsCnamedmarcRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                             `json:"type"`
	JSON dmarcReportEditResponseRecordsCnamedmarcRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsCnamedmarcRecordJSON contains the JSON metadata
// for the struct [DMARCReportEditResponseRecordsCnamedmarcRecord]
type dmarcReportEditResponseRecordsCnamedmarcRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsCnamedmarcRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsCnamedmarcRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsCnamespfRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                           `json:"type"`
	JSON dmarcReportEditResponseRecordsCnamespfRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsCnamespfRecordJSON contains the JSON metadata for
// the struct [DMARCReportEditResponseRecordsCnamespfRecord]
type dmarcReportEditResponseRecordsCnamespfRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsCnamespfRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsCnamespfRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsDKIMRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                       `json:"type"`
	JSON dmarcReportEditResponseRecordsDKIMRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsDKIMRecordJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseRecordsDKIMRecord]
type dmarcReportEditResponseRecordsDKIMRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsDKIMRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsDKIMRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsDMARCRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                        `json:"type"`
	JSON dmarcReportEditResponseRecordsDMARCRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsDMARCRecordJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseRecordsDMARCRecord]
type dmarcReportEditResponseRecordsDMARCRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsDMARCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsDMARCRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportEditResponseRecordsSPFRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                      `json:"type"`
	JSON dmarcReportEditResponseRecordsSPFRecordJSON `json:"-"`
}

// dmarcReportEditResponseRecordsSPFRecordJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseRecordsSPFRecord]
type dmarcReportEditResponseRecordsSPFRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseRecordsSPFRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseRecordsSPFRecordJSON) RawJSON() string {
	return r.raw
}

// DMARC configuration status
type DMARCReportEditResponseStatus string

const (
	DMARCReportEditResponseStatusMissingDMARCReport   DMARCReportEditResponseStatus = "missing-dmarc-report"
	DMARCReportEditResponseStatusMultipleDMARCReports DMARCReportEditResponseStatus = "multiple-dmarc-reports"
	DMARCReportEditResponseStatusMissingDMARCRua      DMARCReportEditResponseStatus = "missing-dmarc-rua"
	DMARCReportEditResponseStatusCNAMEOnDMARCRecord   DMARCReportEditResponseStatus = "cname-on-dmarc-record"
)

func (r DMARCReportEditResponseStatus) IsKnown() bool {
	switch r {
	case DMARCReportEditResponseStatusMissingDMARCReport, DMARCReportEditResponseStatusMultipleDMARCReports, DMARCReportEditResponseStatusMissingDMARCRua, DMARCReportEditResponseStatusCNAMEOnDMARCRecord:
		return true
	}
	return false
}

// Response for GET/PATCH /dmarc-reports
type DMARCReportGetResponse struct {
	// List of approved sending sources (omitted when empty)
	ApprovedSources []DMARCReportGetResponseApprovedSource `json:"approved_sources"`
	// Deprecated, use created_at
	//
	// Deprecated: deprecated
	Created time.Time `json:"created" format:"date-time"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether DMARC reports are enabled
	Enabled bool `json:"enabled"`
	// Deprecated, use modified_at
	//
	// Deprecated: deprecated
	Modified time.Time `json:"modified" format:"date-time"`
	// Last modification timestamp
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// Live DNS records for the zone, grouped by type
	Records DMARCReportGetResponseRecords `json:"records"`
	// Prefix for DMARC RUA addresses (32-char hex string)
	RuaPrefix string `json:"rua_prefix"`
	// Whether to skip the setup wizard
	SkipWizard bool `json:"skip_wizard"`
	// DMARC configuration status
	Status DMARCReportGetResponseStatus `json:"status"`
	// Use `zone_id` instead
	//
	// Deprecated: deprecated
	Tag string `json:"tag"`
	// Zone identifier
	ZoneID string                     `json:"zone_id"`
	JSON   dmarcReportGetResponseJSON `json:"-"`
}

// dmarcReportGetResponseJSON contains the JSON metadata for the struct
// [DMARCReportGetResponse]
type dmarcReportGetResponseJSON struct {
	ApprovedSources apijson.Field
	Created         apijson.Field
	CreatedAt       apijson.Field
	Enabled         apijson.Field
	Modified        apijson.Field
	ModifiedAt      apijson.Field
	Records         apijson.Field
	RuaPrefix       apijson.Field
	SkipWizard      apijson.Field
	Status          apijson.Field
	Tag             apijson.Field
	ZoneID          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DMARCReportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseJSON) RawJSON() string {
	return r.raw
}

// A single approved sending source
type DMARCReportGetResponseApprovedSource struct {
	// Deprecated, use created_at
	//
	// Deprecated: deprecated
	Created time.Time `json:"created" format:"date-time"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The source domain
	Domain string `json:"domain"`
	// Resolved IP addresses from SPF
	IPs []string `json:"ips"`
	// Deprecated, use modified_at
	//
	// Deprecated: deprecated
	Modified time.Time `json:"modified" format:"date-time"`
	// Last modification timestamp
	ModifiedAt time.Time `json:"modified_at" format:"date-time"`
	// Source name (typically same as domain)
	Name string `json:"name"`
	// URL-friendly identifier
	Slug string `json:"slug"`
	// Source UUID
	Tag  string                                   `json:"tag"`
	JSON dmarcReportGetResponseApprovedSourceJSON `json:"-"`
}

// dmarcReportGetResponseApprovedSourceJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseApprovedSource]
type dmarcReportGetResponseApprovedSourceJSON struct {
	Created     apijson.Field
	CreatedAt   apijson.Field
	Domain      apijson.Field
	IPs         apijson.Field
	Modified    apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseApprovedSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseApprovedSourceJSON) RawJSON() string {
	return r.raw
}

// Live DNS records for the zone, grouped by type
type DMARCReportGetResponseRecords struct {
	// BIMI TXT records
	BimiRecords []DMARCReportGetResponseRecordsBimiRecord `json:"bimi_records"`
	// CNAME records for DKIM
	CNAMEDKIMRecords []DMARCReportGetResponseRecordsCnamedkimRecord `json:"cname_dkim_records"`
	// CNAME records at \_dmarc (problematic)
	CNAMEDMARCRecords []DMARCReportGetResponseRecordsCnamedmarcRecord `json:"cname_dmarc_records"`
	// CNAME records for SPF
	CNAMESPFRecords []DMARCReportGetResponseRecordsCnamespfRecord `json:"cname_spf_records"`
	// DKIM TXT records
	DKIMRecords []DMARCReportGetResponseRecordsDKIMRecord `json:"dkim_records"`
	// DMARC TXT records
	DMARCRecords []DMARCReportGetResponseRecordsDMARCRecord `json:"dmarc_records"`
	// SPF TXT records
	SPFRecords []DMARCReportGetResponseRecordsSPFRecord `json:"spf_records"`
	JSON       dmarcReportGetResponseRecordsJSON        `json:"-"`
}

// dmarcReportGetResponseRecordsJSON contains the JSON metadata for the struct
// [DMARCReportGetResponseRecords]
type dmarcReportGetResponseRecordsJSON struct {
	BimiRecords       apijson.Field
	CNAMEDKIMRecords  apijson.Field
	CNAMEDMARCRecords apijson.Field
	CNAMESPFRecords   apijson.Field
	DKIMRecords       apijson.Field
	DMARCRecords      apijson.Field
	SPFRecords        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsBimiRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                      `json:"type"`
	JSON dmarcReportGetResponseRecordsBimiRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsBimiRecordJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseRecordsBimiRecord]
type dmarcReportGetResponseRecordsBimiRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsBimiRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsBimiRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsCnamedkimRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                           `json:"type"`
	JSON dmarcReportGetResponseRecordsCnamedkimRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsCnamedkimRecordJSON contains the JSON metadata for
// the struct [DMARCReportGetResponseRecordsCnamedkimRecord]
type dmarcReportGetResponseRecordsCnamedkimRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsCnamedkimRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsCnamedkimRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsCnamedmarcRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                            `json:"type"`
	JSON dmarcReportGetResponseRecordsCnamedmarcRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsCnamedmarcRecordJSON contains the JSON metadata for
// the struct [DMARCReportGetResponseRecordsCnamedmarcRecord]
type dmarcReportGetResponseRecordsCnamedmarcRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsCnamedmarcRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsCnamedmarcRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsCnamespfRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                          `json:"type"`
	JSON dmarcReportGetResponseRecordsCnamespfRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsCnamespfRecordJSON contains the JSON metadata for
// the struct [DMARCReportGetResponseRecordsCnamespfRecord]
type dmarcReportGetResponseRecordsCnamespfRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsCnamespfRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsCnamespfRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsDKIMRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                      `json:"type"`
	JSON dmarcReportGetResponseRecordsDKIMRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsDKIMRecordJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseRecordsDKIMRecord]
type dmarcReportGetResponseRecordsDKIMRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsDKIMRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsDKIMRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsDMARCRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                       `json:"type"`
	JSON dmarcReportGetResponseRecordsDMARCRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsDMARCRecordJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseRecordsDMARCRecord]
type dmarcReportGetResponseRecordsDMARCRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsDMARCRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsDMARCRecordJSON) RawJSON() string {
	return r.raw
}

// Summary of a single DNS record
type DMARCReportGetResponseRecordsSPFRecord struct {
	// DNS record ID
	ID string `json:"id"`
	// Record content
	Content string `json:"content"`
	// DNS record name
	Name string `json:"name"`
	// Time to live in seconds
	TTL int64 `json:"ttl"`
	// Record type
	Type string                                     `json:"type"`
	JSON dmarcReportGetResponseRecordsSPFRecordJSON `json:"-"`
}

// dmarcReportGetResponseRecordsSPFRecordJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseRecordsSPFRecord]
type dmarcReportGetResponseRecordsSPFRecordJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseRecordsSPFRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseRecordsSPFRecordJSON) RawJSON() string {
	return r.raw
}

// DMARC configuration status
type DMARCReportGetResponseStatus string

const (
	DMARCReportGetResponseStatusMissingDMARCReport   DMARCReportGetResponseStatus = "missing-dmarc-report"
	DMARCReportGetResponseStatusMultipleDMARCReports DMARCReportGetResponseStatus = "multiple-dmarc-reports"
	DMARCReportGetResponseStatusMissingDMARCRua      DMARCReportGetResponseStatus = "missing-dmarc-rua"
	DMARCReportGetResponseStatusCNAMEOnDMARCRecord   DMARCReportGetResponseStatus = "cname-on-dmarc-record"
)

func (r DMARCReportGetResponseStatus) IsKnown() bool {
	switch r {
	case DMARCReportGetResponseStatusMissingDMARCReport, DMARCReportGetResponseStatusMultipleDMARCReports, DMARCReportGetResponseStatusMissingDMARCRua, DMARCReportGetResponseStatusCNAMEOnDMARCRecord:
		return true
	}
	return false
}

type DMARCReportEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Enable or disable DMARC reports for this zone
	Enabled param.Field[bool] `json:"enabled"`
	// Skip the DMARC setup wizard
	SkipWizard param.Field[bool] `json:"skip_wizard"`
}

func (r DMARCReportEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DMARCReportEditResponseEnvelope struct {
	Errors   []DMARCReportEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DMARCReportEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DMARCReportEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response for GET/PATCH /dmarc-reports
	Result DMARCReportEditResponse             `json:"result"`
	JSON   dmarcReportEditResponseEnvelopeJSON `json:"-"`
}

// dmarcReportEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DMARCReportEditResponseEnvelope]
type dmarcReportEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DMARCReportEditResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DMARCReportEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dmarcReportEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dmarcReportEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseEnvelopeErrors]
type dmarcReportEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DMARCReportEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DMARCReportEditResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dmarcReportEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dmarcReportEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DMARCReportEditResponseEnvelopeErrorsSource]
type dmarcReportEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DMARCReportEditResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DMARCReportEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dmarcReportEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dmarcReportEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DMARCReportEditResponseEnvelopeMessages]
type dmarcReportEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DMARCReportEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DMARCReportEditResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dmarcReportEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dmarcReportEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DMARCReportEditResponseEnvelopeMessagesSource]
type dmarcReportEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DMARCReportEditResponseEnvelopeSuccess bool

const (
	DMARCReportEditResponseEnvelopeSuccessTrue DMARCReportEditResponseEnvelopeSuccess = true
)

func (r DMARCReportEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DMARCReportEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DMARCReportGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type DMARCReportGetResponseEnvelope struct {
	Errors   []DMARCReportGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DMARCReportGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DMARCReportGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response for GET/PATCH /dmarc-reports
	Result DMARCReportGetResponse             `json:"result"`
	JSON   dmarcReportGetResponseEnvelopeJSON `json:"-"`
}

// dmarcReportGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DMARCReportGetResponseEnvelope]
type dmarcReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DMARCReportGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DMARCReportGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dmarcReportGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dmarcReportGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseEnvelopeErrors]
type dmarcReportGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DMARCReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DMARCReportGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dmarcReportGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dmarcReportGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DMARCReportGetResponseEnvelopeErrorsSource]
type dmarcReportGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DMARCReportGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code" api:"required"`
	Message          string                                       `json:"message" api:"required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DMARCReportGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dmarcReportGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dmarcReportGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DMARCReportGetResponseEnvelopeMessages]
type dmarcReportGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DMARCReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DMARCReportGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dmarcReportGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dmarcReportGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DMARCReportGetResponseEnvelopeMessagesSource]
type dmarcReportGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DMARCReportGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dmarcReportGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DMARCReportGetResponseEnvelopeSuccess bool

const (
	DMARCReportGetResponseEnvelopeSuccessTrue DMARCReportGetResponseEnvelopeSuccess = true
)

func (r DMARCReportGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DMARCReportGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
