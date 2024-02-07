// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// URLScannerScanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewURLScannerScanService] method
// instead.
type URLScannerScanService struct {
	Options []option.RequestOption
}

// NewURLScannerScanService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewURLScannerScanService(opts ...option.RequestOption) (r *URLScannerScanService) {
	r = &URLScannerScanService{}
	r.Options = opts
	return
}

// Submit a URL to scan. You can also set some options, like the visibility level
// and custom headers. Accounts are limited to 1 new scan every 10 seconds and 8000
// per month. If you need more, please reach out.
func (r *URLScannerScanService) New(ctx context.Context, accountID string, body URLScannerScanNewParams, opts ...option.RequestOption) (res *URLScannerScanNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get URL scan by uuid
func (r *URLScannerScanService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *URLScannerScanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *URLScannerScanService) Har(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *URLScannerScanHarResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/har", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *URLScannerScanService) Screenshot(ctx context.Context, accountID string, scanID string, query URLScannerScanScreenshotParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/screenshot", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type URLScannerScanNewResponse struct {
	Errors   []URLScannerScanNewResponseError   `json:"errors,required"`
	Messages []URLScannerScanNewResponseMessage `json:"messages,required"`
	Result   URLScannerScanNewResponseResult    `json:"result,required"`
	Success  bool                               `json:"success,required"`
	JSON     urlScannerScanNewResponseJSON      `json:"-"`
}

// urlScannerScanNewResponseJSON contains the JSON metadata for the struct
// [URLScannerScanNewResponse]
type urlScannerScanNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewResponseError struct {
	Message string                             `json:"message,required"`
	JSON    urlScannerScanNewResponseErrorJSON `json:"-"`
}

// urlScannerScanNewResponseErrorJSON contains the JSON metadata for the struct
// [URLScannerScanNewResponseError]
type urlScannerScanNewResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewResponseMessage struct {
	Message string                               `json:"message,required"`
	JSON    urlScannerScanNewResponseMessageJSON `json:"-"`
}

// urlScannerScanNewResponseMessageJSON contains the JSON metadata for the struct
// [URLScannerScanNewResponseMessage]
type urlScannerScanNewResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewResponseResult struct {
	// Time when url was submitted for scanning.
	Time time.Time `json:"time,required" format:"date-time"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url,required"`
	// Scan ID.
	Uuid string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                              `json:"visibility,required"`
	JSON       urlScannerScanNewResponseResultJSON `json:"-"`
}

// urlScannerScanNewResponseResultJSON contains the JSON metadata for the struct
// [URLScannerScanNewResponseResult]
type urlScannerScanNewResponseResultJSON struct {
	Time        apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponse struct {
	Errors   []URLScannerScanGetResponseError   `json:"errors,required"`
	Messages []URLScannerScanGetResponseMessage `json:"messages,required"`
	Result   URLScannerScanGetResponseResult    `json:"result,required"`
	// Whether request was successful or not
	Success bool                          `json:"success,required"`
	JSON    urlScannerScanGetResponseJSON `json:"-"`
}

// urlScannerScanGetResponseJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponse]
type urlScannerScanGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseError struct {
	Message string                             `json:"message,required"`
	JSON    urlScannerScanGetResponseErrorJSON `json:"-"`
}

// urlScannerScanGetResponseErrorJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseError]
type urlScannerScanGetResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseMessage struct {
	Message string                               `json:"message,required"`
	JSON    urlScannerScanGetResponseMessageJSON `json:"-"`
}

// urlScannerScanGetResponseMessageJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseMessage]
type urlScannerScanGetResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResult struct {
	Scan URLScannerScanGetResponseResultScan `json:"scan,required"`
	JSON urlScannerScanGetResponseResultJSON `json:"-"`
}

// urlScannerScanGetResponseResultJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseResult]
type urlScannerScanGetResponseResultJSON struct {
	Scan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScan struct {
	Certificates []URLScannerScanGetResponseResultScanCertificate `json:"certificates,required"`
	Geo          URLScannerScanGetResponseResultScanGeo           `json:"geo,required"`
	Meta         URLScannerScanGetResponseResultScanMeta          `json:"meta,required"`
	Page         URLScannerScanGetResponseResultScanPage          `json:"page,required"`
	Performance  []URLScannerScanGetResponseResultScanPerformance `json:"performance,required"`
	Task         URLScannerScanGetResponseResultScanTask          `json:"task,required"`
	Verdicts     URLScannerScanGetResponseResultScanVerdicts      `json:"verdicts,required"`
	// Dictionary of Autonomous System Numbers where ASN's are the keys
	Asns    URLScannerScanGetResponseResultScanAsns    `json:"asns"`
	Domains URLScannerScanGetResponseResultScanDomains `json:"domains"`
	IPs     URLScannerScanGetResponseResultScanIPs     `json:"ips"`
	Links   URLScannerScanGetResponseResultScanLinks   `json:"links"`
	JSON    urlScannerScanGetResponseResultScanJSON    `json:"-"`
}

// urlScannerScanGetResponseResultScanJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScan]
type urlScannerScanGetResponseResultScanJSON struct {
	Certificates apijson.Field
	Geo          apijson.Field
	Meta         apijson.Field
	Page         apijson.Field
	Performance  apijson.Field
	Task         apijson.Field
	Verdicts     apijson.Field
	Asns         apijson.Field
	Domains      apijson.Field
	IPs          apijson.Field
	Links        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanCertificate struct {
	Issuer      string                                             `json:"issuer,required"`
	SubjectName string                                             `json:"subjectName,required"`
	ValidFrom   float64                                            `json:"validFrom,required"`
	ValidTo     float64                                            `json:"validTo,required"`
	JSON        urlScannerScanGetResponseResultScanCertificateJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanCertificateJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseResultScanCertificate]
type urlScannerScanGetResponseResultScanCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanGeo struct {
	// GeoIP continent location
	Continents []string `json:"continents,required"`
	// GeoIP country location
	Locations []string                                   `json:"locations,required"`
	JSON      urlScannerScanGetResponseResultScanGeoJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanGeoJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanGeo]
type urlScannerScanGetResponseResultScanGeoJSON struct {
	Continents  apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMeta struct {
	Processors URLScannerScanGetResponseResultScanMetaProcessors `json:"processors,required"`
	JSON       urlScannerScanGetResponseResultScanMetaJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanMeta]
type urlScannerScanGetResponseResultScanMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessors struct {
	Categories         URLScannerScanGetResponseResultScanMetaProcessorsCategories `json:"categories,required"`
	GoogleSafeBrowsing []string                                                    `json:"google_safe_browsing,required"`
	Phishing           []string                                                    `json:"phishing,required"`
	Rank               URLScannerScanGetResponseResultScanMetaProcessorsRank       `json:"rank,required"`
	Tech               []URLScannerScanGetResponseResultScanMetaProcessorsTech     `json:"tech,required"`
	JSON               urlScannerScanGetResponseResultScanMetaProcessorsJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseResultScanMetaProcessors]
type urlScannerScanGetResponseResultScanMetaProcessorsJSON struct {
	Categories         apijson.Field
	GoogleSafeBrowsing apijson.Field
	Phishing           apijson.Field
	Rank               apijson.Field
	Tech               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsCategories struct {
	Content []URLScannerScanGetResponseResultScanMetaProcessorsCategoriesContent `json:"content,required"`
	Risks   []URLScannerScanGetResponseResultScanMetaProcessorsCategoriesRisk    `json:"risks,required"`
	JSON    urlScannerScanGetResponseResultScanMetaProcessorsCategoriesJSON      `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsCategoriesJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseResultScanMetaProcessorsCategories]
type urlScannerScanGetResponseResultScanMetaProcessorsCategoriesJSON struct {
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessorsCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsCategoriesContent struct {
	ID              int64                                                                  `json:"id,required"`
	Name            string                                                                 `json:"name,required"`
	SuperCategoryID int64                                                                  `json:"super_category_id"`
	JSON            urlScannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON contains
// the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanMetaProcessorsCategoriesContent]
type urlScannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessorsCategoriesContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsCategoriesRisk struct {
	ID              int64                                                               `json:"id,required"`
	Name            string                                                              `json:"name,required"`
	SuperCategoryID int64                                                               `json:"super_category_id,required"`
	JSON            urlScannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseResultScanMetaProcessorsCategoriesRisk]
type urlScannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessorsCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsRank struct {
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                                     `json:"rank"`
	JSON urlScannerScanGetResponseResultScanMetaProcessorsRankJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsRankJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseResultScanMetaProcessorsRank]
type urlScannerScanGetResponseResultScanMetaProcessorsRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessorsRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsTech struct {
	Categories  []URLScannerScanGetResponseResultScanMetaProcessorsTechCategory `json:"categories,required"`
	Confidence  int64                                                           `json:"confidence,required"`
	Evidence    URLScannerScanGetResponseResultScanMetaProcessorsTechEvidence   `json:"evidence,required"`
	Icon        string                                                          `json:"icon,required"`
	Name        string                                                          `json:"name,required"`
	Slug        string                                                          `json:"slug,required"`
	Website     string                                                          `json:"website,required"`
	Description string                                                          `json:"description"`
	JSON        urlScannerScanGetResponseResultScanMetaProcessorsTechJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsTechJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseResultScanMetaProcessorsTech]
type urlScannerScanGetResponseResultScanMetaProcessorsTechJSON struct {
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

func (r *URLScannerScanGetResponseResultScanMetaProcessorsTech) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsTechCategory struct {
	ID       int64                                                             `json:"id,required"`
	Groups   []int64                                                           `json:"groups,required"`
	Name     string                                                            `json:"name,required"`
	Priority int64                                                             `json:"priority,required"`
	Slug     string                                                            `json:"slug,required"`
	JSON     urlScannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseResultScanMetaProcessorsTechCategory]
type urlScannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON struct {
	ID          apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessorsTechCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsTechEvidence struct {
	ImpliedBy []string                                                               `json:"impliedBy,required"`
	Patterns  []URLScannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern `json:"patterns,required"`
	JSON      urlScannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON      `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseResultScanMetaProcessorsTechEvidence]
type urlScannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON struct {
	ImpliedBy   apijson.Field
	Patterns    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanMetaProcessorsTechEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern struct {
	Confidence int64    `json:"confidence,required"`
	Excludes   []string `json:"excludes,required"`
	Implies    []string `json:"implies,required"`
	Match      string   `json:"match,required"`
	// Header or Cookie name when set
	Name    string                                                                   `json:"name,required"`
	Regex   string                                                                   `json:"regex,required"`
	Type    string                                                                   `json:"type,required"`
	Value   string                                                                   `json:"value,required"`
	Version string                                                                   `json:"version,required"`
	JSON    urlScannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern]
type urlScannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON struct {
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

func (r *URLScannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPage struct {
	Asn                   string                                                     `json:"asn,required"`
	AsnLocationAlpha2     string                                                     `json:"asnLocationAlpha2,required"`
	Asnname               string                                                     `json:"asnname,required"`
	Console               []URLScannerScanGetResponseResultScanPageConsole           `json:"console,required"`
	Cookies               []URLScannerScanGetResponseResultScanPageCooky             `json:"cookies,required"`
	Country               string                                                     `json:"country,required"`
	CountryLocationAlpha2 string                                                     `json:"countryLocationAlpha2,required"`
	Domain                string                                                     `json:"domain,required"`
	Headers               []URLScannerScanGetResponseResultScanPageHeader            `json:"headers,required"`
	IP                    string                                                     `json:"ip,required"`
	Js                    URLScannerScanGetResponseResultScanPageJs                  `json:"js,required"`
	SecurityViolations    []URLScannerScanGetResponseResultScanPageSecurityViolation `json:"securityViolations,required"`
	Status                float64                                                    `json:"status,required"`
	Subdivision1Name      string                                                     `json:"subdivision1Name,required"`
	Subdivision2name      string                                                     `json:"subdivision2name,required"`
	URL                   string                                                     `json:"url,required"`
	JSON                  urlScannerScanGetResponseResultScanPageJSON                `json:"-"`
}

// urlScannerScanGetResponseResultScanPageJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanPage]
type urlScannerScanGetResponseResultScanPageJSON struct {
	Asn                   apijson.Field
	AsnLocationAlpha2     apijson.Field
	Asnname               apijson.Field
	Console               apijson.Field
	Cookies               apijson.Field
	Country               apijson.Field
	CountryLocationAlpha2 apijson.Field
	Domain                apijson.Field
	Headers               apijson.Field
	IP                    apijson.Field
	Js                    apijson.Field
	SecurityViolations    apijson.Field
	Status                apijson.Field
	Subdivision1Name      apijson.Field
	Subdivision2name      apijson.Field
	URL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPageConsole struct {
	Category string                                             `json:"category,required"`
	Text     string                                             `json:"text,required"`
	Type     string                                             `json:"type,required"`
	URL      string                                             `json:"url"`
	JSON     urlScannerScanGetResponseResultScanPageConsoleJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanPageConsoleJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseResultScanPageConsole]
type urlScannerScanGetResponseResultScanPageConsoleJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanPageConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPageCooky struct {
	Domain       string                                           `json:"domain,required"`
	Expires      float64                                          `json:"expires,required"`
	HTTPOnly     bool                                             `json:"httpOnly,required"`
	Name         string                                           `json:"name,required"`
	Path         string                                           `json:"path,required"`
	SameParty    bool                                             `json:"sameParty,required"`
	Secure       bool                                             `json:"secure,required"`
	Session      bool                                             `json:"session,required"`
	Size         float64                                          `json:"size,required"`
	SourcePort   float64                                          `json:"sourcePort,required"`
	SourceScheme string                                           `json:"sourceScheme,required"`
	Value        string                                           `json:"value,required"`
	Priority     string                                           `json:"priority"`
	JSON         urlScannerScanGetResponseResultScanPageCookyJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanPageCookyJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanPageCooky]
type urlScannerScanGetResponseResultScanPageCookyJSON struct {
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

func (r *URLScannerScanGetResponseResultScanPageCooky) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPageHeader struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  urlScannerScanGetResponseResultScanPageHeaderJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanPageHeaderJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanPageHeader]
type urlScannerScanGetResponseResultScanPageHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanPageHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPageJs struct {
	Variables []URLScannerScanGetResponseResultScanPageJsVariable `json:"variables,required"`
	JSON      urlScannerScanGetResponseResultScanPageJsJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanPageJsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanPageJs]
type urlScannerScanGetResponseResultScanPageJsJSON struct {
	Variables   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanPageJs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPageJsVariable struct {
	Name string                                                `json:"name,required"`
	Type string                                                `json:"type,required"`
	JSON urlScannerScanGetResponseResultScanPageJsVariableJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanPageJsVariableJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseResultScanPageJsVariable]
type urlScannerScanGetResponseResultScanPageJsVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanPageJsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPageSecurityViolation struct {
	Category string                                                       `json:"category,required"`
	Text     string                                                       `json:"text,required"`
	URL      string                                                       `json:"url,required"`
	JSON     urlScannerScanGetResponseResultScanPageSecurityViolationJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanPageSecurityViolationJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseResultScanPageSecurityViolation]
type urlScannerScanGetResponseResultScanPageSecurityViolationJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanPageSecurityViolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanPerformance struct {
	ConnectEnd                 float64                                            `json:"connectEnd,required"`
	ConnectStart               float64                                            `json:"connectStart,required"`
	DecodedBodySize            float64                                            `json:"decodedBodySize,required"`
	DomainLookupEnd            float64                                            `json:"domainLookupEnd,required"`
	DomainLookupStart          float64                                            `json:"domainLookupStart,required"`
	DomComplete                float64                                            `json:"domComplete,required"`
	DomContentLoadedEventEnd   float64                                            `json:"domContentLoadedEventEnd,required"`
	DomContentLoadedEventStart float64                                            `json:"domContentLoadedEventStart,required"`
	DomInteractive             float64                                            `json:"domInteractive,required"`
	Duration                   float64                                            `json:"duration,required"`
	EncodedBodySize            float64                                            `json:"encodedBodySize,required"`
	EntryType                  string                                             `json:"entryType,required"`
	FetchStart                 float64                                            `json:"fetchStart,required"`
	InitiatorType              string                                             `json:"initiatorType,required"`
	LoadEventEnd               float64                                            `json:"loadEventEnd,required"`
	LoadEventStart             float64                                            `json:"loadEventStart,required"`
	Name                       string                                             `json:"name,required"`
	NextHopProtocol            string                                             `json:"nextHopProtocol,required"`
	RedirectCount              float64                                            `json:"redirectCount,required"`
	RedirectEnd                float64                                            `json:"redirectEnd,required"`
	RedirectStart              float64                                            `json:"redirectStart,required"`
	RequestStart               float64                                            `json:"requestStart,required"`
	ResponseEnd                float64                                            `json:"responseEnd,required"`
	ResponseStart              float64                                            `json:"responseStart,required"`
	SecureConnectionStart      float64                                            `json:"secureConnectionStart,required"`
	StartTime                  float64                                            `json:"startTime,required"`
	TransferSize               float64                                            `json:"transferSize,required"`
	Type                       string                                             `json:"type,required"`
	UnloadEventEnd             float64                                            `json:"unloadEventEnd,required"`
	UnloadEventStart           float64                                            `json:"unloadEventStart,required"`
	WorkerStart                float64                                            `json:"workerStart,required"`
	JSON                       urlScannerScanGetResponseResultScanPerformanceJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanPerformanceJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseResultScanPerformance]
type urlScannerScanGetResponseResultScanPerformanceJSON struct {
	ConnectEnd                 apijson.Field
	ConnectStart               apijson.Field
	DecodedBodySize            apijson.Field
	DomainLookupEnd            apijson.Field
	DomainLookupStart          apijson.Field
	DomComplete                apijson.Field
	DomContentLoadedEventEnd   apijson.Field
	DomContentLoadedEventStart apijson.Field
	DomInteractive             apijson.Field
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

func (r *URLScannerScanGetResponseResultScanPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanTask struct {
	// Submitter location
	ClientLocation string                                            `json:"clientLocation,required"`
	ClientType     URLScannerScanGetResponseResultScanTaskClientType `json:"clientType,required"`
	// URL of the primary request, after all HTTP redirects
	EffectiveURL string                                             `json:"effectiveUrl,required"`
	Errors       []URLScannerScanGetResponseResultScanTaskError     `json:"errors,required"`
	ScannedFrom  URLScannerScanGetResponseResultScanTaskScannedFrom `json:"scannedFrom,required"`
	Status       URLScannerScanGetResponseResultScanTaskStatus      `json:"status,required"`
	Success      bool                                               `json:"success,required"`
	Time         string                                             `json:"time,required"`
	TimeEnd      string                                             `json:"timeEnd,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID
	Uuid       string                                            `json:"uuid,required"`
	Visibility URLScannerScanGetResponseResultScanTaskVisibility `json:"visibility,required"`
	JSON       urlScannerScanGetResponseResultScanTaskJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanTaskJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanTask]
type urlScannerScanGetResponseResultScanTaskJSON struct {
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
	Uuid           apijson.Field
	Visibility     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanTaskClientType string

const (
	URLScannerScanGetResponseResultScanTaskClientTypeSite      URLScannerScanGetResponseResultScanTaskClientType = "Site"
	URLScannerScanGetResponseResultScanTaskClientTypeAutomatic URLScannerScanGetResponseResultScanTaskClientType = "Automatic"
	URLScannerScanGetResponseResultScanTaskClientTypeAPI       URLScannerScanGetResponseResultScanTaskClientType = "Api"
)

type URLScannerScanGetResponseResultScanTaskError struct {
	Message string                                           `json:"message,required"`
	JSON    urlScannerScanGetResponseResultScanTaskErrorJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanTaskErrorJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanTaskError]
type urlScannerScanGetResponseResultScanTaskErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanTaskError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanTaskScannedFrom struct {
	// IATA code of Cloudflare datacenter
	Colo string                                                 `json:"colo,required"`
	JSON urlScannerScanGetResponseResultScanTaskScannedFromJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanTaskScannedFromJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseResultScanTaskScannedFrom]
type urlScannerScanGetResponseResultScanTaskScannedFromJSON struct {
	Colo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanTaskScannedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanTaskStatus string

const (
	URLScannerScanGetResponseResultScanTaskStatusQueued           URLScannerScanGetResponseResultScanTaskStatus = "Queued"
	URLScannerScanGetResponseResultScanTaskStatusInProgress       URLScannerScanGetResponseResultScanTaskStatus = "InProgress"
	URLScannerScanGetResponseResultScanTaskStatusInPostProcessing URLScannerScanGetResponseResultScanTaskStatus = "InPostProcessing"
	URLScannerScanGetResponseResultScanTaskStatusFinished         URLScannerScanGetResponseResultScanTaskStatus = "Finished"
)

type URLScannerScanGetResponseResultScanTaskVisibility string

const (
	URLScannerScanGetResponseResultScanTaskVisibilityPublic   URLScannerScanGetResponseResultScanTaskVisibility = "Public"
	URLScannerScanGetResponseResultScanTaskVisibilityUnlisted URLScannerScanGetResponseResultScanTaskVisibility = "Unlisted"
)

type URLScannerScanGetResponseResultScanVerdicts struct {
	Overall URLScannerScanGetResponseResultScanVerdictsOverall `json:"overall,required"`
	JSON    urlScannerScanGetResponseResultScanVerdictsJSON    `json:"-"`
}

// urlScannerScanGetResponseResultScanVerdictsJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanVerdicts]
type urlScannerScanGetResponseResultScanVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanVerdictsOverall struct {
	Categories []URLScannerScanGetResponseResultScanVerdictsOverallCategory `json:"categories,required"`
	// Please visit https://safebrowsing.google.com/ for more information.
	GsbThreatTypes []string `json:"gsb_threat_types,required"`
	// At least one of our subsystems marked the site as potentially malicious at the
	// time of the scan.
	Malicious bool                                                   `json:"malicious,required"`
	Phishing  []string                                               `json:"phishing,required"`
	JSON      urlScannerScanGetResponseResultScanVerdictsOverallJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanVerdictsOverallJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseResultScanVerdictsOverall]
type urlScannerScanGetResponseResultScanVerdictsOverallJSON struct {
	Categories     apijson.Field
	GsbThreatTypes apijson.Field
	Malicious      apijson.Field
	Phishing       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanVerdictsOverallCategory struct {
	ID              float64                                                        `json:"id,required"`
	Name            string                                                         `json:"name,required"`
	SuperCategoryID float64                                                        `json:"super_category_id,required"`
	JSON            urlScannerScanGetResponseResultScanVerdictsOverallCategoryJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanVerdictsOverallCategoryJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseResultScanVerdictsOverallCategory]
type urlScannerScanGetResponseResultScanVerdictsOverallCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanVerdictsOverallCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Dictionary of Autonomous System Numbers where ASN's are the keys
type URLScannerScanGetResponseResultScanAsns struct {
	// ASN's contacted
	Asn  URLScannerScanGetResponseResultScanAsnsAsn  `json:"asn"`
	JSON urlScannerScanGetResponseResultScanAsnsJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanAsnsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanAsns]
type urlScannerScanGetResponseResultScanAsnsJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanAsns) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ASN's contacted
type URLScannerScanGetResponseResultScanAsnsAsn struct {
	Asn            string                                         `json:"asn,required"`
	Description    string                                         `json:"description,required"`
	LocationAlpha2 string                                         `json:"location_alpha2,required"`
	Name           string                                         `json:"name,required"`
	OrgName        string                                         `json:"org_name,required"`
	JSON           urlScannerScanGetResponseResultScanAsnsAsnJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanAsnsAsnJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanAsnsAsn]
type urlScannerScanGetResponseResultScanAsnsAsnJSON struct {
	Asn            apijson.Field
	Description    apijson.Field
	LocationAlpha2 apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanAsnsAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomains struct {
	ExampleCom URLScannerScanGetResponseResultScanDomainsExampleCom `json:"example.com"`
	JSON       urlScannerScanGetResponseResultScanDomainsJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanDomains]
type urlScannerScanGetResponseResultScanDomainsJSON struct {
	ExampleCom  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleCom struct {
	Categories URLScannerScanGetResponseResultScanDomainsExampleComCategories `json:"categories,required"`
	DNS        []URLScannerScanGetResponseResultScanDomainsExampleComDNS      `json:"dns,required"`
	Name       string                                                         `json:"name,required"`
	Rank       URLScannerScanGetResponseResultScanDomainsExampleComRank       `json:"rank,required"`
	Type       string                                                         `json:"type,required"`
	JSON       urlScannerScanGetResponseResultScanDomainsExampleComJSON       `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseResultScanDomainsExampleCom]
type urlScannerScanGetResponseResultScanDomainsExampleComJSON struct {
	Categories  apijson.Field
	DNS         apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleCom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComCategories struct {
	Inherited URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInherited `json:"inherited,required"`
	Content   []URLScannerScanGetResponseResultScanDomainsExampleComCategoriesContent `json:"content"`
	Risks     []URLScannerScanGetResponseResultScanDomainsExampleComCategoriesRisk    `json:"risks"`
	JSON      urlScannerScanGetResponseResultScanDomainsExampleComCategoriesJSON      `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComCategoriesJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComCategories]
type urlScannerScanGetResponseResultScanDomainsExampleComCategoriesJSON struct {
	Inherited   apijson.Field
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInherited struct {
	Content []URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent `json:"content"`
	From    string                                                                           `json:"from"`
	Risks   []URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk    `json:"risks"`
	JSON    urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON      `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInherited]
type urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent struct {
	ID              int64                                                                              `json:"id,required"`
	Name            string                                                                             `json:"name,required"`
	SuperCategoryID int64                                                                              `json:"super_category_id"`
	JSON            urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent]
type urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk struct {
	ID              int64                                                                           `json:"id,required"`
	Name            string                                                                          `json:"name,required"`
	SuperCategoryID int64                                                                           `json:"super_category_id"`
	JSON            urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk]
type urlScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComCategoriesContent struct {
	ID              int64                                                                     `json:"id,required"`
	Name            string                                                                    `json:"name,required"`
	SuperCategoryID int64                                                                     `json:"super_category_id"`
	JSON            urlScannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComCategoriesContent]
type urlScannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComCategoriesContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComCategoriesRisk struct {
	ID              int64                                                                  `json:"id,required"`
	Name            string                                                                 `json:"name,required"`
	SuperCategoryID int64                                                                  `json:"super_category_id"`
	JSON            urlScannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON contains
// the JSON metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComCategoriesRisk]
type urlScannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComDNS struct {
	Address     string                                                      `json:"address,required"`
	DnssecValid bool                                                        `json:"dnssec_valid,required"`
	Name        string                                                      `json:"name,required"`
	Type        string                                                      `json:"type,required"`
	JSON        urlScannerScanGetResponseResultScanDomainsExampleComDNSJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComDNSJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComDNS]
type urlScannerScanGetResponseResultScanDomainsExampleComDNSJSON struct {
	Address     apijson.Field
	DnssecValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanDomainsExampleComRank struct {
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                                        `json:"rank"`
	JSON urlScannerScanGetResponseResultScanDomainsExampleComRankJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanDomainsExampleComRankJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseResultScanDomainsExampleComRank]
type urlScannerScanGetResponseResultScanDomainsExampleComRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanDomainsExampleComRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanIPs struct {
	IP   URLScannerScanGetResponseResultScanIPsIP   `json:"ip"`
	JSON urlScannerScanGetResponseResultScanIPsJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanIPsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanIPs]
type urlScannerScanGetResponseResultScanIPsJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanIPsIP struct {
	Asn               string                                       `json:"asn,required"`
	AsnDescription    string                                       `json:"asnDescription,required"`
	AsnLocationAlpha2 string                                       `json:"asnLocationAlpha2,required"`
	AsnName           string                                       `json:"asnName,required"`
	AsnOrgName        string                                       `json:"asnOrgName,required"`
	Continent         string                                       `json:"continent,required"`
	GeonameID         string                                       `json:"geonameId,required"`
	IP                string                                       `json:"ip,required"`
	IPVersion         string                                       `json:"ipVersion,required"`
	Latitude          string                                       `json:"latitude,required"`
	LocationAlpha2    string                                       `json:"locationAlpha2,required"`
	LocationName      string                                       `json:"locationName,required"`
	Longitude         string                                       `json:"longitude,required"`
	Subdivision1Name  string                                       `json:"subdivision1Name,required"`
	Subdivision2Name  string                                       `json:"subdivision2Name,required"`
	JSON              urlScannerScanGetResponseResultScanIPsIPJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanIPsIPJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanIPsIP]
type urlScannerScanGetResponseResultScanIPsIPJSON struct {
	Asn               apijson.Field
	AsnDescription    apijson.Field
	AsnLocationAlpha2 apijson.Field
	AsnName           apijson.Field
	AsnOrgName        apijson.Field
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

func (r *URLScannerScanGetResponseResultScanIPsIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanLinks struct {
	Link URLScannerScanGetResponseResultScanLinksLink `json:"link"`
	JSON urlScannerScanGetResponseResultScanLinksJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanLinksJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseResultScanLinks]
type urlScannerScanGetResponseResultScanLinksJSON struct {
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseResultScanLinksLink struct {
	// Outgoing link detected in the DOM
	Href string                                           `json:"href,required"`
	Text string                                           `json:"text,required"`
	JSON urlScannerScanGetResponseResultScanLinksLinkJSON `json:"-"`
}

// urlScannerScanGetResponseResultScanLinksLinkJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseResultScanLinksLink]
type urlScannerScanGetResponseResultScanLinksLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseResultScanLinksLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponse struct {
	Errors   []URLScannerScanHarResponseError   `json:"errors,required"`
	Messages []URLScannerScanHarResponseMessage `json:"messages,required"`
	Result   URLScannerScanHarResponseResult    `json:"result,required"`
	// Whether search request was successful or not
	Success bool                          `json:"success,required"`
	JSON    urlScannerScanHarResponseJSON `json:"-"`
}

// urlScannerScanHarResponseJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponse]
type urlScannerScanHarResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseError struct {
	Message string                             `json:"message,required"`
	JSON    urlScannerScanHarResponseErrorJSON `json:"-"`
}

// urlScannerScanHarResponseErrorJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseError]
type urlScannerScanHarResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseMessage struct {
	Message string                               `json:"message,required"`
	JSON    urlScannerScanHarResponseMessageJSON `json:"-"`
}

// urlScannerScanHarResponseMessageJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseMessage]
type urlScannerScanHarResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResult struct {
	Har  URLScannerScanHarResponseResultHar  `json:"har,required"`
	JSON urlScannerScanHarResponseResultJSON `json:"-"`
}

// urlScannerScanHarResponseResultJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseResult]
type urlScannerScanHarResponseResultJSON struct {
	Har         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHar struct {
	Log  URLScannerScanHarResponseResultHarLog  `json:"log,required"`
	JSON urlScannerScanHarResponseResultHarJSON `json:"-"`
}

// urlScannerScanHarResponseResultHarJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseResultHar]
type urlScannerScanHarResponseResultHarJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLog struct {
	Creator URLScannerScanHarResponseResultHarLogCreator `json:"creator,required"`
	Entries []URLScannerScanHarResponseResultHarLogEntry `json:"entries,required"`
	Pages   []URLScannerScanHarResponseResultHarLogPage  `json:"pages,required"`
	Version string                                       `json:"version,required"`
	JSON    urlScannerScanHarResponseResultHarLogJSON    `json:"-"`
}

// urlScannerScanHarResponseResultHarLogJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseResultHarLog]
type urlScannerScanHarResponseResultHarLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogCreator struct {
	Comment string                                           `json:"comment,required"`
	Name    string                                           `json:"name,required"`
	Version string                                           `json:"version,required"`
	JSON    urlScannerScanHarResponseResultHarLogCreatorJSON `json:"-"`
}

// urlScannerScanHarResponseResultHarLogCreatorJSON contains the JSON metadata for
// the struct [URLScannerScanHarResponseResultHarLogCreator]
type urlScannerScanHarResponseResultHarLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogEntry struct {
	InitialPriority string                                               `json:"_initialPriority,required"`
	InitiatorType   string                                               `json:"_initiator_type,required"`
	Priority        string                                               `json:"_priority,required"`
	RequestID       string                                               `json:"_requestId,required"`
	RequestTime     float64                                              `json:"_requestTime,required"`
	ResourceType    string                                               `json:"_resourceType,required"`
	Cache           interface{}                                          `json:"cache,required"`
	Connection      string                                               `json:"connection,required"`
	Pageref         string                                               `json:"pageref,required"`
	Request         URLScannerScanHarResponseResultHarLogEntriesRequest  `json:"request,required"`
	Response        URLScannerScanHarResponseResultHarLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                                               `json:"serverIPAddress,required"`
	StartedDateTime string                                               `json:"startedDateTime,required"`
	Time            float64                                              `json:"time,required"`
	JSON            urlScannerScanHarResponseResultHarLogEntryJSON       `json:"-"`
}

// urlScannerScanHarResponseResultHarLogEntryJSON contains the JSON metadata for
// the struct [URLScannerScanHarResponseResultHarLogEntry]
type urlScannerScanHarResponseResultHarLogEntryJSON struct {
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

func (r *URLScannerScanHarResponseResultHarLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogEntriesRequest struct {
	BodySize    float64                                                     `json:"bodySize,required"`
	Headers     []URLScannerScanHarResponseResultHarLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                                     `json:"headersSize,required"`
	HTTPVersion string                                                      `json:"httpVersion,required"`
	Method      string                                                      `json:"method,required"`
	URL         string                                                      `json:"url,required"`
	JSON        urlScannerScanHarResponseResultHarLogEntriesRequestJSON     `json:"-"`
}

// urlScannerScanHarResponseResultHarLogEntriesRequestJSON contains the JSON
// metadata for the struct [URLScannerScanHarResponseResultHarLogEntriesRequest]
type urlScannerScanHarResponseResultHarLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogEntriesRequestHeader struct {
	Name  string                                                        `json:"name,required"`
	Value string                                                        `json:"value,required"`
	JSON  urlScannerScanHarResponseResultHarLogEntriesRequestHeaderJSON `json:"-"`
}

// urlScannerScanHarResponseResultHarLogEntriesRequestHeaderJSON contains the JSON
// metadata for the struct
// [URLScannerScanHarResponseResultHarLogEntriesRequestHeader]
type urlScannerScanHarResponseResultHarLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogEntriesResponse struct {
	TransferSize float64                                                      `json:"_transferSize,required"`
	BodySize     float64                                                      `json:"bodySize,required"`
	Content      URLScannerScanHarResponseResultHarLogEntriesResponseContent  `json:"content,required"`
	Headers      []URLScannerScanHarResponseResultHarLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                                      `json:"headersSize,required"`
	HTTPVersion  string                                                       `json:"httpVersion,required"`
	RedirectURL  string                                                       `json:"redirectURL,required"`
	Status       float64                                                      `json:"status,required"`
	StatusText   string                                                       `json:"statusText,required"`
	JSON         urlScannerScanHarResponseResultHarLogEntriesResponseJSON     `json:"-"`
}

// urlScannerScanHarResponseResultHarLogEntriesResponseJSON contains the JSON
// metadata for the struct [URLScannerScanHarResponseResultHarLogEntriesResponse]
type urlScannerScanHarResponseResultHarLogEntriesResponseJSON struct {
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

func (r *URLScannerScanHarResponseResultHarLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogEntriesResponseContent struct {
	MimeType    string                                                          `json:"mimeType,required"`
	Size        float64                                                         `json:"size,required"`
	Compression int64                                                           `json:"compression"`
	JSON        urlScannerScanHarResponseResultHarLogEntriesResponseContentJSON `json:"-"`
}

// urlScannerScanHarResponseResultHarLogEntriesResponseContentJSON contains the
// JSON metadata for the struct
// [URLScannerScanHarResponseResultHarLogEntriesResponseContent]
type urlScannerScanHarResponseResultHarLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogEntriesResponseHeader struct {
	Name  string                                                         `json:"name,required"`
	Value string                                                         `json:"value,required"`
	JSON  urlScannerScanHarResponseResultHarLogEntriesResponseHeaderJSON `json:"-"`
}

// urlScannerScanHarResponseResultHarLogEntriesResponseHeaderJSON contains the JSON
// metadata for the struct
// [URLScannerScanHarResponseResultHarLogEntriesResponseHeader]
type urlScannerScanHarResponseResultHarLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogPage struct {
	ID              string                                                `json:"id,required"`
	PageTimings     URLScannerScanHarResponseResultHarLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                                                `json:"startedDateTime,required"`
	Title           string                                                `json:"title,required"`
	JSON            urlScannerScanHarResponseResultHarLogPageJSON         `json:"-"`
}

// urlScannerScanHarResponseResultHarLogPageJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseResultHarLogPage]
type urlScannerScanHarResponseResultHarLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseResultHarLogPagesPageTimings struct {
	OnContentLoad float64                                                   `json:"onContentLoad,required"`
	OnLoad        float64                                                   `json:"onLoad,required"`
	JSON          urlScannerScanHarResponseResultHarLogPagesPageTimingsJSON `json:"-"`
}

// urlScannerScanHarResponseResultHarLogPagesPageTimingsJSON contains the JSON
// metadata for the struct [URLScannerScanHarResponseResultHarLogPagesPageTimings]
type urlScannerScanHarResponseResultHarLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *URLScannerScanHarResponseResultHarLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewParams struct {
	URL param.Field[string] `json:"url,required"`
	// Set custom headers
	CustomHeaders param.Field[interface{}] `json:"customHeaders"`
	// Take multiple screenshots targeting different device types
	ScreenshotsResolutions param.Field[[]URLScannerScanNewParamsScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[URLScannerScanNewParamsVisibility] `json:"visibility"`
}

func (r URLScannerScanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Device resolutions.
type URLScannerScanNewParamsScreenshotsResolution string

const (
	URLScannerScanNewParamsScreenshotsResolutionDesktop URLScannerScanNewParamsScreenshotsResolution = "desktop"
	URLScannerScanNewParamsScreenshotsResolutionMobile  URLScannerScanNewParamsScreenshotsResolution = "mobile"
	URLScannerScanNewParamsScreenshotsResolutionTablet  URLScannerScanNewParamsScreenshotsResolution = "tablet"
)

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type URLScannerScanNewParamsVisibility string

const (
	URLScannerScanNewParamsVisibilityPublic   URLScannerScanNewParamsVisibility = "Public"
	URLScannerScanNewParamsVisibilityUnlisted URLScannerScanNewParamsVisibility = "Unlisted"
)

type URLScannerScanScreenshotParams struct {
	// Target device type
	Resolution param.Field[URLScannerScanScreenshotParamsResolution] `query:"resolution"`
}

// URLQuery serializes [URLScannerScanScreenshotParams]'s query parameters as
// `url.Values`.
func (r URLScannerScanScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Target device type
type URLScannerScanScreenshotParamsResolution string

const (
	URLScannerScanScreenshotParamsResolutionDesktop URLScannerScanScreenshotParamsResolution = "desktop"
	URLScannerScanScreenshotParamsResolutionMobile  URLScannerScanScreenshotParamsResolution = "mobile"
	URLScannerScanScreenshotParamsResolutionTablet  URLScannerScanScreenshotParamsResolution = "tablet"
)
