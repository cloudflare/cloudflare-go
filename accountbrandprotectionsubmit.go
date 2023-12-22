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
func (r *AccountBrandProtectionSubmitService) PhishingURLScannerSubmitSuspiciousURLForScanning(ctx context.Context, accountIdentifier string, body AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams, opts ...option.RequestOption) (res *PhishingURLSubmitSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PhishingURLSubmitSingleResponse struct {
	Errors   []PhishingURLSubmitSingleResponseError   `json:"errors"`
	Messages []PhishingURLSubmitSingleResponseMessage `json:"messages"`
	Result   PhishingURLSubmitSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success PhishingURLSubmitSingleResponseSuccess `json:"success"`
	JSON    phishingURLSubmitSingleResponseJSON    `json:"-"`
}

// phishingURLSubmitSingleResponseJSON contains the JSON metadata for the struct
// [PhishingURLSubmitSingleResponse]
type phishingURLSubmitSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PhishingURLSubmitSingleResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    phishingURLSubmitSingleResponseErrorJSON `json:"-"`
}

// phishingURLSubmitSingleResponseErrorJSON contains the JSON metadata for the
// struct [PhishingURLSubmitSingleResponseError]
type phishingURLSubmitSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PhishingURLSubmitSingleResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    phishingURLSubmitSingleResponseMessageJSON `json:"-"`
}

// phishingURLSubmitSingleResponseMessageJSON contains the JSON metadata for the
// struct [PhishingURLSubmitSingleResponseMessage]
type phishingURLSubmitSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PhishingURLSubmitSingleResponseResult struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []PhishingURLSubmitSingleResponseResultExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []PhishingURLSubmitSingleResponseResultSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []PhishingURLSubmitSingleResponseResultSubmittedURL `json:"submitted_urls"`
	JSON          phishingURLSubmitSingleResponseResultJSON           `json:"-"`
}

// phishingURLSubmitSingleResponseResultJSON contains the JSON metadata for the
// struct [PhishingURLSubmitSingleResponseResult]
type phishingURLSubmitSingleResponseResultJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PhishingURLSubmitSingleResponseResultExcludedURL struct {
	// URL that was excluded.
	URL  string                                               `json:"url"`
	JSON phishingURLSubmitSingleResponseResultExcludedURLJSON `json:"-"`
}

// phishingURLSubmitSingleResponseResultExcludedURLJSON contains the JSON metadata
// for the struct [PhishingURLSubmitSingleResponseResultExcludedURL]
type phishingURLSubmitSingleResponseResultExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponseResultExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PhishingURLSubmitSingleResponseResultSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                               `json:"url_id"`
	JSON  phishingURLSubmitSingleResponseResultSkippedURLJSON `json:"-"`
}

// phishingURLSubmitSingleResponseResultSkippedURLJSON contains the JSON metadata
// for the struct [PhishingURLSubmitSingleResponseResultSkippedURL]
type phishingURLSubmitSingleResponseResultSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponseResultSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PhishingURLSubmitSingleResponseResultSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                                 `json:"url_id"`
	JSON  phishingURLSubmitSingleResponseResultSubmittedURLJSON `json:"-"`
}

// phishingURLSubmitSingleResponseResultSubmittedURLJSON contains the JSON metadata
// for the struct [PhishingURLSubmitSingleResponseResultSubmittedURL]
type phishingURLSubmitSingleResponseResultSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishingURLSubmitSingleResponseResultSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PhishingURLSubmitSingleResponseSuccess bool

const (
	PhishingURLSubmitSingleResponseSuccessTrue PhishingURLSubmitSingleResponseSuccess = true
)

type AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams struct {
	// URL(s) to filter submissions results by
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r AccountBrandProtectionSubmitPhishingURLScannerSubmitSuspiciousURLForScanningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
