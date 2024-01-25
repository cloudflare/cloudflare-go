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

// AccountBrandProtectionSubmitService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountBrandProtectionSubmitService] method instead.
type AccountBrandProtectionSubmitService struct {
	Options []option.RequestOption
}

// NewAccountBrandProtectionSubmitService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountBrandProtectionSubmitService(opts ...option.RequestOption) (r *AccountBrandProtectionSubmitService) {
	r = &AccountBrandProtectionSubmitService{}
	r.Options = opts
	return
}

// Submit suspicious URL for scanning
func (r *AccountBrandProtectionSubmitService) PhishingURLScannerSubmitSuspiciousURLForScanning(ctx context.Context, accountIdentifier string, body AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams, opts ...option.RequestOption) (res *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse struct {
	Errors   []AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseError   `json:"errors"`
	Messages []AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessage `json:"messages"`
	Result   AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSuccess `json:"success"`
	JSON    accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseJSON    `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseError struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseErrorJSON `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseError]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessage struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessageJSON `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessage]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResult struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURL `json:"submitted_urls"`
	JSON          accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultJSON           `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResult]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURL struct {
	// URL that was excluded.
	URL  string                                                                                                    `json:"url"`
	JSON accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURLJSON `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURLJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURL]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                                                                                    `json:"url_id"`
	JSON  accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURLJSON `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURLJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURL]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                                                                                      `json:"url_id"`
	JSON  accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURLJSON `json:"-"`
}

// accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURLJSON
// contains the JSON metadata for the struct
// [AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURL]
type accountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseResultSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSuccess bool

const (
	AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSuccessTrue AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningResponseSuccess = true
)

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams struct {
	// URL(s) to filter submissions results by
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
