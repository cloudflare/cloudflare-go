// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// ScanService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanService] method instead.
type ScanService struct {
	Options []option.RequestOption
}

// NewScanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScanService(opts ...option.RequestOption) (r *ScanService) {
	r = &ScanService{}
	r.Options = opts
	return
}

// Submit a URL to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/.
func (r *ScanService) New(ctx context.Context, accountID string, body ScanNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env ScanNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Search scans by date and webpages' requests, including full URL (after
// redirects), hostname, and path. <br/> A successful scan will appear in search
// results a few minutes after finishing but may take much longer if the system in
// under load. By default, only successfully completed scans will appear in search
// results, unless searching by `scanId`. Please take into account that older scans
// may be removed from the search index at an unspecified time.
func (r *ScanService) List(ctx context.Context, accountID string, query ScanListParams, opts ...option.RequestOption) (res *ScanListResponse, err error) {
	var env ScanListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get URL scan by uuid
func (r *ScanService) Get(ctx context.Context, accountID string, scanID string, query ScanGetParams, opts ...option.RequestOption) (res *ScanGetResponse, err error) {
	var env ScanGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *ScanService) HAR(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *ScanHARResponse, err error) {
	var env ScanHARResponseEnvelope
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/har", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *ScanService) Screenshot(ctx context.Context, accountID string, scanID string, query ScanScreenshotParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scanId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/screenshot", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ScanListResponse struct {
	Tasks []ScanListResponseTask `json:"tasks,required"`
	JSON  scanListResponseJSON   `json:"-"`
}

// scanListResponseJSON contains the JSON metadata for the struct
// [ScanListResponse]
type scanListResponseJSON struct {
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseTask struct {
	// Alpha-2 country code
	Country string `json:"country,required"`
	// Whether scan was successful or not
	Success bool `json:"success,required"`
	// When scan was submitted (UTC)
	Time time.Time `json:"time,required" format:"date-time"`
	// Scan url (after redirects)
	URL string `json:"url,required"`
	// Scan id
	UUID string `json:"uuid,required" format:"uuid"`
	// Visibility status.
	Visibility string                   `json:"visibility,required"`
	JSON       scanListResponseTaskJSON `json:"-"`
}

// scanListResponseTaskJSON contains the JSON metadata for the struct
// [ScanListResponseTask]
type scanListResponseTaskJSON struct {
	Country     apijson.Field
	Success     apijson.Field
	Time        apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseTaskJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponse struct {
	Scan ScanGetResponseScan `json:"scan,required"`
	JSON scanGetResponseJSON `json:"-"`
}

// scanGetResponseJSON contains the JSON metadata for the struct [ScanGetResponse]
type scanGetResponseJSON struct {
	Scan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScan struct {
	Certificates []ScanGetResponseScanCertificate `json:"certificates,required"`
	Geo          ScanGetResponseScanGeo           `json:"geo,required"`
	Meta         ScanGetResponseScanMeta          `json:"meta,required"`
	Page         ScanGetResponseScanPage          `json:"page,required"`
	Performance  []ScanGetResponseScanPerformance `json:"performance,required"`
	Task         ScanGetResponseScanTask          `json:"task,required"`
	Verdicts     ScanGetResponseScanVerdicts      `json:"verdicts,required"`
	// Dictionary of Autonomous System Numbers where ASN's are the keys
	ASNs    ScanGetResponseScanASNs    `json:"asns"`
	Domains ScanGetResponseScanDomains `json:"domains"`
	IPs     ScanGetResponseScanIPs     `json:"ips"`
	Links   ScanGetResponseScanLinks   `json:"links"`
	JSON    scanGetResponseScanJSON    `json:"-"`
}

// scanGetResponseScanJSON contains the JSON metadata for the struct
// [ScanGetResponseScan]
type scanGetResponseScanJSON struct {
	Certificates apijson.Field
	Geo          apijson.Field
	Meta         apijson.Field
	Page         apijson.Field
	Performance  apijson.Field
	Task         apijson.Field
	Verdicts     apijson.Field
	ASNs         apijson.Field
	Domains      apijson.Field
	IPs          apijson.Field
	Links        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanGetResponseScan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanCertificate struct {
	Issuer      string                             `json:"issuer,required"`
	SubjectName string                             `json:"subjectName,required"`
	ValidFrom   float64                            `json:"validFrom,required"`
	ValidTo     float64                            `json:"validTo,required"`
	JSON        scanGetResponseScanCertificateJSON `json:"-"`
}

// scanGetResponseScanCertificateJSON contains the JSON metadata for the struct
// [ScanGetResponseScanCertificate]
type scanGetResponseScanCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanCertificateJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanGeo struct {
	// GeoIP continent location
	Continents []string `json:"continents,required"`
	// GeoIP country location
	Locations []string                   `json:"locations,required"`
	JSON      scanGetResponseScanGeoJSON `json:"-"`
}

// scanGetResponseScanGeoJSON contains the JSON metadata for the struct
// [ScanGetResponseScanGeo]
type scanGetResponseScanGeoJSON struct {
	Continents  apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanGeoJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMeta struct {
	Processors ScanGetResponseScanMetaProcessors `json:"processors,required"`
	JSON       scanGetResponseScanMetaJSON       `json:"-"`
}

// scanGetResponseScanMetaJSON contains the JSON metadata for the struct
// [ScanGetResponseScanMeta]
type scanGetResponseScanMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessors struct {
	Categories ScanGetResponseScanMetaProcessorsCategories `json:"categories,required"`
	Phishing   []string                                    `json:"phishing,required"`
	Rank       ScanGetResponseScanMetaProcessorsRank       `json:"rank,required"`
	Tech       []ScanGetResponseScanMetaProcessorsTech     `json:"tech,required"`
	JSON       scanGetResponseScanMetaProcessorsJSON       `json:"-"`
}

// scanGetResponseScanMetaProcessorsJSON contains the JSON metadata for the struct
// [ScanGetResponseScanMetaProcessors]
type scanGetResponseScanMetaProcessorsJSON struct {
	Categories  apijson.Field
	Phishing    apijson.Field
	Rank        apijson.Field
	Tech        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsCategories struct {
	Content []URLScannerDomain                                `json:"content,required"`
	Risks   []ScanGetResponseScanMetaProcessorsCategoriesRisk `json:"risks,required"`
	JSON    scanGetResponseScanMetaProcessorsCategoriesJSON   `json:"-"`
}

// scanGetResponseScanMetaProcessorsCategoriesJSON contains the JSON metadata for
// the struct [ScanGetResponseScanMetaProcessorsCategories]
type scanGetResponseScanMetaProcessorsCategoriesJSON struct {
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsCategoriesJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsCategoriesRisk struct {
	ID              int64                                               `json:"id,required"`
	Name            string                                              `json:"name,required"`
	SuperCategoryID int64                                               `json:"super_category_id,required"`
	JSON            scanGetResponseScanMetaProcessorsCategoriesRiskJSON `json:"-"`
}

// scanGetResponseScanMetaProcessorsCategoriesRiskJSON contains the JSON metadata
// for the struct [ScanGetResponseScanMetaProcessorsCategoriesRisk]
type scanGetResponseScanMetaProcessorsCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsCategoriesRiskJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsRank struct {
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                     `json:"rank"`
	JSON scanGetResponseScanMetaProcessorsRankJSON `json:"-"`
}

// scanGetResponseScanMetaProcessorsRankJSON contains the JSON metadata for the
// struct [ScanGetResponseScanMetaProcessorsRank]
type scanGetResponseScanMetaProcessorsRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsRankJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsTech struct {
	Categories  []ScanGetResponseScanMetaProcessorsTechCategory `json:"categories,required"`
	Confidence  int64                                           `json:"confidence,required"`
	Evidence    ScanGetResponseScanMetaProcessorsTechEvidence   `json:"evidence,required"`
	Icon        string                                          `json:"icon,required"`
	Name        string                                          `json:"name,required"`
	Slug        string                                          `json:"slug,required"`
	Website     string                                          `json:"website,required"`
	Description string                                          `json:"description"`
	JSON        scanGetResponseScanMetaProcessorsTechJSON       `json:"-"`
}

// scanGetResponseScanMetaProcessorsTechJSON contains the JSON metadata for the
// struct [ScanGetResponseScanMetaProcessorsTech]
type scanGetResponseScanMetaProcessorsTechJSON struct {
	Categories  apijson.Field
	Confidence  apijson.Field
	Evidence    apijson.Field
	Icon        apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Website     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsTech) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsTechJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsTechCategory struct {
	ID       int64                                             `json:"id,required"`
	Groups   []int64                                           `json:"groups,required"`
	Name     string                                            `json:"name,required"`
	Priority int64                                             `json:"priority,required"`
	Slug     string                                            `json:"slug,required"`
	JSON     scanGetResponseScanMetaProcessorsTechCategoryJSON `json:"-"`
}

// scanGetResponseScanMetaProcessorsTechCategoryJSON contains the JSON metadata for
// the struct [ScanGetResponseScanMetaProcessorsTechCategory]
type scanGetResponseScanMetaProcessorsTechCategoryJSON struct {
	ID          apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsTechCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsTechCategoryJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsTechEvidence struct {
	ImpliedBy []string                                               `json:"impliedBy,required"`
	Patterns  []ScanGetResponseScanMetaProcessorsTechEvidencePattern `json:"patterns,required"`
	JSON      scanGetResponseScanMetaProcessorsTechEvidenceJSON      `json:"-"`
}

// scanGetResponseScanMetaProcessorsTechEvidenceJSON contains the JSON metadata for
// the struct [ScanGetResponseScanMetaProcessorsTechEvidence]
type scanGetResponseScanMetaProcessorsTechEvidenceJSON struct {
	ImpliedBy   apijson.Field
	Patterns    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsTechEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsTechEvidenceJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanMetaProcessorsTechEvidencePattern struct {
	Confidence int64    `json:"confidence,required"`
	Excludes   []string `json:"excludes,required"`
	Implies    []string `json:"implies,required"`
	Match      string   `json:"match,required"`
	// Header or Cookie name when set
	Name    string                                                   `json:"name,required"`
	Regex   string                                                   `json:"regex,required"`
	Type    string                                                   `json:"type,required"`
	Value   string                                                   `json:"value,required"`
	Version string                                                   `json:"version,required"`
	JSON    scanGetResponseScanMetaProcessorsTechEvidencePatternJSON `json:"-"`
}

// scanGetResponseScanMetaProcessorsTechEvidencePatternJSON contains the JSON
// metadata for the struct [ScanGetResponseScanMetaProcessorsTechEvidencePattern]
type scanGetResponseScanMetaProcessorsTechEvidencePatternJSON struct {
	Confidence  apijson.Field
	Excludes    apijson.Field
	Implies     apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Regex       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanMetaProcessorsTechEvidencePattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanMetaProcessorsTechEvidencePatternJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPage struct {
	ASN                   string                                     `json:"asn,required"`
	ASNLocationAlpha2     string                                     `json:"asnLocationAlpha2,required"`
	Asnname               string                                     `json:"asnname,required"`
	Console               []ScanGetResponseScanPageConsole           `json:"console,required"`
	Cookies               []ScanGetResponseScanPageCookie            `json:"cookies,required"`
	Country               string                                     `json:"country,required"`
	CountryLocationAlpha2 string                                     `json:"countryLocationAlpha2,required"`
	Domain                string                                     `json:"domain,required"`
	Headers               []ScanGetResponseScanPageHeader            `json:"headers,required"`
	IP                    string                                     `json:"ip,required"`
	JS                    ScanGetResponseScanPageJS                  `json:"js,required"`
	SecurityViolations    []ScanGetResponseScanPageSecurityViolation `json:"securityViolations,required"`
	Status                float64                                    `json:"status,required"`
	Subdivision1Name      string                                     `json:"subdivision1Name,required"`
	Subdivision2name      string                                     `json:"subdivision2name,required"`
	URL                   string                                     `json:"url,required"`
	JSON                  scanGetResponseScanPageJSON                `json:"-"`
}

// scanGetResponseScanPageJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPage]
type scanGetResponseScanPageJSON struct {
	ASN                   apijson.Field
	ASNLocationAlpha2     apijson.Field
	Asnname               apijson.Field
	Console               apijson.Field
	Cookies               apijson.Field
	Country               apijson.Field
	CountryLocationAlpha2 apijson.Field
	Domain                apijson.Field
	Headers               apijson.Field
	IP                    apijson.Field
	JS                    apijson.Field
	SecurityViolations    apijson.Field
	Status                apijson.Field
	Subdivision1Name      apijson.Field
	Subdivision2name      apijson.Field
	URL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ScanGetResponseScanPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPageConsole struct {
	Category string                             `json:"category,required"`
	Text     string                             `json:"text,required"`
	Type     string                             `json:"type,required"`
	URL      string                             `json:"url"`
	JSON     scanGetResponseScanPageConsoleJSON `json:"-"`
}

// scanGetResponseScanPageConsoleJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPageConsole]
type scanGetResponseScanPageConsoleJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanPageConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageConsoleJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPageCookie struct {
	Domain       string                            `json:"domain,required"`
	Expires      float64                           `json:"expires,required"`
	HTTPOnly     bool                              `json:"httpOnly,required"`
	Name         string                            `json:"name,required"`
	Path         string                            `json:"path,required"`
	SameParty    bool                              `json:"sameParty,required"`
	Secure       bool                              `json:"secure,required"`
	Session      bool                              `json:"session,required"`
	Size         float64                           `json:"size,required"`
	SourcePort   float64                           `json:"sourcePort,required"`
	SourceScheme string                            `json:"sourceScheme,required"`
	Value        string                            `json:"value,required"`
	Priority     string                            `json:"priority"`
	JSON         scanGetResponseScanPageCookieJSON `json:"-"`
}

// scanGetResponseScanPageCookieJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPageCookie]
type scanGetResponseScanPageCookieJSON struct {
	Domain       apijson.Field
	Expires      apijson.Field
	HTTPOnly     apijson.Field
	Name         apijson.Field
	Path         apijson.Field
	SameParty    apijson.Field
	Secure       apijson.Field
	Session      apijson.Field
	Size         apijson.Field
	SourcePort   apijson.Field
	SourceScheme apijson.Field
	Value        apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanGetResponseScanPageCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageCookieJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPageHeader struct {
	Name  string                            `json:"name,required"`
	Value string                            `json:"value,required"`
	JSON  scanGetResponseScanPageHeaderJSON `json:"-"`
}

// scanGetResponseScanPageHeaderJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPageHeader]
type scanGetResponseScanPageHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanPageHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPageJS struct {
	Variables []ScanGetResponseScanPageJSVariable `json:"variables,required"`
	JSON      scanGetResponseScanPageJSJSON       `json:"-"`
}

// scanGetResponseScanPageJSJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPageJS]
type scanGetResponseScanPageJSJSON struct {
	Variables   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanPageJS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageJSJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPageJSVariable struct {
	Name string                                `json:"name,required"`
	Type string                                `json:"type,required"`
	JSON scanGetResponseScanPageJSVariableJSON `json:"-"`
}

// scanGetResponseScanPageJSVariableJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPageJSVariable]
type scanGetResponseScanPageJSVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanPageJSVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageJSVariableJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPageSecurityViolation struct {
	Category string                                       `json:"category,required"`
	Text     string                                       `json:"text,required"`
	URL      string                                       `json:"url,required"`
	JSON     scanGetResponseScanPageSecurityViolationJSON `json:"-"`
}

// scanGetResponseScanPageSecurityViolationJSON contains the JSON metadata for the
// struct [ScanGetResponseScanPageSecurityViolation]
type scanGetResponseScanPageSecurityViolationJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanPageSecurityViolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPageSecurityViolationJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanPerformance struct {
	ConnectEnd                 float64                            `json:"connectEnd,required"`
	ConnectStart               float64                            `json:"connectStart,required"`
	DecodedBodySize            float64                            `json:"decodedBodySize,required"`
	DomainLookupEnd            float64                            `json:"domainLookupEnd,required"`
	DomainLookupStart          float64                            `json:"domainLookupStart,required"`
	DOMComplete                float64                            `json:"domComplete,required"`
	DOMContentLoadedEventEnd   float64                            `json:"domContentLoadedEventEnd,required"`
	DOMContentLoadedEventStart float64                            `json:"domContentLoadedEventStart,required"`
	DOMInteractive             float64                            `json:"domInteractive,required"`
	Duration                   float64                            `json:"duration,required"`
	EncodedBodySize            float64                            `json:"encodedBodySize,required"`
	EntryType                  string                             `json:"entryType,required"`
	FetchStart                 float64                            `json:"fetchStart,required"`
	InitiatorType              string                             `json:"initiatorType,required"`
	LoadEventEnd               float64                            `json:"loadEventEnd,required"`
	LoadEventStart             float64                            `json:"loadEventStart,required"`
	Name                       string                             `json:"name,required"`
	NextHopProtocol            string                             `json:"nextHopProtocol,required"`
	RedirectCount              float64                            `json:"redirectCount,required"`
	RedirectEnd                float64                            `json:"redirectEnd,required"`
	RedirectStart              float64                            `json:"redirectStart,required"`
	RequestStart               float64                            `json:"requestStart,required"`
	ResponseEnd                float64                            `json:"responseEnd,required"`
	ResponseStart              float64                            `json:"responseStart,required"`
	SecureConnectionStart      float64                            `json:"secureConnectionStart,required"`
	StartTime                  float64                            `json:"startTime,required"`
	TransferSize               float64                            `json:"transferSize,required"`
	Type                       string                             `json:"type,required"`
	UnloadEventEnd             float64                            `json:"unloadEventEnd,required"`
	UnloadEventStart           float64                            `json:"unloadEventStart,required"`
	WorkerStart                float64                            `json:"workerStart,required"`
	JSON                       scanGetResponseScanPerformanceJSON `json:"-"`
}

// scanGetResponseScanPerformanceJSON contains the JSON metadata for the struct
// [ScanGetResponseScanPerformance]
type scanGetResponseScanPerformanceJSON struct {
	ConnectEnd                 apijson.Field
	ConnectStart               apijson.Field
	DecodedBodySize            apijson.Field
	DomainLookupEnd            apijson.Field
	DomainLookupStart          apijson.Field
	DOMComplete                apijson.Field
	DOMContentLoadedEventEnd   apijson.Field
	DOMContentLoadedEventStart apijson.Field
	DOMInteractive             apijson.Field
	Duration                   apijson.Field
	EncodedBodySize            apijson.Field
	EntryType                  apijson.Field
	FetchStart                 apijson.Field
	InitiatorType              apijson.Field
	LoadEventEnd               apijson.Field
	LoadEventStart             apijson.Field
	Name                       apijson.Field
	NextHopProtocol            apijson.Field
	RedirectCount              apijson.Field
	RedirectEnd                apijson.Field
	RedirectStart              apijson.Field
	RequestStart               apijson.Field
	ResponseEnd                apijson.Field
	ResponseStart              apijson.Field
	SecureConnectionStart      apijson.Field
	StartTime                  apijson.Field
	TransferSize               apijson.Field
	Type                       apijson.Field
	UnloadEventEnd             apijson.Field
	UnloadEventStart           apijson.Field
	WorkerStart                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ScanGetResponseScanPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanPerformanceJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanTask struct {
	// Submitter location
	ClientLocation string                            `json:"clientLocation,required"`
	ClientType     ScanGetResponseScanTaskClientType `json:"clientType,required"`
	// URL of the primary request, after all HTTP redirects
	EffectiveURL string                             `json:"effectiveUrl,required"`
	Errors       []ScanGetResponseScanTaskError     `json:"errors,required"`
	ScannedFrom  ScanGetResponseScanTaskScannedFrom `json:"scannedFrom,required"`
	Status       ScanGetResponseScanTaskStatus      `json:"status,required"`
	Success      bool                               `json:"success,required"`
	Time         string                             `json:"time,required"`
	TimeEnd      string                             `json:"timeEnd,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID
	UUID       string                            `json:"uuid,required"`
	Visibility ScanGetResponseScanTaskVisibility `json:"visibility,required"`
	JSON       scanGetResponseScanTaskJSON       `json:"-"`
}

// scanGetResponseScanTaskJSON contains the JSON metadata for the struct
// [ScanGetResponseScanTask]
type scanGetResponseScanTaskJSON struct {
	ClientLocation apijson.Field
	ClientType     apijson.Field
	EffectiveURL   apijson.Field
	Errors         apijson.Field
	ScannedFrom    apijson.Field
	Status         apijson.Field
	Success        apijson.Field
	Time           apijson.Field
	TimeEnd        apijson.Field
	URL            apijson.Field
	UUID           apijson.Field
	Visibility     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScanGetResponseScanTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanTaskJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanTaskClientType string

const (
	ScanGetResponseScanTaskClientTypeSite      ScanGetResponseScanTaskClientType = "Site"
	ScanGetResponseScanTaskClientTypeAutomatic ScanGetResponseScanTaskClientType = "Automatic"
	ScanGetResponseScanTaskClientTypeAPI       ScanGetResponseScanTaskClientType = "Api"
)

func (r ScanGetResponseScanTaskClientType) IsKnown() bool {
	switch r {
	case ScanGetResponseScanTaskClientTypeSite, ScanGetResponseScanTaskClientTypeAutomatic, ScanGetResponseScanTaskClientTypeAPI:
		return true
	}
	return false
}

type ScanGetResponseScanTaskError struct {
	Message string                           `json:"message,required"`
	JSON    scanGetResponseScanTaskErrorJSON `json:"-"`
}

// scanGetResponseScanTaskErrorJSON contains the JSON metadata for the struct
// [ScanGetResponseScanTaskError]
type scanGetResponseScanTaskErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanTaskError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanTaskErrorJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanTaskScannedFrom struct {
	// IATA code of Cloudflare datacenter
	Colo string                                 `json:"colo,required"`
	JSON scanGetResponseScanTaskScannedFromJSON `json:"-"`
}

// scanGetResponseScanTaskScannedFromJSON contains the JSON metadata for the struct
// [ScanGetResponseScanTaskScannedFrom]
type scanGetResponseScanTaskScannedFromJSON struct {
	Colo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanTaskScannedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanTaskScannedFromJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanTaskStatus string

const (
	ScanGetResponseScanTaskStatusQueued           ScanGetResponseScanTaskStatus = "Queued"
	ScanGetResponseScanTaskStatusInProgress       ScanGetResponseScanTaskStatus = "InProgress"
	ScanGetResponseScanTaskStatusInPostProcessing ScanGetResponseScanTaskStatus = "InPostProcessing"
	ScanGetResponseScanTaskStatusFinished         ScanGetResponseScanTaskStatus = "Finished"
)

func (r ScanGetResponseScanTaskStatus) IsKnown() bool {
	switch r {
	case ScanGetResponseScanTaskStatusQueued, ScanGetResponseScanTaskStatusInProgress, ScanGetResponseScanTaskStatusInPostProcessing, ScanGetResponseScanTaskStatusFinished:
		return true
	}
	return false
}

type ScanGetResponseScanTaskVisibility string

const (
	ScanGetResponseScanTaskVisibilityPublic   ScanGetResponseScanTaskVisibility = "Public"
	ScanGetResponseScanTaskVisibilityUnlisted ScanGetResponseScanTaskVisibility = "Unlisted"
)

func (r ScanGetResponseScanTaskVisibility) IsKnown() bool {
	switch r {
	case ScanGetResponseScanTaskVisibilityPublic, ScanGetResponseScanTaskVisibilityUnlisted:
		return true
	}
	return false
}

type ScanGetResponseScanVerdicts struct {
	Overall ScanGetResponseScanVerdictsOverall `json:"overall,required"`
	JSON    scanGetResponseScanVerdictsJSON    `json:"-"`
}

// scanGetResponseScanVerdictsJSON contains the JSON metadata for the struct
// [ScanGetResponseScanVerdicts]
type scanGetResponseScanVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanVerdictsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanVerdictsOverall struct {
	Categories []ScanGetResponseScanVerdictsOverallCategory `json:"categories,required"`
	// At least one of our subsystems marked the site as potentially malicious at the
	// time of the scan.
	Malicious bool                                   `json:"malicious,required"`
	Phishing  []string                               `json:"phishing,required"`
	JSON      scanGetResponseScanVerdictsOverallJSON `json:"-"`
}

// scanGetResponseScanVerdictsOverallJSON contains the JSON metadata for the struct
// [ScanGetResponseScanVerdictsOverall]
type scanGetResponseScanVerdictsOverallJSON struct {
	Categories  apijson.Field
	Malicious   apijson.Field
	Phishing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanVerdictsOverallJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanVerdictsOverallCategory struct {
	ID              float64                                        `json:"id,required"`
	Name            string                                         `json:"name,required"`
	SuperCategoryID float64                                        `json:"super_category_id,required"`
	JSON            scanGetResponseScanVerdictsOverallCategoryJSON `json:"-"`
}

// scanGetResponseScanVerdictsOverallCategoryJSON contains the JSON metadata for
// the struct [ScanGetResponseScanVerdictsOverallCategory]
type scanGetResponseScanVerdictsOverallCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanGetResponseScanVerdictsOverallCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanVerdictsOverallCategoryJSON) RawJSON() string {
	return r.raw
}

// Dictionary of Autonomous System Numbers where ASN's are the keys
type ScanGetResponseScanASNs struct {
	// ASN's contacted
	ASN  ScanGetResponseScanASNsASN  `json:"asn"`
	JSON scanGetResponseScanASNsJSON `json:"-"`
}

// scanGetResponseScanASNsJSON contains the JSON metadata for the struct
// [ScanGetResponseScanASNs]
type scanGetResponseScanASNsJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanASNs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanASNsJSON) RawJSON() string {
	return r.raw
}

// ASN's contacted
type ScanGetResponseScanASNsASN struct {
	ASN            string                         `json:"asn,required"`
	Description    string                         `json:"description,required"`
	LocationAlpha2 string                         `json:"location_alpha2,required"`
	Name           string                         `json:"name,required"`
	OrgName        string                         `json:"org_name,required"`
	JSON           scanGetResponseScanASNsASNJSON `json:"-"`
}

// scanGetResponseScanASNsASNJSON contains the JSON metadata for the struct
// [ScanGetResponseScanASNsASN]
type scanGetResponseScanASNsASNJSON struct {
	ASN            apijson.Field
	Description    apijson.Field
	LocationAlpha2 apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ScanGetResponseScanASNsASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanASNsASNJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanDomains struct {
	ExampleCom ScanGetResponseScanDomainsExampleCom `json:"example.com"`
	JSON       scanGetResponseScanDomainsJSON       `json:"-"`
}

// scanGetResponseScanDomainsJSON contains the JSON metadata for the struct
// [ScanGetResponseScanDomains]
type scanGetResponseScanDomainsJSON struct {
	ExampleCom  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanDomainsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanDomainsExampleCom struct {
	Categories ScanGetResponseScanDomainsExampleComCategories `json:"categories,required"`
	DNS        []ScanGetResponseScanDomainsExampleComDNS      `json:"dns,required"`
	Name       string                                         `json:"name,required"`
	Rank       ScanGetResponseScanDomainsExampleComRank       `json:"rank,required"`
	Type       string                                         `json:"type,required"`
	JSON       scanGetResponseScanDomainsExampleComJSON       `json:"-"`
}

// scanGetResponseScanDomainsExampleComJSON contains the JSON metadata for the
// struct [ScanGetResponseScanDomainsExampleCom]
type scanGetResponseScanDomainsExampleComJSON struct {
	Categories  apijson.Field
	DNS         apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanDomainsExampleCom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanDomainsExampleComJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanDomainsExampleComCategories struct {
	Inherited ScanGetResponseScanDomainsExampleComCategoriesInherited `json:"inherited,required"`
	Content   []URLScannerDomain                                      `json:"content"`
	Risks     []URLScannerDomain                                      `json:"risks"`
	JSON      scanGetResponseScanDomainsExampleComCategoriesJSON      `json:"-"`
}

// scanGetResponseScanDomainsExampleComCategoriesJSON contains the JSON metadata
// for the struct [ScanGetResponseScanDomainsExampleComCategories]
type scanGetResponseScanDomainsExampleComCategoriesJSON struct {
	Inherited   apijson.Field
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanDomainsExampleComCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanDomainsExampleComCategoriesJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanDomainsExampleComCategoriesInherited struct {
	Content []URLScannerDomain                                          `json:"content"`
	From    string                                                      `json:"from"`
	Risks   []URLScannerDomain                                          `json:"risks"`
	JSON    scanGetResponseScanDomainsExampleComCategoriesInheritedJSON `json:"-"`
}

// scanGetResponseScanDomainsExampleComCategoriesInheritedJSON contains the JSON
// metadata for the struct
// [ScanGetResponseScanDomainsExampleComCategoriesInherited]
type scanGetResponseScanDomainsExampleComCategoriesInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanDomainsExampleComCategoriesInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanDomainsExampleComCategoriesInheritedJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanDomainsExampleComDNS struct {
	Address     string                                      `json:"address,required"`
	DNSSECValid bool                                        `json:"dnssec_valid,required"`
	Name        string                                      `json:"name,required"`
	Type        string                                      `json:"type,required"`
	JSON        scanGetResponseScanDomainsExampleComDNSJSON `json:"-"`
}

// scanGetResponseScanDomainsExampleComDNSJSON contains the JSON metadata for the
// struct [ScanGetResponseScanDomainsExampleComDNS]
type scanGetResponseScanDomainsExampleComDNSJSON struct {
	Address     apijson.Field
	DNSSECValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanDomainsExampleComDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanDomainsExampleComDNSJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanDomainsExampleComRank struct {
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                        `json:"rank"`
	JSON scanGetResponseScanDomainsExampleComRankJSON `json:"-"`
}

// scanGetResponseScanDomainsExampleComRankJSON contains the JSON metadata for the
// struct [ScanGetResponseScanDomainsExampleComRank]
type scanGetResponseScanDomainsExampleComRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanDomainsExampleComRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanDomainsExampleComRankJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanIPs struct {
	IP   ScanGetResponseScanIPsIP   `json:"ip"`
	JSON scanGetResponseScanIPsJSON `json:"-"`
}

// scanGetResponseScanIPsJSON contains the JSON metadata for the struct
// [ScanGetResponseScanIPs]
type scanGetResponseScanIPsJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanIPsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanIPsIP struct {
	ASN               string                       `json:"asn,required"`
	ASNDescription    string                       `json:"asnDescription,required"`
	ASNLocationAlpha2 string                       `json:"asnLocationAlpha2,required"`
	ASNName           string                       `json:"asnName,required"`
	ASNOrgName        string                       `json:"asnOrgName,required"`
	Continent         string                       `json:"continent,required"`
	GeonameID         string                       `json:"geonameId,required"`
	IP                string                       `json:"ip,required"`
	IPVersion         string                       `json:"ipVersion,required"`
	Latitude          string                       `json:"latitude,required"`
	LocationAlpha2    string                       `json:"locationAlpha2,required"`
	LocationName      string                       `json:"locationName,required"`
	Longitude         string                       `json:"longitude,required"`
	Subdivision1Name  string                       `json:"subdivision1Name,required"`
	Subdivision2Name  string                       `json:"subdivision2Name,required"`
	JSON              scanGetResponseScanIPsIPJSON `json:"-"`
}

// scanGetResponseScanIPsIPJSON contains the JSON metadata for the struct
// [ScanGetResponseScanIPsIP]
type scanGetResponseScanIPsIPJSON struct {
	ASN               apijson.Field
	ASNDescription    apijson.Field
	ASNLocationAlpha2 apijson.Field
	ASNName           apijson.Field
	ASNOrgName        apijson.Field
	Continent         apijson.Field
	GeonameID         apijson.Field
	IP                apijson.Field
	IPVersion         apijson.Field
	Latitude          apijson.Field
	LocationAlpha2    apijson.Field
	LocationName      apijson.Field
	Longitude         apijson.Field
	Subdivision1Name  apijson.Field
	Subdivision2Name  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ScanGetResponseScanIPsIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanIPsIPJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanLinks struct {
	Link ScanGetResponseScanLinksLink `json:"link"`
	JSON scanGetResponseScanLinksJSON `json:"-"`
}

// scanGetResponseScanLinksJSON contains the JSON metadata for the struct
// [ScanGetResponseScanLinks]
type scanGetResponseScanLinksJSON struct {
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanLinksJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseScanLinksLink struct {
	// Outgoing link detected in the DOM
	Href string                           `json:"href,required"`
	Text string                           `json:"text,required"`
	JSON scanGetResponseScanLinksLinkJSON `json:"-"`
}

// scanGetResponseScanLinksLinkJSON contains the JSON metadata for the struct
// [ScanGetResponseScanLinksLink]
type scanGetResponseScanLinksLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseScanLinksLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseScanLinksLinkJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponse struct {
	HAR  ScanHARResponseHAR  `json:"har,required"`
	JSON scanHARResponseJSON `json:"-"`
}

// scanHARResponseJSON contains the JSON metadata for the struct [ScanHARResponse]
type scanHARResponseJSON struct {
	HAR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHAR struct {
	Log  ScanHARResponseHARLog  `json:"log,required"`
	JSON scanHARResponseHARJSON `json:"-"`
}

// scanHARResponseHARJSON contains the JSON metadata for the struct
// [ScanHARResponseHAR]
type scanHARResponseHARJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHAR) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLog struct {
	Creator ScanHARResponseHARLogCreator `json:"creator,required"`
	Entries []ScanHARResponseHARLogEntry `json:"entries,required"`
	Pages   []ScanHARResponseHARLogPage  `json:"pages,required"`
	Version string                       `json:"version,required"`
	JSON    scanHARResponseHARLogJSON    `json:"-"`
}

// scanHARResponseHARLogJSON contains the JSON metadata for the struct
// [ScanHARResponseHARLog]
type scanHARResponseHARLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHARLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogCreator struct {
	Comment string                           `json:"comment,required"`
	Name    string                           `json:"name,required"`
	Version string                           `json:"version,required"`
	JSON    scanHARResponseHARLogCreatorJSON `json:"-"`
}

// scanHARResponseHARLogCreatorJSON contains the JSON metadata for the struct
// [ScanHARResponseHARLogCreator]
type scanHARResponseHARLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHARLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogCreatorJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogEntry struct {
	InitialPriority string                               `json:"_initialPriority,required"`
	InitiatorType   string                               `json:"_initiator_type,required"`
	Priority        string                               `json:"_priority,required"`
	RequestID       string                               `json:"_requestId,required"`
	RequestTime     float64                              `json:"_requestTime,required"`
	ResourceType    string                               `json:"_resourceType,required"`
	Cache           interface{}                          `json:"cache,required"`
	Connection      string                               `json:"connection,required"`
	Pageref         string                               `json:"pageref,required"`
	Request         ScanHARResponseHARLogEntriesRequest  `json:"request,required"`
	Response        ScanHARResponseHARLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                               `json:"serverIPAddress,required"`
	StartedDateTime string                               `json:"startedDateTime,required"`
	Time            float64                              `json:"time,required"`
	JSON            scanHARResponseHARLogEntryJSON       `json:"-"`
}

// scanHARResponseHARLogEntryJSON contains the JSON metadata for the struct
// [ScanHARResponseHARLogEntry]
type scanHARResponseHARLogEntryJSON struct {
	InitialPriority apijson.Field
	InitiatorType   apijson.Field
	Priority        apijson.Field
	RequestID       apijson.Field
	RequestTime     apijson.Field
	ResourceType    apijson.Field
	Cache           apijson.Field
	Connection      apijson.Field
	Pageref         apijson.Field
	Request         apijson.Field
	Response        apijson.Field
	ServerIPAddress apijson.Field
	StartedDateTime apijson.Field
	Time            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanHARResponseHARLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogEntryJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogEntriesRequest struct {
	BodySize    float64                                     `json:"bodySize,required"`
	Headers     []ScanHARResponseHARLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                     `json:"headersSize,required"`
	HTTPVersion string                                      `json:"httpVersion,required"`
	Method      string                                      `json:"method,required"`
	URL         string                                      `json:"url,required"`
	JSON        scanHARResponseHARLogEntriesRequestJSON     `json:"-"`
}

// scanHARResponseHARLogEntriesRequestJSON contains the JSON metadata for the
// struct [ScanHARResponseHARLogEntriesRequest]
type scanHARResponseHARLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHARLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogEntriesRequestHeader struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  scanHARResponseHARLogEntriesRequestHeaderJSON `json:"-"`
}

// scanHARResponseHARLogEntriesRequestHeaderJSON contains the JSON metadata for the
// struct [ScanHARResponseHARLogEntriesRequestHeader]
type scanHARResponseHARLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHARLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogEntriesRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogEntriesResponse struct {
	TransferSize float64                                      `json:"_transferSize,required"`
	BodySize     float64                                      `json:"bodySize,required"`
	Content      ScanHARResponseHARLogEntriesResponseContent  `json:"content,required"`
	Headers      []ScanHARResponseHARLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                      `json:"headersSize,required"`
	HTTPVersion  string                                       `json:"httpVersion,required"`
	RedirectURL  string                                       `json:"redirectURL,required"`
	Status       float64                                      `json:"status,required"`
	StatusText   string                                       `json:"statusText,required"`
	JSON         scanHARResponseHARLogEntriesResponseJSON     `json:"-"`
}

// scanHARResponseHARLogEntriesResponseJSON contains the JSON metadata for the
// struct [ScanHARResponseHARLogEntriesResponse]
type scanHARResponseHARLogEntriesResponseJSON struct {
	TransferSize apijson.Field
	BodySize     apijson.Field
	Content      apijson.Field
	Headers      apijson.Field
	HeadersSize  apijson.Field
	HTTPVersion  apijson.Field
	RedirectURL  apijson.Field
	Status       apijson.Field
	StatusText   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScanHARResponseHARLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogEntriesResponseContent struct {
	MimeType    string                                          `json:"mimeType,required"`
	Size        float64                                         `json:"size,required"`
	Compression int64                                           `json:"compression"`
	JSON        scanHARResponseHARLogEntriesResponseContentJSON `json:"-"`
}

// scanHARResponseHARLogEntriesResponseContentJSON contains the JSON metadata for
// the struct [ScanHARResponseHARLogEntriesResponseContent]
type scanHARResponseHARLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHARLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogEntriesResponseContentJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogEntriesResponseHeader struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  scanHARResponseHARLogEntriesResponseHeaderJSON `json:"-"`
}

// scanHARResponseHARLogEntriesResponseHeaderJSON contains the JSON metadata for
// the struct [ScanHARResponseHARLogEntriesResponseHeader]
type scanHARResponseHARLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseHARLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogEntriesResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogPage struct {
	ID              string                                `json:"id,required"`
	PageTimings     ScanHARResponseHARLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                                `json:"startedDateTime,required"`
	Title           string                                `json:"title,required"`
	JSON            scanHARResponseHARLogPageJSON         `json:"-"`
}

// scanHARResponseHARLogPageJSON contains the JSON metadata for the struct
// [ScanHARResponseHARLogPage]
type scanHARResponseHARLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ScanHARResponseHARLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogPageJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseHARLogPagesPageTimings struct {
	OnContentLoad float64                                   `json:"onContentLoad,required"`
	OnLoad        float64                                   `json:"onLoad,required"`
	JSON          scanHARResponseHARLogPagesPageTimingsJSON `json:"-"`
}

// scanHARResponseHARLogPagesPageTimingsJSON contains the JSON metadata for the
// struct [ScanHARResponseHARLogPagesPageTimings]
type scanHARResponseHARLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanHARResponseHARLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseHARLogPagesPageTimingsJSON) RawJSON() string {
	return r.raw
}

type ScanNewParams struct {
	URL         param.Field[string] `json:"url,required"`
	Customagent param.Field[string] `json:"customagent"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	Referer       param.Field[string]            `json:"referer"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]ScanNewParamsScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[ScanNewParamsVisibility] `json:"visibility"`
}

func (r ScanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Device resolutions.
type ScanNewParamsScreenshotsResolution string

const (
	ScanNewParamsScreenshotsResolutionDesktop ScanNewParamsScreenshotsResolution = "desktop"
	ScanNewParamsScreenshotsResolutionMobile  ScanNewParamsScreenshotsResolution = "mobile"
	ScanNewParamsScreenshotsResolutionTablet  ScanNewParamsScreenshotsResolution = "tablet"
)

func (r ScanNewParamsScreenshotsResolution) IsKnown() bool {
	switch r {
	case ScanNewParamsScreenshotsResolutionDesktop, ScanNewParamsScreenshotsResolutionMobile, ScanNewParamsScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type ScanNewParamsVisibility string

const (
	ScanNewParamsVisibilityPublic   ScanNewParamsVisibility = "Public"
	ScanNewParamsVisibilityUnlisted ScanNewParamsVisibility = "Unlisted"
)

func (r ScanNewParamsVisibility) IsKnown() bool {
	switch r {
	case ScanNewParamsVisibilityPublic, ScanNewParamsVisibilityUnlisted:
		return true
	}
	return false
}

type ScanNewResponseEnvelope struct {
	// URL to api report.
	API     string `json:"api,required"`
	Message string `json:"message,required"`
	// URL to report.
	Result string `json:"result,required"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url,required"`
	// Scan ID.
	UUID string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                         `json:"visibility,required"`
	Options    ScanNewResponseEnvelopeOptions `json:"options"`
	JSON       scanNewResponseEnvelopeJSON    `json:"-"`
}

// scanNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanNewResponseEnvelope]
type scanNewResponseEnvelopeJSON struct {
	API         apijson.Field
	Message     apijson.Field
	Result      apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	Visibility  apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanNewResponseEnvelopeOptions struct {
	Useragent string                             `json:"useragent"`
	JSON      scanNewResponseEnvelopeOptionsJSON `json:"-"`
}

// scanNewResponseEnvelopeOptionsJSON contains the JSON metadata for the struct
// [ScanNewResponseEnvelopeOptions]
type scanNewResponseEnvelopeOptionsJSON struct {
	Useragent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanNewResponseEnvelopeOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanNewResponseEnvelopeOptionsJSON) RawJSON() string {
	return r.raw
}

type ScanListParams struct {
	// Return only scans created by account.
	AccountScans param.Field[bool] `query:"account_scans"`
	// Filter scans by Autonomous System Number (ASN) of _any_ request made by the
	// webpage.
	ASN param.Field[string] `query:"asn"`
	// Filter scans requested before date (inclusive).
	DateEnd param.Field[time.Time] `query:"date_end" format:"date-time"`
	// Filter scans requested after date (inclusive).
	DateStart param.Field[time.Time] `query:"date_start" format:"date-time"`
	// Filter scans by hash of any html/js/css request made by the webpage.
	Hash param.Field[string] `query:"hash"`
	// Filter scans by hostname of _any_ request made by the webpage.
	Hostname param.Field[string] `query:"hostname"`
	// Filter scans by IP address (IPv4 or IPv6) of _any_ request made by the webpage.
	IP param.Field[string] `query:"ip"`
	// Filter scans by malicious verdict.
	IsMalicious param.Field[bool] `query:"is_malicious"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Pagination cursor to get the next set of results.
	NextCursor param.Field[string] `query:"next_cursor"`
	// Filter scans by main page Autonomous System Number (ASN).
	PageASN param.Field[string] `query:"page_asn"`
	// Filter scans by main page hostname (domain of effective URL).
	PageHostname param.Field[string] `query:"page_hostname"`
	// Filter scans by main page IP address (IPv4 or IPv6).
	PageIP param.Field[string] `query:"page_ip"`
	// Filter scans by exact match of effective URL path (also supports suffix search).
	PagePath param.Field[string] `query:"page_path"`
	// Filter scans by submitted or scanned URL
	PageURL param.Field[string] `query:"page_url"`
	// Filter scans by url path of _any_ request made by the webpage.
	Path param.Field[string] `query:"path"`
	// Scan uuid
	ScanID param.Field[string] `query:"scanId" format:"uuid"`
	// Filter scans by URL of _any_ request made by the webpage
	URL param.Field[string] `query:"url"`
}

// URLQuery serializes [ScanListParams]'s query parameters as `url.Values`.
func (r ScanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScanListResponseEnvelope struct {
	Errors   []ScanListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScanListResponseEnvelopeMessages `json:"messages,required"`
	Result   ScanListResponse                   `json:"result,required"`
	// Whether search request was successful or not
	Success bool                         `json:"success,required"`
	JSON    scanListResponseEnvelopeJSON `json:"-"`
}

// scanListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanListResponseEnvelope]
type scanListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseEnvelopeErrors struct {
	Message string                             `json:"message,required"`
	JSON    scanListResponseEnvelopeErrorsJSON `json:"-"`
}

// scanListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScanListResponseEnvelopeErrors]
type scanListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScanListResponseEnvelopeMessages struct {
	Message string                               `json:"message,required"`
	JSON    scanListResponseEnvelopeMessagesJSON `json:"-"`
}

// scanListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ScanListResponseEnvelopeMessages]
type scanListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScanGetParams struct {
	// Whether to return full report (scan summary and network log).
	Full param.Field[bool] `query:"full"`
}

// URLQuery serializes [ScanGetParams]'s query parameters as `url.Values`.
func (r ScanGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ScanGetResponseEnvelope struct {
	Errors   []ScanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ScanGetResponse                   `json:"result,required"`
	// Whether request was successful or not
	Success bool                        `json:"success,required"`
	JSON    scanGetResponseEnvelopeJSON `json:"-"`
}

// scanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanGetResponseEnvelope]
type scanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseEnvelopeErrors struct {
	Message string                            `json:"message,required"`
	JSON    scanGetResponseEnvelopeErrorsJSON `json:"-"`
}

// scanGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScanGetResponseEnvelopeErrors]
type scanGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScanGetResponseEnvelopeMessages struct {
	Message string                              `json:"message,required"`
	JSON    scanGetResponseEnvelopeMessagesJSON `json:"-"`
}

// scanGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ScanGetResponseEnvelopeMessages]
type scanGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseEnvelope struct {
	Errors   []ScanHARResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScanHARResponseEnvelopeMessages `json:"messages,required"`
	Result   ScanHARResponse                   `json:"result,required"`
	// Whether search request was successful or not
	Success bool                        `json:"success,required"`
	JSON    scanHARResponseEnvelopeJSON `json:"-"`
}

// scanHARResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanHARResponseEnvelope]
type scanHARResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseEnvelopeErrors struct {
	Message string                            `json:"message,required"`
	JSON    scanHARResponseEnvelopeErrorsJSON `json:"-"`
}

// scanHARResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ScanHARResponseEnvelopeErrors]
type scanHARResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScanHARResponseEnvelopeMessages struct {
	Message string                              `json:"message,required"`
	JSON    scanHARResponseEnvelopeMessagesJSON `json:"-"`
}

// scanHARResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ScanHARResponseEnvelopeMessages]
type scanHARResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanHARResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanHARResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScanScreenshotParams struct {
	// Target device type.
	Resolution param.Field[ScanScreenshotParamsResolution] `query:"resolution"`
}

// URLQuery serializes [ScanScreenshotParams]'s query parameters as `url.Values`.
func (r ScanScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Target device type.
type ScanScreenshotParamsResolution string

const (
	ScanScreenshotParamsResolutionDesktop ScanScreenshotParamsResolution = "desktop"
	ScanScreenshotParamsResolutionMobile  ScanScreenshotParamsResolution = "mobile"
	ScanScreenshotParamsResolutionTablet  ScanScreenshotParamsResolution = "tablet"
)

func (r ScanScreenshotParamsResolution) IsKnown() bool {
	switch r {
	case ScanScreenshotParamsResolutionDesktop, ScanScreenshotParamsResolutionMobile, ScanScreenshotParamsResolutionTablet:
		return true
	}
	return false
}
