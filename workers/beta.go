// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// BetaService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaService] method instead.
type BetaService struct {
	Options []option.RequestOption
	Workers *BetaWorkerService
}

// NewBetaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaService(opts ...option.RequestOption) (r *BetaService) {
	r = &BetaService{}
	r.Options = opts
	r.Workers = NewBetaWorkerService(opts...)
	return
}
