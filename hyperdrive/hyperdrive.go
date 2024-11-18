// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperdrive

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// HyperdriveService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHyperdriveService] method instead.
type HyperdriveService struct {
	Options []option.RequestOption
	Configs *ConfigService
}

// NewHyperdriveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHyperdriveService(opts ...option.RequestOption) (r *HyperdriveService) {
	r = &HyperdriveService{}
	r.Options = opts
	r.Configs = NewConfigService(opts...)
	return
}
