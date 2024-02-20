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
func (r *BrandProtectionSubmitService) New(ctx context.Context, accountID string, body BrandProtectionSubmitNewParams, opts ...option.RequestOption) (res *BrandProtectionSubmitNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env BrandProtectionSubmitNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BrandProtectionSubmitNewResponse struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []BrandProtectionSubmitNewResponseExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []BrandProtectionSubmitNewResponseSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []BrandProtectionSubmitNewResponseSubmittedURL `json:"submitted_urls"`
	JSON          brandProtectionSubmitNewResponseJSON           `json:"-"`
}

// brandProtectionSubmitNewResponseJSON contains the JSON metadata for the struct
// [BrandProtectionSubmitNewResponse]
type brandProtectionSubmitNewResponseJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitNewResponseExcludedURL struct {
	// URL that was excluded.
	URL  string                                          `json:"url"`
	JSON brandProtectionSubmitNewResponseExcludedURLJSON `json:"-"`
}

// brandProtectionSubmitNewResponseExcludedURLJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitNewResponseExcludedURL]
type brandProtectionSubmitNewResponseExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponseExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitNewResponseSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                          `json:"url_id"`
	JSON  brandProtectionSubmitNewResponseSkippedURLJSON `json:"-"`
}

// brandProtectionSubmitNewResponseSkippedURLJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitNewResponseSkippedURL]
type brandProtectionSubmitNewResponseSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponseSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitNewResponseSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                            `json:"url_id"`
	JSON  brandProtectionSubmitNewResponseSubmittedURLJSON `json:"-"`
}

// brandProtectionSubmitNewResponseSubmittedURLJSON contains the JSON metadata for
// the struct [BrandProtectionSubmitNewResponseSubmittedURL]
type brandProtectionSubmitNewResponseSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponseSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitNewParams struct {
	// URL(s) to filter submissions results by
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r BrandProtectionSubmitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BrandProtectionSubmitNewResponseEnvelope struct {
	Errors   []BrandProtectionSubmitNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []BrandProtectionSubmitNewResponseEnvelopeMessages `json:"messages,required"`
	Result   BrandProtectionSubmitNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success BrandProtectionSubmitNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    brandProtectionSubmitNewResponseEnvelopeJSON    `json:"-"`
}

// brandProtectionSubmitNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [BrandProtectionSubmitNewResponseEnvelope]
type brandProtectionSubmitNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    brandProtectionSubmitNewResponseEnvelopeErrorsJSON `json:"-"`
}

// brandProtectionSubmitNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [BrandProtectionSubmitNewResponseEnvelopeErrors]
type brandProtectionSubmitNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BrandProtectionSubmitNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    brandProtectionSubmitNewResponseEnvelopeMessagesJSON `json:"-"`
}

// brandProtectionSubmitNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [BrandProtectionSubmitNewResponseEnvelopeMessages]
type brandProtectionSubmitNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BrandProtectionSubmitNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BrandProtectionSubmitNewResponseEnvelopeSuccess bool

const (
	BrandProtectionSubmitNewResponseEnvelopeSuccessTrue BrandProtectionSubmitNewResponseEnvelopeSuccess = true
)
