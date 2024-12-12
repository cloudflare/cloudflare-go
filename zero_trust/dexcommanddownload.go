// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// DEXCommandDownloadService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandDownloadService] method instead.
type DEXCommandDownloadService struct {
	Options []option.RequestOption
}

// NewDEXCommandDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXCommandDownloadService(opts ...option.RequestOption) (r *DEXCommandDownloadService) {
	r = &DEXCommandDownloadService{}
	r.Options = opts
	return
}
