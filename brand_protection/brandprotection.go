// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// BrandProtectionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrandProtectionService] method instead.
type BrandProtectionService struct {
	Options     []option.RequestOption
	Queries     *QueryService
	Matches     *MatchService
	Logos       *LogoService
	LogoMatches *LogoMatchService
}

// NewBrandProtectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrandProtectionService(opts ...option.RequestOption) (r *BrandProtectionService) {
	r = &BrandProtectionService{}
	r.Options = opts
	r.Queries = NewQueryService(opts...)
	r.Matches = NewMatchService(opts...)
	r.Logos = NewLogoService(opts...)
	r.LogoMatches = NewLogoMatchService(opts...)
	return
}

// Submit suspicious URL for scanning.
func (r *BrandProtectionService) Submit(ctx context.Context, params BrandProtectionSubmitParams, opts ...option.RequestOption) (res *Submit, err error) {
	var env BrandProtectionSubmitResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets phishing details about a URL.
func (r *BrandProtectionService) URLInfo(ctx context.Context, params BrandProtectionURLInfoParams, opts ...option.RequestOption) (res *Info, err error) {
	var env BrandProtectionURLInfoResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Info struct {
	// List of categorizations applied to this submission.
	Categorizations []InfoCategorization `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []URLInfoModelResults `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []RuleMatch `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus ScanStatus `json:"scan_status"`
	// For internal use.
	ScreenshotDownloadSignature string `json:"screenshot_download_signature"`
	// For internal use.
	ScreenshotPath string `json:"screenshot_path"`
	// URL that was submitted.
	URL  string   `json:"url"`
	JSON infoJSON `json:"-"`
}

// infoJSON contains the JSON metadata for the struct [Info]
type infoJSON struct {
	Categorizations             apijson.Field
	ModelResults                apijson.Field
	RuleMatches                 apijson.Field
	ScanStatus                  apijson.Field
	ScreenshotDownloadSignature apijson.Field
	ScreenshotPath              apijson.Field
	URL                         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *Info) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r infoJSON) RawJSON() string {
	return r.raw
}

type InfoCategorization struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                 `json:"verification_status"`
	JSON               infoCategorizationJSON `json:"-"`
}

// infoCategorizationJSON contains the JSON metadata for the struct
// [InfoCategorization]
type infoCategorizationJSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *InfoCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r infoCategorizationJSON) RawJSON() string {
	return r.raw
}

type RuleMatch struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string        `json:"name"`
	JSON ruleMatchJSON `json:"-"`
}

// ruleMatchJSON contains the JSON metadata for the struct [RuleMatch]
type ruleMatchJSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleMatchJSON) RawJSON() string {
	return r.raw
}

// Status of the most recent scan found.
type ScanStatus struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64          `json:"submission_id"`
	JSON         scanStatusJSON `json:"-"`
}

// scanStatusJSON contains the JSON metadata for the struct [ScanStatus]
type scanStatusJSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScanStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanStatusJSON) RawJSON() string {
	return r.raw
}

type Submit struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []SubmitExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned.
	SkippedURLs []SubmitSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []SubmitSubmittedURL `json:"submitted_urls"`
	JSON          submitJSON           `json:"-"`
}

// submitJSON contains the JSON metadata for the struct [Submit]
type submitJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Submit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submitJSON) RawJSON() string {
	return r.raw
}

type SubmitExcludedURL struct {
	// URL that was excluded.
	URL  string                `json:"url"`
	JSON submitExcludedURLJSON `json:"-"`
}

// submitExcludedURLJSON contains the JSON metadata for the struct
// [SubmitExcludedURL]
type submitExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmitExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submitExcludedURLJSON) RawJSON() string {
	return r.raw
}

type SubmitSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                `json:"url_id"`
	JSON  submitSkippedURLJSON `json:"-"`
}

// submitSkippedURLJSON contains the JSON metadata for the struct
// [SubmitSkippedURL]
type submitSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmitSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submitSkippedURLJSON) RawJSON() string {
	return r.raw
}

type SubmitSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                  `json:"url_id"`
	JSON  submitSubmittedURLJSON `json:"-"`
}

// submitSubmittedURLJSON contains the JSON metadata for the struct
// [SubmitSubmittedURL]
type submitSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmitSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submitSubmittedURLJSON) RawJSON() string {
	return r.raw
}

type URLInfoModelResults struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// This is the score that is outputted by the model for this submission.
	ModelScore float64                 `json:"model_score"`
	JSON       urlInfoModelResultsJSON `json:"-"`
}

// urlInfoModelResultsJSON contains the JSON metadata for the struct
// [URLInfoModelResults]
type urlInfoModelResultsJSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLInfoModelResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlInfoModelResultsJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// URL(s) to filter submissions results by.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r BrandProtectionSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandProtectionSubmitResponseEnvelope struct {
	Errors   []BrandProtectionSubmitResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionSubmitResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success BrandProtectionSubmitResponseEnvelopeSuccess `json:"success,required"`
	Result  Submit                                       `json:"result"`
	JSON    brandProtectionSubmitResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrandProtectionSubmitResponseEnvelope]
type brandProtectionSubmitResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           BrandProtectionSubmitResponseEnvelopeErrorsSource `json:"source"`
	JSON             brandProtectionSubmitResponseEnvelopeErrorsJSON   `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitResponseEnvelopeErrors]
type brandProtectionSubmitResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    brandProtectionSubmitResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [BrandProtectionSubmitResponseEnvelopeErrorsSource]
type brandProtectionSubmitResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           BrandProtectionSubmitResponseEnvelopeMessagesSource `json:"source"`
	JSON             brandProtectionSubmitResponseEnvelopeMessagesJSON   `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitResponseEnvelopeMessages]
type brandProtectionSubmitResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    brandProtectionSubmitResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [BrandProtectionSubmitResponseEnvelopeMessagesSource]
type brandProtectionSubmitResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BrandProtectionSubmitResponseEnvelopeSuccess bool

const (
	BrandProtectionSubmitResponseEnvelopeSuccessTrue BrandProtectionSubmitResponseEnvelopeSuccess = true
)

func (r BrandProtectionSubmitResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BrandProtectionSubmitResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BrandProtectionURLInfoParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Submission URL(s) to filter submission results by.
	URL param.Field[[]string] `query:"url"`
	// Submission ID(s) to filter submission results by.
	URLID param.Field[[]int64] `query:"url_id"`
}

// URLQuery serializes [BrandProtectionURLInfoParams]'s query parameters as
// `url.Values`.
func (r BrandProtectionURLInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type BrandProtectionURLInfoResponseEnvelope struct {
	Errors   []BrandProtectionURLInfoResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionURLInfoResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success BrandProtectionURLInfoResponseEnvelopeSuccess `json:"success,required"`
	Result  Info                                          `json:"result"`
	JSON    brandProtectionURLInfoResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrandProtectionURLInfoResponseEnvelope]
type brandProtectionURLInfoResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           BrandProtectionURLInfoResponseEnvelopeErrorsSource `json:"source"`
	JSON             brandProtectionURLInfoResponseEnvelopeErrorsJSON   `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BrandProtectionURLInfoResponseEnvelopeErrors]
type brandProtectionURLInfoResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    brandProtectionURLInfoResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [BrandProtectionURLInfoResponseEnvelopeErrorsSource]
type brandProtectionURLInfoResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           BrandProtectionURLInfoResponseEnvelopeMessagesSource `json:"source"`
	JSON             brandProtectionURLInfoResponseEnvelopeMessagesJSON   `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BrandProtectionURLInfoResponseEnvelopeMessages]
type brandProtectionURLInfoResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    brandProtectionURLInfoResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [BrandProtectionURLInfoResponseEnvelopeMessagesSource]
type brandProtectionURLInfoResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BrandProtectionURLInfoResponseEnvelopeSuccess bool

const (
	BrandProtectionURLInfoResponseEnvelopeSuccessTrue BrandProtectionURLInfoResponseEnvelopeSuccess = true
)

func (r BrandProtectionURLInfoResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BrandProtectionURLInfoResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
