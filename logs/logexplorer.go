// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// LogExplorerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogExplorerService] method instead.
type LogExplorerService struct {
	Options  []option.RequestOption
	Query    *LogExplorerQueryService
	Datasets *LogExplorerDatasetService
}

// NewLogExplorerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogExplorerService(opts ...option.RequestOption) (r *LogExplorerService) {
	r = &LogExplorerService{}
	r.Options = opts
	r.Query = NewLogExplorerQueryService(opts...)
	r.Datasets = NewLogExplorerDatasetService(opts...)
	return
}
