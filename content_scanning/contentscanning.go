// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ContentScanningService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentScanningService] method instead.
type ContentScanningService struct {
	Options  []option.RequestOption
	Payloads *PayloadService
	Settings *SettingService
}

// NewContentScanningService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContentScanningService(opts ...option.RequestOption) (r *ContentScanningService) {
	r = &ContentScanningService{}
	r.Options = opts
	r.Payloads = NewPayloadService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}
