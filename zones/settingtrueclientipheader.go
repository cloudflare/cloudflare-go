// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingTrueClientIPHeaderService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingTrueClientIPHeaderService] method instead.
type SettingTrueClientIPHeaderService struct {
	Options []option.RequestOption
}

// NewSettingTrueClientIPHeaderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSettingTrueClientIPHeaderService(opts ...option.RequestOption) (r *SettingTrueClientIPHeaderService) {
	r = &SettingTrueClientIPHeaderService{}
	r.Options = opts
	return
}
