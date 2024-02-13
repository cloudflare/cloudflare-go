// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// BrandProtectionSubmitService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBrandProtectionSubmitService]
// method instead.
type BrandProtectionSubmitService struct {
	Options []option.RequestOption
}

// NewBrandProtectionSubmitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrandProtectionSubmitService(opts ...option.RequestOption) (r *BrandProtectionSubmitService) {
	r = &BrandProtectionSubmitService{}
	r.Options = opts
	return
}

// Submit suspicious URL for scanning
func (r *BrandProtectionSubmitService) PhishingURLScannerSubmitSuspiciousURLForScanning(ctx context.Context, accountID string, body BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams, opts ...option.RequestOption) (res *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelope
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURL `json:"submitted_urls"`
	JSON          brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseJSON           `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURL struct {
	// URL that was excluded.
	URL  string                                                                                       `json:"url"`
	JSON brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURLJSON `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURLJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURL]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                                                                       `json:"url_id"`
	JSON  brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURLJSON `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURLJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURL]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                                                                         `json:"url_id"`
	JSON  brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURLJSON `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURLJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURL]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams struct {
	// URL(s) to filter submissions results by
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelope struct {
	Errors   []BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessages `json:"messages,required"`
	Result   BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeSuccess `json:"success,required"`
	JSON    brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelope]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrors struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrorsJSON `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrors]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessages struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessagesJSON `json:"-"`
}

// brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessages]
type brandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeSuccess bool

const (
	BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeSuccessTrue BrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseEnvelopeSuccess = true
)
