// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// PDFService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPDFService] method instead.
type PDFService struct {
	Options []option.RequestOption
}

// NewPDFService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPDFService(opts ...option.RequestOption) (r *PDFService) {
	r = &PDFService{}
	r.Options = opts
	return
}
