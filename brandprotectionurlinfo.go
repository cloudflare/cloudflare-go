// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// BrandProtectionURLInfoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBrandProtectionURLInfoService]
// method instead.
type BrandProtectionURLInfoService struct {
	Options []option.RequestOption
}

// NewBrandProtectionURLInfoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrandProtectionURLInfoService(opts ...option.RequestOption) (r *BrandProtectionURLInfoService) {
	r = &BrandProtectionURLInfoService{}
	r.Options = opts
	return
}

// Get results for a URL scan
func (r *BrandProtectionURLInfoService) PhishingURLInformationGetResultsForAURLScan(ctx context.Context, accountID string, query BrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams, opts ...option.RequestOption) (res *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelope
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse struct {
	// List of categorizations applied to this submission.
	Categorizations []BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorization `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResult `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatch `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatus `json:"scan_status"`
	// For internal use.
	ScreenshotDownloadSignature string `json:"screenshot_download_signature"`
	// For internal use.
	ScreenshotPath string `json:"screenshot_path"`
	// URL that was submitted.
	URL  string                                                                        `json:"url"`
	JSON brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseJSON struct {
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

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorization struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                                                                                      `json:"verification_status"`
	JSON               brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorizationJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorizationJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorization]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorizationJSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResult struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                                                                                  `json:"model_score"`
	JSON       brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResultJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResultJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResult]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResultJSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseModelResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatch struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                                                                                 `json:"name"`
	JSON brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatchJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatchJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatch]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatchJSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseRuleMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the most recent scan found.
type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatus struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                                                                                   `json:"submission_id"`
	JSON         brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatusJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatusJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatus]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatusJSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseScanStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams struct {
	URL        param.Field[string]                                                                            `query:"url"`
	URLIDParam param.Field[BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam] `query:"url_id_param"`
}

// URLQuery serializes
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams]'s
// query parameters as `url.Values`.
func (r BrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam struct {
	// Submission ID(s) to filter submission results by.
	URLID param.Field[int64] `query:"url_id"`
}

// URLQuery serializes
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam]'s
// query parameters as `url.Values`.
func (r BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelope struct {
	Errors   []BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessages `json:"messages,required"`
	Result   BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeSuccess `json:"success,required"`
	JSON    brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelope]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrors struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrorsJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrors]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessages struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessagesJSON `json:"-"`
}

// brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessages]
type brandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeSuccess bool

const (
	BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeSuccessTrue BrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseEnvelopeSuccess = true
)
