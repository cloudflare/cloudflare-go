// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// URLScannerService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLScannerService] method instead.
type URLScannerService struct {
	Options    []option.RequestOption
	Responses  *ResponseService
	Scans      *ScanService
	DOM        *DOMService
	HAR        *HARService
	Result     *ResultService
	Screenshot *ScreenshotService
}

// NewURLScannerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewURLScannerService(opts ...option.RequestOption) (r *URLScannerService) {
	r = &URLScannerService{}
	r.Options = opts
	r.Responses = NewResponseService(opts...)
	r.Scans = NewScanService(opts...)
	r.DOM = NewDOMService(opts...)
	r.HAR = NewHARService(opts...)
	r.Result = NewResultService(opts...)
	r.Screenshot = NewScreenshotService(opts...)
	return
}

// Submit URLs to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/ and
// take into account scans submitted in bulk have lower priority and may take
// longer to finish.
func (r *URLScannerService) Bulk(ctx context.Context, accountID string, body URLScannerBulkParams, opts ...option.RequestOption) (res *[]URLScannerBulkResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required accountId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type URLScannerDomain struct {
	ID              int64                `json:"id,required"`
	Name            string               `json:"name,required"`
	SuperCategoryID int64                `json:"super_category_id"`
	JSON            urlScannerDomainJSON `json:"-"`
}

// urlScannerDomainJSON contains the JSON metadata for the struct
// [URLScannerDomain]
type urlScannerDomainJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLScannerDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlScannerDomainJSON) RawJSON() string {
	return r.raw
}

type URLScannerBulkResponse struct {
	// URL to api report.
	API string `json:"api,required"`
	// URL to report.
	Result string `json:"result,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID.
	UUID string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                        `json:"visibility,required"`
	Options    URLScannerBulkResponseOptions `json:"options"`
	JSON       urlScannerBulkResponseJSON    `json:"-"`
}

// urlScannerBulkResponseJSON contains the JSON metadata for the struct
// [URLScannerBulkResponse]
type urlScannerBulkResponseJSON struct {
	API         apijson.Field
	Result      apijson.Field
	URL         apijson.Field
	UUID        apijson.Field
	Visibility  apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlScannerBulkResponseJSON) RawJSON() string {
	return r.raw
}

type URLScannerBulkResponseOptions struct {
	Useragent string                            `json:"useragent"`
	JSON      urlScannerBulkResponseOptionsJSON `json:"-"`
}

// urlScannerBulkResponseOptionsJSON contains the JSON metadata for the struct
// [URLScannerBulkResponseOptions]
type urlScannerBulkResponseOptionsJSON struct {
	Useragent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLScannerBulkResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlScannerBulkResponseOptionsJSON) RawJSON() string {
	return r.raw
}

type URLScannerBulkParams struct {
	// List of urls to scan (up to a 100).
	Body []URLScannerBulkParamsBody `json:"body,required"`
}

func (r URLScannerBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type URLScannerBulkParamsBody struct {
	URL         param.Field[string] `json:"url,required"`
	Customagent param.Field[string] `json:"customagent"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	Referer       param.Field[string]            `json:"referer"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]URLScannerBulkParamsBodyScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[URLScannerBulkParamsBodyVisibility] `json:"visibility"`
}

func (r URLScannerBulkParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Device resolutions.
type URLScannerBulkParamsBodyScreenshotsResolution string

const (
	URLScannerBulkParamsBodyScreenshotsResolutionDesktop URLScannerBulkParamsBodyScreenshotsResolution = "desktop"
	URLScannerBulkParamsBodyScreenshotsResolutionMobile  URLScannerBulkParamsBodyScreenshotsResolution = "mobile"
	URLScannerBulkParamsBodyScreenshotsResolutionTablet  URLScannerBulkParamsBodyScreenshotsResolution = "tablet"
)

func (r URLScannerBulkParamsBodyScreenshotsResolution) IsKnown() bool {
	switch r {
	case URLScannerBulkParamsBodyScreenshotsResolutionDesktop, URLScannerBulkParamsBodyScreenshotsResolutionMobile, URLScannerBulkParamsBodyScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type URLScannerBulkParamsBodyVisibility string

const (
	URLScannerBulkParamsBodyVisibilityPublic   URLScannerBulkParamsBodyVisibility = "Public"
	URLScannerBulkParamsBodyVisibilityUnlisted URLScannerBulkParamsBodyVisibility = "Unlisted"
)

func (r URLScannerBulkParamsBodyVisibility) IsKnown() bool {
	switch r {
	case URLScannerBulkParamsBodyVisibilityPublic, URLScannerBulkParamsBodyVisibilityUnlisted:
		return true
	}
	return false
}
