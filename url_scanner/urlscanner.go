// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
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
