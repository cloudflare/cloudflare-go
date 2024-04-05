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
func (r *BrandProtectionService) Submit(ctx context.Context, params BrandProtectionSubmitParams, opts ...option.RequestOption) (res *Submit, err error) {
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
func (r *BrandProtectionService) URLInfo(ctx context.Context, params BrandProtectionURLInfoParams, opts ...option.RequestOption) (res *Info, err error) {
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

type Info struct {
	// List of categorizations applied to this submission.
	Categorizations []UnnamedSchemaRef209db30ed499548152d6f3bccf720b54 `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []UnnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014 `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []UnnamedSchemaRef3e10ea08deb8102a27500f986488c1e8 `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus UnnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869 `json:"scan_status"`
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
	ExcludedURLs []UnnamedSchemaRef767c0981cf47f45f0c766253dbd18669 `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []UnnamedSchemaRef44e66100b948bfe723054c56b6144766 `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []UnnamedSchemaRef39419d70e2399b28b15cd660afd342fb `json:"submitted_urls"`
	JSON          submitJSON                                         `json:"-"`
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

type UnnamedSchemaRef209db30ed499548152d6f3bccf720b54 struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                                               `json:"verification_status"`
	JSON               unnamedSchemaRef209db30ed499548152d6f3bccf720b54JSON `json:"-"`
}

// unnamedSchemaRef209db30ed499548152d6f3bccf720b54JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef209db30ed499548152d6f3bccf720b54]
type unnamedSchemaRef209db30ed499548152d6f3bccf720b54JSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UnnamedSchemaRef209db30ed499548152d6f3bccf720b54) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef209db30ed499548152d6f3bccf720b54JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef39419d70e2399b28b15cd660afd342fb struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                                `json:"url_id"`
	JSON  unnamedSchemaRef39419d70e2399b28b15cd660afd342fbJSON `json:"-"`
}

// unnamedSchemaRef39419d70e2399b28b15cd660afd342fbJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef39419d70e2399b28b15cd660afd342fb]
type unnamedSchemaRef39419d70e2399b28b15cd660afd342fbJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef39419d70e2399b28b15cd660afd342fb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef39419d70e2399b28b15cd660afd342fbJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef3e10ea08deb8102a27500f986488c1e8 struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                                               `json:"name"`
	JSON unnamedSchemaRef3e10ea08deb8102a27500f986488c1e8JSON `json:"-"`
}

// unnamedSchemaRef3e10ea08deb8102a27500f986488c1e8JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef3e10ea08deb8102a27500f986488c1e8]
type unnamedSchemaRef3e10ea08deb8102a27500f986488c1e8JSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef3e10ea08deb8102a27500f986488c1e8) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef3e10ea08deb8102a27500f986488c1e8JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef44e66100b948bfe723054c56b6144766 struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                                `json:"url_id"`
	JSON  unnamedSchemaRef44e66100b948bfe723054c56b6144766JSON `json:"-"`
}

// unnamedSchemaRef44e66100b948bfe723054c56b6144766JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef44e66100b948bfe723054c56b6144766]
type unnamedSchemaRef44e66100b948bfe723054c56b6144766JSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef44e66100b948bfe723054c56b6144766) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef44e66100b948bfe723054c56b6144766JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef767c0981cf47f45f0c766253dbd18669 struct {
	// URL that was excluded.
	URL  string                                               `json:"url"`
	JSON unnamedSchemaRef767c0981cf47f45f0c766253dbd18669JSON `json:"-"`
}

// unnamedSchemaRef767c0981cf47f45f0c766253dbd18669JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef767c0981cf47f45f0c766253dbd18669]
type unnamedSchemaRef767c0981cf47f45f0c766253dbd18669JSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef767c0981cf47f45f0c766253dbd18669) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef767c0981cf47f45f0c766253dbd18669JSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014 struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                                              `json:"model_score"`
	JSON       unnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014JSON `json:"-"`
}

// unnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014]
type unnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014JSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef9b4c9779a35b172cb69c71389ebc7014JSON) RawJSON() string {
	return r.raw
}

// Status of the most recent scan found.
type UnnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869 struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                                                `json:"submission_id"`
	JSON         unnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869JSON `json:"-"`
}

// unnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869]
type unnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869JSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UnnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefA64e2a18a86750b6bd72cdf37ecfd869JSON) RawJSON() string {
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Submit                                                    `json:"result,required"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   Info                                                      `json:"result,required"`
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
