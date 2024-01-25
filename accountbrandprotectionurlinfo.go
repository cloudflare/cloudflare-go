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

// AccountBrandProtectionURLInfoService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountBrandProtectionURLInfoService] method instead.
type AccountBrandProtectionURLInfoService struct {
	Options []option.RequestOption
}

// NewAccountBrandProtectionURLInfoService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountBrandProtectionURLInfoService(opts ...option.RequestOption) (r *AccountBrandProtectionURLInfoService) {
	r = &AccountBrandProtectionURLInfoService{}
	r.Options = opts
	return
}

// Get results for a URL scan
func (r *AccountBrandProtectionURLInfoService) PhishingURLInformationGetResultsForAURLScan(ctx context.Context, accountIdentifier string, query AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams, opts ...option.RequestOption) (res *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse struct {
	Errors   []AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseError   `json:"errors"`
	Messages []AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessage `json:"messages"`
	Result   AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseSuccess `json:"success"`
	JSON    accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseJSON    `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseError struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseErrorJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseError]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessage struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessageJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessage]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResult struct {
	// List of categorizations applied to this submission.
	Categorizations []AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorization `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResult `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatch `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatus `json:"scan_status"`
	// For internal use.
	ScreenshotDownloadSignature string `json:"screenshot_download_signature"`
	// For internal use.
	ScreenshotPath string `json:"screenshot_path"`
	// URL that was submitted.
	URL  string                                                                                     `json:"url"`
	JSON accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResult]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultJSON struct {
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

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorization struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                                                                                                   `json:"verification_status"`
	JSON               accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorizationJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorizationJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorization]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorizationJSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResult struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                                                                                               `json:"model_score"`
	JSON       accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResultJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResultJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResult]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResultJSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultModelResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatch struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                                                                                              `json:"name"`
	JSON accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatchJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatchJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatch]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatchJSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultRuleMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the most recent scan found.
type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatus struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                                                                                                `json:"submission_id"`
	JSON         accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatusJSON `json:"-"`
}

// accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatusJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatus]
type accountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatusJSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseResultScanStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseSuccess bool

const (
	AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseSuccessTrue AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanResponseSuccess = true
)

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams struct {
	URL        param.Field[string]                                                                                   `query:"url"`
	URLIDParam param.Field[AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam] `query:"url_id_param"`
}

// URLQuery serializes
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams]'s
// query parameters as `url.Values`.
func (r AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAURLScanParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam struct {
	// Submission ID(s) to filter submission results by.
	URLID param.Field[int64] `query:"url_id"`
}

// URLQuery serializes
// [AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam]'s
// query parameters as `url.Values`.
func (r AccountBrandProtectionURLInfoPhishingURLInformationGetResultsForAurlScanParamsURLIDParam) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
