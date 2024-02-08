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
	var env URLScannerScanNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get URL scan by uuid
func (r *URLScannerScanService) Get(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *URLScannerScanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env URLScannerScanGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *URLScannerScanService) Har(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *URLScannerScanHarResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env URLScannerScanHarResponseEnvelope
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/har", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
	// Time when url was submitted for scanning.
	Time time.Time `json:"time,required" format:"date-time"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url,required"`
	// Scan ID.
	Uuid string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                        `json:"visibility,required"`
	JSON       urlScannerScanNewResponseJSON `json:"-"`
}

// urlScannerScanNewResponseJSON contains the JSON metadata for the struct
// [URLScannerScanNewResponse]
type urlScannerScanNewResponseJSON struct {
	Time        apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponse struct {
	Scan URLScannerScanGetResponseScan `json:"scan,required"`
	JSON urlScannerScanGetResponseJSON `json:"-"`
}

// urlScannerScanGetResponseJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponse]
type urlScannerScanGetResponseJSON struct {
	Scan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScan struct {
	Certificates []URLScannerScanGetResponseScanCertificate `json:"certificates,required"`
	Geo          URLScannerScanGetResponseScanGeo           `json:"geo,required"`
	Meta         URLScannerScanGetResponseScanMeta          `json:"meta,required"`
	Page         URLScannerScanGetResponseScanPage          `json:"page,required"`
	Performance  []URLScannerScanGetResponseScanPerformance `json:"performance,required"`
	Task         URLScannerScanGetResponseScanTask          `json:"task,required"`
	Verdicts     URLScannerScanGetResponseScanVerdicts      `json:"verdicts,required"`
	// Dictionary of Autonomous System Numbers where ASN's are the keys
	Asns    URLScannerScanGetResponseScanAsns    `json:"asns"`
	Domains URLScannerScanGetResponseScanDomains `json:"domains"`
	IPs     URLScannerScanGetResponseScanIPs     `json:"ips"`
	Links   URLScannerScanGetResponseScanLinks   `json:"links"`
	JSON    urlScannerScanGetResponseScanJSON    `json:"-"`
}

// urlScannerScanGetResponseScanJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScan]
type urlScannerScanGetResponseScanJSON struct {
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

func (r *URLScannerScanGetResponseScan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanCertificate struct {
	Issuer      string                                       `json:"issuer,required"`
	SubjectName string                                       `json:"subjectName,required"`
	ValidFrom   float64                                      `json:"validFrom,required"`
	ValidTo     float64                                      `json:"validTo,required"`
	JSON        urlScannerScanGetResponseScanCertificateJSON `json:"-"`
}

// urlScannerScanGetResponseScanCertificateJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanCertificate]
type urlScannerScanGetResponseScanCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanGeo struct {
	// GeoIP continent location
	Continents []string `json:"continents,required"`
	// GeoIP country location
	Locations []string                             `json:"locations,required"`
	JSON      urlScannerScanGetResponseScanGeoJSON `json:"-"`
}

// urlScannerScanGetResponseScanGeoJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanGeo]
type urlScannerScanGetResponseScanGeoJSON struct {
	Continents  apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMeta struct {
	Processors URLScannerScanGetResponseScanMetaProcessors `json:"processors,required"`
	JSON       urlScannerScanGetResponseScanMetaJSON       `json:"-"`
}

// urlScannerScanGetResponseScanMetaJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanMeta]
type urlScannerScanGetResponseScanMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessors struct {
	Categories         URLScannerScanGetResponseScanMetaProcessorsCategories `json:"categories,required"`
	GoogleSafeBrowsing []string                                              `json:"google_safe_browsing,required"`
	Phishing           []string                                              `json:"phishing,required"`
	Rank               URLScannerScanGetResponseScanMetaProcessorsRank       `json:"rank,required"`
	Tech               []URLScannerScanGetResponseScanMetaProcessorsTech     `json:"tech,required"`
	JSON               urlScannerScanGetResponseScanMetaProcessorsJSON       `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseScanMetaProcessors]
type urlScannerScanGetResponseScanMetaProcessorsJSON struct {
	Categories         apijson.Field
	GoogleSafeBrowsing apijson.Field
	Phishing           apijson.Field
	Rank               apijson.Field
	Tech               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsCategories struct {
	Content []URLScannerScanGetResponseScanMetaProcessorsCategoriesContent `json:"content,required"`
	Risks   []URLScannerScanGetResponseScanMetaProcessorsCategoriesRisk    `json:"risks,required"`
	JSON    urlScannerScanGetResponseScanMetaProcessorsCategoriesJSON      `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsCategoriesJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseScanMetaProcessorsCategories]
type urlScannerScanGetResponseScanMetaProcessorsCategoriesJSON struct {
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessorsCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsCategoriesContent struct {
	ID              int64                                                            `json:"id,required"`
	Name            string                                                           `json:"name,required"`
	SuperCategoryID int64                                                            `json:"super_category_id"`
	JSON            urlScannerScanGetResponseScanMetaProcessorsCategoriesContentJSON `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsCategoriesContentJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseScanMetaProcessorsCategoriesContent]
type urlScannerScanGetResponseScanMetaProcessorsCategoriesContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessorsCategoriesContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsCategoriesRisk struct {
	ID              int64                                                         `json:"id,required"`
	Name            string                                                        `json:"name,required"`
	SuperCategoryID int64                                                         `json:"super_category_id,required"`
	JSON            urlScannerScanGetResponseScanMetaProcessorsCategoriesRiskJSON `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsCategoriesRiskJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseScanMetaProcessorsCategoriesRisk]
type urlScannerScanGetResponseScanMetaProcessorsCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessorsCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsRank struct {
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                               `json:"rank"`
	JSON urlScannerScanGetResponseScanMetaProcessorsRankJSON `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsRankJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseScanMetaProcessorsRank]
type urlScannerScanGetResponseScanMetaProcessorsRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessorsRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsTech struct {
	Categories  []URLScannerScanGetResponseScanMetaProcessorsTechCategory `json:"categories,required"`
	Confidence  int64                                                     `json:"confidence,required"`
	Evidence    URLScannerScanGetResponseScanMetaProcessorsTechEvidence   `json:"evidence,required"`
	Icon        string                                                    `json:"icon,required"`
	Name        string                                                    `json:"name,required"`
	Slug        string                                                    `json:"slug,required"`
	Website     string                                                    `json:"website,required"`
	Description string                                                    `json:"description"`
	JSON        urlScannerScanGetResponseScanMetaProcessorsTechJSON       `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsTechJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseScanMetaProcessorsTech]
type urlScannerScanGetResponseScanMetaProcessorsTechJSON struct {
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

func (r *URLScannerScanGetResponseScanMetaProcessorsTech) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsTechCategory struct {
	ID       int64                                                       `json:"id,required"`
	Groups   []int64                                                     `json:"groups,required"`
	Name     string                                                      `json:"name,required"`
	Priority int64                                                       `json:"priority,required"`
	Slug     string                                                      `json:"slug,required"`
	JSON     urlScannerScanGetResponseScanMetaProcessorsTechCategoryJSON `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsTechCategoryJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseScanMetaProcessorsTechCategory]
type urlScannerScanGetResponseScanMetaProcessorsTechCategoryJSON struct {
	ID          apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessorsTechCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsTechEvidence struct {
	ImpliedBy []string                                                         `json:"impliedBy,required"`
	Patterns  []URLScannerScanGetResponseScanMetaProcessorsTechEvidencePattern `json:"patterns,required"`
	JSON      urlScannerScanGetResponseScanMetaProcessorsTechEvidenceJSON      `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsTechEvidenceJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseScanMetaProcessorsTechEvidence]
type urlScannerScanGetResponseScanMetaProcessorsTechEvidenceJSON struct {
	ImpliedBy   apijson.Field
	Patterns    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanMetaProcessorsTechEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanMetaProcessorsTechEvidencePattern struct {
	Confidence int64    `json:"confidence,required"`
	Excludes   []string `json:"excludes,required"`
	Implies    []string `json:"implies,required"`
	Match      string   `json:"match,required"`
	// Header or Cookie name when set
	Name    string                                                             `json:"name,required"`
	Regex   string                                                             `json:"regex,required"`
	Type    string                                                             `json:"type,required"`
	Value   string                                                             `json:"value,required"`
	Version string                                                             `json:"version,required"`
	JSON    urlScannerScanGetResponseScanMetaProcessorsTechEvidencePatternJSON `json:"-"`
}

// urlScannerScanGetResponseScanMetaProcessorsTechEvidencePatternJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseScanMetaProcessorsTechEvidencePattern]
type urlScannerScanGetResponseScanMetaProcessorsTechEvidencePatternJSON struct {
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

func (r *URLScannerScanGetResponseScanMetaProcessorsTechEvidencePattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPage struct {
	Asn                   string                                               `json:"asn,required"`
	AsnLocationAlpha2     string                                               `json:"asnLocationAlpha2,required"`
	Asnname               string                                               `json:"asnname,required"`
	Console               []URLScannerScanGetResponseScanPageConsole           `json:"console,required"`
	Cookies               []URLScannerScanGetResponseScanPageCooky             `json:"cookies,required"`
	Country               string                                               `json:"country,required"`
	CountryLocationAlpha2 string                                               `json:"countryLocationAlpha2,required"`
	Domain                string                                               `json:"domain,required"`
	Headers               []URLScannerScanGetResponseScanPageHeader            `json:"headers,required"`
	IP                    string                                               `json:"ip,required"`
	Js                    URLScannerScanGetResponseScanPageJs                  `json:"js,required"`
	SecurityViolations    []URLScannerScanGetResponseScanPageSecurityViolation `json:"securityViolations,required"`
	Status                float64                                              `json:"status,required"`
	Subdivision1Name      string                                               `json:"subdivision1Name,required"`
	Subdivision2name      string                                               `json:"subdivision2name,required"`
	URL                   string                                               `json:"url,required"`
	JSON                  urlScannerScanGetResponseScanPageJSON                `json:"-"`
}

// urlScannerScanGetResponseScanPageJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanPage]
type urlScannerScanGetResponseScanPageJSON struct {
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

func (r *URLScannerScanGetResponseScanPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPageConsole struct {
	Category string                                       `json:"category,required"`
	Text     string                                       `json:"text,required"`
	Type     string                                       `json:"type,required"`
	URL      string                                       `json:"url"`
	JSON     urlScannerScanGetResponseScanPageConsoleJSON `json:"-"`
}

// urlScannerScanGetResponseScanPageConsoleJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanPageConsole]
type urlScannerScanGetResponseScanPageConsoleJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanPageConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPageCooky struct {
	Domain       string                                     `json:"domain,required"`
	Expires      float64                                    `json:"expires,required"`
	HTTPOnly     bool                                       `json:"httpOnly,required"`
	Name         string                                     `json:"name,required"`
	Path         string                                     `json:"path,required"`
	SameParty    bool                                       `json:"sameParty,required"`
	Secure       bool                                       `json:"secure,required"`
	Session      bool                                       `json:"session,required"`
	Size         float64                                    `json:"size,required"`
	SourcePort   float64                                    `json:"sourcePort,required"`
	SourceScheme string                                     `json:"sourceScheme,required"`
	Value        string                                     `json:"value,required"`
	Priority     string                                     `json:"priority"`
	JSON         urlScannerScanGetResponseScanPageCookyJSON `json:"-"`
}

// urlScannerScanGetResponseScanPageCookyJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanPageCooky]
type urlScannerScanGetResponseScanPageCookyJSON struct {
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

func (r *URLScannerScanGetResponseScanPageCooky) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPageHeader struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  urlScannerScanGetResponseScanPageHeaderJSON `json:"-"`
}

// urlScannerScanGetResponseScanPageHeaderJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanPageHeader]
type urlScannerScanGetResponseScanPageHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanPageHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPageJs struct {
	Variables []URLScannerScanGetResponseScanPageJsVariable `json:"variables,required"`
	JSON      urlScannerScanGetResponseScanPageJsJSON       `json:"-"`
}

// urlScannerScanGetResponseScanPageJsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanPageJs]
type urlScannerScanGetResponseScanPageJsJSON struct {
	Variables   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanPageJs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPageJsVariable struct {
	Name string                                          `json:"name,required"`
	Type string                                          `json:"type,required"`
	JSON urlScannerScanGetResponseScanPageJsVariableJSON `json:"-"`
}

// urlScannerScanGetResponseScanPageJsVariableJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseScanPageJsVariable]
type urlScannerScanGetResponseScanPageJsVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanPageJsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPageSecurityViolation struct {
	Category string                                                 `json:"category,required"`
	Text     string                                                 `json:"text,required"`
	URL      string                                                 `json:"url,required"`
	JSON     urlScannerScanGetResponseScanPageSecurityViolationJSON `json:"-"`
}

// urlScannerScanGetResponseScanPageSecurityViolationJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseScanPageSecurityViolation]
type urlScannerScanGetResponseScanPageSecurityViolationJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanPageSecurityViolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanPerformance struct {
	ConnectEnd                 float64                                      `json:"connectEnd,required"`
	ConnectStart               float64                                      `json:"connectStart,required"`
	DecodedBodySize            float64                                      `json:"decodedBodySize,required"`
	DomainLookupEnd            float64                                      `json:"domainLookupEnd,required"`
	DomainLookupStart          float64                                      `json:"domainLookupStart,required"`
	DomComplete                float64                                      `json:"domComplete,required"`
	DomContentLoadedEventEnd   float64                                      `json:"domContentLoadedEventEnd,required"`
	DomContentLoadedEventStart float64                                      `json:"domContentLoadedEventStart,required"`
	DomInteractive             float64                                      `json:"domInteractive,required"`
	Duration                   float64                                      `json:"duration,required"`
	EncodedBodySize            float64                                      `json:"encodedBodySize,required"`
	EntryType                  string                                       `json:"entryType,required"`
	FetchStart                 float64                                      `json:"fetchStart,required"`
	InitiatorType              string                                       `json:"initiatorType,required"`
	LoadEventEnd               float64                                      `json:"loadEventEnd,required"`
	LoadEventStart             float64                                      `json:"loadEventStart,required"`
	Name                       string                                       `json:"name,required"`
	NextHopProtocol            string                                       `json:"nextHopProtocol,required"`
	RedirectCount              float64                                      `json:"redirectCount,required"`
	RedirectEnd                float64                                      `json:"redirectEnd,required"`
	RedirectStart              float64                                      `json:"redirectStart,required"`
	RequestStart               float64                                      `json:"requestStart,required"`
	ResponseEnd                float64                                      `json:"responseEnd,required"`
	ResponseStart              float64                                      `json:"responseStart,required"`
	SecureConnectionStart      float64                                      `json:"secureConnectionStart,required"`
	StartTime                  float64                                      `json:"startTime,required"`
	TransferSize               float64                                      `json:"transferSize,required"`
	Type                       string                                       `json:"type,required"`
	UnloadEventEnd             float64                                      `json:"unloadEventEnd,required"`
	UnloadEventStart           float64                                      `json:"unloadEventStart,required"`
	WorkerStart                float64                                      `json:"workerStart,required"`
	JSON                       urlScannerScanGetResponseScanPerformanceJSON `json:"-"`
}

// urlScannerScanGetResponseScanPerformanceJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanPerformance]
type urlScannerScanGetResponseScanPerformanceJSON struct {
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

func (r *URLScannerScanGetResponseScanPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanTask struct {
	// Submitter location
	ClientLocation string                                      `json:"clientLocation,required"`
	ClientType     URLScannerScanGetResponseScanTaskClientType `json:"clientType,required"`
	// URL of the primary request, after all HTTP redirects
	EffectiveURL string                                       `json:"effectiveUrl,required"`
	Errors       []URLScannerScanGetResponseScanTaskError     `json:"errors,required"`
	ScannedFrom  URLScannerScanGetResponseScanTaskScannedFrom `json:"scannedFrom,required"`
	Status       URLScannerScanGetResponseScanTaskStatus      `json:"status,required"`
	Success      bool                                         `json:"success,required"`
	Time         string                                       `json:"time,required"`
	TimeEnd      string                                       `json:"timeEnd,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID
	Uuid       string                                      `json:"uuid,required"`
	Visibility URLScannerScanGetResponseScanTaskVisibility `json:"visibility,required"`
	JSON       urlScannerScanGetResponseScanTaskJSON       `json:"-"`
}

// urlScannerScanGetResponseScanTaskJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanTask]
type urlScannerScanGetResponseScanTaskJSON struct {
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

func (r *URLScannerScanGetResponseScanTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanTaskClientType string

const (
	URLScannerScanGetResponseScanTaskClientTypeSite      URLScannerScanGetResponseScanTaskClientType = "Site"
	URLScannerScanGetResponseScanTaskClientTypeAutomatic URLScannerScanGetResponseScanTaskClientType = "Automatic"
	URLScannerScanGetResponseScanTaskClientTypeAPI       URLScannerScanGetResponseScanTaskClientType = "Api"
)

type URLScannerScanGetResponseScanTaskError struct {
	Message string                                     `json:"message,required"`
	JSON    urlScannerScanGetResponseScanTaskErrorJSON `json:"-"`
}

// urlScannerScanGetResponseScanTaskErrorJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanTaskError]
type urlScannerScanGetResponseScanTaskErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanTaskError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanTaskScannedFrom struct {
	// IATA code of Cloudflare datacenter
	Colo string                                           `json:"colo,required"`
	JSON urlScannerScanGetResponseScanTaskScannedFromJSON `json:"-"`
}

// urlScannerScanGetResponseScanTaskScannedFromJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseScanTaskScannedFrom]
type urlScannerScanGetResponseScanTaskScannedFromJSON struct {
	Colo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanTaskScannedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanTaskStatus string

const (
	URLScannerScanGetResponseScanTaskStatusQueued           URLScannerScanGetResponseScanTaskStatus = "Queued"
	URLScannerScanGetResponseScanTaskStatusInProgress       URLScannerScanGetResponseScanTaskStatus = "InProgress"
	URLScannerScanGetResponseScanTaskStatusInPostProcessing URLScannerScanGetResponseScanTaskStatus = "InPostProcessing"
	URLScannerScanGetResponseScanTaskStatusFinished         URLScannerScanGetResponseScanTaskStatus = "Finished"
)

type URLScannerScanGetResponseScanTaskVisibility string

const (
	URLScannerScanGetResponseScanTaskVisibilityPublic   URLScannerScanGetResponseScanTaskVisibility = "Public"
	URLScannerScanGetResponseScanTaskVisibilityUnlisted URLScannerScanGetResponseScanTaskVisibility = "Unlisted"
)

type URLScannerScanGetResponseScanVerdicts struct {
	Overall URLScannerScanGetResponseScanVerdictsOverall `json:"overall,required"`
	JSON    urlScannerScanGetResponseScanVerdictsJSON    `json:"-"`
}

// urlScannerScanGetResponseScanVerdictsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanVerdicts]
type urlScannerScanGetResponseScanVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanVerdictsOverall struct {
	Categories []URLScannerScanGetResponseScanVerdictsOverallCategory `json:"categories,required"`
	// Please visit https://safebrowsing.google.com/ for more information.
	GsbThreatTypes []string `json:"gsb_threat_types,required"`
	// At least one of our subsystems marked the site as potentially malicious at the
	// time of the scan.
	Malicious bool                                             `json:"malicious,required"`
	Phishing  []string                                         `json:"phishing,required"`
	JSON      urlScannerScanGetResponseScanVerdictsOverallJSON `json:"-"`
}

// urlScannerScanGetResponseScanVerdictsOverallJSON contains the JSON metadata for
// the struct [URLScannerScanGetResponseScanVerdictsOverall]
type urlScannerScanGetResponseScanVerdictsOverallJSON struct {
	Categories     apijson.Field
	GsbThreatTypes apijson.Field
	Malicious      apijson.Field
	Phishing       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanVerdictsOverallCategory struct {
	ID              float64                                                  `json:"id,required"`
	Name            string                                                   `json:"name,required"`
	SuperCategoryID float64                                                  `json:"super_category_id,required"`
	JSON            urlScannerScanGetResponseScanVerdictsOverallCategoryJSON `json:"-"`
}

// urlScannerScanGetResponseScanVerdictsOverallCategoryJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseScanVerdictsOverallCategory]
type urlScannerScanGetResponseScanVerdictsOverallCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanVerdictsOverallCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Dictionary of Autonomous System Numbers where ASN's are the keys
type URLScannerScanGetResponseScanAsns struct {
	// ASN's contacted
	Asn  URLScannerScanGetResponseScanAsnsAsn  `json:"asn"`
	JSON urlScannerScanGetResponseScanAsnsJSON `json:"-"`
}

// urlScannerScanGetResponseScanAsnsJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanAsns]
type urlScannerScanGetResponseScanAsnsJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanAsns) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ASN's contacted
type URLScannerScanGetResponseScanAsnsAsn struct {
	Asn            string                                   `json:"asn,required"`
	Description    string                                   `json:"description,required"`
	LocationAlpha2 string                                   `json:"location_alpha2,required"`
	Name           string                                   `json:"name,required"`
	OrgName        string                                   `json:"org_name,required"`
	JSON           urlScannerScanGetResponseScanAsnsAsnJSON `json:"-"`
}

// urlScannerScanGetResponseScanAsnsAsnJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanAsnsAsn]
type urlScannerScanGetResponseScanAsnsAsnJSON struct {
	Asn            apijson.Field
	Description    apijson.Field
	LocationAlpha2 apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanAsnsAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomains struct {
	ExampleCom URLScannerScanGetResponseScanDomainsExampleCom `json:"example.com"`
	JSON       urlScannerScanGetResponseScanDomainsJSON       `json:"-"`
}

// urlScannerScanGetResponseScanDomainsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanDomains]
type urlScannerScanGetResponseScanDomainsJSON struct {
	ExampleCom  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleCom struct {
	Categories URLScannerScanGetResponseScanDomainsExampleComCategories `json:"categories,required"`
	DNS        []URLScannerScanGetResponseScanDomainsExampleComDNS      `json:"dns,required"`
	Name       string                                                   `json:"name,required"`
	Rank       URLScannerScanGetResponseScanDomainsExampleComRank       `json:"rank,required"`
	Type       string                                                   `json:"type,required"`
	JSON       urlScannerScanGetResponseScanDomainsExampleComJSON       `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseScanDomainsExampleCom]
type urlScannerScanGetResponseScanDomainsExampleComJSON struct {
	Categories  apijson.Field
	DNS         apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleCom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComCategories struct {
	Inherited URLScannerScanGetResponseScanDomainsExampleComCategoriesInherited `json:"inherited,required"`
	Content   []URLScannerScanGetResponseScanDomainsExampleComCategoriesContent `json:"content"`
	Risks     []URLScannerScanGetResponseScanDomainsExampleComCategoriesRisk    `json:"risks"`
	JSON      urlScannerScanGetResponseScanDomainsExampleComCategoriesJSON      `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComCategoriesJSON contains the JSON
// metadata for the struct
// [URLScannerScanGetResponseScanDomainsExampleComCategories]
type urlScannerScanGetResponseScanDomainsExampleComCategoriesJSON struct {
	Inherited   apijson.Field
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComCategoriesInherited struct {
	Content []URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContent `json:"content"`
	From    string                                                                     `json:"from"`
	Risks   []URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRisk    `json:"risks"`
	JSON    urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedJSON      `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedJSON contains
// the JSON metadata for the struct
// [URLScannerScanGetResponseScanDomainsExampleComCategoriesInherited]
type urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComCategoriesInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContent struct {
	ID              int64                                                                        `json:"id,required"`
	Name            string                                                                       `json:"name,required"`
	SuperCategoryID int64                                                                        `json:"super_category_id"`
	JSON            urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContentJSON `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContentJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContent]
type urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRisk struct {
	ID              int64                                                                     `json:"id,required"`
	Name            string                                                                    `json:"name,required"`
	SuperCategoryID int64                                                                     `json:"super_category_id"`
	JSON            urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRiskJSON `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRiskJSON
// contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRisk]
type urlScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComCategoriesInheritedRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComCategoriesContent struct {
	ID              int64                                                               `json:"id,required"`
	Name            string                                                              `json:"name,required"`
	SuperCategoryID int64                                                               `json:"super_category_id"`
	JSON            urlScannerScanGetResponseScanDomainsExampleComCategoriesContentJSON `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComCategoriesContentJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseScanDomainsExampleComCategoriesContent]
type urlScannerScanGetResponseScanDomainsExampleComCategoriesContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComCategoriesContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComCategoriesRisk struct {
	ID              int64                                                            `json:"id,required"`
	Name            string                                                           `json:"name,required"`
	SuperCategoryID int64                                                            `json:"super_category_id"`
	JSON            urlScannerScanGetResponseScanDomainsExampleComCategoriesRiskJSON `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComCategoriesRiskJSON contains the
// JSON metadata for the struct
// [URLScannerScanGetResponseScanDomainsExampleComCategoriesRisk]
type urlScannerScanGetResponseScanDomainsExampleComCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComDNS struct {
	Address     string                                                `json:"address,required"`
	DnssecValid bool                                                  `json:"dnssec_valid,required"`
	Name        string                                                `json:"name,required"`
	Type        string                                                `json:"type,required"`
	JSON        urlScannerScanGetResponseScanDomainsExampleComDNSJSON `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComDNSJSON contains the JSON metadata
// for the struct [URLScannerScanGetResponseScanDomainsExampleComDNS]
type urlScannerScanGetResponseScanDomainsExampleComDNSJSON struct {
	Address     apijson.Field
	DnssecValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanDomainsExampleComRank struct {
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                                  `json:"rank"`
	JSON urlScannerScanGetResponseScanDomainsExampleComRankJSON `json:"-"`
}

// urlScannerScanGetResponseScanDomainsExampleComRankJSON contains the JSON
// metadata for the struct [URLScannerScanGetResponseScanDomainsExampleComRank]
type urlScannerScanGetResponseScanDomainsExampleComRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanDomainsExampleComRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanIPs struct {
	IP   URLScannerScanGetResponseScanIPsIP   `json:"ip"`
	JSON urlScannerScanGetResponseScanIPsJSON `json:"-"`
}

// urlScannerScanGetResponseScanIPsJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanIPs]
type urlScannerScanGetResponseScanIPsJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanIPsIP struct {
	Asn               string                                 `json:"asn,required"`
	AsnDescription    string                                 `json:"asnDescription,required"`
	AsnLocationAlpha2 string                                 `json:"asnLocationAlpha2,required"`
	AsnName           string                                 `json:"asnName,required"`
	AsnOrgName        string                                 `json:"asnOrgName,required"`
	Continent         string                                 `json:"continent,required"`
	GeonameID         string                                 `json:"geonameId,required"`
	IP                string                                 `json:"ip,required"`
	IPVersion         string                                 `json:"ipVersion,required"`
	Latitude          string                                 `json:"latitude,required"`
	LocationAlpha2    string                                 `json:"locationAlpha2,required"`
	LocationName      string                                 `json:"locationName,required"`
	Longitude         string                                 `json:"longitude,required"`
	Subdivision1Name  string                                 `json:"subdivision1Name,required"`
	Subdivision2Name  string                                 `json:"subdivision2Name,required"`
	JSON              urlScannerScanGetResponseScanIPsIPJSON `json:"-"`
}

// urlScannerScanGetResponseScanIPsIPJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanIPsIP]
type urlScannerScanGetResponseScanIPsIPJSON struct {
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

func (r *URLScannerScanGetResponseScanIPsIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanLinks struct {
	Link URLScannerScanGetResponseScanLinksLink `json:"link"`
	JSON urlScannerScanGetResponseScanLinksJSON `json:"-"`
}

// urlScannerScanGetResponseScanLinksJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseScanLinks]
type urlScannerScanGetResponseScanLinksJSON struct {
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseScanLinksLink struct {
	// Outgoing link detected in the DOM
	Href string                                     `json:"href,required"`
	Text string                                     `json:"text,required"`
	JSON urlScannerScanGetResponseScanLinksLinkJSON `json:"-"`
}

// urlScannerScanGetResponseScanLinksLinkJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseScanLinksLink]
type urlScannerScanGetResponseScanLinksLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseScanLinksLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponse struct {
	Har  URLScannerScanHarResponseHar  `json:"har,required"`
	JSON urlScannerScanHarResponseJSON `json:"-"`
}

// urlScannerScanHarResponseJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponse]
type urlScannerScanHarResponseJSON struct {
	Har         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHar struct {
	Log  URLScannerScanHarResponseHarLog  `json:"log,required"`
	JSON urlScannerScanHarResponseHarJSON `json:"-"`
}

// urlScannerScanHarResponseHarJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseHar]
type urlScannerScanHarResponseHarJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLog struct {
	Creator URLScannerScanHarResponseHarLogCreator `json:"creator,required"`
	Entries []URLScannerScanHarResponseHarLogEntry `json:"entries,required"`
	Pages   []URLScannerScanHarResponseHarLogPage  `json:"pages,required"`
	Version string                                 `json:"version,required"`
	JSON    urlScannerScanHarResponseHarLogJSON    `json:"-"`
}

// urlScannerScanHarResponseHarLogJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseHarLog]
type urlScannerScanHarResponseHarLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogCreator struct {
	Comment string                                     `json:"comment,required"`
	Name    string                                     `json:"name,required"`
	Version string                                     `json:"version,required"`
	JSON    urlScannerScanHarResponseHarLogCreatorJSON `json:"-"`
}

// urlScannerScanHarResponseHarLogCreatorJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseHarLogCreator]
type urlScannerScanHarResponseHarLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogEntry struct {
	InitialPriority string                                         `json:"_initialPriority,required"`
	InitiatorType   string                                         `json:"_initiator_type,required"`
	Priority        string                                         `json:"_priority,required"`
	RequestID       string                                         `json:"_requestId,required"`
	RequestTime     float64                                        `json:"_requestTime,required"`
	ResourceType    string                                         `json:"_resourceType,required"`
	Cache           interface{}                                    `json:"cache,required"`
	Connection      string                                         `json:"connection,required"`
	Pageref         string                                         `json:"pageref,required"`
	Request         URLScannerScanHarResponseHarLogEntriesRequest  `json:"request,required"`
	Response        URLScannerScanHarResponseHarLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                                         `json:"serverIPAddress,required"`
	StartedDateTime string                                         `json:"startedDateTime,required"`
	Time            float64                                        `json:"time,required"`
	JSON            urlScannerScanHarResponseHarLogEntryJSON       `json:"-"`
}

// urlScannerScanHarResponseHarLogEntryJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseHarLogEntry]
type urlScannerScanHarResponseHarLogEntryJSON struct {
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

func (r *URLScannerScanHarResponseHarLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogEntriesRequest struct {
	BodySize    float64                                               `json:"bodySize,required"`
	Headers     []URLScannerScanHarResponseHarLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                               `json:"headersSize,required"`
	HTTPVersion string                                                `json:"httpVersion,required"`
	Method      string                                                `json:"method,required"`
	URL         string                                                `json:"url,required"`
	JSON        urlScannerScanHarResponseHarLogEntriesRequestJSON     `json:"-"`
}

// urlScannerScanHarResponseHarLogEntriesRequestJSON contains the JSON metadata for
// the struct [URLScannerScanHarResponseHarLogEntriesRequest]
type urlScannerScanHarResponseHarLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogEntriesRequestHeader struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  urlScannerScanHarResponseHarLogEntriesRequestHeaderJSON `json:"-"`
}

// urlScannerScanHarResponseHarLogEntriesRequestHeaderJSON contains the JSON
// metadata for the struct [URLScannerScanHarResponseHarLogEntriesRequestHeader]
type urlScannerScanHarResponseHarLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogEntriesResponse struct {
	TransferSize float64                                                `json:"_transferSize,required"`
	BodySize     float64                                                `json:"bodySize,required"`
	Content      URLScannerScanHarResponseHarLogEntriesResponseContent  `json:"content,required"`
	Headers      []URLScannerScanHarResponseHarLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                                `json:"headersSize,required"`
	HTTPVersion  string                                                 `json:"httpVersion,required"`
	RedirectURL  string                                                 `json:"redirectURL,required"`
	Status       float64                                                `json:"status,required"`
	StatusText   string                                                 `json:"statusText,required"`
	JSON         urlScannerScanHarResponseHarLogEntriesResponseJSON     `json:"-"`
}

// urlScannerScanHarResponseHarLogEntriesResponseJSON contains the JSON metadata
// for the struct [URLScannerScanHarResponseHarLogEntriesResponse]
type urlScannerScanHarResponseHarLogEntriesResponseJSON struct {
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

func (r *URLScannerScanHarResponseHarLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogEntriesResponseContent struct {
	MimeType    string                                                    `json:"mimeType,required"`
	Size        float64                                                   `json:"size,required"`
	Compression int64                                                     `json:"compression"`
	JSON        urlScannerScanHarResponseHarLogEntriesResponseContentJSON `json:"-"`
}

// urlScannerScanHarResponseHarLogEntriesResponseContentJSON contains the JSON
// metadata for the struct [URLScannerScanHarResponseHarLogEntriesResponseContent]
type urlScannerScanHarResponseHarLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogEntriesResponseHeader struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  urlScannerScanHarResponseHarLogEntriesResponseHeaderJSON `json:"-"`
}

// urlScannerScanHarResponseHarLogEntriesResponseHeaderJSON contains the JSON
// metadata for the struct [URLScannerScanHarResponseHarLogEntriesResponseHeader]
type urlScannerScanHarResponseHarLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogPage struct {
	ID              string                                          `json:"id,required"`
	PageTimings     URLScannerScanHarResponseHarLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                                          `json:"startedDateTime,required"`
	Title           string                                          `json:"title,required"`
	JSON            urlScannerScanHarResponseHarLogPageJSON         `json:"-"`
}

// urlScannerScanHarResponseHarLogPageJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseHarLogPage]
type urlScannerScanHarResponseHarLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseHarLogPagesPageTimings struct {
	OnContentLoad float64                                             `json:"onContentLoad,required"`
	OnLoad        float64                                             `json:"onLoad,required"`
	JSON          urlScannerScanHarResponseHarLogPagesPageTimingsJSON `json:"-"`
}

// urlScannerScanHarResponseHarLogPagesPageTimingsJSON contains the JSON metadata
// for the struct [URLScannerScanHarResponseHarLogPagesPageTimings]
type urlScannerScanHarResponseHarLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *URLScannerScanHarResponseHarLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewParams struct {
	URL param.Field[string] `json:"url,required"`
	// Set custom headers
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
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

type URLScannerScanNewResponseEnvelope struct {
	Errors   []URLScannerScanNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []URLScannerScanNewResponseEnvelopeMessages `json:"messages,required"`
	Result   URLScannerScanNewResponse                   `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	JSON     urlScannerScanNewResponseEnvelopeJSON       `json:"-"`
}

// urlScannerScanNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [URLScannerScanNewResponseEnvelope]
type urlScannerScanNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewResponseEnvelopeErrors struct {
	Message string                                      `json:"message,required"`
	JSON    urlScannerScanNewResponseEnvelopeErrorsJSON `json:"-"`
}

// urlScannerScanNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [URLScannerScanNewResponseEnvelopeErrors]
type urlScannerScanNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanNewResponseEnvelopeMessages struct {
	Message string                                        `json:"message,required"`
	JSON    urlScannerScanNewResponseEnvelopeMessagesJSON `json:"-"`
}

// urlScannerScanNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [URLScannerScanNewResponseEnvelopeMessages]
type urlScannerScanNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseEnvelope struct {
	Errors   []URLScannerScanGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []URLScannerScanGetResponseEnvelopeMessages `json:"messages,required"`
	Result   URLScannerScanGetResponse                   `json:"result,required"`
	// Whether request was successful or not
	Success bool                                  `json:"success,required"`
	JSON    urlScannerScanGetResponseEnvelopeJSON `json:"-"`
}

// urlScannerScanGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [URLScannerScanGetResponseEnvelope]
type urlScannerScanGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseEnvelopeErrors struct {
	Message string                                      `json:"message,required"`
	JSON    urlScannerScanGetResponseEnvelopeErrorsJSON `json:"-"`
}

// urlScannerScanGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseEnvelopeErrors]
type urlScannerScanGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanGetResponseEnvelopeMessages struct {
	Message string                                        `json:"message,required"`
	JSON    urlScannerScanGetResponseEnvelopeMessagesJSON `json:"-"`
}

// urlScannerScanGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [URLScannerScanGetResponseEnvelopeMessages]
type urlScannerScanGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseEnvelope struct {
	Errors   []URLScannerScanHarResponseEnvelopeErrors   `json:"errors,required"`
	Messages []URLScannerScanHarResponseEnvelopeMessages `json:"messages,required"`
	Result   URLScannerScanHarResponse                   `json:"result,required"`
	// Whether search request was successful or not
	Success bool                                  `json:"success,required"`
	JSON    urlScannerScanHarResponseEnvelopeJSON `json:"-"`
}

// urlScannerScanHarResponseEnvelopeJSON contains the JSON metadata for the struct
// [URLScannerScanHarResponseEnvelope]
type urlScannerScanHarResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseEnvelopeErrors struct {
	Message string                                      `json:"message,required"`
	JSON    urlScannerScanHarResponseEnvelopeErrorsJSON `json:"-"`
}

// urlScannerScanHarResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseEnvelopeErrors]
type urlScannerScanHarResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLScannerScanHarResponseEnvelopeMessages struct {
	Message string                                        `json:"message,required"`
	JSON    urlScannerScanHarResponseEnvelopeMessagesJSON `json:"-"`
}

// urlScannerScanHarResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [URLScannerScanHarResponseEnvelopeMessages]
type urlScannerScanHarResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerScanHarResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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
