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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
	Categorizations []shared.UnnamedSchemaRef13 `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []shared.UnnamedSchemaRef14 `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []shared.UnnamedSchemaRef15 `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus shared.UnnamedSchemaRef16 `json:"scan_status"`
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

type IntelPhishingURLSubmit struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []shared.UnnamedSchemaRef10 `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []shared.UnnamedSchemaRef11 `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []shared.UnnamedSchemaRef12 `json:"submitted_urls"`
	JSON          intelPhishingURLSubmitJSON  `json:"-"`
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   IntelPhishingURLSubmit       `json:"result,required"`
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
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   IntelPhishingURLInfo         `json:"result,required"`
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
