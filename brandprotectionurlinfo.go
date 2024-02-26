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
func (r *BrandProtectionURLInfoService) Get(ctx context.Context, params BrandProtectionURLInfoGetParams, opts ...option.RequestOption) (res *BrandProtectionURLInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BrandProtectionURLInfoGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BrandProtectionURLInfoGetResponse struct {
	// List of categorizations applied to this submission.
	Categorizations []BrandProtectionURLInfoGetResponseCategorization `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []BrandProtectionURLInfoGetResponseModelResult `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []BrandProtectionURLInfoGetResponseRuleMatch `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus BrandProtectionURLInfoGetResponseScanStatus `json:"scan_status"`
	// For internal use.
	ScreenshotDownloadSignature string `json:"screenshot_download_signature"`
	// For internal use.
	ScreenshotPath string `json:"screenshot_path"`
	// URL that was submitted.
	URL  string                                `json:"url"`
	JSON brandProtectionURLInfoGetResponseJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseJSON contains the JSON metadata for the struct
// [BrandProtectionURLInfoGetResponse]
type brandProtectionURLInfoGetResponseJSON struct {
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

func (r *BrandProtectionURLInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoGetResponseCategorization struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                                              `json:"verification_status"`
	JSON               brandProtectionURLInfoGetResponseCategorizationJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseCategorizationJSON contains the JSON metadata
// for the struct [BrandProtectionURLInfoGetResponseCategorization]
type brandProtectionURLInfoGetResponseCategorizationJSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoGetResponseModelResult struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                                          `json:"model_score"`
	JSON       brandProtectionURLInfoGetResponseModelResultJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseModelResultJSON contains the JSON metadata for
// the struct [BrandProtectionURLInfoGetResponseModelResult]
type brandProtectionURLInfoGetResponseModelResultJSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseModelResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoGetResponseRuleMatch struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                                         `json:"name"`
	JSON brandProtectionURLInfoGetResponseRuleMatchJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseRuleMatchJSON contains the JSON metadata for
// the struct [BrandProtectionURLInfoGetResponseRuleMatch]
type brandProtectionURLInfoGetResponseRuleMatchJSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseRuleMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the most recent scan found.
type BrandProtectionURLInfoGetResponseScanStatus struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                                           `json:"submission_id"`
	JSON         brandProtectionURLInfoGetResponseScanStatusJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseScanStatusJSON contains the JSON metadata for
// the struct [BrandProtectionURLInfoGetResponseScanStatus]
type brandProtectionURLInfoGetResponseScanStatusJSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseScanStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoGetParams struct {
	// Identifier
	AccountID  param.Field[string]                                    `path:"account_id,required"`
	URL        param.Field[string]                                    `query:"url"`
	URLIDParam param.Field[BrandProtectionURLInfoGetParamsURLIDParam] `query:"url_id_param"`
}

// URLQuery serializes [BrandProtectionURLInfoGetParams]'s query parameters as
// `url.Values`.
func (r BrandProtectionURLInfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandProtectionURLInfoGetParamsURLIDParam struct {
	// Submission ID(s) to filter submission results by.
	URLID param.Field[int64] `query:"url_id"`
}

// URLQuery serializes [BrandProtectionURLInfoGetParamsURLIDParam]'s query
// parameters as `url.Values`.
func (r BrandProtectionURLInfoGetParamsURLIDParam) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BrandProtectionURLInfoGetResponseEnvelope struct {
	Errors   []BrandProtectionURLInfoGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionURLInfoGetResponseEnvelopeMessages `json:"messages,required"`
	Result   BrandProtectionURLInfoGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BrandProtectionURLInfoGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    brandProtectionURLInfoGetResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionURLInfoGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrandProtectionURLInfoGetResponseEnvelope]
type brandProtectionURLInfoGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    brandProtectionURLInfoGetResponseEnvelopeErrorsJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [BrandProtectionURLInfoGetResponseEnvelopeErrors]
type brandProtectionURLInfoGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionURLInfoGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    brandProtectionURLInfoGetResponseEnvelopeMessagesJSON `json:"-"`
}

// brandProtectionURLInfoGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BrandProtectionURLInfoGetResponseEnvelopeMessages]
type brandProtectionURLInfoGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionURLInfoGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BrandProtectionURLInfoGetResponseEnvelopeSuccess bool

const (
	BrandProtectionURLInfoGetResponseEnvelopeSuccessTrue BrandProtectionURLInfoGetResponseEnvelopeSuccess = true
)
