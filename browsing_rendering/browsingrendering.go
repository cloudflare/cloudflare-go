// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browsing_rendering

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// BrowsingRenderingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBrowsingRenderingService] method instead.
type BrowsingRenderingService struct {
	Options    []option.RequestOption
	Content    *ContentService
	Pdf        *PdfService
	Scrape     *ScrapeService
	Screenshot *ScreenshotService
	Snapshot   *SnapshotService
}

// NewBrowsingRenderingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBrowsingRenderingService(opts ...option.RequestOption) (r *BrowsingRenderingService) {
	r = &BrowsingRenderingService{}
	r.Options = opts
	r.Content = NewContentService(opts...)
	r.Pdf = NewPdfService(opts...)
	r.Scrape = NewScrapeService(opts...)
	r.Screenshot = NewScreenshotService(opts...)
	r.Snapshot = NewSnapshotService(opts...)
	return
}
