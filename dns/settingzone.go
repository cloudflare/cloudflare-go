// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// SettingZoneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingZoneService] method instead.
type SettingZoneService struct {
	Options []option.RequestOption
	Views   *SettingZoneViewService
}

// NewSettingZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZoneService(opts ...option.RequestOption) (r *SettingZoneService) {
	r = &SettingZoneService{}
	r.Options = opts
	r.Views = NewSettingZoneViewService(opts...)
	return
}
