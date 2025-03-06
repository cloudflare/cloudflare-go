// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ScreenshotService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScreenshotService] method instead.
type ScreenshotService struct {
	Options []option.RequestOption
}

// NewScreenshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScreenshotService(opts ...option.RequestOption) (r *ScreenshotService) {
	r = &ScreenshotService{}
	r.Options = opts
	return
}
