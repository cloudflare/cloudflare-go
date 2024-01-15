// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountUrlscannerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountUrlscannerService] method
// instead.
type AccountUrlscannerService struct {
	Options    []option.RequestOption
	Scan       *AccountUrlscannerScanService
	Har        *AccountUrlscannerHarService
	Screenshot *AccountUrlscannerScreenshotService
}

// NewAccountUrlscannerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountUrlscannerService(opts ...option.RequestOption) (r *AccountUrlscannerService) {
	r = &AccountUrlscannerService{}
	r.Options = opts
	r.Scan = NewAccountUrlscannerScanService(opts...)
	r.Har = NewAccountUrlscannerHarService(opts...)
	r.Screenshot = NewAccountUrlscannerScreenshotService(opts...)
	return
}
