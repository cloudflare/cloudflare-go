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
	Categorizations []interface{} `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []interface{} `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []interface{} `json:"rule_matches"`
	ScanStatus  interface{}   `json:"scan_status"`
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

type Submit struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []interface{} `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []interface{} `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []interface{} `json:"submitted_urls"`
	JSON          submitJSON    `json:"-"`
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
	Errors   []interface{} `json:"errors,required"`
	Messages []interface{} `json:"messages,required"`
	// Whether the API call was successful
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
	Errors   []interface{} `json:"errors,required"`
	Messages []interface{} `json:"messages,required"`
	// Whether the API call was successful
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
