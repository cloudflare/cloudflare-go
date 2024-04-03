// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// BrandProtectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBrandProtectionService] method
// instead.
type BrandProtectionService struct {
	Options []option.RequestOption
}

// NewBrandProtectionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBrandProtectionService(opts ...option.RequestOption) (r *BrandProtectionService) {
	r = &BrandProtectionService{}
	r.Options = opts
	return
}

// Submit suspicious URL for scanning
func (r *BrandProtectionService) Submit(ctx context.Context, params BrandProtectionSubmitParams, opts ...option.RequestOption) (res *IntelPhishingURLSubmit, err error) {
	opts = append(r.Options[:], opts...)
	var env BrandProtectionSubmitResponseEnvelope
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get results for a URL scan
func (r *BrandProtectionService) URLInfo(ctx context.Context, params BrandProtectionURLInfoParams, opts ...option.RequestOption) (res *IntelPhishingURLInfo, err error) {
	opts = append(r.Options[:], opts...)
	var env BrandProtectionURLInfoResponseEnvelope
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelPhishingURLInfo struct {
	// List of categorizations applied to this submission.
	Categorizations []IntelPhishingURLInfoCategorization `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []IntelPhishingURLInfoModelResult `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []IntelPhishingURLInfoRuleMatch `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus IntelPhishingURLInfoScanStatus `json:"scan_status"`
	// For internal use.
	ScreenshotDownloadSignature string `json:"screenshot_download_signature"`
	// For internal use.
	ScreenshotPath string `json:"screenshot_path"`
	// URL that was submitted.
	URL  string                   `json:"url"`
	JSON intelPhishingURLInfoJSON `json:"-"`
}

// intelPhishingURLInfoJSON contains the JSON metadata for the struct
// [IntelPhishingURLInfo]
type intelPhishingURLInfoJSON struct {
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

func (r *IntelPhishingURLInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLInfoJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLInfoCategorization struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                                 `json:"verification_status"`
	JSON               intelPhishingURLInfoCategorizationJSON `json:"-"`
}

// intelPhishingURLInfoCategorizationJSON contains the JSON metadata for the struct
// [IntelPhishingURLInfoCategorization]
type intelPhishingURLInfoCategorizationJSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IntelPhishingURLInfoCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLInfoCategorizationJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLInfoModelResult struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                             `json:"model_score"`
	JSON       intelPhishingURLInfoModelResultJSON `json:"-"`
}

// intelPhishingURLInfoModelResultJSON contains the JSON metadata for the struct
// [IntelPhishingURLInfoModelResult]
type intelPhishingURLInfoModelResultJSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPhishingURLInfoModelResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLInfoModelResultJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLInfoRuleMatch struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                            `json:"name"`
	JSON intelPhishingURLInfoRuleMatchJSON `json:"-"`
}

// intelPhishingURLInfoRuleMatchJSON contains the JSON metadata for the struct
// [IntelPhishingURLInfoRuleMatch]
type intelPhishingURLInfoRuleMatchJSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPhishingURLInfoRuleMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLInfoRuleMatchJSON) RawJSON() string {
	return r.raw
}

// Status of the most recent scan found.
type IntelPhishingURLInfoScanStatus struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                              `json:"submission_id"`
	JSON         intelPhishingURLInfoScanStatusJSON `json:"-"`
}

// intelPhishingURLInfoScanStatusJSON contains the JSON metadata for the struct
// [IntelPhishingURLInfoScanStatus]
type intelPhishingURLInfoScanStatusJSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IntelPhishingURLInfoScanStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLInfoScanStatusJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLSubmit struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []IntelPhishingURLSubmitExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []IntelPhishingURLSubmitSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []IntelPhishingURLSubmitSubmittedURL `json:"submitted_urls"`
	JSON          intelPhishingURLSubmitJSON           `json:"-"`
}

// intelPhishingURLSubmitJSON contains the JSON metadata for the struct
// [IntelPhishingURLSubmit]
type intelPhishingURLSubmitJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IntelPhishingURLSubmit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLSubmitJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLSubmitExcludedURL struct {
	// URL that was excluded.
	URL  string                                `json:"url"`
	JSON intelPhishingURLSubmitExcludedURLJSON `json:"-"`
}

// intelPhishingURLSubmitExcludedURLJSON contains the JSON metadata for the struct
// [IntelPhishingURLSubmitExcludedURL]
type intelPhishingURLSubmitExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPhishingURLSubmitExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLSubmitExcludedURLJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLSubmitSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                `json:"url_id"`
	JSON  intelPhishingURLSubmitSkippedURLJSON `json:"-"`
}

// intelPhishingURLSubmitSkippedURLJSON contains the JSON metadata for the struct
// [IntelPhishingURLSubmitSkippedURL]
type intelPhishingURLSubmitSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPhishingURLSubmitSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLSubmitSkippedURLJSON) RawJSON() string {
	return r.raw
}

type IntelPhishingURLSubmitSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                  `json:"url_id"`
	JSON  intelPhishingURLSubmitSubmittedURLJSON `json:"-"`
}

// intelPhishingURLSubmitSubmittedURLJSON contains the JSON metadata for the struct
// [IntelPhishingURLSubmitSubmittedURL]
type intelPhishingURLSubmitSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelPhishingURLSubmitSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelPhishingURLSubmitSubmittedURLJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// URL(s) to filter submissions results by
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r BrandProtectionSubmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandProtectionSubmitResponseEnvelope struct {
	Errors   []BrandProtectionSubmitResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionSubmitResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelPhishingURLSubmit                          `json:"result,required"`
	// Whether the API call was successful
	Success BrandProtectionSubmitResponseEnvelopeSuccess `json:"success,required"`
	JSON    brandProtectionSubmitResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrandProtectionSubmitResponseEnvelope]
type brandProtectionSubmitResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    brandProtectionSubmitResponseEnvelopeErrorsJSON `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitResponseEnvelopeErrors]
type brandProtectionSubmitResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionSubmitResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    brandProtectionSubmitResponseEnvelopeMessagesJSON `json:"-"`
}

// brandProtectionSubmitResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitResponseEnvelopeMessages]
type brandProtectionSubmitResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionSubmitResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// Identifier
	AccountID  param.Field[string]                                 `path:"account_id,required"`
	URL        param.Field[string]                                 `query:"url"`
	URLIDParam param.Field[BrandProtectionURLInfoParamsURLIDParam] `query:"url_id_param"`
}

// URLQuery serializes [BrandProtectionURLInfoParams]'s query parameters as
// `url.Values`.
func (r BrandProtectionURLInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandProtectionURLInfoParamsURLIDParam struct {
	// Submission ID(s) to filter submission results by.
	URLID param.Field[int64] `query:"url_id"`
}

// URLQuery serializes [BrandProtectionURLInfoParamsURLIDParam]'s query parameters
// as `url.Values`.
func (r BrandProtectionURLInfoParamsURLIDParam) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandProtectionURLInfoResponseEnvelope struct {
	Errors   []BrandProtectionURLInfoResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionURLInfoResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelPhishingURLInfo                             `json:"result,required"`
	// Whether the API call was successful
	Success BrandProtectionURLInfoResponseEnvelopeSuccess `json:"success,required"`
	JSON    brandProtectionURLInfoResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrandProtectionURLInfoResponseEnvelope]
type brandProtectionURLInfoResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    brandProtectionURLInfoResponseEnvelopeErrorsJSON `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [BrandProtectionURLInfoResponseEnvelopeErrors]
type brandProtectionURLInfoResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type BrandProtectionURLInfoResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    brandProtectionURLInfoResponseEnvelopeMessagesJSON `json:"-"`
}

// brandProtectionURLInfoResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BrandProtectionURLInfoResponseEnvelopeMessages]
type brandProtectionURLInfoResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r brandProtectionURLInfoResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
